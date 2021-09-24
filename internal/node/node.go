package node

import (
	"container/list"
	"context"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
)

// Node Emission Events
const (
	Event_STATUS = "status"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Emitter is the event emitter for this node
	*emitter.Emitter

	// StateMachine state
	*state.StateMachine

	// Host and context
	*host.SNRHost

	// Properties
	ctx context.Context

	// Queue - the transfer queue
	queue *list.List

	// Profile - the node's profile
	profile *common.Profile

	// TransferProtocol - the transfer protocol
	*transfer.TransferProtocol

	// ExchangeProtocol - the exchange protocol
	*exchange.ExchangeProtocol

	// LobbyProtocol - The lobby protocol
	*lobby.LobbyProtocol
}

// NewNode Creates a node with its implemented protocols
func NewNode(ctx context.Context, host *host.SNRHost, loc *common.Location) *Node {
	// Initialize Node
	node := &Node{
		Emitter: emitter.New(2048),
		SNRHost: host,
		ctx:     ctx,
		queue:   list.New(),
	}

	// Set Transfer Protocol
	node.TransferProtocol = transfer.NewProtocol(host, node.Emitter)
	node.Emit(Event_STATUS, true, "Transfer Protocol Set")

	// Set Exchange Protocol
	exch, err := exchange.NewProtocol(ctx, host, node.Emitter)
	if err != nil {
		logger.Error("Failed to start ExchangeProtocol", zap.Error(err))
		return node
	}
	node.Emit(Event_STATUS, true, "Exchange Protocol Set")
	node.ExchangeProtocol = exch

	// Set Lobby Protocol
	lobby, err := lobby.NewProtocol(host, loc, node.Emitter)
	if err != nil {
		logger.Error("Failed to start LobbyProtocol", zap.Error(err))
		return node
	}
	node.Emit(Event_STATUS, true, "Lobby Protocol Set")
	node.LobbyProtocol = lobby
	return node
}

// Peer method returns the peer of the node
func (n *Node) Peer() *common.Peer {
	// Find PublicKey Buffer
	pubBuf, err := crypto.MarshalPublicKey(n.SNRHost.PublicKey())
	if err != nil {
		logger.Error("Failed to marshal public key", zap.Error(err))
		return nil
	}

	// Return Peer
	return &common.Peer{
		SName:     n.profile.SName,
		Status:    common.Peer_ONLINE,
		Info:      device.Info(),
		Profile:   n.profile,
		PublicKey: pubBuf,
	}
}

// Edit method updates Node's profile
func (n *Node) Edit(p *common.Profile) error {
	// Set Profile
	n.profile = p

	// Push Update to Exchange
	err := n.ExchangeProtocol.Update(n.Peer())
	if err != nil {
		logger.Error("Failed to update Exchange", zap.Error(err))
		return err
	}

	return nil
}

// Supply a transfer item to the queue
func (n *Node) Supply(paths []string) error {
	// Create Transfer
	tr := common.Payload{
		Metadata: n.NewMetadata(),
	}

	// Initialize Transfer Items and add iterate over paths
	items := make([]*common.Payload_Item, len(paths))
	for _, path := range paths {
		// Check if path is a url
		if common.IsUrl(path) {
			items = append(items, common.NewTransferUrlItem(path))
		} else {
			// Create File Item
			item, err := common.NewTransferFileItem(path)
			if err != nil {
				logger.Error("Failed to edit Profile", zap.Error(err))
				n.Emit(Event_STATUS, err)
				return err
			}

			// Add item to transfer
			items = append(items, item)
		}
	}

	// Add items to transfer
	tr.Items = items
	n.queue.PushBack(&tr)
	return nil
}

// Share a peer to have a transfer
func (n *Node) Share(peer *common.Peer) error {
	// Create Invite Request
	req := &transfer.InviteRequest{
		Payload:  n.queue.Front().Value.(*common.Payload),
		Metadata: n.NewMetadata(),
		To:       peer,
		From:     n.Peer(),
	}

	// Fetch Peer ID from Exchange
	id, err := n.ExchangeProtocol.Find(peer.GetSName())
	if err != nil {
		logger.Error("Failed to search peer", zap.Error(err))
		return err
	}

	// Invite peer
	err = n.TransferProtocol.Request(id, req)
	if err != nil {
		logger.Error("Failed to invite peer", zap.Error(err))
		n.Emit(Event_STATUS, err)
		return err
	}
	return nil
}

// Respond to an invite request
func (n *Node) Respond(decs bool) error {
	// Create Invite Response
	var resp *transfer.InviteResponse
	if decs {
		resp = &transfer.InviteResponse{Success: true}
	} else {
		resp = &transfer.InviteResponse{Success: false}
	}

	// Respond on TransferProtocol
	err := n.TransferProtocol.Respond(resp)
	if err != nil {
		logger.Error("Failed to respond to invite", zap.Error(err))
		n.Emit(Event_STATUS, err)
		return err
	}
	return nil
}
