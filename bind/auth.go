package sonr

import (
	"bufio"
	"bytes"
	"context"
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
	go sn.authStream.readLoop()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(id peer.ID) error {
	// Start New Auth Stream
	ctx := context.Background()
	stream, err := sn.host.NewStream(ctx, id, protocol.ID("/sonr/auth"))
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
	go sn.authStream.readLoop()
	return nil
}

// ^ write Message on Stream ^ //
func (asc *authStreamConn) write(authMsg *pb.AuthMessage) error {
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
func (asc *authStreamConn) readLoop() error {
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
func (asc *authStreamConn) handleMessage(message *pb.AuthMessage) {
	// ** Contains Data **
	// Convert Protobuf to bytes
	authRaw, err := proto.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	// Check Message Subject
	switch message.Subject {
	// @ Request to Invite
	case pb.AuthMessage_REQUEST:
		// Callback the Invitation
		callbackRef := *asc.self.callback
		callbackRef.OnInvited(authRaw)

	// @ Peer Accepted Response to Invite
	case pb.AuthMessage_ACCEPT:
		fmt.Println("Auth Accepted")
		// Callback to Proxies
		callbackRef := *asc.self.callback
		callbackRef.OnResponded(authRaw)

	// @ Peer Accepted Response to Invite
	case pb.AuthMessage_DECLINE:
		fmt.Println("Auth Declined")
		// Callback to Proxies
		callbackRef := *asc.self.callback
		callbackRef.OnResponded(authRaw)

	// ! Invalid Subject
	default:
		fmt.Println("Not a subject", message.Subject)
	}
}
