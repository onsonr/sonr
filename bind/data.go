package sonr

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/encoding/protojson"
)

// ^ Handle Incoming File Buffer ^ //
func (sn *Node) HandleTransferStream(stream network.Stream) {
	// Create/Set Transfer Stream
	sn.dataStream = dataStreamConn{
		stream: stream,
		self:   sn,
	}
	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.dataStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewTransferStream(ctx context.Context, id peer.ID) error {
	// Start New Transfer Stream
	stream, err := sn.host.NewStream(ctx, id, protocol.ID("/sonr/transfer"))
	if err != nil {
		return err
	}

	// Create/Set Transfer Stream
	sn.dataStream = dataStreamConn{
		stream: stream,
		self:   sn,
	}

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go sn.dataStream.Read()
	return nil
}

// ^ Write Message on Stream ^ //
func (tsc *dataStreamConn) Write(authMsg *pb.AuthMessage) error {
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
func (tsc *dataStreamConn) Read() error {
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
func (tsc *dataStreamConn) handleMessage(data string) {
	// Convert Json to Protobuf
	fmt.Println("Json String: ", data)
}
