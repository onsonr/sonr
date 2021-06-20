package host

import (
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	dsc "github.com/libp2p/go-libp2p-discovery"
	psub "github.com/libp2p/go-libp2p-pubsub"
	swr "github.com/libp2p/go-libp2p-swarm"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
)

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

//interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// ** ─── HostNode Connection Methods ────────────────────────────────────────────────────────
// @ Bootstrap begins bootstrap with peers
func (h *hostNode) Bootstrap() *md.SonrError {
	// Create Bootstrapper Info
	bootstrappers, err := getBootstrapAddrInfo()
	if err != nil {
		return md.NewError(err, md.ErrorMessage_BOOTSTRAP)
	}

	// Bootstrap DHT
	if err := h.kdht.Bootstrap(h.ctxHost); err != nil {
		return md.NewError(err, md.ErrorMessage_BOOTSTRAP)
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := h.host.Connect(h.ctxHost, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(h.kdht)
	dsc.Advertise(h.ctxHost, routingDiscovery, h.point, dscl.TTL(time.Second*4))
	h.disc = routingDiscovery

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctxHost, h.host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_PUBSUB)
	}
	h.pubsub = ps
	go h.handleDHTPeers(routingDiscovery)
	return nil
}

func (h *hostNode) MDNS() error {
	ser, err := discovery.NewMdnsService(h.ctxHost, h.host, util.REFRESH_INTERVAL, h.point)
	if err != nil {
		return err
	}
	h.mdns = ser
	//register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	ser.RegisterNotifee(n)
	go h.handleMDNSPeers(n.PeerChan)
	return nil
}

// # handleDHTPeers: Connects to Peers in DHT
func (h *hostNode) handleDHTPeers(routingDiscovery *dsc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			h.ctxHost,
			h.point,
		)
		if err != nil {
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != h.host.ID() {
				// Connect to Peer
				if err := h.host.Connect(h.ctxHost, pi); err != nil {
					// Remove Peer Reference
					h.host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := h.host.Network().(*swr.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
		}

		// Refresh table every 5 seconds
		md.GetState().NeedsWait()
	}
}

func (h *hostNode) handleMDNSPeers(peerChan chan peer.AddrInfo) {
	for {
		pi := <-peerChan
		if err := h.host.Connect(h.ctxHost, pi); err != nil {
			// Remove Peer Reference
			h.host.Peerstore().ClearAddrs(pi.ID)
			if sw, ok := h.host.Network().(*swr.Swarm); ok {
				sw.Backoff().Clear(pi.ID)
			}
		}
		md.GetState().NeedsWait()
	}
}
