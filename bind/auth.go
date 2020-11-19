package sonr

import (
	"bufio"
	"fmt"
	"io/ioutil"

	"github.com/libp2p/go-libp2p-core/network"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
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
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		readWriter: buffrw,
		stream:     stream,
		callback:   sn.Callback,
		self:       sn,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	// Create new Buffer
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		readWriter: buffrw,
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

	// Convert Request to Proto Binary
	data, err := proto.Marshal(authMsg)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}

	// Write Protobuf
	_, err = asc.readWriter.Write(data)
	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}

	// Write buffered data
	err = asc.readWriter.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer")
		return err
	}
	return nil
}

// ^ Read Data from Msgio ^ //
func (asc *authStreamConn) Read() {
	for {
		// ** Read the Buffer **
		authMsg := pb.AuthMessage{}
		data, _ := ioutil.ReadAll(asc.readWriter)
		err := proto.Unmarshal(data, &authMsg)
		if err != nil {
			fmt.Println("unmarshaling error: ", err)
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
		fmt.Printf("%s.\n", authMsg.Subject)
	}
}
