package host

import (
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	dsc "github.com/libp2p/go-libp2p-discovery"

	swr "github.com/libp2p/go-libp2p-swarm"
	md "github.com/sonr-io/core/pkg/models"
)

// ^ Set Stream Handler for Host ^
func (h *HostNode) HandleStream(pid protocol.ID, handler network.StreamHandler) {
	h.Host.SetStreamHandler(pid, handler)
}


// ^ Start Stream for Host ^
func (h *HostNode) StartStream(p peer.ID, pid protocol.ID) (network.Stream, error) {
	return h.Host.NewStream(h.ctx, p, pid)
}

// @ handleDHTPeers: Connects to Peers in DHT
func (h *HostNode) handleDHTPeers(routingDiscovery *dsc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			h.ctx,
			h.Point,
		)
		if err != nil {
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != h.Host.ID() {
				// Connect to Peer
				if err := h.Host.Connect(h.ctx, pi); err != nil {
					// Remove Peer Reference
					h.Host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := h.Host.Network().(*swr.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
		}

		// Refresh table every 5 seconds
		md.GetState().NeedsWait()
		time.Sleep(REFRESH_DURATION)
	}
}
