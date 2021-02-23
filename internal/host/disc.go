package host

import (
	"context"
	"log"
	"sync"

	disco "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"
	swarm "github.com/libp2p/go-libp2p-swarm"
	disc "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/multiformats/go-multiaddr"
	"github.com/sonr-io/core/internal/lifecycle"
)

// @ discNotifee gets notified when we find a new peer via mDNS discovery ^
// ^ Connects to Rendevouz Nodes then handles discovery ^
func (hc *HostConfig) StartBootstrap(ctx context.Context, h host.Host) {
	// Begin Discovery
	hc.Routing = discovery.NewRoutingDiscovery(hc.DHT)
	discovery.Advertise(ctx, hc.Routing, hc.Point, disco.TTL(hc.Interval))

	// Connect to defined nodes
	var wg sync.WaitGroup

	for _, maddrString := range hc.Bootstrap.P2P.RDVP {
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
	go hc.handleKademliaDiscovery(ctx, h)
}

// ^ Find Peers from Routing Discovery ^ //
func (hc *HostConfig) StartMDNS(ctx context.Context, h host.Host) error {
	// setup mDNS discovery to find local peers
	var err error
	hc.MDNS, err = disc.NewMdnsService(ctx, h, hc.Interval, hc.Point)
	if err != nil {
		return err
	}

	// Create Discovery Notifier
	n := DiscNotifee{h: h, ctx: ctx}
	hc.MDNS.RegisterNotifee(&n)
	return nil
}

// ^ Handles Peers that appear on DHT ^
func (hc *HostConfig) handleKademliaDiscovery(ctx context.Context, h host.Host) {
	// Find Peers
	peerChan, err := hc.FindPeers(ctx, 15)
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
