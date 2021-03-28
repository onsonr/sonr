package node

import (
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	libp2p_cicuit "github.com/libp2p/go-libp2p-circuit"
	"github.com/libp2p/go-libp2p-core/crypto"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	libp2p_quic "github.com/libp2p/go-libp2p-quic-transport"
	swr "github.com/libp2p/go-libp2p-swarm"
	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/internal/network"
)

// ^ Start Begins Assigning Host Parameters ^
func (n *Node) Start(key crypto.PrivKey) error {
	// IP Address
	ip4 := net.IPv4()
	ip6 := net.IPv6()

	// Start Host
	h, err := libp2p.New(
		n.ctx,

		// default tpt + quic
		libp2p.DefaultTransports,
		libp2p.Transport(libp2p_quic.NewTransport),

		// Nat & Relay service
		libp2p.EnableNATService(),
		libp2p.DefaultStaticRelays(),
		libp2p.EnableRelay(libp2p_cicuit.OptHop),

		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip6/%s/tcp/0", ip6),
		),

		libp2p.Identity(key),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create DHT
			kdht, err := dht.New(n.ctx, h)
			if err != nil {
				return nil, err
			}

			// Set DHT
			n.kdht = kdht
			return kdht, err
		}),
	)

	// Set Host for Node
	if err != nil {
		return err
	}
	n.Host = h
	return nil
}

// ^ Bootstrap begins bootstrap with peers ^
func (n *Node) Bootstrap() error {
	// Create Bootstrapper Info
	bootstrappers, err := net.GetBootstrapAddrInfo()
	if err != nil {
		return err
	}

	// Bootstrap DHT
	if err := n.kdht.Bootstrap(n.ctx); err != nil {
		return err
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := n.Host.Connect(n.ctx, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(n.kdht)
	dsc.Advertise(n.ctx, routingDiscovery, n.router.MajorPoint(), dscl.TTL(time.Second*4))

	// Create Pub Sub
	ps, err := psub.NewGossipSub(n.ctx, n.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return err
	}
	n.pubsub = ps
	go n.handleDHTPeers(routingDiscovery)
	return nil
}

// ^ handleDHTPeers: Connects to Peers in DHT ^
func (n *Node) handleDHTPeers(routingDiscovery *dsc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			n.ctx,
			n.router.MajorPoint(),
		)
		if err != nil {
			n.call.Error(err, "Finding DHT Peers")
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != n.Host.ID() {
				// Connect to Peer
				if err := n.Host.Connect(n.ctx, pi); err != nil {
					// Remove Peer Reference
					n.Host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := n.Host.Network().(*swr.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
		}

		// Refresh table every 4 seconds
		md.GetState().NeedsWait()
		time.Sleep(time.Second * 4)
	}
}
