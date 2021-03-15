package host

import (
	"context"
	"time"

	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	md "github.com/sonr-io/core/pkg/models"
)

type HostOptions struct {
	OLC          string
	Directories  *md.Directories
	Connectivity md.Connectivity
}

// ^ NewHost: Creates a host with: (MDNS, TCP, QUIC on UDP) ^
func NewHost(ctx context.Context, opts HostOptions) (host.Host, error) {
	// @1. Established Required Data
	var idht *dht.IpfsDHT
	point := "/sonr/" + opts.OLC
	ipv4 := IPv4()
	ipv6 := IPv6()

	// @2. Get Private Key
	// privKey, err := getKeys(dir)
	// if err != nil {
	// 	return nil, err
	// }

	// @3. Create Libp2p Host
	h, err := libp2p.New(ctx,
		// Identity
		// libp2p.Identity(privKey),

		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ipv4),
			fmt.Sprintf("/ip6/%s/tcp/0", ipv6),

			fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4),
			fmt.Sprintf("/ip6/%s/udp/0/quic", ipv6)),

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

		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create New IDHT
			var err error
			idht, err = dht.New(ctx, h)
			if err != nil {
				return nil, err
			}
			// We use a rendezvous point "meet me here" to announce our location.
			// This is like telling your friends to meet you at the Eiffel Tower.

			//go startBootstrap(ctx, h, idht, point)
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
		libp2p.EnableNATService(),
	)
	if err != nil {
		sentry.CaptureException(err)
	}

	// setup local mDNS discovery
	if opts.Connectivity == md.Connectivity_WiFi {
		err = startMDNS(ctx, h, point)
	} else {
		startBootstrap(ctx, h, idht, point)
	}

	return h, err
}
