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
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Standard Node Implementation
	api.NodeImpl

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
func NewNode(ctx context.Context, options ...NodeOption) (api.NodeImpl, *api.InitializeResponse, error) {
	// Set Node Options
	opts := defaultNodeOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Create Node
	node := &Node{
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

func (n *Node) OnDecision(event *api.DecisionEvent) {
	n.decisionEvents <- event
}

func (n *Node) OnInvite(event *api.InviteEvent) {
	n.inviteEvents <- event
}

func (n *Node) OnRefresh(event *api.RefreshEvent) {
	n.refreshEvents <- event
}

func (n *Node) OnProgress(event *api.ProgressEvent) {
	n.progressEvents <- event
}

func (n *Node) OnComplete(event *api.CompleteEvent) {
	n.completeEvents <- event
}
