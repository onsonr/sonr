package host

import (
	"crypto/rand"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/routing"
	dsc "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	cmgr "github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/sonrhq/core/pkg/node/config"
)

//
// Node Setup and Initialization
//

// defaultNode creates a new node with default options
func defaultNode(config *config.Config) *hostImpl {
	return &hostImpl{
		mdnsPeerChan: make(chan peer.AddrInfo),
		topics:       make(map[string]*ps.Topic),
		ctx:          config.Context.Ctx,
		config:       config,
		callback:     config.Callback,
	}
}

// initializeNode initializes the node
func initializeHost(hn *hostImpl) error {
	var err error
	// findPrivKey returns the private key for the host.
	findPrivKey := func() (crypto.PrivKey, error) {
		privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
		if err == nil {
			return privKey, nil
		}
		return nil, err
	}

	// Create Connection Manager
	connMgr, err := cmgr.NewConnManager(10, 40)
	if err != nil {
		return err
	}

	// Fetch the private key.
	hn.privKey, err = findPrivKey()
	if err != nil {
		return err
	}
	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.ConnectionManager(connMgr),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			hn.IpfsDHT, err = dht.New(hn.ctx, h)
			if err != nil {
				return nil, err
			}
			return hn.IpfsDHT, nil
		}),
	)
	if err != nil {
		return err
	}
	return nil
}

// setupRoutingDiscovery is a Helper Method to initialize the DHT Discovery
func setupRoutingDiscovery(hn *hostImpl) error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	routingDiscovery.Advertise(hn.ctx, hn.config.Context.Rendevouz)

	// Create Pub Sub
	hn.PubSub, err = ps.NewGossipSub(hn.ctx, hn.host, ps.WithDiscovery(routingDiscovery))
	if err != nil {
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, hn.config.Context.Rendevouz)
	if err != nil {
		return err
	}
	return nil
}
