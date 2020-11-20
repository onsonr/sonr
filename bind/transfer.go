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
)

// ^ Data Stream Struct ^ //
type transferStreamConn struct {
	stream   network.Stream
	callback Callback
	self     *Node
}

// ^ Handle Incoming File Buffer ^ //
func (sn *Node) HandleTransferStream(stream network.Stream) {
	// Create/Set Transfer Stream
	sn.TransferStream = transferStreamConn{
		stream:   stream,
		callback: sn.Callback,
		self:     sn,
	}
	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.TransferStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewTransferStream(id peer.ID) error {
	// Start New Transfer Stream
	stream, err := sn.Host.NewStream(sn.CTX, id, protocol.ID("/sonr/transfer"))
	if err != nil {
		return err
	}

	// Create/Set Transfer Stream
	sn.TransferStream = transferStreamConn{
		stream:   stream,
		callback: sn.Callback,
		self:     sn,
	}

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.TransferStream.Read()
	return nil
}

// ^ Write Message on Stream ^ //
func (tsc *transferStreamConn) Write(authMsg *pb.AuthMessage) error {
	// Initialize Writer
	writer := bufio.NewWriter(tsc.stream)
	fmt.Println("Auth Msg Struct: ", authMsg)

	// Convert to String
	json, err := protojson.Marshal(authMsg)
	if err != nil {
		fmt.Println(err)
	}

	_, err = writer.WriteString(fmt.Sprintf("%s\n", string(json)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Write buffered data
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// ^ Read Data from Msgio ^ //
func (tsc *transferStreamConn) Read() error {
	for {
		// ** Read the Buffer **
		data, err := bufio.NewReader(tsc.stream).ReadString('\n')
		// Connection closed, deregister client
		if err == io.EOF {
			return nil
		}
		// Buffer Error
		if err != nil {
			fmt.Println(err)
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
		tsc.handleMessage(data)
	}
}

// ^ Handle Received Message ^ //
func (tsc *transferStreamConn) handleMessage(data string) {
	// Convert Json to Protobuf
	fmt.Println("Json String: ", data)
}
