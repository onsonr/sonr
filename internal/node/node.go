package node

import (
	"context"
	"net"
	"strings"
	sync "sync"
	"time"

	"git.mills.io/prologic/bitcask"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/wallet"
	"github.com/sonr-io/core/pkg/common"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Standard Node Implementation
	api.NodeImpl
	clientStub  *NodeMotorStub
	highwayStub *NodeHighwayStub
	mode        StubMode

	// Host and context
	host     *host.SNRHost
	listener net.Listener

	// Properties
	ctx   context.Context
	store *bitcask.Bitcask
	state *api.State
	once  sync.Once

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
func NewNode(ctx context.Context, l net.Listener, options ...Option) (api.NodeImpl, *api.InitializeResponse, error) {
	// Set Node Options
	opts := defaultNodeOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Initialize Host
	host, err := host.NewHost(ctx, host.WithConnection(opts.connection))
	if err != nil {
		logger.Errorf("%s - Failed to initialize host", err)
		return nil, api.NewInitialzeResponse(nil, false), err
	}

	// Create Node
	node := &Node{
		ctx:            ctx,
		listener:       l,
		host:           host,
		decisionEvents: make(chan *api.DecisionEvent),
		refreshEvents:  make(chan *api.RefreshEvent),
		inviteEvents:   make(chan *api.InviteEvent),
		mailEvents:     make(chan *api.MailboxEvent),
		progressEvents: make(chan *api.ProgressEvent),
		completeEvents: make(chan *api.CompleteEvent),
	}

	// Open Store with profileBuf
	err = node.openStore(ctx, opts)
	if err != nil {
		logger.Errorf("%s - Failed to open database", err)
		return node, api.NewInitialzeResponse(nil, false), err
	}

	// Initialize Stub
	err = opts.Apply(ctx, node)
	if err != nil {
		logger.Errorf("%s - Failed to initialize stub", err)
		return nil, api.NewInitialzeResponse(nil, false), err
	}
	// Begin Background Tasks
	return node, api.NewInitialzeResponse(node.Profile, false), nil
}

// GetState returns the current state of the API
func (n *Node) GetState() *api.State {
	n.once.Do(func() {
		chn := make(chan bool)
		close(chn)
		n.state = &api.State{Chn: chn}
	})
	return n.state
}

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	// Get Profile
	profile, err := n.Profile()
	if err != nil {
		logger.Warn("Failed to get profile from Memory store, using DefaultProfile.", err)
	}

	// Get Public Key
	pubKey, err := wallet.Sonr.GetSnrPubKey(wallet.Account)
	if err != nil {
		logger.Errorf("%s - Failed to get Public Key", err)
		return nil, err
	}

	// Marshal Public Key
	pubBuf, err := pubKey.Buffer()
	if err != nil {
		logger.Errorf("%s - Failed to marshal public key", err)
		return nil, err
	}

	stat, err := common.Stat()
	if err != nil {
		logger.Errorf("%s - Failed to get device stat", err)
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
			HostName: stat["hostName"],
			Os:       stat["os"],
			Id:       stat["id"],
			Arch:     stat["arch"],
		},
	}, nil
}

// Close closes the node
func (n *Node) Close() {
	// Close Client Stub
	if n.mode.HasMotor() {
		if err := n.clientStub.Close(); err != nil {
			logger.Errorf("%s - Failed to close Client Stub, ", err)
		}
	}

	// Close Highway Stub
	if n.mode.IsHighway() {

	}

	// Close Store
	if err := n.store.Close(); err != nil {
		logger.Errorf("%s - Failed to close store, ", err)
	}

	// Close Host
	if err := n.host.Close(); err != nil {
		logger.Errorf("%s - Failed to close host, ", err)
	}
}
