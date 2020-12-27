package sonr

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sf "github.com/sonr-io/core/internal/file"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/lifecycle"
	"github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	tr "github.com/sonr-io/core/internal/transfer"
	"google.golang.org/protobuf/proto"
)

// @ Maximum Files in Node Cache
const maxFileBufferSize = 5

// ^ Interface: Callback is implemented from Plugin to receive updates ^
type Callback interface {
	OnRefreshed(data []byte)
	OnInvited(data []byte)
	OnResponded(data []byte)
	OnQueued(data []byte)
	OnProgress(data float32)
	OnReceived(data []byte)
	OnTransmitted(data []byte)
	OnError(data []byte)
}

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	olc     string
	peer    *md.Peer
	contact *md.Contact

	// Networking Properties
	ctx    context.Context
	host   host.Host
	pubSub *pubsub.PubSub

	// Data Properties
	files       []*sf.SafeMetadata
	directories *md.Directories

	// References
	call     Callback
	lobby    *lobby.Lobby
	peerConn *tr.PeerConnection
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(reqBytes []byte, call Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	node := new(Node)
	node.ctx = context.Background()
	node.call, node.files = call, make([]*sf.SafeMetadata, maxFileBufferSize)

	// ** Unmarshal Request **
	reqMsg := md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		log.Println(err)
		node.error(err, "NewNode")
		return nil
	}

	// @1. Create Host and Start Discovery
	node.host, err = sh.NewHost(node.ctx, reqMsg.Olc)
	if err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @3. Set Node User Information
	if err = node.setInfo(&reqMsg); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @4. Setup Connection w/ Lobby and Set Stream Handlers
	if err = node.setConnection(node.ctx); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// ** Callback Node User Information ** //
	return node
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Pause() {
	log.Println("Sonr Paused.")
	lifecycle.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Resume() {
	log.Println("Sonr Resumed.")
	lifecycle.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Stop() {
	log.Println("Sonr Stopped.")
	sn.host.Close()
}

// ^ error Callback with error instance, and method ^
func (sn *Node) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	// Send Callback
	sn.call.OnError(bytes)

	// Log In Core
	log.Fatalf("[Error] At Method %s : %s", err.Error(), method)
}
