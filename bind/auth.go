package sonr

import (
	"bufio"
	"fmt"

	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-core/network"
	pb "github.com/sonr-io/core/pkg/models"
	io "github.com/tessellator/protoio"
)

// ^ Auth Stream Struct ^ //
type authStreamConn struct {
	self       *Node
	readWriter *bufio.ReadWriter
	stream     network.Stream
	callback   Callback
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	// Create a buffer stream for non blocking read and write.
	rwr := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		readWriter: rwr,
		stream:     stream,
		callback:   sn.Callback,
		self:       sn,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	// Create a buffer stream for non blocking read and write.
	rwr := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		readWriter: rwr,
		stream:     stream,
		callback:   sn.Callback,
		self:       sn,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Write Message on Stream ^ //
func (asc *authStreamConn) Write(authMsg *pb.AuthMessage) error {
	fmt.Println("Auth Msg Struct: ", authMsg)
	marshaledBytes, err := proto.Marshal(authMsg)
	length := len(marshaledBytes)

	bytesWritten, err := io.Write(asc.readWriter, authMsg)
	if err != nil {
		fmt.Println("Write() returned err: ", err)
		return err
	}
	if bytesWritten != int64(length+4) {
		fmt.Println("Write() did not return correct number of bytes written")
	}
	asc.readWriter.Flush()
	return nil
}

// ^ Read Data from Msgio ^ //
func (asc *authStreamConn) Read() {
	for {
		// ** Read the Buffer **
		authMsg := pb.AuthMessage{}
		readBytes, err := ioutil.ReadAll(asc.readWriter)
		err = proto.Unmarshal(readBytes, &authMsg)
		if err != nil {
			fmt.Println("Error unmarshaling msg: ", err)
		}

		// ** Contains Data **
		// Construct message
		fmt.Println("Received String Message:", authMsg.String())
		if authMsg.Subject != pb.AuthMessage_NONE {
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
			}
		}
		// ! Invalid Subject
		fmt.Println("Not a subject", authMsg.Subject)
	}
}
