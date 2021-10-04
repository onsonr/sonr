package node

import (
	"container/list"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/internal/store"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Emitter is the event emitter for this node
	*state.Emitter

	// Host and context
	host *host.SNRHost

	// Properties
	ctx context.Context

	// Persistent Database
	store *store.Store

	// Queue - the transfer queue
	queue *list.List

	// Type - the type of node
	nodeType NodeType

	// TransferProtocol - the transfer protocol
	*transfer.TransferProtocol

	// ExchangeProtocol - the exchange protocol
	*exchange.ExchangeProtocol

	// LobbyProtocol - The lobby protocol
	*lobby.LobbyProtocol
}

// NewNode Creates a node with its implemented protocols
func NewNode(ctx context.Context, opts ...NodeOption) (*Node, *InitializeResponse, error) {
	// Set Node Options
	config := defaultNodeOptions()
	for _, opt := range opts {
		opt(config)
	}

	// Get Device KeyChain Account Key
	privKey, err := device.KeyChain.GetPrivKey(keychain.Account)
	if err != nil {
		return nil, nil, logger.Error("Failed to get private key", err)
	}

	// Initialize Host
	host, err := host.NewHost(ctx, config.GetConnection(), privKey)
	if err != nil {
		return nil, nil, logger.Error("Failed to initialize host", err)
	}

	// Create Node
	node := &Node{
		Emitter:  state.NewEmitter(2048),
		host:     host,
		ctx:      ctx,
		queue:    list.New(),
		nodeType: config.GetNodeType(),
	}

	// Check Config for ClientNode
	if config.isClient {
		node.startClientService(ctx, config.GetLocation(), config.GetProfile())
	}

	// Check Config for HighwayNode
	if config.isHighway {
		// Get Client, Secret Keys
		cKey, sKey, err := config.GetNBKeys()
		if err != nil {
			return nil, nil, logger.Error("Failed to start Highway Service", err)
		}
		node.startHighwayService(ctx, cKey, sKey)
	}

	// Start Background Refresh and Return Response
	go node.pushAutomaticPings()
	return node, node.newInitResponse(nil), nil
}

// Edit method updates Node's profile
func (n *Node) Edit(p *common.Profile) error {
	// Set Profile and Fetch User Peer
	p.LastModified = time.Now().Unix()
	err := n.store.SetProfile(p)
	if err != nil {
		return err
	}

	// Get Peer
	peer, err := n.Peer()
	if err != nil {
		return logger.Error("Failed to get Node Peer Object", err)
	}

	// Push Update to Exchange
	err = n.ExchangeProtocol.Update(peer)
	if err != nil {
		return logger.Error("Failed to update Exchange", err)
	}

	// Push Update to Lobby
	err = n.LobbyProtocol.Update(peer)
	if err != nil {
		return logger.Error("Failed to update Lobby", err)
	}
	return nil
}

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	// Get Public Key
	pubKey, err := device.KeyChain.GetSnrPubKey(keychain.Account)
	if err != nil {
		return nil, logger.Error("Failed to get Public Key", err)
	}

	// Find PublicKey Buffer
	deviceStat, err := device.Stat()
	if err != nil {
		return nil, logger.Error("Failed to get device Stat", err)
	}

	// Marshal Public Key
	pubBuf, err := pubKey.Buffer()
	if err != nil {
		return nil, logger.Error("Failed to marshal public key", err)
	}

	// Get Profile
	profile, err := n.Profile()
	if err != nil {
		return nil, err
	}

	// Return Peer
	return &common.Peer{
		SName:     strings.ToLower(profile.SName),
		Status:    common.Peer_ONLINE,
		Profile:   profile,
		PublicKey: pubBuf,
		Device: &common.Peer_Device{
			HostName: deviceStat.HostName,
			Os:       deviceStat.Os,
			Id:       deviceStat.Id,
			Arch:     deviceStat.Arch,
		},
	}, nil
}

// Profile method returns the profile of the node
func (n *Node) Profile() (*common.Profile, error) {
	pro, err := n.store.GetProfile()
	if err != nil {
		return nil, logger.Error("Failed to retreive Profile", err)
	}
	return pro, nil
}

