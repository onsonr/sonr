package sonr

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/host"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
)

// Callback returns updates from p2p
type Callback interface {
	OnRefreshed(s []byte)
	OnInvited([]byte) //TODO add thumbnail
	OnResponded(decison bool)
	OnProgressed([]byte)
	OnCompleted([]byte)
}

// Start begins the mobile host
func Start(olc string, device string, contact string, call Callback) *Node {
	// Create Context and Node - Begin Setuo
	ctx := context.Background()
	node := new(Node)
	node.CTX = ctx
	node.Callback = call

	// Create Host
	var err error
	node.Host, err = host.NewBasicHost(&ctx)
	if err != nil {
		fmt.Println("Error Creating Host: ", err)
		return nil
	}
	fmt.Println("Host Created")

	// Set Host to Node
	node.Host.SetStreamHandler(protocol.ID("/sonr/auth"), node.HandleAuthStream)

	// Set Profile
	node.Profile = pb.Profile{
		HostId: node.Host.ID().String(),
		Olc:    olc,
		Device: device,
	}

	// Set Contact
	err = jsonpb.UnmarshalString(contact, &node.Contact)
	if err != nil {
		fmt.Println("Error Unmarshalling Contact Data into Buffer: ", err)
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
	lob, err := lobby.Enter(ctx, call, ps, node.GetPeerInfo(), olc)
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
	sn.AuthStream.stream.Close()
	sn.Lobby.End()
	sn.Host.Close()
}
