package node

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	dht "github.com/libp2p/go-libp2p-kad-dht"

	ps "github.com/libp2p/go-libp2p-pubsub"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/sonr-hq/sonr/core/common"
)

var (
	libp2pRendevouz = "/sonr/rendevouz/0.9.2"
)

// Node returns the Node for the Main Node
type Node interface {
	// Connect to a peer
	Connect(pi peer.AddrInfo) error

	// ID returns the node.ID of the host
	ID() (ID, error)

	// Host returns the Host
	Host() host.Host

	// JoinTopic joins a pubsub topic
	JoinTopic(topic string, opts ...ps.TopicOpt) (*ps.Topic, error)

	// NewStream opens a new stream to a peer
	NewStream(pid peer.ID, pids ...protocol.ID) (network.Stream, error)

	// Joins a Channel interface with an underlying pubsub topic and event handler
	Join(name string, opts ...ChannelOption) (*Channel, error)

	// Pubsub returns the pubsub of the node
	Pubsub() *ps.PubSub

	// Send(target string, data interface{}, protocol protocol.ID) error
	Send(id peer.ID, p protocol.ID, data []byte) error

	// SetStreamHandler sets the handler for a protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)
}

// hostImpl type - a p2p host implementing one or more p2p protocols
type hostImpl struct {
	// Standard Node Implementation
	callback common.MotorCallback
	host     host.Host

	// Host and context
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx context.Context

	idht *dht.IpfsDHT
	ps   *pubsub.PubSub
}

// NewMotor Creates a Sonr libp2p Host with the given config
func New(ctx context.Context, options ...NodeOption) (Node, error) {
	// Default options
	var err error
	config := defaultNodeConfig()
	for _, o := range options {
		err = o(config)
		if err != nil {
			return nil, err
		}
	}

	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		mdnsPeerChan: make(chan peer.AddrInfo),
	}
	hn.host, err = libp2p.New(
		libp2p.Identity(config.GetPrivateKey()),
		libp2p.ConnectionManager(config.ConnManager),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(hn.Router),
	)
	if err != nil {
		return nil, err
	}

	if err := hn.idht.Bootstrap(context.Background()); err != nil {
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range config.BootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		}
	}

	// setup dht peer discovery
	err = hn.createDHTDiscovery()
	if err != nil {
		return nil, err
	}
	go hn.Serve()
	return hn, nil
}

// Serve handles incoming peer Addr Info
func (hn *hostImpl) Serve() {
	for {
		select {
		case mdnsPI := <-hn.mdnsPeerChan:
			if err := hn.host.Connect(context.Background(), mdnsPI); err != nil {
				continue
			}

		case dhtPI := <-hn.dhtPeerChan:
			if err := hn.host.Connect(context.Background(), dhtPI); err != nil {
				continue
			}
		case <-hn.ctx.Done():
			return
		}
	}
}
