package host

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
)

// ^ IPv4 returns the non loopback local IP of the host as IPv4 ^
func IPv4() string {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv4Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		} else {
			ipv4Ref = "0.0.0.0"
		}
	}
	// No IPv4 Found
	return ipv4Ref
}

// ^ IPv4 returns the non loopback local IP of the host as IPv4 ^
func IPv6() string {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv6Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv6 := addr.To16(); ipv6 != nil {
			ipv6Ref = ipv6.String()
		} else {
			ipv6Ref = "::"
		}
	}
	// No IPv4 Found
	return ipv6Ref
}

// ^ Creates Libp2p Host With Relay ^ //
func (sh *SonrHost) hostWithRelay() (host.Host, error) {
	// @3. Create Libp2p Host
	h, err := libp2p.New(sh.ctx,
		// Identity
		libp2p.Identity(sh.privateKey),

		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", sh.IPv4),
			"/ip6/::/tcp/0",

			fmt.Sprintf("/ip4/%s/udp/0/quic", sh.IPv4),
			"/ip6/::/udp/0/quic"),

		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),

		// support secio connections
		libp2p.Security(secio.ID, secio.New),

		// support QUIC
		libp2p.Transport(libp2pquic.NewTransport),

		// support any other default transports (TCP)
		libp2p.DefaultTransports,

		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			20,          // HighWater,
			time.Minute, // GracePeriod
		)),

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),

		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create New IDHT
			idht, err := dht.New(sh.ctx, h)
			if err != nil {
				return nil, err
			}

			go func(givenDht *dht.IpfsDHT) {
				// Connect to bootstrap nodes
				var wg sync.WaitGroup
				for _, peerAddr := range dht.DefaultBootstrapPeers {
					peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
					wg.Add(1)
					go func() {
						defer wg.Done()
						// We ignore errors as some bootstrap peers may be down
						h.Connect(sh.ctx, *peerinfo) //nolint
					}()

				}
				wg.Wait()

				// We use a rendezvous point "meet me here" to announce our location.
				// This is like telling your friends to meet you at the Eiffel Tower.
				routingDiscovery := discovery.NewRoutingDiscovery(givenDht)
				discovery.Advertise(sh.ctx, routingDiscovery, sh.Point)
				go handleKademliaDiscovery(sh.ctx, h, routingDiscovery, sh.Point)
			}(idht)
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
		libp2p.EnableNATService(),
	)
	return h, err
}

// ^ Creates Libp2p Host Without Relay ^ //
func (sh *SonrHost) hostWithoutRelay() (host.Host, error) {
	h, err := libp2p.New(sh.ctx,
		// Identity
		libp2p.Identity(sh.privateKey),

		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", sh.IPv4),
			"/ip6/::/tcp/0",

			fmt.Sprintf("/ip4/%s/udp/0/quic", sh.IPv4),
			"/ip6/::/udp/0/quic"),

		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),

		// support secio connections
		libp2p.Security(secio.ID, secio.New),

		// support QUIC
		libp2p.Transport(libp2pquic.NewTransport),

		// support any other default transports (TCP)
		libp2p.DefaultTransports,

		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			20,          // HighWater,
			time.Minute, // GracePeriod
		)),

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),

		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create New IDHT
			idht, err := dht.New(sh.ctx, h)
			if err != nil {
				return nil, err
			}

			go func(givenDht *dht.IpfsDHT) {
				// Connect to bootstrap nodes
				var wg sync.WaitGroup
				for _, peerAddr := range dht.DefaultBootstrapPeers {
					peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
					wg.Add(1)
					go func() {
						defer wg.Done()
						// We ignore errors as some bootstrap peers may be down
						h.Connect(sh.ctx, *peerinfo) //nolint
					}()

				}
				wg.Wait()

				// We use a rendezvous point "meet me here" to announce our location.
				// This is like telling your friends to meet you at the Eiffel Tower.
				routingDiscovery := discovery.NewRoutingDiscovery(givenDht)
				discovery.Advertise(sh.ctx, routingDiscovery, sh.Point)
				go handleKademliaDiscovery(sh.ctx, h, routingDiscovery, sh.Point)
			}(idht)
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableNATService(),
	)
	return h, err
}
