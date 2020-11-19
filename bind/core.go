package sonr

import (
	"context"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/host"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// Callback returns updates from p2p
type Callback interface {
	OnRefreshed(s []byte)
	OnProcessed(fileId string)
	OnInvited([]byte)
	OnResponded(decison bool)
	OnProgressed([]byte)
	OnCompleted([]byte)
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
	node.Contact = pb.Contact{
		FirstName:  connEvent.Contact.FirstName,
		LastName:   connEvent.Contact.LastName,
		ProfilePic: connEvent.Contact.ProfilePic,
	}

	// Set Profile
	node.Profile = pb.Profile{
		HostId: node.Host.ID().String(),
		Olc:    connEvent.Olc,
		Device: connEvent.Device,
	}

	// Initialize Datastore for File Queue
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		fmt.Println("Failed to create file queue")
	}
	node.FileQueue = db

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
	lob, err := lobby.Enter(ctx, call, ps, node.GetPeerInfo(), connEvent.Olc)
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
