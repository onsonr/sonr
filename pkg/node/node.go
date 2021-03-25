package node

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	discLimit "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	disc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	net "github.com/sonr-io/core/internal/network"
	dt "github.com/sonr-io/core/pkg/data"
	tpc "github.com/sonr-io/core/pkg/topic"
	tr "github.com/sonr-io/core/pkg/transfer"
	sf "github.com/sonr-io/core/internal/file"
)

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx context.Context

	// Networking Properties
	host   host.Host
	kdht   *dht.IpfsDHT
	pubsub *pubsub.PubSub
	router *net.ProtocolRouter
	call   dt.NodeCallback

	// Data Handlers
	incoming *tr.IncomingFile
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(opts *net.HostOptions, call dt.NodeCallback) *Node {
	// Create Context and Set Node Properties
	node := new(Node)
	node.ctx = context.Background()
	node.call = call

	// Set Protocol Router
	node.router = net.NewProtocolRouter(opts.ConnRequest)

	// Start Host

	return node
}

// ^ Connect Begins Assigning Host Parameters ^
func (n *Node) Connect(opts *net.HostOptions) error {
	var err error

	// IP Address
	ip4 := net.IPv4()
	ip6 := net.IPv6()

	// Start Host
	h, err := libp2p.New(
		n.ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip6/%s/tcp/0", ip6)),
		libp2p.Identity(opts.PrivateKey),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			20,          // HighWater,
			time.Minute, // GracePeriod
		)),
	)
	if err != nil {
		return err
	}
	n.host = h

	// Create Pub Sub
	ps, err := pubsub.NewGossipSub(n.ctx, n.host)
	if err != nil {
		return err
	}
	n.pubsub = ps
	return nil
}

// ^ Bootstrap begins bootstrap with peers ^
func (n *Node) Bootstrap(opts *net.HostOptions, fs *sf.FileSystem) error {
	// Create Bootstrapper Info
	bootstrappers := opts.GetBootstrapAddrInfo()

	// Set DHT
	kadDHT, err := dht.New(
		n.ctx,
		n.host,
		dht.BootstrapPeers(bootstrappers...),
	)
	if err != nil {
		return err
	}

	// Return Connected
	n.kdht = kadDHT

	// Bootstrap DHT
	if err := n.kdht.Bootstrap(n.ctx); err != nil {
		return err
	}

	// Connect to bootstrap nodes, if any
	hasBootstrapped := false
	for _, pi := range bootstrappers {
		if err := n.host.Connect(n.ctx, pi); err == nil {
			hasBootstrapped = true
			break
		}
	}

	// Check if Bootstrapping Occurred
	if !hasBootstrapped {
		return err
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := disc.NewRoutingDiscovery(n.kdht)
	disc.Advertise(n.ctx, routingDiscovery, n.router.MajorPoint(), discLimit.TTL(time.Second*2))
	go n.handleDHTPeers(routingDiscovery)
	return nil
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLobby(name string) (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.host, n.pubsub, n.router.Topic(name), n.router, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLocal() (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.host, n.pubsub, n.router.LocalTopic(), n.router, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}
