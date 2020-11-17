package sonr

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/host"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/user"
)

// Callback returns updates from p2p
type Callback interface {
	OnRefreshed(s string)
	OnInvited(s string)
	OnAccepted(s string)
	OnDenied(s string)
	OnProgressed(s string)
	OnCompleted(s string)
}

// Start begins the mobile host
func Start(olc string, device string, contact string, tempDir string, call Callback) *Node {
	// Create Context and Node - Begin Setuo
	ctx := context.Background()
	node := new(Node)
	node.ctx = ctx
	node.Callback = call

	// Create Host
	var err error
	node.host, err = host.NewBasicHost(&ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Host Created")

	// Set Host to Node
	node.host.SetStreamHandler(protocol.ID("/sonr/auth"), node.HandleAuthStream)
	node.peerID = node.host.ID().String()

	// Set Profile
	node.profile = user.Profile{
		ID:     node.host.ID().String(),
		OLC:    olc,
		Device: device,
	}

	// Set Contact
	node.contact = user.SetContact(contact)

	// setup local mDNS discovery
	err = initMDNSDiscovery(ctx, *node, call)
	if err != nil {
		panic(err)
	}
	fmt.Println("MDNS Started")

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, node.host)
	if err != nil {
		panic(err)
	}
	fmt.Println("GossipSub Created")

	// Enter location lobby
	lob, err := lobby.Enter(ctx, call, ps, node.GetPeer(), olc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Lobby Joined")
	node.lobby = *lob

	// Return Node
	return node
}
