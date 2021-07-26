package host

import (
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	peerstore "github.com/libp2p/go-libp2p-core/peerstore"
	dsc "github.com/libp2p/go-libp2p-discovery"
	psub "github.com/libp2p/go-libp2p-pubsub"
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
	// Add Host Address to Peerstore
	h.host.Peerstore().AddAddrs(h.ID(), h.host.Addrs(), peerstore.PermanentAddrTTL)
	// Create Bootstrapper Info
	bootstrappers, err := BootstrapAddrInfo()
	if err != nil {
		return md.NewError(err, md.ErrorEvent_BOOTSTRAP)
	}

	// Bootstrap DHT
	if err := h.kdht.Bootstrap(h.ctxHost); err != nil {
		return md.NewError(err, md.ErrorEvent_BOOTSTRAP)
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
	dsc.Advertise(h.ctxHost, routingDiscovery, util.HOST_RENDEVOUZ_POINT, dscl.TTL(time.Second*4))
	h.disc = routingDiscovery

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctxHost, h.host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return md.NewError(err, md.ErrorEvent_HOST_PUBSUB)
	}

	// Handle DHT
	h.pubsub = ps
	go h.handleDHTPeers(routingDiscovery)
	return nil
}

// @ Method Begins MDNS Discovery
func (h *hostNode) MDNS() error {
	// Logging
	md.LogActivate("MDNS")

	// Create MDNS Service
	ser, err := discovery.NewMdnsService(h.ctxHost, h.host, util.REFRESH_INTERVAL, util.HOST_RENDEVOUZ_POINT)
	if err != nil {
		return err
	}
	h.mdns = ser

	// Register Notifier
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	// Handle Events
	ser.RegisterNotifee(n)
	go h.handleMDNSPeers(n.PeerChan)
	return nil
}

// # Helper Method checks if Peer AddrInfo is Unknown
func (h *hostNode) checkUnknown(pi peer.AddrInfo) bool {
	// Iterate and Check
	if len(h.host.Peerstore().Addrs(pi.ID)) > 0 {
		return false
	}

	// Add to PeerStore
	h.host.Peerstore().AddAddrs(pi.ID, pi.Addrs, time.Minute*4)
	return true
}

// # handleDHTPeers: Connects to Peers in DHT
func (h *hostNode) handleDHTPeers(routingDiscovery *dsc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			h.ctxHost,
			util.HOST_RENDEVOUZ_POINT,
		)
		if err != nil {
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if h.checkUnknown(pi) {
				// Connect to Peer
				if err := h.host.Connect(h.ctxHost, pi); err != nil {
					h.host.Peerstore().ClearAddrs(pi.ID)
					continue
				}
			}
		}
	}
}

// # Handle MDNS Peers: Connect to Local MDNS Peers
func (h *hostNode) handleMDNSPeers(peerChan chan peer.AddrInfo) {
	for {
		pi := <-peerChan
		// Validate not Self
		if h.checkUnknown(pi) {
			// Connect to Peer
			if err := h.host.Connect(h.ctxHost, pi); err != nil {
				h.host.Peerstore().ClearAddrs(pi.ID)
				continue
			}
		}
	}
}
