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
type OnInvited func(data []byte)
type OnResponded func(data []byte)
type OnError func(err error, method string)

// Struct to Implement Node Callback Methods
type StreamCallback struct {
	Invited   OnInvited
	Responded OnResponded
	Error     OnError
}

// ^ Struct: Holds/Handles Stream for Authentication  ^ //
type AuthStreamConn struct {
	Call   StreamCallback
	Self   *pb.Peer // Set on Config
	id     string
	stream network.Stream
	peer   *pb.Peer
}

// ^ Start New Stream ^ //
func (asc *AuthStreamConn) NewStream(ctx context.Context, h host.Host, id peer.ID, peer *pb.Peer) error {
	// Create New Auth Stream
	stream, err := h.NewStream(ctx, id, protocol.ID("/sonr/auth"))
	if err != nil {
		return err
	}

	// Set Stream
	asc.stream = stream
	asc.id = stream.ID()
	asc.peer = peer

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go asc.read()
	return nil
}

// ^ Handle Incoming Stream ^ //
func (asc *AuthStreamConn) HandleStream(stream network.Stream) {
	// Set Stream
	asc.id = stream.ID()
	asc.stream = stream

	// Print Stream Info
	fmt.Println("Stream Info: ", stream.Stat())

	// Initialize Routine
	go asc.read()
}

// ^ read Data from Msgio ^ //
func (asc *AuthStreamConn) read() error {
	// Read Length Fixed Bytes
	mrw := msgio.NewReadWriter(asc.stream)
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

	asc.handleMessage(protoMsg)
	return nil
}

// ^ Handle Received Message ^ //
func (asc *AuthStreamConn) handleMessage(msg *pb.AuthMessage) {
	// ** Convert Protobuf to bytes **
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}

	// ** Check Message Subject **
	switch msg.Event {
	// @1. Request to Invite
	case pb.AuthMessage_REQUEST:
		asc.Call.Invited(msgBytes)

	// @2. Peer Accepted Response to Invite
	case pb.AuthMessage_ACCEPT:
		asc.Call.Responded(msgBytes)

	// @3. Peer Declined Response to Invite
	case pb.AuthMessage_DECLINE:
		asc.Call.Responded(msgBytes)

	// ! Invalid Subject
	default:
		asc.Call.Error(err, "handleMessage")
	}
}

// ^ writeAuthMessage Message on Stream ^ //
func (asc *AuthStreamConn) SendInvite(from *pb.Peer, meta *pb.Metadata) error {
	// @1. Create Message
	reqMsg := &pb.AuthMessage{
		Event:    pb.AuthMessage_REQUEST,
		From:     from,
		Metadata: meta,
	}

	// Convert to bytes
	bytes, err := proto.Marshal(reqMsg)
	if err != nil {
		return err
	}

	// Initialize Writer
	writer := msgio.NewWriter(asc.stream)

	// Add Msg to buffer
	if err := writer.WriteMsg(bytes); err != nil {
		return err
	}
	return nil
}

// ^ writeAuthMessage Message on Stream ^ //
func (asc *AuthStreamConn) SendResponse(from *pb.Peer, decision bool) error {
	// ** Validate Stream exists **
	if asc.stream == nil {
		err := errors.New("Auth Stream hasnt been set")
		return err
	}
	//@1. Create Message
	respMsg := &pb.AuthMessage{
		From: from,
	}

	//@2. Check Decision
	if decision == true {
		// User Accepted
		respMsg.Event = pb.AuthMessage_ACCEPT // Set Event
	} else {
		// User Declined
		respMsg.Event = pb.AuthMessage_DECLINE // Set Event
	}

	//@3. Convert Protobuf to bytes
	bytes, err := proto.Marshal(respMsg)
	if err != nil {
		return err
	}

	// @4. Initialize Writer and Write to Stream
	writer := msgio.NewWriter(asc.stream)
	if err := writer.WriteMsg(bytes); err != nil {
		return err
	}
	return nil
}
