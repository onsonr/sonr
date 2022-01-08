package node

import (
	"context"
	"net"

	"git.mills.io/prologic/bitcask"
	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	common "github.com/sonr-io/core/common"

	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/sonr-io/core/wallet"
	"google.golang.org/protobuf/proto"

	ps "github.com/libp2p/go-libp2p-pubsub"
)

var (
	logger   = golog.Default.Child("core/node")
	ctx      context.Context
	instance NodeImpl
)

// NodeImpl returns the NodeImpl for the Main Node
type NodeImpl interface {
	// AuthenticateId returns the Authenticated ID
	AuthenticateId(id *wallet.UUID) (bool, error)

	// AuthenticateMessage authenticates a message
	AuthenticateMessage(msg proto.Message, metadata *common.Metadata) bool

	// Close closes the node
	Close()

	// Connect to a peer
	Connect(pi peer.AddrInfo) error

	// HasRouting returns true if the node has routing
	HasRouting() error

	// HostID returns the ID of the Host
	HostID() peer.ID

	// Join subsrcibes to a topic
	Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error)

	// NewStream opens a new stream to a peer
	NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error)

	// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
	NeedsWait()

	// Pause tells all of goroutines to pause execution
	Pause()

	// Peer returns the peer of the node
	Peer() (*common.Peer, error)

	// Profile returns the profile of the node from Local Store
	Profile() (*common.Profile, error)

	// Resume tells all of goroutines to resume execution
	Resume()

	// Role returns the role of the node
	Role() Role
	
	// Router returns the routing.Router
	Router(h host.Host) (routing.PeerRouting, error)

	// SendMessage sends a message to a peer
	SendMessage(id peer.ID, p protocol.ID, data proto.Message) error

	// SetStreamHandler sets the handler for a protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)

	// SignData signs the data with the private key
	SignData(data []byte) ([]byte, error)

	// SignMessage signs a message with the node's private key
	SignMessage(message proto.Message) ([]byte, error)

	// VerifyData verifies the data signature
	VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool
}

// node type - a p2p host implementing one or more p2p protocols
type node struct {
	// Standard Node Implementation
	host.Host
	NodeImpl
	mode Role

	// Host and context
	connection   common.Connection
	listener     net.Listener
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx   context.Context
	store *bitcask.Bitcask
	*dht.IpfsDHT
	*ps.PubSub

	// State
	flag   uint64
	Chn    chan bool
	status HostStatus
}

// NewMotor Creates a node with its implemented protocols
func NewMotor(ctx context.Context, l net.Listener, options ...Option) (NodeImpl, error) {
	// Initialize DHT
	opts := defaultOptions()
	node, err := opts.Apply(ctx, options...)
	if err != nil {
		return nil, err
	}

	// Start Host
	node.Host, err = libp2p.New(ctx,
		libp2p.Identity(node.privKey),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			opts.LowWater,    // Lowwater
			opts.HighWater,   // HighWater,
			opts.GracePeriod, // GracePeriod
		)),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(node.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		logger.Errorf("%s - NewHost: Failed to create libp2p host", err)
		return nil, err
	}
	node.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := node.Bootstrap(context.Background()); err != nil {
		logger.Errorf("%s - Failed to Bootstrap KDHT to Host", err)
		node.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range opts.BootstrapPeers {
		if err := node.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := node.createDHTDiscovery(opts); err != nil {
		logger.Fatal("Could not start DHT Discovery", err)
		node.SetStatus(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for MDNS
	node.createMdnsDiscovery(opts)
	node.SetStatus(Status_READY)
	go node.Serve()

	// Open Store with profileBuf
	return node, nil
}

// NewHighway Creates a node with its implemented protocols
func NewHighway(ctx context.Context, l net.Listener, options ...Option) (NodeImpl, error) {
	// Initialize DHT
	opts := defaultOptions()
	node, err := opts.Apply(ctx, options...)
	if err != nil {
		return nil, err
	}

	// Start Host
	node.Host, err = libp2p.New(ctx,
		libp2p.Identity(node.privKey),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			opts.LowWater,    // Lowwater
			opts.HighWater,   // HighWater,
			opts.GracePeriod, // GracePeriod
		)),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(node.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		logger.Errorf("%s - NewHost: Failed to create libp2p host", err)
		return nil, err
	}
	node.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := node.Bootstrap(context.Background()); err != nil {
		logger.Errorf("%s - Failed to Bootstrap KDHT to Host", err)
		node.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range opts.BootstrapPeers {
		if err := node.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := node.createDHTDiscovery(opts); err != nil {
		logger.Fatal("Could not start DHT Discovery", err)
		node.SetStatus(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for MDNS
	node.createMdnsDiscovery(opts)
	node.SetStatus(Status_READY)
	go node.Serve()

	return node, nil
}

// HostID returns the ID of the Host
func (n *node) HostID() peer.ID {
	return n.Host.ID()
}

// Role returns the role of the node
func (n *node) Role() Role {
	return n.mode
}
