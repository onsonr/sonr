package node

import (
	"container/list"
	"context"
	"errors"
	"strings"
	"time"

	"git.mills.io/prologic/bitcask"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/state"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	common.NodeImpl
	
	// Emitter is the event emitter for this node
	*state.Emitter

	// Host and context
	host *host.SNRHost

	// Properties
	ctx context.Context

	// Persistent Database
	store *bitcask.Bitcask

	// Node Stub Interface
	stub NodeStub

	// Queue - the transfer queue
	queue *list.List

	// Channels
	// TransferProtocol - decisionEvents
	decisionEvents chan *api.DecisionEvent

	// LobbyProtocol - refreshEvents
	refreshEvents chan *api.RefreshEvent

	// MailboxProtocol - mailEvents
	mailEvents chan *api.MailboxEvent

	// TransferProtocol - inviteEvents
	inviteEvents chan *api.InviteEvent

	// TransferProtocol - progressEvents
	progressEvents chan *api.ProgressEvent

	// TransferProtocol - completeEvents
	completeEvents chan *api.CompleteEvent
}

// NewNode Creates a node with its implemented protocols
func NewNode(ctx context.Context, options ...NodeOption) (common.NodeImpl, *api.InitializeResponse, error) {
	// Set Node Options
	opts := defaultNodeOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Create Node
	node := &Node{
		Emitter:        state.NewEmitter(2048),
		ctx:            ctx,
		queue:          list.New(),
		decisionEvents: make(chan *api.DecisionEvent),
		refreshEvents:  make(chan *api.RefreshEvent),
		inviteEvents:   make(chan *api.InviteEvent),
		mailEvents:     make(chan *api.MailboxEvent),
		progressEvents: make(chan *api.ProgressEvent),
		completeEvents: make(chan *api.CompleteEvent),
	}

	// Initialize Host
	host, err := host.NewHost(ctx, node.Emitter, host.WithConnection(opts.connection))
	if err != nil {
		logger.Error("Failed to initialize host", err)
		return nil, api.NewInitialzeResponse(nil, false), err
	}
	node.host = host

	// Open Store with profileBuf
	err = node.openStore(ctx, opts)
	if err != nil {
		logger.Error("Failed to open database", err)
		return node, api.NewInitialzeResponse(nil, false), err
	}

	// Initialize Stub
	err = opts.Apply(node.ctx, node)
	if err != nil {
		logger.Error("Failed to initialize stub", err)
		return nil, api.NewInitialzeResponse(nil, false), err
	}

	// Begin Background Tasks
	go node.Serve(ctx)
	return node, api.NewInitialzeResponse(node.Profile, false), nil
}

// Close closes the node
func (n *Node) Close() {
	// Close Store
	if err := n.store.Close(); err != nil {
		logger.Error("Failed to close store", err)
	}

	// Close Stub
	if err := n.stub.Close(); err != nil {
		logger.Error("Failed to close host", err)
	}

	// Close Host
	if err := n.host.Close(); err != nil {
		logger.Error("Failed to close host", err)
	}
}

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	// Get Profile
	profile, err := n.Profile()
	if err != nil {
		logger.Warn("Failed to get profile from Memory store, using DefaultProfile.", err)
	}

	// Get Public Key
	pubKey, err := device.KeyChain.GetSnrPubKey(keychain.Account)
	if err != nil {
		logger.Error("Failed to get Public Key", err)
		return nil, err
	}

	// Marshal Public Key
	pubBuf, err := pubKey.Buffer()
	if err != nil {
		logger.Error("Failed to marshal public key", err)
		return nil, err
	}

	stat, err := device.Stat()
	if err != nil {
		logger.Error("Failed to get device stat", err)
		return nil, err
	}
	// Return Peer
	return &common.Peer{
		SName:        strings.ToLower(profile.GetSName()),
		Status:       common.Peer_ONLINE,
		Profile:      profile,
		PublicKey:    pubBuf,
		PeerID:       n.host.ID().String(),
		LastModified: time.Now().Unix(),
		Device: &common.Peer_Device{
			HostName: stat.HostName,
			Os:       stat.Os,
			Id:       stat.Id,
			Arch:     stat.Arch,
		},
	}, nil
}

// Supply a transfer item to the queue
func (n *Node) Supply(paths []string) error {
	// Get Profile
	profile, err := n.Profile()
	if err != nil {
		return err
	}

	// Create Transfer
	payload, err := common.NewPayload(profile, paths)
	if err != nil {
		logger.Error("Failed to Supply Paths", err)
		return err
	}

	// Add items to transfer
	n.queue.PushBack(payload)
	return nil
}

// Serve handles the emitter events.
func (n *Node) Serve(ctx context.Context) {
	logger.Info("🍦  Serving Node event channels...")
	for {
		select {
		// LobbyProtocol: ListRefresh
		case e := <-n.On(lobby.Event_LIST_REFRESH):
			event := e.Args[0].(*api.RefreshEvent)
			n.refreshEvents <- event
		// TransferProtocol: Invited
		case e := <-n.On(transfer.Event_INVITED):
			event := e.Args[0].(*api.InviteEvent)
			n.inviteEvents <- event
		// TransferProtocol: Responded
		case e := <-n.On(transfer.Event_RESPONDED):
			event := e.Args[0].(*api.DecisionEvent)
			n.decisionEvents <- event
		// TransferProtocol: Progress
		case e := <-n.On(transfer.Event_PROGRESS):
			event := e.Args[0].(*api.ProgressEvent)
			n.progressEvents <- event
		// TransferProtocol: Completed
		case e := <-n.On(transfer.Event_COMPLETED):
			event := e.Args[0].(*api.CompleteEvent)
			n.completeEvents <- event
		case <-ctx.Done():
			n.Close()
			return
		}
	}
}

// Share a peer to have a transfer
func (n *Node) NewRequest(to *common.Peer) (peer.ID, *transfer.InviteRequest, error) {
	// Fetch Element from Queue
	elem := n.queue.Front()
	if elem != nil {
		// Get Payload
		payload := n.queue.Remove(elem).(*common.Payload)

		// Create new Metadata
		meta, err := device.KeyChain.CreateMetadata(n.host.ID())
		if err != nil {
			logger.Error("Failed to create new metadata for Shared Invite", err)
			return "", nil, err
		}

		// Fetch User Peer
		from, err := n.Peer()
		if err != nil {
			logger.Error("Failed to get Node Peer Object", err)
			return "", nil, err
		}

		// Create Invite Request
		req := &transfer.InviteRequest{
			Payload:  payload,
			Metadata: api.SignedMetadataToProto(meta),
			To:       to,
			From:     from,
		}

		// Fetch Peer ID from Public Key
		toId, err := to.Libp2pID()
		if err != nil {
			logger.Error("Failed to fetch peer id from public key", err)
			return "", nil, err
		}
		return toId, req, nil
	}
	return "", nil, errors.New("No items in Transfer Queue.")
}

// Respond to an invite request
func (n *Node) NewResponse(decs bool, to *common.Peer) (peer.ID, *transfer.InviteResponse, error) {
	// Get Peer
	from, err := n.Peer()
	if err != nil {
		logger.Error("Failed to get Node Peer Object", err)
		return "", nil, err
	}

	// Create new Metadata
	meta, err := device.KeyChain.CreateMetadata(n.host.ID())
	if err != nil {
		logger.Error("Failed to create new metadata for Shared Invite", err)
		return "", nil, err
	}

	// Create Invite Response
	resp := &transfer.InviteResponse{
		Decision: decs,
		Metadata: api.SignedMetadataToProto(meta),
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Error("Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
}
