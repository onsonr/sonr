package sonr

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
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
	Profile pb.Profile
	Contact pb.Contact

	// Networking Properties
	host       host.Host
	authStream authStreamConn
	dataStream dataStreamConn

	// References
	callback *Callback
	lobby    *lobby.Lobby
	files    []*pb.Metadata
}

// ^ Struct: Holds/Handles Stream for Authentication  ^ //
type authStreamConn struct {
	stream network.Stream
	self   *Node
}

// ^ Struct: Holds/Handles Stream for Data Transfer  ^ //
type dataStreamConn struct {
	stream network.Stream
	self   *Node
}

// ** Error Callback to Plugin with error **
func (sn *Node) sendError(err error, method string) {
	// Create Error Struct
	errorMsg := pb.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		fmt.Println("ERROR CALLBACK ERROR: ", err)
	}

	// Check and callback
	if sn.callback != nil {
		// Reference
		callRef := *sn.callback
		callRef.OnError(bytes)
	}
}
