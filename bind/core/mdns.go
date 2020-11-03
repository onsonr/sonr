package core

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
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
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(ctx context.Context, h host.Host, call SonrCallback) error {
	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, h, DiscoveryInterval, DiscoveryServiceTag)
	if err != nil {
		return err
	}

	n := discoveryNotifee{h: h, call: call}
	disc.RegisterNotifee(&n)
	return nil
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	// Connect to Peer
	err := n.h.Connect(context.Background(), pi)

	// Log Error
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}

	// Update Store with current peer count minus user peer
	n.updateStore(pi, len(n.h.Peerstore().Peers())-1)
}

// updateStore checks if store has been updated with new values
func (n *discoveryNotifee) updateStore(pi peer.AddrInfo, prevPeersCount int) {
	// Get Peers as Slice
	peers := n.h.Peerstore().Peers()

	// Remove User Peer
	for i, v := range peers {
		if v == n.h.ID() {
			peers = append(peers[:i], peers[i+1:]...)
			break
		}
	}

	// Create JSON from the instance data.
	b, _ := json.Marshal(peers)

	// Callback to frontend
	n.call.OnNewPeer(string(b))
}
