package host

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	swarm "github.com/libp2p/go-libp2p-swarm"
	disc "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/sonr-io/core/internal/lifecycle"
)

// @ discoveryInterval is how often we re-publish our mDNS records.
const discoveryInterval = time.Second * 4

// @ discNotifee gets notified when we find a new peer via mDNS discovery ^
type discNotifee struct {
	h   host.Host
	ctx context.Context
}

// ^ startMDNS creates an mDNS discovery service and attaches it to the libp2p Host. ^
func startMDNS(ctx context.Context, h host.Host, point string) error {
	// setup mDNS discovery to find local peers
	disc, err := disc.NewMdnsService(ctx, h, discoveryInterval, point)
	if err != nil {
		return err
	}

	// Create Discovery Notifier
	n := discNotifee{h: h, ctx: ctx}
	disc.RegisterNotifee(&n)
	return nil
}

// ^ Connects to Rendevouz Nodes then handles discovery ^
func startRendevouz(ctx context.Context, host host.Host, point string) error {
	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	kademliaDHT, err := dht.New(ctx, host)
	if err != nil {
		return err
	}

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	log.Println("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		log.Println(err)
	}

	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	go func() {
		var wg sync.WaitGroup
		for _, peerAddr := range dht.DefaultBootstrapPeers {
			peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
			wg.Add(1)

			defer wg.Done()
			if err := host.Connect(ctx, *peerinfo); err != nil {
				log.Println(err)
			}
			lifecycle.GetState().NeedsWait()
		}
		wg.Wait()

		// We use a rendezvous point "meet me here" to announce our location.
		// This is like telling your friends to meet you at the Eiffel Tower.
		log.Println("Announcing ourselves...")
		routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
		discovery.Advertise(ctx, routingDiscovery, point)
		log.Println("Successfully announced!")
		go handleKademliaDiscovery(ctx, host, routingDiscovery, point)
	}()
	return nil
}

// ^ Handles Peers that appear on DHT ^
func handleKademliaDiscovery(ctx context.Context, h host.Host, disc *discovery.RoutingDiscovery, point string) {
	// Find Peers
	peerChan, err := disc.FindPeers(ctx, point, discovery.Limit(15)) //nolint
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
func (n *discNotifee) HandlePeerFound(pi peer.AddrInfo) {
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
