package host

import (
	"context"
	"time"

	"log"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
)

// ^ NewHost: Creates a host with: (MDNS, TCP, QUIC on UDP) ^
func NewHost(ctx context.Context, config HostConfig) (host.Host, error) {
	h, err := libp2p.New(ctx,
		// Identity
		libp2p.Identity(config.PrivateKey),

		// Add listening Addresses
		libp2p.ListenAddrs(config.UDPv4, config.UDPv6, config.TCPv4, config.TCPv6),

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
			idht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}
			go config.StartBootstrap(ctx, h)
			return idht, err
		}),

		// Let this host use relays and advertise itself on relays if
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		log.Fatalln("Error starting node: ", err)
	}

	// setup local mDNS discovery
	err = config.StartMDNS(ctx, h)
	return h, err
}

// ^ NewHost: Creates a host with: (MDNS Only) ^
func NewMDNSHost(ctx context.Context, config HostConfig) (host.Host, error) {
	h, err := libp2p.New(ctx,
		// Identity
		libp2p.Identity(config.PrivateKey),

		// Add listening Addresses
		libp2p.ListenAddrs(config.UDPv4, config.UDPv6, config.TCPv4, config.TCPv6),

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

		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableNATService(),
	)
	if err != nil {
		log.Fatalln("Error starting node: ", err)
	}

	// setup local mDNS discovery
	err = config.StartMDNS(ctx, h)
	return h, err
}
