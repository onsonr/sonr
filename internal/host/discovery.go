package host

import (
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	peerstore "github.com/libp2p/go-libp2p-core/peerstore"
	dsc "github.com/libp2p/go-libp2p-discovery"
	psub "github.com/libp2p/go-libp2p-pubsub"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/tools/net"
	"github.com/sonr-io/core/tools/state"
)

// Bootstrap MDNS Peer Discovery Interval
const REFRESH_INTERVAL = time.Second * 4

// Libp2p Host Rendevouz Point
const HOST_RENDEVOUZ_POINT = "/sonr/rendevouz/0.9.2"

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

//interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// ** ─── HostNode Connection Methods ────────────────────────────────────────────────────────
// Bootstrap begins bootstrap with peers
func (h *SNRHost) Bootstrap() error {
	// Add Host Address to Peerstore
	h.Peerstore().AddAddrs(h.ID(), h.Addrs(), peerstore.PermanentAddrTTL)
	// Create Bootstrapper Info
	bootstrappers, err := net.BootstrapAddrInfo()
	if err != nil {
		return errors.Wrap(err, "Failed to get Bootstrapper AddrInfo")
	}

	// Bootstrap DHT
	if err := h.kdht.Bootstrap(h.ctxHost); err != nil {
		return errors.Wrap(err, "Failed to Bootstrap KDHT to Host")
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

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctxHost, h.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return errors.Wrap(err, "Failed to Create new Gossip Sub")
	}

	// Handle DHT Peers
	h.pubsub = ps
	peersChan, err := routingDiscovery.FindPeers(h.ctxHost, HOST_RENDEVOUZ_POINT, dscl.TTL(REFRESH_INTERVAL))
	if err != nil {
		return errors.Wrap(err, "Failed to create FindPeers Discovery channel")
	}
	go h.handleDiscoveredPeers(peersChan)
	return nil
}

// Method Begins MDNS Discovery
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

// Helper Method checks if Peer AddrInfo is Unknown
func (h *SNRHost) checkUnknown(pi peer.AddrInfo) bool {
	// Iterate and Check
	if len(h.Peerstore().Addrs(pi.ID)) > 0 {
		return false
	}

	// Add to PeerStore
	h.Peerstore().AddAddrs(pi.ID, pi.Addrs, time.Minute*4)
	return true
}

// Handle MDNS Peers: Connect to Local MDNS Peers
// Params: **Read Only** Peer AddrInfo Channel
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
		state.GetState()
	}
}
