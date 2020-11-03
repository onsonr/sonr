package core

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	sonrHost "github.com/sonr-io/p2p/pkg/host"
)

// Start begins the mobile host
func Start(olc string) *SonrNode {
	// Create Context handle events
	ctx := context.Background()

	// Create Host
	host := sonrHost.NewBasicHost(ctx)

	// setup local mDNS discovery
	err := setupMDNsDiscovery(ctx, host)
	if err != nil {
		panic(err)
	}

	return &SonrNode{
		OLC:    olc,
		PeerID: host.ID().String(),
		Host:   host,
	}
}

// Join makes Node join a lobby
func (sn *SonrNode) Join(ctx context.Context, call MessageCallback) *Lobby {
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, sn.Host)
	if err != nil {
		panic(err)
	}

	// Join Lobby Create Copy
	lob := JoinLobby(ctx, ps, sn.Host.ID(), sn.OLC)

	return &Lobby{
		ctx:      lob.ctx,
		ps:       lob.ps,
		topic:    lob.topic,
		sub:      lob.sub,
		selfID:   lob.selfID,
		OLC:      lob.OLC,
		Callback: call,
		messages: lob.messages,
	}
}

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "sonr-mdns"

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

// SetupMDNsDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupMDNsDiscovery(ctx context.Context, h host.Host) error {
	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, h, DiscoveryInterval, DiscoveryServiceTag)
	if err != nil {
		return err
	}

	n := discoveryNotifee{h: h}
	disc.RegisterNotifee(&n)
	return nil
}
