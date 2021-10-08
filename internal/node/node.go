package node

import (
	"container/list"
	"context"
	"strings"
	"time"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/internet"
	"github.com/sonr-io/core/tools/state"
	bolt "go.etcd.io/bbolt"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Emitter is the event emitter for this node
	*state.Emitter

	// Host and context
	host *host.SNRHost

	// TCP Listener for incoming connections
	listener *internet.TCPListener

	// Properties
	ctx context.Context

	// Node Options
	options nodeOptions

	// Persistent Database
	store *bolt.DB

	// Queue - the transfer queue
	queue *list.List

	// Channels
	// TransferProtocol - decisionEvents
	decisionEvents chan *common.DecisionEvent

	// LobbyProtocol - refreshEvents
	refreshEvents chan *common.RefreshEvent

	// MailboxProtocol - mailEvents
	mailEvents chan *common.MailboxEvent

	// TransferProtocol - inviteEvents
	inviteEvents chan *common.InviteEvent

	// TransferProtocol - progressEvents
	progressEvents chan *common.ProgressEvent

	// TransferProtocol - completeEvents
	completeEvents chan *common.CompleteEvent
}

// NewNode Creates a node with its implemented protocols
func NewNode(ctx context.Context, em *state.Emitter, options ...NodeOption) (*Node, *InitializeResponse, error) {
	// Set Node Options
	opts := defaultNodeOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Open TCP Port
	l, err := internet.NewTCPListener(ctx)
	if err != nil {
		logger.Error("Failed to open TCP Port", err)
		return nil, nil, err
	}

	// Initialize Host
	host, err := host.NewHost(ctx, l, em, host.WithConnection(opts.connection))
	if err != nil {
		logger.Error("Failed to initialize host", err)
		return nil, nil, err
	}

	// Wait for Host to be Ready
	time.Sleep(250 * time.Millisecond)

	// Create Node
	node := &Node{
		Emitter:        em,
		host:           host,
		ctx:            ctx,
		queue:          list.New(),
		listener:       l,
		decisionEvents: make(chan *common.DecisionEvent),
		refreshEvents:  make(chan *common.RefreshEvent),
		inviteEvents:   make(chan *common.InviteEvent),
		mailEvents:     make(chan *common.MailboxEvent),
		progressEvents: make(chan *common.ProgressEvent),
		completeEvents: make(chan *common.CompleteEvent),
	}

	// Initialize Node by Type
	err = opts.Apply(ctx, node)
	if err != nil {
		logger.Error("Failed to apply options", err)
		return nil, nil, err
	}

	// Open Database
	err = node.openStore(ctx, host, node.Emitter)
	if err != nil {
		logger.Error("Failed to open database", err)
		return nil, nil, err
	}

	// Begin Background Tasks
	go node.Serve(ctx)
	return node, node.newInitResponse(nil), nil
}

// Close closes the node
func (n *Node) Close() {
	// Close Host
	if err := n.host.Close(); err != nil {
		logger.Error("Failed to close host", err)
	}

	// Close Store
	if err := n.store.Close(); err != nil {
		logger.Error("Failed to close store", err)
	}
}

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	// Get Profile
	profile, err := n.Profile()
	if err != nil {
		return nil, err
	}

	// Get Public Key
	pubKey, err := device.KeyChain.GetSnrPubKey(keychain.Account)
	if err != nil {
		logger.Error("Failed to get Public Key", err)
		return nil, err
	}

	// Find PublicKey Buffer
	deviceStat, err := device.Stat()
	if err != nil {
		logger.Error("Failed to get device Stat", err)
		return nil, err
	}

	// Marshal Public Key
	pubBuf, err := pubKey.Buffer()
	if err != nil {
		logger.Error("Failed to marshal public key", err)
		return nil, err
	}

	// Return Peer
	return &common.Peer{
		SName:     strings.ToLower(profile.GetSName()),
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

// Stat returns the Node info as StatResponse
func (n *Node) Stat() (*StatResponse, error) {
	// Define Error StatResponse
	sendErr := func(err error) (*StatResponse, error) {
		logger.Error("Failed to get Host Stat ", err)
		return &StatResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Get Profile
	profile, err := n.Profile()
	if err != nil {
		return sendErr(err)
	}

	// Get Host Stats
	hStat, err := n.host.Stat()
	if err != nil {
		return sendErr(err)
	}

	// Get Device Stat
	dStat, err := device.Stat()
	if err != nil {
		return sendErr(err)
	}

	// Return StatResponse
	return &StatResponse{
		SName:   profile.SName,
		Profile: profile,
		Network: &StatResponse_Network{
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

// Serve handles the emitter events.
func (n *Node) Serve(ctx context.Context) {
	for {
		select {
		// LobbyProtocol: ListRefresh
		case e := <-n.On(lobby.Event_LIST_REFRESH):
			event := e.Args[0].(*common.RefreshEvent)
			n.refreshEvents <- event
		// TransferProtocol: Invited
		case e := <-n.On(transfer.Event_INVITED):
			event := e.Args[0].(*common.InviteEvent)
			n.inviteEvents <- event
		// TransferProtocol: Responded
		case e := <-n.On(transfer.Event_RESPONDED):
			event := e.Args[0].(*common.DecisionEvent)
			n.decisionEvents <- event
		// TransferProtocol: Progress
		case e := <-n.On(transfer.Event_PROGRESS):
			event := e.Args[0].(*common.ProgressEvent)
			n.progressEvents <- event
		// TransferProtocol: Completed
		case e := <-n.On(transfer.Event_COMPLETED):
			event := e.Args[0].(*common.CompleteEvent)
			n.completeEvents <- event
		case <-ctx.Done():
			n.Close()
			return
		}
	}
}
