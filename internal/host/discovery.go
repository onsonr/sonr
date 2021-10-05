package host

import (
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	dsc "github.com/libp2p/go-libp2p-discovery"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/net"
)

const (
	// HOST_RENDEVOUZ_POINT is the rendezvous point for the host
	HOST_RENDEVOUZ_POINT = "/sonr/rendevouz/0.9.2"

	// REFRESH_INTERVAL is the interval for refreshing the discovery
	REFRESH_INTERVAL = time.Second * 4

	// TTL_DURATION is the duration for TTL for the discovery
	TTL_DURATION = time.Minute * 2
)

// discoveryNotifee is a Notifee for the Discovery Service
type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

// HandlePeerFound is to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// ** ─── HostNode Connection Methods ────────────────────────────────────────────────────────
// Bootstrap begins bootstrap with peers
func (h *SNRHost) Bootstrap() error {
	// Create Bootstrapper Info
	bootstrappers, err := net.BootstrapAddrInfo()
	if err != nil {
		return logger.Error("Failed to get Bootstrapper AddrInfo", err)
	}

	// Bootstrap DHT
	if err := h.kdht.Bootstrap(h.ctxHost); err != nil {
		return logger.Error("Failed to Bootstrap KDHT to Host", err)
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := h.Connect(h.ctxHost, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(h.kdht)
	dsc.Advertise(h.ctxHost, routingDiscovery, HOST_RENDEVOUZ_POINT, dscl.TTL(REFRESH_INTERVAL))
	h.disc = routingDiscovery

	// Handle DHT Peers
	peersChan, err := routingDiscovery.FindPeers(h.ctxHost, HOST_RENDEVOUZ_POINT, dscl.TTL(REFRESH_INTERVAL))
	if err != nil {
		return logger.Error("Failed to create FindPeers Discovery channel", err)
	}
	go h.handleDiscoveredPeers(peersChan)
	return nil
}

// MDNS Method Begins MDNS Discovery
func (h *SNRHost) MDNS() error {
	// Create MDNS Service
	ser, err := discovery.NewMdnsService(h.ctxHost, h.Host, REFRESH_INTERVAL, HOST_RENDEVOUZ_POINT)
	if err != nil {
		return err
	}
	h.mdns = ser

	// Register Notifier
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	// Handle Events
	ser.RegisterNotifee(n)
	go h.handleDiscoveredPeers(n.PeerChan)
	return nil
}

// handleDiscoveredPeers Connect to Peers that are discovered
func (h *SNRHost) handleDiscoveredPeers(peerChan <-chan peer.AddrInfo) {
	for {
		select {
		case pi := <-peerChan:
			// Validate not Self
			if h.checkUnknown(pi) {
				// Connect to Peer
				if err := h.Connect(h.ctxHost, pi); err != nil {
					h.Peerstore().ClearAddrs(pi.ID)
					continue
				}
			}
		case <-h.ctxHost.Done():
			return
		}
	}
}

// checkUnknown is a Helper Method checks if Peer AddrInfo is Unknown
func (h *SNRHost) checkUnknown(pi peer.AddrInfo) bool {
	// Iterate and Check
	if len(h.Peerstore().Addrs(pi.ID)) > 0 {
		return false
	}

	// Add to PeerStore
	h.Peerstore().AddAddrs(pi.ID, pi.Addrs, TTL_DURATION)
	return true
}
