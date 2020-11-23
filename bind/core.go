package sonr

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	sf "github.com/sonr-io/core/internal/file"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/lobby"
	pb "github.com/sonr-io/core/internal/models"
	st "github.com/sonr-io/core/internal/stream"
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
	OnProgress(data []byte)
	OnCompleted(data []byte)
	OnError(data []byte)
}

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Public Properties
	HostID string
	Peer   *pb.Peer

	// Networking Properties
	ctx        context.Context
	host       host.Host
	authStream st.AuthStreamConn
	dataStream st.DataStreamConn

	// References
	call  Callback
	lobby *lobby.Lobby
	files []*sf.SafeMeta
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(reqBytes []byte, call Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	node := new(Node)
	node.ctx = context.Background()
	node.call, node.files = call, make([]*sf.SafeMeta, maxFileBufferSize)

	// ** Unmarshal Request **
	reqMsg := pb.RequestMessage{}
	err := proto.Unmarshal(reqBytes, &reqMsg)
	if err != nil {
		fmt.Println(err)
		node.Error(err, "NewNode")
		return nil
	}

	// @1. Create Host and Set Stream Handlers
	node.host, err = sh.NewHost(node.ctx)
	if err != nil {
		node.Error(err, "NewNode")
		return nil
	}
	node.HostID = node.host.ID().String()
	node.initStreams()

	// @3. Set Node User Information
	if err = node.setPeer(&reqMsg); err != nil {
		node.Error(err, "NewNode")
		return nil
	}

	// @4. Setup Discovery w/ Lobby
	if err = node.setDiscovery(node.ctx, &reqMsg); err != nil {
		node.Error(err, "NewNode")
		return nil
	}

	// ** Callback Node User Information ** //
	return node
}

// ** Error Callback to Plugin with error **
func (sn *Node) Error(err error, method string) {
	// Log In Core
	fmt.Println(fmt.Sprintf("[Error] At Method %s : %s", err.Error(), method))

	// Create Error ProtoBuf
	errorMsg := pb.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		fmt.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Check and callback
	if sn.call != nil {
		// Reference
		sn.call.OnError(bytes)
	}
}
