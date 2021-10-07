package host

import (
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	dsc "github.com/libp2p/go-libp2p-discovery"
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
// Setup Bootstraps the DHT, Sets IpfsDHT and Host, and starts discovery
func (h *SNRHost) Setup() (*SNRHost, error) {
	// Check if DHT/Host are already setup after delay
	time.Sleep(2 * time.Second)
	if h.IpfsDHT == nil || h.Host == nil {
		h.readyChan <- false
		return nil, errors.Wrap(ErrRoutingNotSet, "Failed to Bootstrap")
	}

	// Bootstrap DHT
	if err := h.Bootstrap(h.ctx); err != nil {
		logger.Error("Failed to Bootstrap KDHT to Host", err)
		h.readyChan <- false
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range h.opts.BootstrapPeers {
		if err := h.Connect(h.ctx, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := h.initDHTDiscovery(); err != nil {
		logger.Fatal("Could not start DHT Discovery", err)
		h.readyChan <- false
		return nil, err
	}

	// Initialize Discovery for MDNS
	h.initMdnsDiscovery()
	h.readyChan <- true
	return h, nil
}

// Discover begins Discovery with peers for MDNS and DHT
func (h *SNRHost) initMdnsDiscovery() {
	// Verify if MDNS is Enabled
	if !h.opts.Connection.IsMdnsCompatible() {
		logger.Error("Failed to Start MDNS Discovery ", ErrMDNSInvalidConn)
		return
	}

	// Create MDNS Service
	ser := mdns.NewMdnsService(h.Host, h.opts.Rendezvous)
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	// Handle Events
	ser.RegisterNotifee(n)
	go h.handleDiscoveredPeers(n.PeerChan)
}

// initDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (h *SNRHost) initDHTDiscovery() error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(h.IpfsDHT)
	dsc.Advertise(h.ctx, routingDiscovery, h.opts.Rendezvous, dscl.TTL(h.opts.TTL))

	// Create Pub Sub
	h.PubSub, err = psub.NewGossipSub(h.ctx, h.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		logger.Error("Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
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
