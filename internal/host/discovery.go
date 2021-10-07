package host

import (
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	mdns "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/pkg/errors"
)

// discoveryNotifee is a Notifee for the Discovery Service
type discoveryNotifee struct {
	mdns.Notifee
	PeerChan chan peer.AddrInfo
}

// HandlePeerFound is to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// ** ─── HostNode Connection Methods ────────────────────────────────────────────────────────
// Bootstrap connects to the IpfsDHT and Bootstraps the DHT
func (h *SNRHost) Bootstrap(dht *dht.IpfsDHT, host host.Host) error {
	if dht == nil {
		return errors.Wrap(ErrDHTNotFound, "Failed to Bootstrap")
	}
	if host == nil {
		return errors.Wrap(ErrHostNotSet, "Failed to Bootstrap")
	}

	// Set Properties
	h.IpfsDHT = dht
	h.Host = host

	// Bootstrap DHT
	if err := h.IpfsDHT.Bootstrap(h.ctx); err != nil {
		logger.Error("Failed to Bootstrap KDHT to Host", err)
		return err
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range h.opts.BootstrapPeers {
		if err := h.Connect(h.ctx, pi); err != nil {
			continue
		} else {
			break
		}
	}
	return nil
}

// Discover begins Discovery with peers for MDNS and DHT
func (h *SNRHost) Discover() error {
	// Start MDNS Discovery
	if h.opts.Connection.IsMdnsCompatible() {
		// Create MDNS Service
		ser := mdns.NewMdnsService(h.Host, h.opts.Rendezvous)

		// Register Notifier
		n := &discoveryNotifee{}
		n.PeerChan = make(chan peer.AddrInfo)

		// Handle Events
		ser.RegisterNotifee(n)
		go h.handleDiscoveredPeers(n.PeerChan)
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(h.IpfsDHT)
	dsc.Advertise(h.ctx, routingDiscovery, h.opts.Rendezvous, dscl.TTL(h.opts.TTL))

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctx, h.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		logger.Error("Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	h.PubSub = ps
	peersChan, err := routingDiscovery.FindPeers(h.ctx, h.opts.Rendezvous, dscl.TTL(h.opts.TTL))
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
	h.Peerstore().AddAddrs(pi.ID, pi.Addrs, h.opts.TTL)
	return true
}