// Recents method returns the recent peers of the node
func (n *Node) Recents() (store.RecentsHistory, error) {
	rec, err := n.store.GetRecents()
	if err != nil {
		return nil, logger.Error("Failed to get recents", err)
	}
	return rec, nil
}

// Supply a transfer item to the queue
func (n *Node) Supply(paths []string) error {
	// Get Profile
	profile, err := n.store.GetProfile()
	if err != nil {
		return err
	}

	// Create Transfer
	payload, err := common.NewPayload(profile, paths)
	if err != nil {
		return logger.Error("Failed to Supply Paths", err)
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
		id, err := device.KeyChain.CreateUUID()
		if err != nil {
			return logger.Error("Failed to create new id for Shared Invite", err)
		}

		// Create new Metadata
		meta, err := device.KeyChain.CreateMetadata(n.host.ID())
		if err != nil {
			return logger.Error("Failed to create new metadata for Shared Invite", err)
		}

		// Fetch User Peer
		from, err := n.Peer()
		if err != nil {
			return logger.Error("Failed to get Node Peer Object", err)
		}

		// Create Invite Request
		req := &transfer.InviteRequest{
			Payload:  payload,
			Metadata: common.SignedMetadataToProto(meta),
			To:       to,
			From:     from,
			Uuid:     common.SignedUUIDToProto(id),
		}

		// Fetch Peer ID from Public Key
		toId, err := to.PeerID()
		if err != nil {
			return logger.Error("Failed to fetch peer id from public key", err)
		}

		// Request Peer to Transfer File
		err = n.TransferProtocol.Request(toId, req)
		if err != nil {
			return logger.Error("Failed to invite peer", err)
		}
		return nil
	}
	return errors.New("No items in Transfer Queue.")
}

// Respond to an invite request
func (n *Node) Respond(decs bool, to *common.Peer) error {
	// Create new Metadata
	meta, err := device.KeyChain.CreateMetadata(n.host.ID())
	if err != nil {
		logger.Error("Failed to create new metadata for Shared Invite", err)
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
		Metadata: common.SignedMetadataToProto(meta),
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.PeerID()
	if err != nil {
		logger.Error("Failed to fetch peer id from public key", err)
		return err
	}

	// Respond on TransferProtocol
	err = n.TransferProtocol.Respond(toId, resp)
	if err != nil {
		logger.Error("Failed to respond to invite", err)
		return err
	}
	return nil
}

// Stat returns the Node info as StatResponse
func (n *Node) Stat() (*StatResponse, error) {
	// Get Profile
	profile, err := n.store.GetProfile()
	if err != nil {
		return &StatResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Get Host Stats
	hStat, err := n.host.Stat()
	if err != nil {
		return &StatResponse{
			Success: false,
			Error:   err.Error(),
			SName:   profile.SName,
			Profile: profile,
		}, logger.Error("Failed to get Host Stat", err)
	}

	// Get Device Stat
	dStat, err := device.Stat()
	if err != nil {
		return &StatResponse{
			Success: false,
			Error:   err.Error(),
			SName:   profile.SName,
			Profile: profile,
			Network: &StatResponse_Network{
				PublicKey: hStat.PublicKey,
				PeerID:    hStat.PeerID,
				Multiaddr: hStat.MultAddr,
			},
		}, logger.Error("Failed to get Device Stat", err)
	}

	// Return StatResponse
	return &StatResponse{
		SName:   profile.SName,
		Profile: profile,
		Network: &StatResponse_Network{
			PublicKey: hStat.PublicKey,
			PeerID:    hStat.PeerID,
			Multiaddr: hStat.MultAddr,
		},
		Device: &StatResponse_Device{
			Id:        dStat.Id,
			Name:      dStat.HostName,
			Os:        dStat.Os,
			Arch:      dStat.Arch,
			IsDesktop: dStat.IsDesktop,
			IsMobile:  dStat.IsMobile,
		},
	}, nil
}

func (n *Node) pushAutomaticPings() {
	for {
		select {
		case <-time.After(time.Second * 5):
			p, err := n.Peer()
			if err != nil {
				logger.Error("Failed to push Auto Ping", err)
				continue
			}
			n.LobbyProtocol.Update(p)
			n.ExchangeProtocol.Update(p)

		case <-n.ctx.Done():
			return
		}
	}
}
