package sonr

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	// Create/Set Auth Stream
	sn.authStream = authStreamConn{
		stream: stream,
		self:   sn,
	}
	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.authStream.readAuthStreamLoop()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(peerId peer.ID) error {
	// Start New Auth Stream
	ctx := context.Background()
	stream, err := sn.host.NewStream(ctx, peerId, protocol.ID("/sonr/auth"))
	if err != nil {
		return err
	}

	// Create/Set Auth Stream
	sn.authStream = authStreamConn{
		stream: stream,
		self:   sn,
	}

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.authStream.readAuthStreamLoop()
	return nil
}

// ^ writeAuthMessage Message on Stream ^ //
func (asc *authStreamConn) writeAuthMessage(authMsg *pb.AuthMessage) error {
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
func (asc *authStreamConn) readAuthStreamLoop() error {
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
		asc.handleAuthMessage(message)
	}
}

// ^ Handle Received Message ^ //
func (asc *authStreamConn) handleAuthMessage(msg *pb.AuthMessage) {
	// ** Contains Data **
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}

	// ** Check Message Subject **
	switch msg.Subject {
	// @1. Request to Invite
	case pb.AuthMessage_REQUEST:
		asc.self.call.OnInvited(msgBytes)

	// @2. Peer Accepted Response to Invite
	case pb.AuthMessage_ACCEPT:
		asc.self.call.OnResponded(msgBytes)

	// @3. Peer Declined Response to Invite
	case pb.AuthMessage_DECLINE:
		asc.self.call.OnResponded(msgBytes)

	// ! Invalid Subject
	default:
		err := errors.New(fmt.Sprintf("Not a subject: %s", msg.Subject))
		asc.self.Error(err, "handleMessage")
	}
}
