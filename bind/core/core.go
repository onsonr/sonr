package core

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
)

// SonrCallback returns updates from p2p
type SonrCallback interface {
	OnMessage(s string)
	OnNewPeer(s string)
}

// Start begins the mobile host
func Start(olc string, call SonrCallback) *SonrNode {
	// Create Context and Node - Begin Setuo
	ctx := context.Background()
	node := new(SonrNode)
	node.OLC = olc

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
	node.Host = host
	node.PeerID = host.ID().String()

	// setup local mDNS discovery
	err = setupDiscovery(ctx, *node, call)
	if err != nil {
		panic(err)
	}

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, node.Host)
	if err != nil {
		panic(err)
	}

	// Enter location lobby
	lob, err := sonrLobby.Enter(ctx, call, ps, node.Host.ID(), olc)
	if err != nil {
		panic(err)
	}
	node.Lobby = *lob

	// Return Node
	return node
}
