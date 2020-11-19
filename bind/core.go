package sonr

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/host"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// Callback returns updates from p2p
type Callback interface {
	OnConnected(data []byte)
	OnRefreshed(data []byte)
	OnProcessed(data []byte)
	OnInvited(data []byte)
	OnResponded(data []byte)
	OnTransferring(data []byte)
	OnCompleted(data []byte)
}

// Start begins the mobile host
func Start(data []byte, call Callback) *Node {
	// Create Context and Node - Begin Setuo
	ctx := context.Background()
	node := new(Node)
	node.CTX = ctx
	node.Callback = call

	// Unmarshal Connection Event
	connEvent := pb.ConnectEvent{}
	err := proto.Unmarshal(data, &connEvent)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}

	// Create Host
	node.Host, err = host.NewHost(&ctx)
	if err != nil {
		fmt.Println("Error Creating Host: ", err)
		return nil
	}
	fmt.Println("Host Created: ", node.Host.Addrs())

	// Set Handler
	node.Host.SetStreamHandler(protocol.ID("/sonr/auth"), node.HandleAuthStream)

	// Set Contact
	err = node.setUser(&connEvent)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize Datastore for File Queue
	err = node.setStore()
	if err != nil {
		fmt.Println(err)
	}

	// setup local mDNS discovery
	err = initMDNSDiscovery(ctx, node, call)
	if err != nil {
		panic(err)
	}
	fmt.Println("MDNS Started")

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, node.Host)
	if err != nil {
		panic(err)
	}
	fmt.Println("GossipSub Created")

	// Enter location lobby
	lob, err := lobby.Enter(ctx, call, ps, node.getPeerInfo(), connEvent.Olc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Lobby Joined")
	node.Lobby = *lob

	// Return Node
	return node
}

// Exit Ends Communication
func (sn *Node) Exit() {
	sn.Lobby.End()
	sn.Host.Close()
}
