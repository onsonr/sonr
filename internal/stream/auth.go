package stream

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

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
	go asc.readLoop()
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
	go asc.readLoop()
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
	//

	// Convert to Bytes
	bytes, err := proto.Marshal(reqMsg)
	if err != nil {
		fmt.Println(err)
	}

	//@2. Initialize Writer
	fmt.Println("Auth Msg Struct: ", reqMsg)
	writer := bufio.NewWriter(asc.stream)

	// Write to Stram
	_, err = writer.Write(bytes)
	if err != nil {
		fmt.Println("Auth Stream Outgoing Write Error: ", err)
		return err
	}

	//@3. Write buffered data
	err = writer.Flush()
	if err != nil {
		fmt.Println("Auth Stream Outgoing Flush Error: ", err)
		return err
	}
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

	// Convert to Bytes
	bytes, err := proto.Marshal(respMsg)
	if err != nil {
		fmt.Println(err)
	}

	//@2. Initialize Writer
	fmt.Println("Auth Msg Struct: ", respMsg)
	writer := bufio.NewWriter(asc.stream)

	// Write to Stram
	_, err = writer.Write(bytes)
	if err != nil {
		fmt.Println("Auth Stream Outgoing Write Error: ", err)
		return err
	}

	//@3. Write buffered data
	err = writer.Flush()
	if err != nil {
		fmt.Println("Auth Stream Outgoing Flush Error: ", err)
		return err
	}
	return nil
}

// ^ read Data from Msgio ^ //
func (asc *AuthStreamConn) readLoop() error {
	source := bufio.NewReader(asc.stream)
	buffer := new(bytes.Buffer)
	for {
		// Create Source Reader and Dest Writer
		fmt.Println("Received message")

		// Copy Bytes from reader to writer
		_, err := io.Copy(buffer, source)
		fmt.Println("Copying Bytes to buffer")
		if err != nil {
			fmt.Println("Copying Error")
			return err
		}

		// Create Message from Buffer
		message := &pb.AuthMessage{}
		fmt.Println("Unmarshalling bytes into Message")
		if err := proto.Unmarshal(buffer.Bytes(), message); err != nil {
			log.Fatalln("Failed to parse auth message:", err)
			return err
		}

		// Handle Messages Struct
		asc.handleMessage(message)
	}
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
