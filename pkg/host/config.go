package host

import (
	"context"
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
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// Default configuration
var (
	// defaultBootstrapMultiaddrs is the default list of bootstrap nodes
	defaultBootstrapMultiaddrs = []string{
		// IPFS Bootstrapper nodes.
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",

		// IPFS Cluster Pinning nodes
		// "/ip4/138.201.67.219/tcp/4001/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		// "/ip4/138.201.67.219/udp/4001/quic/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		// "/ip4/138.201.67.220/tcp/4001/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		// "/ip4/138.201.67.220/udp/4001/quic/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		// "/ip4/138.201.68.74/tcp/4001/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		// "/ip4/138.201.68.74/udp/4001/quic/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		// "/ip4/94.130.135.167/tcp/4001/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
		// "/ip4/94.130.135.167/udp/4001/quic/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",

		// You can add more nodes here, for example, another IPFS node you might have running locally, mine was:
		// "/ip4/127.0.0.1/tcp/4010/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
		// "/ip4/127.0.0.1/udp/4010/quic/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
	}

	// defaultCallback is the default callback for the motor
	defaultCallback = common.DefaultCallback()

	// defaultRendezvousString is the default rendezvous string for the motor
	defaultRendezvousString = "sonr"
)

//
// Options for the node
//

// NodeOption is a function that configures a Node
type NodeOption func(*P2PHost) error

// WithBootstrapMultiaddrs sets the bootstrap nodes
func WithBootstrapMultiaddrs(addrs []string) NodeOption {
	return func(n *P2PHost) error {
		n.bootstrappers = addrs
		return nil
	}
}

// SetPeerIds sets the peer ids for the node
func SetPeerIds(peerIds ...peer.ID) NodeOption {
	return func(c *P2PHost) error {
		if len(peerIds) > 0 {
			c.mpcPeerIds = peerIds
		}
		return nil
	}
}

// WithNodeCallback sets the callback for the motor
func WithNodeCallback(callback common.NodeCallback) NodeOption {
	return func(c *P2PHost) error {
		c.callback = callback
		return nil
	}
}

// WithPartyId sets the party id for the node. This is to be replaced by the User defined label for the device
func WithPartyId(partyId string) NodeOption {
	return func(c *P2PHost) error {
		c.partyId = party.ID(partyId)
		return nil
	}
}

// WithWalletShare sets the wallet share for the node
func WithWalletShare(walletShare common.WalletShare) NodeOption {
	return func(c *P2PHost) error {
		c.walletShare = walletShare
		return nil
	}
}

//
// Node Setup and Initialization
//

// defaultNode creates a new node with default options
func defaultNode(ctx context.Context) *P2PHost {
	return &P2PHost{
		mdnsPeerChan:  make(chan peer.AddrInfo),
		topics:        make(map[string]*ps.Topic),
		ctx:           ctx,
		bootstrappers: defaultBootstrapMultiaddrs,
		partyId:       party.ID("current"),
	}
}

// initializeNode initializes the node
func initializeHost(hn *P2PHost) error {
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
	hn.connMgr, err = cmgr.NewConnManager(10, 40)
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
		libp2p.ConnectionManager(hn.connMgr),
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
func setupRoutingDiscovery(hn *P2PHost) error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	routingDiscovery.Advertise(hn.ctx, "sonr")

	// Create Pub Sub
	hn.PubSub, err = ps.NewGossipSub(hn.ctx, hn.host, ps.WithDiscovery(routingDiscovery))
	if err != nil {
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, "sonr")
	if err != nil {
		return err
	}
	return nil
}
