package core

import (
	"context"
	"encoding/json"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/p2p/pkg/lobby"
	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
	"github.com/sonr-io/p2p/pkg/user"
)

// SonrCallback returns updates from p2p
type SonrCallback interface {
	OnMessage(s string)
	OnRefresh(s string)
	OnRequested(s string)
	OnAccepted(s string)
	OnDenied(s string)
	OnProgress(s string)
	OnComplete(s string)
}

// Start begins the mobile host
func Start(data string, call SonrCallback) *SonrNode {
	// Create Context and Node - Begin Setuo
	ctx := context.Background()
	node := new(SonrNode)

	// Retrieve Connection Request
	cm := new(lobby.ConnectRequest)
	err := json.Unmarshal([]byte(data), cm)
	if err != nil {
		println("Invalid Request")
		panic(err)
	}

	// Create Host
	host, err := initBasicHost(ctx)
	if err != nil {
		panic(err)
	}
	node.Host = host
	node.PeerID = host.ID().String()

	// setup local mDNS discovery
	err = initMDNSDiscovery(ctx, *node, call)
	if err != nil {
		panic(err)
	}

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, node.Host)
	if err != nil {
		panic(err)
	}

	// Enter location lobby
	lob, err := sonrLobby.Enter(ctx, call, ps, node.Host.ID(), cm.OLC)
	if err != nil {
		panic(err)
	}
	node.Lobby = *lob

	// Set Node User
	println("Go Contact Result ", cm.Contact)
	contact := user.NewContact(cm.Contact)
	profile := user.NewProfile(node.Host.ID().String(), cm.OLC, cm.Device)
	node.Contact = contact
	node.Profile = profile

	// Return Node
	return node
}
