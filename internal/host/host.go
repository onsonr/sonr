package host

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	tor "berty.tech/go-libp2p-tor-transport"
	config "berty.tech/go-libp2p-tor-transport/config"
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
func NewHost(ctx context.Context, dirs *md.Directories, olc string) (host.Host, string, error) {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv4Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		}
	}

	// @2. Get Host Requirements
	point := "sonr-kademlia+" + olc
	torTransport, err := tor.NewBuilder( // Create a builder
		config.EnableEmbeded,
		config.DoSlowStart,
		config.AllowTcpDial, // Some Configurator are already ready to use.
		config.SetSetupTimeout(time.Minute),
		config.SetTemporaryDirectory(dirs.Temporary),
	)
	if err != nil {
		return nil, "", err
	}

	// @2. Create Libp2p Host
	h, err := libp2p.New(ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ipv4Ref),
			"/ip6/::/tcp/0",

			fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4Ref),
			"/ip6/::/udp/0/quic"),

		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),

		// support secio connections
		libp2p.Security(secio.ID, secio.New),

		// support QUIC
		libp2p.Transport(quic.NewTransport),

		// support TOR - Mobile Networks
		libp2p.Transport(torTransport),

		// support any other default transports (TCP)
		libp2p.DefaultTransports,

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),

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

			// We use a rendezvous point "meet me here" to announce our location.
			// This is like telling your friends to meet you at the Eiffel Tower.
			log.Println("Announcing ourselves...")
			routingDiscovery := discovery.NewRoutingDiscovery(idht)
			discovery.Advertise(ctx, routingDiscovery, point, discovery.TTL((time.Second * 3)))
			log.Println("Successfully announced!")
			go handleKademliaDiscovery(ctx, h, routingDiscovery, point)
			log.Println("Waiting for Peers...")
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
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
	peerRefreshTicker := time.NewTicker(time.Second * 3)
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
