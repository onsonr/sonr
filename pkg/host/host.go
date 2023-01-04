package host

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/routing"
	cmgr "github.com/libp2p/go-libp2p/p2p/net/connmgr"

	// mplex "github.com/libp2p/go-libp2p-mplex"
	ps "github.com/libp2p/go-libp2p-pubsub"
	/// direct "github.com/libp2p/go-libp2p-webrtc-direct"
	// "github.com/pion/webrtc/v3"
)

// hostImpl type - a p2p host implementing one or more p2p protocols
type hostImpl struct {
	// Standard Node Implementation
	host    host.Host
	accAddr string

	// Host and context
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx context.Context

	*dht.IpfsDHT
	*ps.PubSub
}

// New Creates a Sonr libp2p Host with the given config
func New(ctx context.Context) (*hostImpl, error) {
	var err error
	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
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
	hn.privKey, err = findPrivKey()
	if err != nil {
		return nil, err
	}

	// Create Connection Manager
	cnnmgr, err := cmgr.NewConnManager(10, 40)
	if err != nil {
		return nil, err
	}

	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.ConnectionManager(cnnmgr),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(hn.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		return nil, err
	}

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pistr := range defaultBootstrapMultiaddrs {
		pi, err := peer.AddrInfoFromString(pistr)
		if err != nil {
			continue
		}

		if err := hn.Connect(*pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(); err != nil {
		return nil, err
	}
	return hn, nil
}

// Host returns the host of the node
func (hn *hostImpl) Host() host.Host {
	return hn.host
}

// HostID returns the ID of the Host
func (n *hostImpl) HostID() peer.ID {
	return n.host.ID()
}

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *hostImpl) Connect(pi peer.AddrInfo) error {
	// Check if host is ready
	if !hn.HasRouting() {
		return fmt.Errorf("Host does not have routing")
	}

	// Call Underlying Host to Connect
	return hn.host.Connect(hn.ctx, pi)
}

// HandlePeerFound is to be called when new  peer is found
func (hn *hostImpl) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// HasRouting returns no-error if the host is ready for connect
func (h *hostImpl) HasRouting() bool {
	return h.IpfsDHT != nil && h.host != nil
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *hostImpl) Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
	// Check if PubSub is Set
	if hn.PubSub == nil {
		return nil, errors.New("Join: Pubsub has not been set on SNRHost")
	}

	// Check if topic is valid
	if topic == "" {
		return nil, errors.New("Join: Empty topic string provided to Join for host.Pubsub")
	}

	// Call Underlying Pubsub to Connect
	return hn.PubSub.Join(topic, opts...)
}

// NewStream opens a new stream to the peer with given peer id
func (n *hostImpl) NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return n.host.NewStream(ctx, pid, pids...)
}

// NewTopic creates a new topic
func (n *hostImpl) NewTopic(name string, opts ...ps.TopicOpt) (*ps.Topic, *ps.TopicEventHandler, *ps.Subscription, error) {
	// Check if PubSub is Set
	if n.PubSub == nil {
		return nil, nil, nil, errors.New("NewTopic: Pubsub has not been set on SNRHost")
	}

	// Call Underlying Pubsub to Connect
	t, err := n.Join(name, opts...)
	if err != nil {
		return nil, nil, nil, err
	}

	// Create Event Handler
	h, err := t.EventHandler()
	if err != nil {
		return nil, nil, nil, err
	}

	// Create Subscriber
	s, err := t.Subscribe()
	if err != nil {
		return nil, nil, nil, err
	}
	return t, h, s, nil
}

// Router returns the host node Peer Routing Function
func (hn *hostImpl) Router(h host.Host) (routing.PeerRouting, error) {
	// Create DHT
	kdht, err := dht.New(hn.ctx, h)
	if err != nil {
		return nil, err
	}

	// Set Properties
	hn.IpfsDHT = kdht

	// Setup Properties
	return hn.IpfsDHT, nil
}

// PubSub returns the host node PubSub Function
func (hn *hostImpl) Pubsub() *ps.PubSub {
	return hn.PubSub
}

// Routing returns the host node Peer Routing Function
func (hn *hostImpl) Routing() routing.Routing {
	return hn.IpfsDHT
}

// SetStreamHandler sets the handler for a given protocol
func (n *hostImpl) SetStreamHandler(protocol protocol.ID, handler network.StreamHandler) {
	n.host.SetStreamHandler(protocol, handler)
}
