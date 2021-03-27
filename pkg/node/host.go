package node

import (
	"fmt"
	"time"

	// Imported
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	swr "github.com/libp2p/go-libp2p-swarm"
	tls "github.com/libp2p/go-libp2p-tls"

	// Local
	net "github.com/sonr-io/core/internal/network"
	dt "github.com/sonr-io/core/pkg/data"
	tpc "github.com/sonr-io/core/pkg/topic"
)

// ^ Start Begins Assigning Host Parameters ^
func (n *Node) Start(key crypto.PrivKey) error {
	// IP Address
	ip4 := net.IPv4()
	ip6 := net.IPv6()

	// Start Host
	h, err := libp2p.New(
		n.ctx,
		libp2p.Identity(key),
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip6/%s/tcp/0", ip6)),
		// support TLS connections
		libp2p.Security(tls.ID, tls.New),
		libp2p.DefaultTransports,
		libp2p.NATPortMap(),
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

	// Create Pub Sub
	ps, err := psub.NewGossipSub(n.ctx, n.Host)
	if err != nil {
		return err
	}
	n.pubsub = ps
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
	go n.handleDHTPeers(routingDiscovery)
	return nil
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLocal() (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.pubsub, n.router.LocalTopic(), n.router, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Close Ends All Network Communication ^
func (n *Node) Close() {
	n.Host.Close()
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
		dt.GetState().NeedsWait()
		time.Sleep(time.Second * 4)
	}
}
