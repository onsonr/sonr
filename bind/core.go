package sonr

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	sf "github.com/sonr-io/core/internal/file"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/lobby"
	md "github.com/sonr-io/core/internal/models"
	str "github.com/sonr-io/core/internal/stream"
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
	OnCompleted(data []byte)
	OnError(data []byte)
}

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Public Properties
	HostID      string
	Peer        *md.Peer
	directories *md.Directories

	// Networking Properties
	ctx        context.Context
	host       host.Host
	authStream str.AuthStreamConn
	dataStream str.DataStreamConn

	// Data Properties
	files []*sf.SafeFile

	// References
	callbackRef Callback
	lobby       *lobby.Lobby
	peerConn    *tr.PeerConnection
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(reqBytes []byte, call Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	node := new(Node)
	node.ctx = context.Background()
	node.callbackRef, node.files = call, make([]*sf.SafeFile, maxFileBufferSize)

	// ** Unmarshal Request **
	reqMsg := md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		fmt.Println(err)
		node.error(err, "NewNode")
		return nil
	}

	// @1. Create Host and Set Stream Handlers
	node.host, node.HostID, err = sh.NewHost(node.ctx)
	if err != nil {
		node.error(err, "NewNode")
		return nil
	}
	node.setStreams()

	// @3. Set Node User Information
	if err = node.setPeer(&reqMsg); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// @4. Setup Discovery w/ Lobby
	if err = node.setDiscovery(node.ctx, &reqMsg); err != nil {
		node.error(err, "NewNode")
		return nil
	}

	// ** Callback Node User Information ** //
	return node
}
