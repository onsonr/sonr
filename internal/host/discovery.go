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
func (hn *SNRHost) Setup() (*SNRHost, error) {
	// Check if DHT/Host are already setup after delay
	time.Sleep(200 * time.Millisecond)
	if hn.IpfsDHT == nil || hn.Host == nil {
		hn.SetStatus(Status_FAIL)
		return nil, errors.Wrap(ErrRoutingNotSet, "Failed to Bootstrap")
	}

	// Bootstrap DHT
	if err := hn.Bootstrap(hn.ctx); err != nil {
		logger.Error("Failed to Bootstrap KDHT to Host", err)
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range hn.opts.BootstrapPeers {
		if err := hn.Connect(hn.ctx, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.initDHTDiscovery(); err != nil {
		logger.Fatal("Could not start DHT Discovery", err)
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for MDNS
	hn.initMdnsDiscovery()
	hn.SetStatus(Status_READY)
	return hn, nil
}

// Discover begins Discovery with peers for MDNS and DHT
func (hn *SNRHost) initMdnsDiscovery() {
	// Verify if MDNS is Enabled
	if !hn.opts.Connection.IsMdnsCompatible() {
		logger.Error("Failed to Start MDNS Discovery ", ErrMDNSInvalidConn)
		return
	}

	// Create MDNS Service
	ser := mdns.NewMdnsService(hn.Host, hn.opts.Rendezvous)
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	// Handle Events
	ser.RegisterNotifee(n)
	go hn.handleDiscoveredPeers(n.PeerChan)
}

// initDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *SNRHost) initDHTDiscovery() error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, hn.opts.Rendezvous, dscl.TTL(hn.opts.TTL))

	// Create Pub Sub
	hn.PubSub, err = psub.NewGossipSub(hn.ctx, hn.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Error("Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	peersChan, err := routingDiscovery.FindPeers(hn.ctx, hn.opts.Rendezvous, dscl.TTL(hn.opts.TTL))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Error("Failed to create FindPeers Discovery channel", err)
		return err
	}
	go hn.handleDiscoveredPeers(peersChan)
	return nil
}

// handleDiscoveredPeers Connect to Peers that are discovered
func (hn *SNRHost) handleDiscoveredPeers(peerChan <-chan peer.AddrInfo) {
	for {
		select {
		case pi := <-peerChan:
			// Validate not Self
			if hn.checkUnknown(pi) {
				// Connect to Peer
				if err := hn.Connect(hn.ctx, pi); err != nil {
					hn.Peerstore().ClearAddrs(pi.ID)
					continue
				}
			}
		case <-hn.ctx.Done():
			return
		}
	}
}

// checkUnknown is a Helper Method checks if Peer AddrInfo is Unknown
func (hn *SNRHost) checkUnknown(pi peer.AddrInfo) bool {
	// Iterate and Check
	if len(hn.Peerstore().Addrs(pi.ID)) > 0 {
		return false
	}

	// Add to PeerStore
	hn.Peerstore().AddAddrs(pi.ID, pi.Addrs, hn.opts.TTL)
	return true
}
