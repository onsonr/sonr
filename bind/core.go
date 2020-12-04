package sonr

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	sonrFile "github.com/sonr-io/core/internal/file"
	sonrHost "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/lobby"
	sonrModel "github.com/sonr-io/core/internal/models"
	sonrStream "github.com/sonr-io/core/internal/stream"
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
	Peer        *sonrModel.Peer
	directories *sonrModel.Directories

	// Networking Properties
	ctx        context.Context
	host       host.Host
	authStream sonrStream.AuthStreamConn
	dataStream sonrStream.DataStreamConn
	peerConn   *tr.PeerConnection

	// Data Properties
	files []*sonrFile.SafeFile

	// References
	callbackRef Callback
	lobby       *lobby.Lobby
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(reqBytes []byte, call Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	node := new(Node)
	node.ctx = context.Background()
	node.callbackRef, node.files = call, make([]*sonrFile.SafeFile, maxFileBufferSize)

	// ** Unmarshal Request **
	reqMsg := sonrModel.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		fmt.Println(err)
		node.error(err, "NewNode")
		return nil
	}

	// @1. Create Host and Set Stream Handlers
	node.host, node.HostID, err = sonrHost.NewHost(node.ctx)
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
