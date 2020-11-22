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
func (asc *AuthStreamConn) Send(authMsg *pb.AuthMessage) error {
	// Initialize Writer
	writer := bufio.NewWriter(asc.stream)
	fmt.Println("Auth Msg Struct: ", authMsg)

	// Convert to Bytes
	bytes, err := proto.Marshal(authMsg)
	if err != nil {
		fmt.Println(err)
	}

	// Write to Stram
	_, err = writer.Write(bytes)
	if err != nil {
		fmt.Println("Auth Stream Outgoing Write Error: ", err)
		return err
	}

	// Write buffered data
	err = writer.Flush()
	if err != nil {
		fmt.Println("Auth Stream Outgoing Flush Error: ", err)
		return err
	}
	return nil
}

// ^ read Data from Msgio ^ //
func (asc *AuthStreamConn) readLoop() error {
	for {
		// Create Source Reader and Dest Writer
		source := bufio.NewReader(asc.stream)
		buffer := new(bytes.Buffer)

		// Copy Bytes from reader to writer
		if _, err := io.Copy(buffer, source); err != nil {
			fmt.Println("Auth Stream Incoming Copy Error: ", err)
			return err
		}

		// Create Message from Buffer
		message := &pb.AuthMessage{}
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
		asc.Call.Invited(msgBytes)

	// @2. Peer Accepted Response to Invite
	case pb.AuthMessage_ACCEPT:
		asc.Call.Responded(msgBytes)

	// @3. Peer Declined Response to Invite
	case pb.AuthMessage_DECLINE:
		asc.Call.Responded(msgBytes)

	// ! Invalid Subject
	default:
		err := errors.New(fmt.Sprintf("Not a subject: %s", msg.Event))
		asc.Call.Error(err, "handleMessage")
	}
}
