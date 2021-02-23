package host

import (
	"context"
	"sync"
	"time"

	"log"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	disco "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	swarm "github.com/libp2p/go-libp2p-swarm"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	disc "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/multiformats/go-multiaddr"
	"github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
)

const Interval = time.Second * 4

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
			go StartBootstrap(ctx, h, idht, &config)
			return idht, err
		}),

		// Let this host use relays and advertise itself on relays if
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		log.Fatalln("Error starting node: ", err)
	}

	// Check for Wifi
	if config.Connectivity == md.ConnectionRequest_Wifi {
		// setup local mDNS discovery
		err = StartMDNS(ctx, h, &config)
	}
	return h, err
}

// @ discNotifee gets notified when we find a new peer via mDNS discovery ^
// ^ Connects to Rendevouz Nodes then handles discovery ^
func StartBootstrap(ctx context.Context, h host.Host, idht *dht.IpfsDHT, hc *HostConfig) {
	// Begin Discovery
	routing := discovery.NewRoutingDiscovery(idht)
	discovery.Advertise(ctx, routing, hc.Point, disco.TTL(Interval))

	// Connect to defined nodes
	var wg sync.WaitGroup
	bootstrap, err := getBootstrap()
	if err != nil {
		print(err)
		return
	}
	for _, maddrString := range bootstrap.P2P.RDVP {
		maddr, err := multiaddr.NewMultiaddr(maddrString.Maddr)
		if err != nil {
			log.Println(err)
		}
		wg.Add(1)
		peerinfo, _ := peer.AddrInfoFromP2pAddr(maddr)
		h.Connect(ctx, *peerinfo) //nolint
		wg.Done()
		lifecycle.GetState().NeedsWait()
	}
	wg.Wait()
	go handleKademliaDiscovery(ctx, h, routing, hc)
}

// ^ Find Peers from Routing Discovery ^ //
func StartMDNS(ctx context.Context, h host.Host, hc *HostConfig) error {
	// setup mDNS discovery to find local peers
	var err error
	mdns, err := disc.NewMdnsService(ctx, h, Interval, hc.Point)
	if err != nil {
		return err
	}

	// Create Discovery Notifier
	n := DiscNotifee{h: h, ctx: ctx}
	mdns.RegisterNotifee(&n)
	return nil
}

// ^ Handles Peers that appear on DHT ^
func handleKademliaDiscovery(ctx context.Context, h host.Host, routing *discovery.RoutingDiscovery, hc *HostConfig) {
	// Find Peers
	peerChan, err := routing.FindPeers(ctx, hc.Point, disco.Limit(15))
	if err != nil {
		log.Println("Failed to get DHT Peer Channel: ", err)
		return
	}

	// Start Routing Discovery
	for {
		select {
		case pi := <-peerChan:
			var wg sync.WaitGroup
			if pi.ID == h.ID() {
				continue
			} else {
				wg.Add(1)
				err := h.Connect(ctx, pi)
				checkConnErr(err, pi.ID, h)
			}
			wg.Wait()
		case <-ctx.Done():
			return
		}
		lifecycle.GetState().NeedsWait()
	}
}

// ^ HandlePeerFound connects to peers discovered via mDNS. ^
func (n *DiscNotifee) HandlePeerFound(pi peer.AddrInfo) {
	// Connect to Peer
	err := n.h.Connect(n.ctx, pi)
	checkConnErr(err, pi.ID, n.h)
	lifecycle.GetState().NeedsWait()
}

// ^ Helper: Checks for Connect Error ^
func checkConnErr(err error, id peer.ID, h host.Host) {
	if err != nil {
		log.Printf("error connecting to peer %s: %s\n", id.Pretty(), err)
		h.Peerstore().ClearAddrs(id)

		if sw, ok := h.Network().(*swarm.Swarm); ok {
			sw.Backoff().Clear(id)
		}
	}
}
