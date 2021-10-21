package node

import (
	"context"
	"errors"
	"strings"
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
	clientStub  *ClientNodeStub
	highwayStub *HighwayNodeStub
	mode        NodeStubMode

	// Host and context
	host *host.SNRHost

	// Properties
	ctx        context.Context
	isTerminal bool
	store      *bitcask.Bitcask

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
func NewNode(ctx context.Context, options ...Option) (api.NodeImpl, *api.InitializeResponse, error) {
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
	host, err := host.NewHost(ctx, host.WithConnection(opts.connection), host.WithTerminal(opts.isTerminal))
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
	return node, api.NewInitialzeResponse(node.Profile, false), nil
}

// Close closes the node
func (n *Node) Close() {
	// Close Client Stub
	if n.mode.IsClient() {
		if err := n.clientStub.Close(); err != nil {
			logger.Error("Failed to close Client Stub, ", err)
		}
	}

	// Close Highway Stub
	if n.mode.IsHighway() {
		if err := n.highwayStub.Close(); err != nil {
			logger.Error("Failed to close Highway Stub, ", err)
		}
	}

	// Close Store
	if err := n.store.Close(); err != nil {
		logger.Error("Failed to close store, ", err)
	}

	// Close Host
	if err := n.host.Close(); err != nil {
		logger.Error("Failed to close host, ", err)
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
	pubKey, err := wallet.Primary.GetSnrPubKey(wallet.Account)
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

	stat, err := common.Stat()
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
			HostName: stat["hostName"],
			Os:       stat["os"],
			Id:       stat["id"],
			Arch:     stat["arch"],
		},
	}, nil
}

// OnDecision is callback for NodeImpl for decisionEvents
func (n *Node) OnDecision(event *api.DecisionEvent) {
	n.printTerminal(event.Title(), event.Message())
	n.decisionEvents <- event
}

// OnInvite is callback for NodeImpl for inviteEvents
func (n *Node) OnInvite(event *api.InviteEvent) {
	n.printTerminal(event.Title(), event.Message())
	n.promptTerminal("Accept Invite", func(result bool) {
		if n.mode.IsClient() {
			n.clientStub.Respond(n.ctx, &api.RespondRequest{
				Decision: result,
				Peer:     event.GetFrom(),
			})
		}
	})
	n.inviteEvents <- event
}

// OnRefresh is callback for NodeImpl for refreshEvents
func (n *Node) OnRefresh(event *api.RefreshEvent) {
	n.refreshEvents <- event
}

// OnProgress is callback for NodeImpl for progressEvents
func (n *Node) OnProgress(event *api.ProgressEvent) {
	n.progressTerminal(event)
	n.progressEvents <- event
}

// OnComplete is callback for NodeImpl for completeEvents
func (n *Node) OnComplete(event *api.CompleteEvent) {
	n.printTerminal(event.Title(), event.Message())
	n.completeEvents <- event
}

// NewSNID returns a new SNID
func (n *Node) NewSNID(sname string) (*wallet.SNID, error) {
	// Check if SNID is empty
	if len(sname) == 0 {
		return nil, errors.New("SName not provided.")
	}

	// Find Records
	recs, err := n.host.LookupTXT(n.ctx, sname)
	if err != nil {
		return nil, err
	}

	// Get Name from Records
	rec, err := recs.GetNameRecord()
	if err != nil {
		return nil, err
	}

	// Get Pub Key
	pubKey, err := rec.PubKeyBuffer()
	if err != nil {
		return nil, err
	}

	// Get Peer ID
	id, err := rec.PeerID()
	if err != nil {
		return nil, err
	}

	// Return SNID
	return &wallet.SNID{
		Domain: sname,
		PeerID: id.String(),
		PubKey: pubKey,
	}, nil
}
