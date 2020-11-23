package stream

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// Define Function Types
type OnProgressed func(data []byte)
type OnComplete func(data []byte)

// Struct to Implement Node Callback Methods
type DataCallback struct {
	Progressed OnProgressed
	Completed  OnComplete
	Error      OnError
}

// ^ Struct: Holds/Handles Stream for Authentication  ^ //
type DataStreamConn struct {
	Call DataCallback
	Self *pb.Peer
	File sf.SonrFile

	id     string
	data   *pb.Metadata
	remote *pb.Peer
	stream network.Stream
}

// ^ Start New Stream ^ //
func (dsc *DataStreamConn) Transfer(ctx context.Context, h host.Host, id peer.ID, r *pb.Peer, tf *sf.TransferFile) error {
	// Create New Auth Stream
	stream, err := h.NewStream(ctx, id, protocol.ID("/sonr/data"))
	if err != nil {
		return err
	}

	// Set Stream
	dsc.stream = stream
	dsc.id = stream.ID()
	dsc.remote = r

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go dsc.writeFileToStream(tf)
	return nil
}

// ^ Handle Incoming Stream ^ //
func (dsc *DataStreamConn) HandleStream(stream network.Stream) {
	// Set Stream
	dsc.stream = stream
	dsc.id = stream.ID()

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	mrw := msgio.NewReader(dsc.stream)
	go dsc.read(mrw)
}

// ^ read Data from Msgio ^ //
func (dsc *DataStreamConn) read(mrw msgio.ReadCloser) error {
	// Read Length Fixed Bytes
	lengthBytes, err := mrw.ReadMsg()
	if err != nil {
		return err
	}

	// Unmarshal Bytes into Proto
	protoMsg := &pb.Block{}
	err = proto.Unmarshal(lengthBytes, protoMsg)
	if err != nil {
		return err
	}

	dsc.handleBlock(protoMsg)
	return nil
}

// ^ Handle Received Message ^ //
func (dsc *DataStreamConn) handleBlock(msg *pb.Block) {
	// Verify Bytes Remaining
	fmt.Println("Current ", msg.Current, "Total ", msg.Total)

	if msg.Current < msg.Total {
		// Add Block to Buffer
		dsc.File.AddBlock(msg.Data)
	}

	// Save File on Buffer Complete
	if msg.Current == msg.Total {
		// Add Block to Buffer
		dsc.File.AddBlock(msg.Data)
		fmt.Println("Completed All Blocks, Save the File")
	}
}

func (dsc *DataStreamConn) writeFileToStream(tf *sf.TransferFile) error {
	// Retreive Transfer Blocks
	tf.Generate()

	// Create Delay to allow processing
	time.Sleep(time.Second)
	blocks := tf.Blocks()

	// Initialize Writer
	writer := msgio.NewWriter(dsc.stream)

	// Iterate through blocks and write to message
	for _, block := range blocks {
		// Convert to bytes
		bytes, err := proto.Marshal(block)
		if err != nil {
			return err
		}

		// Add Msg to buffer
		if err := writer.WriteMsg(bytes); err != nil {
			return err
		}
	}
	return nil
}
