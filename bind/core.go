package sonr

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/pkg/host"
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
	OnError(data []byte)
}

// Start begins the mobile host
func Start(data []byte, call Callback) *Node {
	// Create Context and Node - Begin Setuo
	node := new(Node)
	node.CTX = context.Background()
	node.Call = call

	// Unmarshal Connection Event
	connEvent := pb.ConnectEvent{}
	err := proto.Unmarshal(data, &connEvent)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}

	// @1. Create Host
	node.Host, err = host.NewHost(&node.CTX)
	if err != nil {
		fmt.Println("Error Creating Host: ", err)
		return nil
	}
	fmt.Println("Host Created: ", node.Host.Addrs())

	// @2. Set Stream Handlers
	node.Host.SetStreamHandler(protocol.ID("/sonr/auth"), node.HandleAuthStream)

	// @3. Set Node User Information
	err = node.setUser(&connEvent)
	if err != nil {
		fmt.Println(err)
	}

	// @4. Initialize Datastore for File Queue
	err = node.setStore()
	if err != nil {
		fmt.Println(err)
	}

	// @5. Setup Discovery
	err = node.setDiscovery()
	if err != nil {
		fmt.Println(err)
	}

	// @6. Enter Lobby
	err = node.setLobby(&connEvent)
	if err != nil {
		fmt.Println(err)
	}

	// ** Callback Node User Information ** //
	call.OnConnected(node.getUser())
	return node
}

// Exit Ends Communication
func (sn *Node) Exit() {
	sn.Lobby.End()
	sn.Host.Close()
}
