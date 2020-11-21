package sonr

import (
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
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
