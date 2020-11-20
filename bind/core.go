package sonr

import (
	"context"

	"github.com/libp2p/go-libp2p-core/protocol"
	sh "github.com/sonr-io/core/pkg/host"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// Callback returns updates from p2p
type Callback interface {
	OnEvent(data []byte)
	OnProgress(data []byte)
	OnError(data []byte)
}

// Start begins the mobile host
func Start(data []byte, call Callback) *Node {
	// Create Context and Node - Begin Setuo
	ctx := context.Background()
	node := new(Node)
	node.CTX = ctx
	node.Call = call

	// Unmarshal Connection Event
	connEvent := pb.ConnectEvent{}
	node.BytesToProto(data, &connEvent)

	// @1. Create Host
	var err error
	node.Host, err = sh.NewHost(&node.CTX)
	if err != nil {
		node.NewError(err, 5, pb.Error_NETWORK)
		return nil
	}

	// @2. Set Stream Handlers
	node.Host.SetStreamHandler(protocol.ID("/sonr/auth"), node.HandleAuthStream)

	// @3. Set Node User Information
	node.setUser(&connEvent)

	// @4. Initialize Datastore for File Queue
	node.setStore()

	// @5. Setup Discovery
	node.setDiscovery()

	// @6. Enter Lobby
	node.setLobby(&connEvent)

	// ** Callback Node User Information ** //
	node.Callback(pb.Callback_CONNECTED, node.GetUser())
	return node
}

// ^ Sends generic protobuf with subject ^
func (sn *Node) Callback(event pb.Callback_Event, providedData []byte) {
	// Create Callback Protobuf
	callback := &pb.Callback{
		On:   event,
		Data: providedData,
	}

	// Convert to bytes
	raw, err := proto.Marshal(callback)
	if err != nil {
		sn.NewError(err, 4, pb.Error_BYTES)
	}

	// Send Generic callback
	sn.Call.OnEvent(raw)
}

// ^ Exit Ends Communication ^
func (sn *Node) Exit() {
	sn.Lobby.End()
	sn.Host.Close()
}
