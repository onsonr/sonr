package transfer

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Define Function Types
type OnProtobuf func(data []byte)
type OnError func(err error, method string)

// Struct to Implement Node Callback Methods
type PeerCallback struct {
	Invited   OnProtobuf
	Responded OnProtobuf
	Completed OnProtobuf
	Error     OnError
}

// ^ Struct: Holds/Handles GRPC Calls and Handles Data Stream  ^ //
type PeerConnection struct {
	// Peer/Self Info
	Self *pb.Peer
	Peer *pb.Peer

	// Stream Info
	auth   *AuthHandler
	host   host.Host
	stream network.Stream

	// Data Handling
	Call PeerCallback
}

// ^ Initialize sets up new Peer Connection handler ^
func Initialize(host host.Host, callback PeerCallback) (*PeerConnection, error) {
	// Initialize Parameters into PeerConnection
	peerConn := &PeerConnection{
		Call: callback,
		host: host,
	}

	// Set Handlers
	host.SetStreamHandler(protocol.ID("/sonr/data/transfer"), peerConn.HandleTransfer)

	// Create Auth Handler
	peerConn.auth = NewAuthHandler(peerConn)
	return peerConn, nil
}

func (dsc *PeerConnection) Invite(peerID peer.ID, info *pb.Peer, sm *sf.SafeFile) {
	// @2. Create Invite Message
	reqMsg := &pb.AuthMessage{
		Event:    pb.AuthMessage_REQUEST,
		From:     info,
		Metadata: sm.Metadata(),
	}

	// Convert to bytes
	bytes, err := proto.Marshal(reqMsg)
	if err != nil {
		dsc.Call.Error(err, "Invite")
		log.Fatalln(err)
	}

	// Send GRPC Call
	err = dsc.auth.sendInvite(peerID, bytes)
	if err != nil {
		dsc.Call.Error(err, "Invite")
		log.Fatalln(err)
	}
}

// ^ Start New Stream ^ //
func (dsc *PeerConnection) Transfer(ctx context.Context, peerID peer.ID, sm *sf.SafeFile) {
	// Create New Auth Stream
	stream, err := dsc.host.NewStream(ctx, peerID, protocol.ID("/sonr/data/transfer"))
	if err != nil {
		dsc.Call.Error(err, "Transfer")
		log.Fatalln(err)
	}

	// Set Stream
	dsc.stream = stream

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	writer := msgio.NewWriter(dsc.stream)
	go dsc.writeTransferStream(writer, sm)
}
