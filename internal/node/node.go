package node

import (
	"container/list"
	"context"
	"errors"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
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
	*state.Emitter

	// Host and context
	host *host.SNRHost

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
func NewNode(ctx context.Context, host *host.SNRHost, loc *common.Location) (*Node, error) {
	// Initialize Node
	node := &Node{
		Emitter: state.NewEmitter(2048),
		host:    host,
		ctx:     ctx,
		queue:   list.New(),
	}
	// Set Transfer Protocol
	node.TransferProtocol = transfer.NewProtocol(ctx, host, node.Emitter)

	// Set Exchange Protocol
	exch, err := exchange.NewProtocol(ctx, host, node.Emitter)
	if err != nil {
		logger.Error("Failed to start ExchangeProtocol", zap.Error(err))
		return nil, err
	}
	node.ExchangeProtocol = exch

	// Set Lobby Protocol
	lobby, err := lobby.NewProtocol(host, loc, node.Emitter)
	if err != nil {
		logger.Error("Failed to start LobbyProtocol", zap.Error(err))
		return nil, err
	}
	node.LobbyProtocol = lobby
	return node, nil
}

// Edit method updates Node's profile
func (n *Node) Edit(p *common.Profile) error {
	// Set Profile and Fetch User Peer
	n.profile = p
	peer, err := n.Peer()
	if err != nil {
		return err
	}

	// Push Update to Exchange
	err = n.ExchangeProtocol.Update(peer)
	if err != nil {
		logger.Error("Failed to update Exchange", zap.Error(err))
		return err
	}

	// Push Update to Lobby
	err = n.LobbyProtocol.Update(peer)
	if err != nil {
		logger.Error("Failed to update Lobby", zap.Error(err))
		return err
	}
	return nil
}

// Supply a transfer item to the queue
func (n *Node) Supply(paths []string) error {
	// Create Transfer
	payload, err := common.NewPayload(n.profile, paths)
	if err != nil {
		logger.Error("Failed to Supply Paths", zap.Error(err))
		return err
	}

	// Add items to transfer
	n.queue.PushBack(payload)
	return nil
}

// Share a peer to have a transfer
func (n *Node) Share(to *common.Peer) error {
	// Fetch Element from Queue
	elem := n.queue.Front()
	if elem != nil {
		// Get Payload
		payload := n.queue.Remove(elem).(*common.Payload)

		// Create New ID for Invite
		id, err := n.host.NewID()
		if err != nil {
			logger.Error("Failed to create new id for Shared Invite", zap.Error(err))
			return err
		}

		// Create new Metadata
		meta, err := n.host.NewMetadata()
		if err != nil {
			logger.Error("Failed to create new metadata for Shared Invite", zap.Error(err))
			return err
		}

		// Fetch User Peer
		from, err := n.Peer()
		if err != nil {
			return err
		}

		// Create Invite Request
		req := &transfer.InviteRequest{
			Payload:  payload,
			Metadata: meta,
			To:       to,
			From:     from,
			Uuid:     id,
		}

		// Fetch Peer ID from Public Key
		toId, err := to.PeerID()
		if err != nil {
			logger.Error("Failed to fetch peer id from public key", zap.Error(err))
			return err
		}

		// Request Peer to Transfer File
		err = n.TransferProtocol.Request(toId, req)
		if err != nil {
			logger.Error("Failed to invite peer", zap.Error(err))
			n.Emit(Event_STATUS, err)
			return err
		}
		return nil
	}
	return errors.New("No items in Transfer Queue.")
}

// Respond to an invite request
func (n *Node) Respond(decs bool, to *common.Peer) error {
	// Create new Metadata
	meta, err := n.host.NewMetadata()
	if err != nil {
		logger.Error("Failed to create new metadata for Shared Invite", zap.Error(err))
		return err
	}

	// Fetch User Peer
	from, err := n.Peer()
	if err != nil {
		return err
	}

	// Create Invite Response
	resp := &transfer.InviteResponse{
		Decision: decs,
		Metadata: meta,
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.PeerID()
	if err != nil {
		logger.Error("Failed to fetch peer id from public key", zap.Error(err))
		return err
	}

	// Respond on TransferProtocol
	err = n.TransferProtocol.Respond(toId, resp)
	if err != nil {
		logger.Error("Failed to respond to invite", zap.Error(err))
		n.Emit(Event_STATUS, err)
		return err
	}
	return nil
}
