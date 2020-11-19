package sonr

import (
	"bufio"
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p-core/network"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ^ Auth Stream Struct ^ //
type authStreamConn struct {
	stream   network.Stream
	callback Callback
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		stream:   stream,
		callback: sn.Callback,
	}
	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		stream:   stream,
		callback: sn.Callback,
	}

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Write Message on Stream ^ //
func (asc *authStreamConn) Write(authMsg *pb.AuthMessage) error {
	// Initialize Writer
	writer := bufio.NewWriter(asc.stream)
	fmt.Println("Auth Msg Struct: ", authMsg)

	// Convert to String
	json, err := protojson.Marshal(authMsg)
	if err != nil {
		fmt.Println("Error Marshalling json: ", err)
	}

	// Write Message with "Delimiter"=(Seperator for Message Values)
	_, err = writer.WriteString(fmt.Sprintf("%s\n", string(json)))
	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}

	// Write buffered data
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer")
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
			fmt.Println("Error reading from buffer")
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
	// Convert Bytes to Json
	fmt.Println("Json String: ", data)
	authMsg := pb.AuthMessage{}
	err := protojson.Unmarshal([]byte(data), &authMsg)
	if err != nil {
		fmt.Println("Error unmarshaling msg into json: ", err)
	}

	// ** Contains Data **
	// Check Message Subject
	switch authMsg.Subject {
	// @ Request to Invite
	case pb.AuthMessage_REQUEST:
		// Retreive Values
		data, err := proto.Marshal(&authMsg)
		if err != nil {
			fmt.Println("Error Marshaling RefreshMessage ", err)
		}

		// Callback the Invitation
		asc.callback.OnInvited(data)

	// @ Peer Accepted Response to Invite
	case pb.AuthMessage_ACCEPT:
		fmt.Println("Auth Accepted")
		// Callback to Proxies
		asc.callback.OnResponded(true)

	// @ Peer Accepted Response to Invite
	case pb.AuthMessage_DECLINE:
		fmt.Println("Auth Declined")
		// Callback to Proxies
		asc.callback.OnResponded(false)

	// ! Invalid Subject
	default:
		fmt.Println("Not a subject", authMsg.Subject)
	}
}
