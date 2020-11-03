package core

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Second

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "sonr-mdns"

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h    host.Host
	call SonrCallback
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
func setupDiscovery(ctx context.Context, h host.Host, call SonrCallback) error {
	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, h, DiscoveryInterval, DiscoveryServiceTag)
	if err != nil {
		return err
	}

	// Create Discovery Notifier
	n := discoveryNotifee{h: h, call: call}
	disc.RegisterNotifee(&n)
	return nil
}

// Get Slice of Peers minus User
func (n *discoveryNotifee) getPeersWithoutUser() peer.IDSlice {
	// Get Peers as Slice
	slice := n.h.Peerstore().Peers()

	// Remove User Peer
	removeFromSlice(slice, n.h.ID())

	// Return Slice
	return slice
}

// Get Slice of Peers minus User
func removeFromSlice(slice peer.IDSlice, value peer.ID) peer.IDSlice {
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
	peers := n.getPeersWithoutUser()

	// Remove Disconnected Peers
	for _, peerID := range peers {
		// Check State
		status := n.h.Network().Connectedness(peerID)

		// Remove From Store if NotConnected
		if status == network.NotConnected {
			// Remove from List
			removeFromSlice(peers, peerID)

			// Close Connection
			n.h.Network().ClosePeer(peerID)
		}
	}

	// Create JSON from the instance data.
	b, err := json.Marshal(peers)
	if err != nil {
		fmt.Printf("error formatting json")
	}

	// Callback to frontend
	n.call.OnNewPeer(string(b))
}
