package host

import (
	"context"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

// @ discoveryInterval is how often we re-publish our mDNS records.
const discoveryInterval = time.Second

// @ discoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const discoveryServiceTag = "sonr-mdns"

// @ discoveryNotifee gets notified when we find a new peer via mDNS discovery ^
type discoveryNotifee struct {
	h host.Host
}

// ^ startMDNS creates an mDNS discovery service and attaches it to the libp2p Host. ^
func startMDNS(ctx context.Context, h host.Host) error {
	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, h, discoveryInterval, discoveryServiceTag)
	if err != nil {
		return err
	}

	// Create Discovery Notifier
	n := discoveryNotifee{h: h}
	disc.RegisterNotifee(&n)
	return nil
}

// HandlePeerFound connects to peers discovered via mDNS.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	// Connect to Peer
	err := n.h.Connect(context.Background(), pi)

	// Log Error
	if err != nil {
		log.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}
