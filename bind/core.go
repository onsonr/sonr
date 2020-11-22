package sonr

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
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
	OnQueued(status bool)
	OnProgress(data []byte)
	OnError(data []byte)
}

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Public Properties
	HostID peer.ID
	Peer   *pb.Peer

	// Networking Properties
	ctx        context.Context
	host       host.Host
	authStream st.AuthStreamConn

	// References
	call  Callback
	lobby *lobby.Lobby
	files []*pb.Metadata
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
