package host

import (
	"context"
	"log"
	"sync"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

// @ discoveryDHTTag is used in our mDNS advertisements to discover other chat peers.
const defaultDHTTag = "sonr-dht+"

// ^ startDHT creates an mDNS discovery service and attaches it to the libp2p Host. ^
func startDHT(ctx context.Context, h host.Host, olc string) error {
	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	discTag := defaultDHTTag + olc
	kademliaDHT, err := dht.New(ctx, h)
	if err != nil {
		return err
	}

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	//logger.Debug("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		return err
	}

	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	var wg sync.WaitGroup
	for _, peerAddr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(ctx, *peerinfo); err != nil {
				log.Fatalln(err)
			}
		}()
	}
	wg.Wait()

	// We use a rendezvous point "meet me here" to announce our location.
	// This is like telling your friends to meet you at the Eiffel Tower.
	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(ctx, routingDiscovery, discTag)
	go handleDHTPeers(ctx, h, discTag, routingDiscovery)
	return nil
}

func handleDHTPeers(ctx context.Context, h host.Host, discTag string, disc *discovery.RoutingDiscovery) {
	// Find Peers with DHT Tag
	peerChan, err := disc.FindPeers(context.Background(), discTag)
	if err != nil {
		panic(err)
	}

	// Iterate Through List
	for peer := range peerChan {
		if peer.ID == h.ID() {
			continue
		}

		// Connect to Peer
		err := h.Connect(context.Background(), peer)

		// Log Error
		if err != nil {
			log.Printf("error connecting to peer %s: %s\n", peer.ID.Pretty(), err)
		}

	}
	select {}
}
