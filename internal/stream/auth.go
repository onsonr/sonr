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
type AuthCallback struct {
	Invited   OnInvited
	Responded OnResponded
	Error     OnError
}

// ^ Struct: Holds/Handles Stream for Authentication  ^ //
type AuthStreamConn struct {
	Call     AuthCallback
	Self     *pb.Peer
	Peer     *pb.Peer
	Metadata *pb.Metadata

	id     string
	stream network.Stream
}

// ^ Handle Incoming Stream ^ //
func (asc *AuthStreamConn) HandleStream(stream network.Stream) {
	// Set Stream
	asc.id = stream.ID()
	asc.stream = stream

	// Print Stream Info
	fmt.Println("Stream Info: ", stream.Stat())

	// Initialize Routine
	mrw := msgio.NewReader(asc.stream)
	go asc.read(mrw)
}

// ^ Start New Stream ^ //
func (asc *AuthStreamConn) Invite(ctx context.Context, h host.Host, id peer.ID, to *pb.Peer) error {
	// ** Set Peer ** //
	asc.Peer = to

	//@1. Create New Auth Stream
	stream, err := h.NewStream(ctx, id, protocol.ID("/sonr/auth"))
	if err != nil {
		return err
	}

	// Set Stream
	asc.stream = stream
	asc.id = stream.ID()

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	mrw := msgio.NewReader(asc.stream)
	go asc.read(mrw)

	// @2. Create Invite Message
	reqMsg := &pb.AuthMessage{
		Event:    pb.AuthMessage_REQUEST,
		From:     asc.Self,
		Metadata: asc.Metadata,
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

// ^ read Data from Msgio ^ //
func (asc *AuthStreamConn) read(mrw msgio.ReadCloser) error {
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
		asc.Metadata = msg.Metadata
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

// ^ Send Accept Message on Stream ^ //
func (asc *AuthStreamConn) Accept() error {
	// ** Validate Stream exists **
	if asc.stream == nil {
		err := errors.New("Auth Stream hasnt been set")
		return err
	}
	//@1. Create Message
	respMsg := &pb.AuthMessage{
		From:  asc.Self,
		Event: pb.AuthMessage_ACCEPT,
	}

	//@2. Convert Protobuf to bytes
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

// ^ Send Decline Message on Stream ^ //
func (asc *AuthStreamConn) Decline() error {
	// ** Validate Stream exists **
	if asc.stream == nil {
		err := errors.New("Auth Stream hasnt been set")
		return err
	}
	//@1. Create Message
	respMsg := &pb.AuthMessage{
		From:  asc.Self,
		Event: pb.AuthMessage_DECLINE,
	}

	//@2. Convert Protobuf to bytes
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
