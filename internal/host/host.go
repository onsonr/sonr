package host

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	quic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	md "github.com/sonr-io/core/internal/models"
)

// ^ NewHost: Creates a host with: (MDNS, TCP, QUIC on UDP) ^
func NewHost(ctx context.Context, conReq *md.ConnectionRequest) (host.Host, string, error) {

	// @2. Get Host Requirements
	point := "sonr-kademlia+" + conReq.Olc
	ipv4 := GetIPv4()

	// @2. Create Libp2p Host
	h, err := libp2p.New(ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0/060214", ipv4),
			"/ip6/::/tcp/0/052006",

			fmt.Sprintf("/ip4/%s/udp/0/quic/021769", ipv4),
			"/ip6/::/udp/0/quic/091175"),

		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),

		// support secio connections
		libp2p.Security(secio.ID, secio.New),

		// support QUIC
		libp2p.Transport(quic.NewTransport),

		// support TOR - Mobile Networks
		//libp2p.Transport(torTransport),

		// support any other default transports (TCP)
		libp2p.DefaultTransports,

		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create New IDHT
			idht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}

			// Connect to bootstrap nodes
			var wg sync.WaitGroup
			for _, peerAddr := range dht.DefaultBootstrapPeers {
				peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
				wg.Add(1)
				go func() {
					defer wg.Done()
					if err := h.Connect(ctx, *peerinfo); err != nil {
						log.Println(err)
					} else {
						log.Println("Connection established with bootstrap node:", *peerinfo)
					}
				}()
			}
			wg.Wait()

			// Start DHT Discovery
			routingDiscovery := discovery.NewRoutingDiscovery(idht)
			discovery.Advertise(ctx, routingDiscovery, point, discovery.TTL((time.Second * 1)))

			// Handle Peers
			go handleKademliaDiscovery(ctx, h, routingDiscovery, point)
			return idht, err
		}),

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.EnableNATService(),

		// Let this host use relays and advertise itself on relays if behind NAT
		libp2p.EnableAutoRelay(),
	)

	// setup local mDNS discovery
	// err = startMDNS(ctx, h, olc)
	fmt.Println("MDNS Started")
	return h, h.ID().String(), err
}

// ^ Handles Peers that appear on DHT ^
func handleKademliaDiscovery(ctx context.Context, h host.Host, disc *discovery.RoutingDiscovery, point string) {
	// Timer checks to dispose of peers
	peerRefreshTicker := time.NewTicker(time.Second * 1)
	defer peerRefreshTicker.Stop()

	// Start Routing Discovery
	for {
		select {
		case <-peerRefreshTicker.C:
			peerChan, err := disc.FindPeers(ctx, point)
			if err != nil {
				log.Println("Failed to find peers: ", err)
				return
			}
			for peer := range peerChan {
				if peer.ID == h.ID() {
					continue
				} else {
					log.Println("Found peer:", peer)
					err := h.Connect(ctx, peer)
					if err != nil {
						log.Println("Error occurred connecting to peer: ", err)
					}
				}

			}
		case <-ctx.Done():
			return
		}
	}
}
