package host

import (
	"time"

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
// Bootstrap begins bootstrap with peers
func (h *hostNode) Bootstrap(deviceId string) *md.SonrError {
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
	p, d := util.DHT_OPTS()
	routingDiscovery := dsc.NewRoutingDiscovery(h.kdht)
	dsc.Advertise(h.ctxHost, routingDiscovery, p, d)
	h.disc = routingDiscovery

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctxHost, h.host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return md.NewError(err, md.ErrorEvent_HOST_PUBSUB)
	}

	// Handle DHT Peers
	h.pubsub = ps
	peersChan, err := routingDiscovery.FindPeers(h.ctxHost, p, d)
	if err != nil {
		return md.NewError(err, md.ErrorEvent_HOST_PUBSUB)
	}
	go h.handleDiscoveredPeers(peersChan)
	return nil
}

// Method Begins MDNS Discovery
func (h *hostNode) MDNS() error {
	// Logging
	md.LogActivate("MDNS")

	// Create MDNS Service
	d, p := util.MDNS_OPTS()
	ser, err := discovery.NewMdnsService(h.ctxHost, h.host, d, p)
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

// # Handle MDNS Peers: Connect to Local MDNS Peers
// Params: **Read Only** Peer AddrInfo Channel
func (h *hostNode) handleDiscoveredPeers(peerChan <-chan peer.AddrInfo) {
	for {
		select {
		case pi := <-peerChan:
			// Validate not Self
			if h.checkUnknown(pi) {
				// Connect to Peer
				if err := h.host.Connect(h.ctxHost, pi); err != nil {
					h.host.Peerstore().ClearAddrs(pi.ID)
					continue
				}
			}
		case <-h.ctxHost.Done():
			return
		}
		md.GetState()
	}
}
