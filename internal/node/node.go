package node

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	rt "github.com/libp2p/go-libp2p/core/routing"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/sonr-io/sonr/pkg/common"
)

var (
	libp2pRendevouz = "/sonr/rendevouz/0.9.2"
)

// Node returns the Node for the Main Node
type Node interface {
	// Address returns the account address of the underlying wallet for the host
	Address() string

	// Connect to a peer
	Connect(pi peer.AddrInfo) error

	// HasRouting returns true if the node has routing
	HasRouting() bool

	// Host returns the Host
	Host() host.Host

	// HostID returns the ID of the Host
	HostID() peer.ID

	// Join subsrcibes to a topic
	Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error)

	// NewStream opens a new stream to a peer
	NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error)

	// NewChannel joins a Channel interface with an underlying pubsub topic and event handler
	NewChannel(ctx context.Context, name string, opts ...ps.TopicOpt) (*Channel, error)

	// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
	NeedsWait()

	// Pubsub returns the pubsub of the node
	Pubsub() *ps.PubSub

	// Router returns the routing.Router
	Router(h host.Host) (rt.PeerRouting, error)

	// Routing returns the routing.Routing
	Routing() rt.Routing

	// Send(ctx context.Context, target string, data interface{}, protocol protocol.ID) error
	Send(id peer.ID, p protocol.ID, data []byte) error

	// SetStreamHandler sets the handler for a protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)
}

// hostImpl type - a p2p host implementing one or more p2p protocols
type hostImpl struct {
	// Standard Node Implementation
	callback common.MotorCallback
	host     host.Host

	accAddr string

	// Host and context
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx context.Context

	*dht.IpfsDHT
	*pubsub.PubSub

	// State
	fsm *SFSM
}

// NewMotor Creates a Sonr libp2p Host with the given config
func NewMotor(ctx context.Context, cb common.MotorCallback) (Node, error) {
	var err error
	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		fsm:          NewFSM(ctx),
		mdnsPeerChan: make(chan peer.AddrInfo),
	}
	// findPrivKey returns the private key for the host.
	findPrivKey := func() (crypto.PrivKey, error) {
		privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
		if err == nil {

			return privKey, nil
		}
		return nil, err
	}
	// Fetch the private key.
	privKey, err := findPrivKey()
	if err != nil {
		return nil, err
	}

	// Create Connection Manager
	connmgr, err := connmgr.NewConnManager(
		100, // Lowwater
		400, // HighWater,
		connmgr.WithGracePeriod(time.Minute),
	)

	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(privKey),
		libp2p.ConnectionManager(connmgr),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(hn.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}
	hn.fsm.SetState(Status_CONNECTING)

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {
		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(); err != nil {
		// Check if we need to close the listener
		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}

	hn.fsm.SetState(Status_READY)
	go hn.Serve()
	return hn, nil
}
