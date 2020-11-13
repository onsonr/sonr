package sonr

import (
	"context"
	"encoding/json"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/host"
	"github.com/sonr-io/core/pkg/lobby"
)

// Callback returns updates from p2p
type Callback interface {
	OnMessage(s string)
	OnRefresh(s string)
	OnRequested(s string)
	OnAccepted(s string)
	OnDenied(s string)
	OnProgress(s string)
	OnComplete(s string)
}

// Start begins the mobile host
func Start(data string, call Callback) *Node {
	// Create Context and Node - Begin Setuo
	ctx := context.Background()
	node := new(Node)

	// Retrieve Connection Request
	cm := new(lobby.ConnectRequest)
	err := json.Unmarshal([]byte(data), cm)
	if err != nil {
		println("Invalid Request")
		panic(err)
	}

	// Create Host
	h, err := host.NewHost(&ctx)
	// Check for Error
	if err != nil {
		panic(err)
	}
	println("Host Created")
	node.Host = h
	node.PeerID = h.ID().String()

	// Set User data to node
	err = node.SetUser(*cm)
	if err != nil {
		println("Cannot unmarshal contact")
	}

	// setup local mDNS discovery
	err = initMDNSDiscovery(ctx, *node, call)
	if err != nil {
		panic(err)
	}
	println("MDNS Started")

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, node.Host)
	if err != nil {
		panic(err)
	}
	println("GossipSub Created")

	// Enter location lobby
	lob, err := lobby.Enter(ctx, call, ps, node.Host.ID(), node.Contact.FirstName, node.Contact.LastName, node.Profile.Device, node.Contact.ProfilePic, cm.OLC)
	if err != nil {
		panic(err)
	}
	println("Lobby Joined")
	node.Lobby = *lob

	// Return Node
	return node
}
