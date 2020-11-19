package sonr

import (
	"fmt"

	io "github.com/gogo/protobuf/io"
	"github.com/libp2p/go-libp2p-core/network"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Auth Stream Struct ^ //
type authStreamConn struct {
	self     *Node
	writer   io.WriteCloser
	reader   io.ReadCloser
	stream   network.Stream
	callback Callback
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	// Create a buffer stream for non blocking read and write.
	protow := io.NewDelimitedWriter(stream)
	protor := io.NewDelimitedReader(stream, 16000)

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		reader:   protor,
		writer:   protow,
		stream:   stream,
		callback: sn.Callback,
		self:     sn,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	// Create a buffer stream for non blocking read and write.
	protow := io.NewDelimitedWriter(stream)
	protor := io.NewDelimitedReader(stream, 16000)

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		reader:   protor,
		writer:   protow,
		stream:   stream,
		callback: sn.Callback,
		self:     sn,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Write Message on Stream ^ //
func (asc *authStreamConn) Write(authMsg *pb.AuthMessage) error {
	fmt.Println("Auth Msg Struct: ", authMsg)
	// Write Protobuf
	err := asc.writer.WriteMsg(authMsg)
	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}

	// Write buffered data
	err = asc.writer.Close()
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
		authMsg := &pb.AuthMessage{}
		err := asc.reader.ReadMsg(authMsg)
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
				data, err := proto.Marshal(authMsg)
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
