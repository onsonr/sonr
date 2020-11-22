package stream

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
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
	id     string
	stream network.Stream
}

// ^ Start New Stream ^ //
func (asc *AuthStreamConn) New(ctx context.Context, h host.Host, id peer.ID) error {
	// Create New Auth Stream
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
	go asc.read()
	return nil
}

// ^ Handle Incoming Stream ^ //
func (asc *AuthStreamConn) SetStream(stream network.Stream) {
	// Set Stream
	asc.stream = stream
	asc.id = stream.ID()

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go asc.read()
}

// ^ read Data from Msgio ^ //
func (asc *AuthStreamConn) read() error {
	// Find Message Size
	buf := make([]byte, 4) // 32 bits
	if _, err := io.ReadFull(asc.stream, buf); err != nil {
		fmt.Println("Error getting byte size: ", err)
		return err
	}
	n, size := binary.Varint(buf)
	if n == 0 {
		return errors.New("Buffer is too small")
	}
	if n < 0 {
		
	}

	// Create Bytes from message
	bytes := make([]byte, size)
	if _, err := io.ReadFull(asc.stream, bytes); err != nil {
		fmt.Println("Error reading full: ", err)
		return err
	}

	// Unmarshal into protobuf
	message := &pb.AuthMessage{}
	err := proto.Unmarshal(bytes, message)
	if err != nil {
		fmt.Println("Error Unmarshaling message: ", err)
		return err
	}

	asc.handleMessage(message)
	return nil
}

// ^ Handle Received Message ^ //
func (asc *AuthStreamConn) handleMessage(msg *pb.AuthMessage) {
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
		asc.Call.Invited(msgBytes)

	// @2. Peer Accepted Response to Invite
	case pb.AuthMessage_ACCEPT:
		fmt.Println("Handling Message received Accept: ", msg.String())
		asc.Call.Responded(msgBytes)

	// @3. Peer Declined Response to Invite
	case pb.AuthMessage_DECLINE:
		fmt.Println("Handling Message received Decline: ", msg.String())
		asc.Call.Responded(msgBytes)

	// ! Invalid Subject
	default:
		err := errors.New(fmt.Sprintf("Not a subject: %s", msg.Event))
		asc.Call.Error(err, "handleMessage")
	}
}

// ^ writeAuthMessage Message on Stream ^ //
func (asc *AuthStreamConn) SendInvite(from *pb.Peer, to *pb.Peer, meta *pb.Metadata) error {
	// @1. Create Message
	reqMsg := &pb.AuthMessage{
		Event:    pb.AuthMessage_REQUEST,
		From:     from,
		To:       to,
		Metadata: meta,
	}

	// Convert to bytes
	bytes, err := proto.Marshal(reqMsg)
	if err != nil {
		return err
	}

	// Initialize Writer
	writer := bufio.NewWriter(asc.stream)

	// Add Varint and write to buffer
	writeWithVarint(writer, bytes)
	return nil
}

// ^ writeAuthMessage Message on Stream ^ //
func (asc *AuthStreamConn) SendResponse(from *pb.Peer, to *pb.Peer, decision bool) error {
	//@1. Create Message
	respMsg := &pb.AuthMessage{
		From: from,
		To:   to,
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
	writer := bufio.NewWriter(asc.stream)

	// Add Varint and write to buffer
	writeWithVarint(writer, bytes)
	return nil
}

// ^ Adds Varint to Byte Buffer ^ //
func writeWithVarint(w io.Writer, msg []byte) error {
	buf := make([]byte, 8)
	binary.PutVarint(buf, int64(len(msg)))

	if _, err := w.Write(buf); err != nil {
		return err
	}

	if _, err := w.Write(msg); err != nil {
		return err
	}
	return nil
}
