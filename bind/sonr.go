package sonr

import (
	"context"
	"encoding/json"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/lobby"
	sonrLobby "github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/user"
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
	host, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			100,         // Lowwater
			400,         // HighWater,
			time.Minute, // GracePeriod
		)))
	println("Host Created")

	// Check for Error
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
	println("MDNS Started")

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, node.Host)
	if err != nil {
		panic(err)
	}
	println("GossipSub Created")

	// Enter location lobby
	lob, err := sonrLobby.Enter(ctx, call, ps, node.Host.ID(), cm.OLC)
	if err != nil {
		panic(err)
	}
	println("Lobby Joined")
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
