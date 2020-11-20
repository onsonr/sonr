package sonr

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

// discoveryInterval is how often we re-publish our mDNS records.
const discoveryInterval = time.Second

// discoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const discoveryServiceTag = "sonr-mdns"

type HostCallback interface {
	OnEvent(data []byte)
	OnError(data []byte)
}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h    host.Host
	call HostCallback
}

// initMDNSDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
func initMDNSDiscovery(ctx context.Context, h host.Host, call HostCallback) error {
	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, h, discoveryInterval, discoveryServiceTag)
	if err != nil {
		return err
	}

	// Create Discovery Notifier
	n := discoveryNotifee{h: h, call: call}
	disc.RegisterNotifee(&n)
	return nil
}

// Get Slice of Peers minus User
func (n *discoveryNotifee) GetPeersAsSlice() peer.IDSlice {
	// Get Peers as Slice
	peers := n.h.Peerstore().Peers()

	// Remove User Peer
	peers = removeIDFromSlice(peers, n.h.ID())

	// Return Slice
	return peers
}

// Get Slice of Peers minus User
func removeIDFromSlice(slice peer.IDSlice, value peer.ID) peer.IDSlice {
	// Remove User Peer
	for i, v := range slice {
		if v == value {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	// Return Slice
	return slice
}

// HandlePeerFound connects to peers discovered via mDNS.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	// Connect to Peer
	err := n.h.Connect(context.Background(), pi)

	// Log Error
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}

	// Get Peers as Slice
	peers := n.GetPeersAsSlice()

	// Remove Disconnected Peers
	for _, peerID := range peers {
		// Check State
		status := n.h.Network().Connectedness(peerID)

		// Remove From Store if NotConnected
		if status == network.NotConnected {
			// Disconnect
			n.h.Network().ClosePeer(peerID)
		}
	}
}
