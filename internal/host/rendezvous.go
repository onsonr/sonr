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
)

// @ defaultRendezvousTag is used in our mDNS advertisements to discover other chat peers.
const defaultRendezvousTag = "sonr-rendezvous+"

// ^ startRendezvous creates an KAD-DHT discovery service and attaches it to the libp2p Host. ^
func startRendezvous(ctx context.Context, h host.Host, kademliaDHT *dht.IpfsDHT, olc string) error {
	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	point := defaultRendezvousTag + olc
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
	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(ctx, routingDiscovery, point, discovery.TTL((time.Second * 3)))
	log.Println("Successfully announced!")
	go handleKademliaDiscovery(ctx, h, routingDiscovery, point)
	return nil
}

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
				}
				log.Println("Found peer:", peer)
			}
		case <-ctx.Done():
			return
		}
	}
}
