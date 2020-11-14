package host

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	peer "github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	routing "github.com/libp2p/go-libp2p-routing"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	"github.com/multiformats/go-multiaddr"
)

// SonrHost packages peer channel and host ref
type SonrHost struct {
	Host    host.Host
	Channel <-chan peer.AddrInfo
}

// NewHost creates new host, sets it up, then returns it
func NewHost(ctx *context.Context) (host.Host, error) {
	// Create Host
	h, err := libp2p.New(*ctx,
		// Multiple listen addresses
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/0",
			"/ip6/::/tcp/0",

			"/ip4/0.0.0.0/udp/0/quic",
			"/ip6/::/udp/0/quic", // a UDP endpoint for the QUIC transport
		),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support secio connections
		libp2p.Security(secio.ID, secio.New),
		// support QUIC - experimental
		libp2p.Transport(libp2pquic.NewTransport),
		// support any other default transports (TCP)
		libp2p.DefaultTransports,
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(connmgr.NewConnManager(
			100,         // Lowwater
			400,         // HighWater,
			time.Minute, // GracePeriod
		)),
		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Reference IPFS DHT
			var idht *dht.IpfsDHT
			var err error

			// Create DHT
			idht, err = dht.New(*ctx, h)
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
	)

	// Check for error
	if err != nil {
		return nil, err
	}

	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	kademliaDHT, err := dht.New(*ctx, h)
	if err != nil {
		panic(err)
	}

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	println("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(*ctx); err != nil {
		panic(err)
	}

	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	var wg sync.WaitGroup
	for _, peerAddr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(*ctx, *peerinfo); err != nil {
				println(err)
			} else {
				println("Connection established with bootstrap node")
			}
		}()
	}
	wg.Wait()

	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(*ctx, routingDiscovery, "sonr-dht")

	peerChan, err := routingDiscovery.FindPeers(*ctx, "sonr-dht")
	if err != nil {
		panic(err)
	}

	sh := &SonrHost{
		Host:    h,
		Channel: peerChan,
	}

	go sh.handlePeers()

	return h, nil
}

// NewBasicHost creates a host without any options
func NewBasicHost(ctx *context.Context) (host.Host, error) {
	// Find IPv4 Address
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	var ipv4Ref string
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
			ipv4Ref = ipv4.String()
		}
	}

	// Create Multi Address
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", ipv4Ref, 0))

	// Create Libp2p Host
	h, err := libp2p.New(*ctx, libp2p.ListenAddrs(
		sourceMultiAddr),
	)
	return h, err
}

func (sh *SonrHost) handlePeers() {
	for peer := range sh.Channel {
		if peer.ID == sh.Host.ID() {
			continue
		}
		println("Found peer:", peer.String())

		err := sh.Host.Connect(context.Background(), peer)

		if err != nil {
			println("Error connecting to peer")
			continue
		} else {
			println("Connected to:", peer.String())
		}
	}
}
