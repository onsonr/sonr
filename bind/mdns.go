package sonr

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

// discoveryInterval is how often we re-publish our mDNS records.
const discoveryInterval = time.Second

// discoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const discoveryServiceTag = "sonr-mdns"

// initMDNSDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
func initMDNSDiscovery(ctx context.Context, sn Node, call Callback) error {
	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, sn.Host, discoveryInterval, discoveryServiceTag)
	if err != nil {
		return err
	}

	// Create Discovery Notifier
	n := discoveryNotifee{sn: sn, call: call}
	disc.RegisterNotifee(&n)
	return nil
}

// Get Slice of Peers minus User
func (n *discoveryNotifee) getPeersAsSlice() peer.IDSlice {
	// Get Peers as Slice
	peers := n.sn.Host.Peerstore().Peers()

	// Remove User Peer
	peers = removeIDFromSlice(peers, n.sn.Host.ID())

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
	err := n.sn.Host.Connect(context.Background(), pi)

	// Log Error
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}

	// Get Peers as Slice
	peers := n.getPeersAsSlice()

	// Remove Disconnected Peers
	for _, peerID := range peers {
		// Check State
		status := n.sn.Host.Network().Connectedness(peerID)

		// Remove From Store if NotConnected
		if status == network.NotConnected {
			// Disconnect
			n.sn.Host.Network().ClosePeer(peerID)
		}
	}

	// Callback to frontend
	n.sendCallback(peers)
}

// updateStore checks if store has been updated with new values
func (n *discoveryNotifee) sendCallback(peers peer.IDSlice) {
	// Remove Disconnected Peers
	for _, peerID := range peers {
		// Check State
		status := n.sn.Host.Network().Connectedness(peerID)

		// Remove From Store if NotConnected
		if status == network.NotConnected {
			// Remove from List
			peers = removeIDFromSlice(peers, peerID)
		}
	}

	// Create JSON from the instance data.
	b, err := json.Marshal(peers)
	if err != nil {
		fmt.Printf("error formatting json")
	}

	// Callback to frontend
	n.call.OnRefresh(string(b))
}
