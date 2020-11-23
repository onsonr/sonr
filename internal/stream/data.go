package stream

import (
	"context"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	msgio "github.com/libp2p/go-msgio"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// Define Function Types
type OnProgressed func(data []byte)

// Struct to Implement Node Callback Methods
type DataCallback struct {
	Progressed OnProgressed
	Error      OnError
}

// ^ Struct: Holds/Handles Stream for Authentication  ^ //
type DataStreamConn struct {
	Call StreamCallback
	Self *pb.Peer

	id     string
	data   *pb.Metadata
	remote *pb.Peer
	stream network.Stream
}

// ^ Start New Stream ^ //
func (dsc *DataStreamConn) NewStream(ctx context.Context, h host.Host, id peer.ID, r *pb.Peer, s *pb.Peer) error {
	// Create New Auth Stream
	stream, err := h.NewStream(ctx, id, protocol.ID("/sonr/auth"))
	if err != nil {
		return err
	}

	// Set Stream
	dsc.stream = stream
	dsc.id = stream.ID()

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go dsc.read()
	return nil
}

// ^ Handle Incoming Stream ^ //
func (dsc *DataStreamConn) HandleStream(stream network.Stream) {
	// Set Stream
	dsc.stream = stream
	dsc.id = stream.ID()

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go dsc.read()
}

// ^ read Data from Msgio ^ //
func (dsc *DataStreamConn) read() error {
	// Read Length Fixed Bytes
	mrw := msgio.NewReadWriter(dsc.stream)
	lengthBytes, err := mrw.ReadMsg()
	if err != nil {
		return err
	}

	// Unmarshal Bytes into Proto
	protoMsg := &pb.AuthMessage{}
	err = proto.Unmarshal(lengthBytes, protoMsg)
	if err != nil {
		return err
	}

	dsc.handleBlock(protoMsg)
	return nil
}

// ^ Handle Received Message ^ //
func (dsc *DataStreamConn) handleBlock(msg *pb.AuthMessage) {
	// ** Contains Data **
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}

	// ** Check Message Subject **
	switch msg.Event {
	// @1. Request to Invite
	case pb.AuthMessage_REQUEST:
		fmt.Println("Handling Message received Request: ", msg.String())
		dsc.Call.Invited(msgBytes)

	// @2. Peer Accepted Response to Invite
	case pb.AuthMessage_ACCEPT:
		fmt.Println("Handling Message received Accept: ", msg.String())
		dsc.Call.Responded(msgBytes)

	// @3. Peer Declined Response to Invite
	case pb.AuthMessage_DECLINE:
		fmt.Println("Handling Message received Decline: ", msg.String())
		dsc.Call.Responded(msgBytes)

	// ! Invalid Subject
	default:
		err := errors.New(fmt.Sprintf("Not a subject: %s", msg.Event))
		dsc.Call.Error(err, "handleMessage")
	}
}

// // ^ writeAuthMessage Message on Stream ^ //
// func (dsc *DataStreamConn) Send(to *pb.Peer, item *sf.Item) error {
// 	// @1. Retreive Blocks in Data
// 	if !item.IsReady {
// 		err := errors.New("ItemFile blocks arent ready")
// 		return err
// 	}

// 	blocks := item.GetBlocks()

// 	// Initialize Writer
// 	writer := msgio.NewWriter(dsc.stream)

// 	// Add Msg to buffer
// 	if err := writer.WriteMsg(bytes); err != nil {
// 		return err
// 	}
// 	return nil
// }

// ^ writeAuthMessage Message on Stream ^ //
func (dsc *DataStreamConn) SendComplete(from *pb.Peer, to *pb.Peer, decision bool) error {
	//@1. Create Message
	respMsg := &pb.AuthMessage{
		From: from,
	}

	// ** Check Decision **
	if decision == true {
		// User Accepted
		respMsg.Event = pb.AuthMessage_ACCEPT // Set Event
	} else {
		// @ User Declined
		respMsg.Event = pb.AuthMessage_DECLINE // Set Event
	}

	// Convert to bytes
	bytes, err := proto.Marshal(respMsg)
	if err != nil {
		return err
	}

	// Initialize Writer
	writer := msgio.NewWriter(dsc.stream)

	// Add Msg to buffer
	if err := writer.WriteMsg(bytes); err != nil {
		return err
	}
	return nil
}
