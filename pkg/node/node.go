package node

import (
	"context"
	"fmt"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	discLimit "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	disc "github.com/libp2p/go-libp2p-discovery"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/internal/network"
	tr "github.com/sonr-io/core/pkg/transfer"

	//dq "github.com/sonr-io/core/pkg/user"
	sf "github.com/sonr-io/core/internal/file"
)

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx  context.Context
	opts *net.HostOptions

	// Networking Properties
	host   host.Host
	kdht   *dht.IpfsDHT
	pubsub *pubsub.PubSub
	router *net.ProtocolRouter

	call dt.NodeCallback

	// Data
	// Data Handlers
	incoming *tr.IncomingFile

	// Peers
	auth  *AuthService
	local *TopicManager
	// major    *TopicManager
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(opts *net.HostOptions, call dt.NodeCallback) *Node {
	// Create Context and Set Node Properties
	node := new(Node)
	node.ctx = context.Background()
	node.call = call
	node.opts = opts

	// Set File System
	node.router = net.NewProtocolRouter(opts.ConnRequest)

	// IP Address
	ip4 := net.IPv4()
	ip6 := net.IPv6()

	// Start Host
	h, err := libp2p.New(
		node.ctx,
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
		node.call.Connected(false)
		return nil
	}

	// Create GRPC Client/Server
	rpcServer := rpc.NewServer(h, node.router.Auth())

	// Create AuthService
	ath := AuthService{
		call:   node.call,
		respCh: make(chan *md.AuthReply, 1),
	}

	// Register Service
	err = rpcServer.Register(&ath)
	if err != nil {
		return nil
	}

	// Set RPC Services
	node.auth = &ath
	node.host = h
	return node
}

// ^ Init Begins Assigning Host Parameters ^
func (n *Node) Init(opts *net.HostOptions) bool {
	// Create Pub Sub
	ps, err := pubsub.NewGossipSub(n.ctx, n.host)
	if err != nil {
		sentry.CaptureException(err)
		n.call.Connected(false)
		return false
	}
	n.pubsub = ps

	// Create Bootstrapper Info
	bootstrappers := opts.GetBootstrapAddrInfo()
	kadDHT, err := dht.New(
		n.ctx,
		n.host,
		dht.BootstrapPeers(bootstrappers...),
	)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Error while Creating routing DHT"))
		n.call.Error(err, "Error while Creating routing DHT")
		n.call.Ready(false)
		return false
	}

	// Return Connected
	n.kdht = kadDHT
	n.call.Connected(true)
	return true
}

// ^ Bootstrap begins bootstrap with peers ^
func (n *Node) Bootstrap(opts *net.HostOptions, fs *sf.FileSystem, gp dt.ReturnPeer, gpb dt.ReturnBuf) bool {
	// Create Bootstrapper Info
	bootstrappers := opts.GetBootstrapAddrInfo()

	// Bootstrap DHT
	if err := n.kdht.Bootstrap(n.ctx); err != nil {
		sentry.CaptureException(errors.Wrap(err, "Error while Bootstrapping DHT"))
		n.call.Error(err, "Error while Bootstrapping DHT")
		n.call.Ready(false)
		return false
	}

	// Connect to bootstrap nodes, if any
	hasBootstrapped := false
	for _, pi := range bootstrappers {
		if err := n.host.Connect(n.ctx, pi); err == nil {
			hasBootstrapped = true
		}
		dt.GetState().NeedsWait()
	}

	// Check if Bootstrapping Occurred
	if !hasBootstrapped {
		sentry.CaptureException(errors.New("Failed to connect to any bootstrap peers"))
		return false
	} else {
		n.call.Ready(true)
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := disc.NewRoutingDiscovery(n.kdht)
	disc.Advertise(n.ctx, routingDiscovery, n.router.GloalPoint(), discLimit.TTL(time.Second*2))
	go n.handleDHTPeers(routingDiscovery)

	// Join Local Lobby Point
	var err error
	if n.local, err = n.JoinTopic(n.router.LocalTopic(), n.router.LocalTopicExchange(), gp, gpb); err != nil {
		sentry.CaptureException(err)
		n.call.Error(err, "Joining Lobby")
		n.call.Ready(false)
		return false
	}
	return true
}

// ^ User Node Info ^ //
// @ ID Returns Host ID
func (n *Node) ID() peer.ID {
	return n.host.ID()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Pause() {
	// Check if Response Is Invited
	dt.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Resume() {
	dt.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Close() {
	n.host.Close()
}
