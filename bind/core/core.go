package core

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
)

// DiscoveryInterval is how often we re-publish our mDNS records.
const discoveryInterval = time.Duration(10) * time.Second

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const discoveryServiceTag = "sonr-mdns"

// MessageCallback returns message from lobby
type MessageCallback interface {
	OnMessage(s string)
}

// Start begins the mobile host
func Start(olc string, call MessageCallback) *SonrNode {
	// Create Context handle events
	ctx := context.Background()

	// Create Host
	host, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			100,         // Lowwater
			400,         // HighWater,
			time.Minute, // GracePeriod
		)))
	if err != nil {
		panic(err)
	}

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		panic(err)
	}

	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, host, discoveryInterval, discoveryServiceTag)
	if err != nil {
		panic(err)
	}

	n := discoveryNotifee{h: host, ctx: ctx}
	disc.RegisterNotifee(&n)

	lob, err := sonrLobby.Enter(ctx, call, ps, host.ID(), olc)

	return &SonrNode{
		OLC:    olc,
		PeerID: host.ID().String(),
		Host:   host,
		Lobby:  lob,
	}
}

// RefreshPeers returns peers as string
func (sn *SonrNode) RefreshPeers() string {
	peers := sn.Host.Peerstore().Peers()
	// Create JSON from the instance data.
	// ... Ignore errors.
	b, _ := json.Marshal(peers)
	// Convert bytes to string.
	s := string(b)
	return s
}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h   host.Host
	ctx context.Context
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	err := n.h.Connect(n.ctx, pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}
