package host

import (
	"errors"
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	dsc "github.com/libp2p/go-libp2p-discovery"
	psub "github.com/libp2p/go-libp2p-pubsub"
)

var (
	ErrDHTNotFound = errors.New("DHT has not been set by Routing Function")
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
func (hn *SNRHost) checkDhtSet() error {
	if hn.IpfsDHT == nil {
		return ErrDHTNotFound
	}
	return nil
}

// Bootstrap begins bootstrap with peers
func (h *SNRHost) Bootstrap() error {
	// Check DHT Set
	time.Sleep(3 * time.Second)
	if err := h.checkDhtSet(); err != nil {
		logger.Error("Host DHT was never set", err)
		return err
	}

	// Bootstrap DHT
	if err := h.IpfsDHT.Bootstrap(h.ctx); err != nil {
		logger.Error("Failed to Bootstrap KDHT to Host", err)
		return err
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range h.opts.bootstrapPeers {
		if err := h.Connect(h.ctx, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(h.IpfsDHT)
	dsc.Advertise(h.ctx, routingDiscovery, h.opts.rendezvous, dscl.TTL(h.opts.ttl))
	h.disc = routingDiscovery

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctx, h.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		logger.Error("Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	h.PubSub = ps
	peersChan, err := routingDiscovery.FindPeers(h.ctx, h.opts.rendezvous, dscl.TTL(h.opts.ttl))
	if err != nil {
		logger.Error("Failed to create FindPeers Discovery channel", err)
		return err
	}
	go h.handleDiscoveredPeers(peersChan)
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
				if err := h.Connect(h.ctx, pi); err != nil {
					h.Peerstore().ClearAddrs(pi.ID)
					continue
				}
			}
		case <-h.ctx.Done():
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
	h.Peerstore().AddAddrs(pi.ID, pi.Addrs, h.opts.ttl)
	return true
}
