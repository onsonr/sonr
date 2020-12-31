package host

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"
	swarm "github.com/libp2p/go-libp2p-swarm"
	disc "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/multiformats/go-multiaddr"
	"github.com/sonr-io/core/internal/lifecycle"
)

// @ discoveryInterval is how often we re-publish our mDNS records.
const discoveryInterval = time.Second

// @ discNotifee gets notified when we find a new peer via mDNS discovery ^
type discNotifee struct {
	h   host.Host
	ctx context.Context
}

// ^ Connects to Rendevouz Nodes then handles discovery ^
func connectRendevouzNodes(ctx context.Context, h host.Host, disc *discovery.RoutingDiscovery, point string) {
	// Connect to defined nodes
	var wg sync.WaitGroup

	for _, maddrString := range config.P2P.RDVP {
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
	print("Connected to bootstrap peers")
	go handleKademliaDiscovery(ctx, h, disc, point)

}

// ^ Handles Peers that appear on DHT ^
func handleKademliaDiscovery(ctx context.Context, h host.Host, disc *discovery.RoutingDiscovery, point string) {
	// Timer checks to dispose of peers
	peerChan, err := disc.FindPeers(ctx, point, discovery.Limit(15)) //nolint
	if err != nil {
		log.Println("Failed to get DHT Peer Channel: ", err)
		return
	}

	// Start Routing Discovery
	for {
		select {
		case peer := <-peerChan:
			var wg sync.WaitGroup
			if peer.ID == h.ID() {
				continue
			} else {
				wg.Add(1)
				h.Connect(ctx, peer) //nolint
			}
			wg.Wait()
		case <-ctx.Done():
			return
		}
		lifecycle.GetState().NeedsWait()
	}
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

// HandlePeerFound connects to peers discovered via mDNS.
func (n *discNotifee) HandlePeerFound(pi peer.AddrInfo) {
	// Connect to Peer
	err := n.h.Connect(n.ctx, pi)

	// Log Error for connection
	if err != nil {
		log.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
		n.h.Peerstore().ClearAddrs(pi.ID)

		if sw, ok := n.h.Network().(*swarm.Swarm); ok {
			sw.Backoff().Clear(pi.ID)
		}
	}
	lifecycle.GetState().NeedsWait()
}
