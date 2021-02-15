package host

import (
	"context"
	"time"

	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	md "github.com/sonr-io/core/internal/models"
)

type State int

const (
	Stopped State = iota
	Paused
	Running
)

// ^ NewHost: Creates a host with: (MDNS, TCP, QUIC on UDP) ^
func NewHost(ctx context.Context, dir *md.Directories, olc string) (host.Host, error) {
	// @1. Established Required Data
	point := "/sonr/" + olc
	ipv4 := IPv4()

	// @2. Get Private Key
	privKey, err := getKeys(dir)
	if err != nil {
		return nil, err
	}

	// @3. Create Libp2p Host
	h, err := libp2p.New(ctx,
		// Identity
		libp2p.Identity(privKey),

		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ipv4),
			"/ip6/::/tcp/0",

			fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4),
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
			idht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}
			// We use a rendezvous point "meet me here" to announce our location.
			// This is like telling your friends to meet you at the Eiffel Tower.

			go startBootstrap(ctx, h, idht, point)
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
		libp2p.EnableNATService(),
	)
	if err != nil {
		log.Fatalln("Error starting node: ", err)
	}

	// setup local mDNS discovery
	err = startMDNS(ctx, h, point)
	return h, err
}

// ^ NewHost: Creates a host with: (MDNS Only) ^
func NewMDNSHost(ctx context.Context, dir *md.Directories, olc string) (host.Host, error) {
	// @1. Established Required Data
	ipv4 := IPv4()

	// @2. Get Private Key
	privKey, err := getKeys(dir)
	if err != nil {
		return nil, err
	}

	// @3. Create Libp2p Host
	h, err := libp2p.New(ctx,
		// Identity
		libp2p.Identity(privKey),

		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ipv4),
			"/ip6/::/tcp/0",

			fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4),
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

		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableNATService(),
	)
	if err != nil {
		log.Fatalln("Error starting node: ", err)
	}

	// setup local mDNS discovery
	err = startMDNS(ctx, h, olc)
	return h, err
}
