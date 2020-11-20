package sonr

import (
	"bufio"
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ^ Auth Stream Struct ^ //
type authStreamConn struct {
	stream   network.Stream
	callback Callback
	self     *Node
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		stream:   stream,
		callback: sn.Call,
		self:     sn,
	}
	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(id peer.ID) error {
	// Start New Auth Stream
	stream, err := sn.Host.NewStream(sn.CTX, id, protocol.ID("/sonr/auth"))
	if err != nil {
		return err
	}

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		stream:   stream,
		callback: sn.Call,
		self:     sn,
	}

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.AuthStream.Read()
	return nil
}

// ^ Write Message on Stream ^ //
func (asc *authStreamConn) Write(authMsg *pb.AuthMessage) error {
	// Initialize Writer
	writer := bufio.NewWriter(asc.stream)
	fmt.Println("Auth Msg Struct: ", authMsg)

	// Convert to String
	json, err := protojson.Marshal(authMsg)
	if err != nil {
		fmt.Printf("Error: %s, %s", err, pb.Error_JSON)
	}

	// Write Message with "Delimiter"=(Seperator for Message Values)
	_, err = writer.WriteString(fmt.Sprintf("%s\n", string(json)))
	if err != nil {
		fmt.Printf("Error: %s, %s", err, pb.Error_BUFFER)
		return err
	}

	// Write buffered data
	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error: %s, %s", err, pb.Error_BUFFER)
		return err
	}
	return nil
}

// ^ Read Data from Msgio ^ //
func (asc *authStreamConn) Read() error {
	for {
		// ** Read the Buffer **
		data, err := bufio.NewReader(asc.stream).ReadString('\n')
		// Connection closed, deregister client
		if err == io.EOF {
			return nil
		}
		// Buffer Error
		if err != nil {
			fmt.Printf("Error: %s, %s", err, pb.Error_BUFFER)
			return err
		}

		// Empty String
		if data == "" {
			return nil
		}

		// End of Message
		if data == "\n" {
			return nil
		}

		// @ Handle it
		asc.handleMessage(data)
	}
}

// ^ Handle Received Message ^ //
func (asc *authStreamConn) handleMessage(data string) {
	// Convert Json to Protobuf
	fmt.Println("Json String: ", data)
	authMsg := pb.AuthMessage{}
	err := protojson.Unmarshal([]byte(data), &authMsg)
	if err != nil {
		fmt.Printf("Error: %s, %s", err, pb.Error_PROTO)
	}

	// Convert Protobuf to bytes
	authRaw, err := proto.Marshal(&authMsg)
	if err != nil {
		fmt.Printf("Error: %s, %s", err, pb.Error_BYTES)
	}

	// ** Contains Data **
	// Check Message Subject
	switch authMsg.Subject {
	// @ Request to Invite
	case pb.AuthMessage_REQUEST:
		// Callback the Invitation
		asc.self.Callback(pb.Callback_INVITED, authRaw)

	// @ Peer Accepted Response to Invite
	case pb.AuthMessage_ACCEPT:
		fmt.Println("Auth Accepted")
		// Callback to Proxies
		asc.self.Callback(pb.Callback_RESPONDED, authRaw)

	// @ Peer Accepted Response to Invite
	case pb.AuthMessage_DECLINE:
		fmt.Println("Auth Declined")
		// Callback to Proxies
		asc.self.Callback(pb.Callback_RESPONDED, authRaw)

	// ! Invalid Subject
	default:
		fmt.Println("Not a subject", authMsg.Subject)
		fmt.Printf("Error: %s, %s", err, pb.Error_PEER)
	}
}
