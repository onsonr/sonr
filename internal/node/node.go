package node

import (
	"context"
	"strings"
	"time"

	"git.mills.io/prologic/bitcask"
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
	// Standard Node Implementation
	common.NodeImpl

	// Emitter is the event emitter for this node
	*state.Emitter

	// Host and context
	host *host.SNRHost

	// Properties
	ctx context.Context

	store *bitcask.Bitcask

	// Node Stub Interface
	stub NodeStub

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
		decisionEvents: make(chan *api.DecisionEvent),
		refreshEvents:  make(chan *api.RefreshEvent),
		inviteEvents:   make(chan *api.InviteEvent),
		mailEvents:     make(chan *api.MailboxEvent),
		progressEvents: make(chan *api.ProgressEvent),
		completeEvents: make(chan *api.CompleteEvent),
	}

	// Initialize Host
	host, err := host.NewHost(ctx, host.WithConnection(opts.connection))
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
	err = opts.Apply(ctx, node)
	if err != nil {
		logger.Error("Failed to initialize stub", err)
		return nil, api.NewInitialzeResponse(nil, false), err
	}

	// Begin Background Tasks
	go node.Serve(ctx)
	return node, api.NewInitialzeResponse(node.GetProfile, false), nil
}

// Close closes the node
func (n *Node) Close() {
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
	profile, err := n.GetProfile()
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
			HostName: stat[device.StatKey_HostName],
			Os:       stat[device.StatKey_Os],
			Id:       stat[device.StatKey_Id],
			Arch:     stat[device.StatKey_Arch],
		},
	}, nil
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
		}
	}
}
