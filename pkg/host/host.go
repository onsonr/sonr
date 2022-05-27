package host

import (
	"context"
	"crypto/rand"

	"github.com/kataras/go-events"
	"github.com/libp2p/go-libp2p"
	cmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	mplex "github.com/libp2p/go-libp2p-mplex"
	ps "github.com/libp2p/go-libp2p-pubsub"
	direct "github.com/libp2p/go-libp2p-webrtc-direct"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pion/webrtc/v3"
	"github.com/sonr-io/sonr/pkg/config"
)

// hostImpl type - a p2p host implementing one or more p2p protocols
type hostImpl struct {
	// Standard Node Implementation
	host   host.Host
	config *config.Config
	events events.EventEmmiter

	// Host and context
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx context.Context

	*dht.IpfsDHT
	*ps.PubSub

	// State
	flag   uint64
	Chn    chan bool
	status HostStatus
}

// NewDefaultHost Creates a Sonr libp2p Host with the given config
func NewDefaultHost(ctx context.Context, c *config.Config) (SonrHost, error) {
	var err error
	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		status:       Status_IDLE,
		mdnsPeerChan: make(chan peer.AddrInfo),
		config:       c,
		events:       events.New(),
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
	cnnmgr, err := cmgr.NewConnManager(c.Libp2pLowWater, c.Libp2pHighWater)
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
	hn.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range c.Libp2pBootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(c); err != nil {
		// Check if we need to close the listener
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for MDNS
	if !c.Libp2pMdnsDisabled && hn.Role() != config.Role_HIGHWAY {
		// hn.createMdnsDiscovery(config)
	}

	hn.SetStatus(Status_READY)
	go hn.Serve()
	return hn, nil
}

// NewWasmHost Creates a Sonr libp2p Host with the given config and wasm module
func NewWasmHost(ctx context.Context, c *config.Config) (SonrHost, error) {
	var err error
	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		status:       Status_IDLE,
		mdnsPeerChan: make(chan peer.AddrInfo),
		config:       c,
		events:       events.New(),
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

	maddr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/9090/http/p2p-webrtc-direct")
	if err != nil {
		return nil, err
	}

	transport := direct.NewTransport(
		webrtc.Configuration{},
		new(mplex.Transport),
	)
	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.Routing(hn.Router),
		libp2p.Transport(transport),
		libp2p.ListenAddrs(maddr),
		libp2p.DisableRelay(),
	)
	if err != nil {
		return nil, err
	}
	hn.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {

		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range c.Libp2pBootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(c); err != nil {
		// Check if we need to close the listener
		hn.SetStatus(Status_FAIL)

		return nil, err
	}

	hn.SetStatus(Status_READY)
	go hn.Serve()
	return hn, nil
}

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *hostImpl) createDHTDiscovery(c *config.Config) error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, c.Libp2pRendezvous, c.Libp2pTTL)

	// Create Pub Sub
	hn.PubSub, err = ps.NewGossipSub(hn.ctx, hn.host, ps.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, c.Libp2pRendezvous, c.Libp2pTTL)
	if err != nil {
		hn.SetStatus(Status_FAIL)
		return err
	}
	hn.SetStatus(Status_READY)
	return nil
}

// TODO Migrate MDNS Service to latesat libp2p spec
// // createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
// func (hn *hostImpl) createMdnsDiscovery(c *config.Config) {
// 	if hn.Role() == device.Role_MOTOR {
// 		// Create MDNS Service
// 		ser := mdns.NewMdnsService(hn.host, c.Libp2pRendezvous)

// 		// Handle Events
// 		ser.RegisterNotifee(hn)
// 	}
// }
