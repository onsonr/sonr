package stream

import (
	"context"
	"fmt"

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
type OnProgressed func(float32)
type OnComplete func(data []byte)

// Struct to Implement Node Callback Methods
type DataCallback struct {
	Progressed OnProgressed
	Completed  OnComplete
	Error      OnError
}

// ^ Struct: Holds/Handles Stream for Authentication  ^ //
type DataStreamConn struct {
	// Properties
	Call DataCallback
	Host host.Host
	File sf.SonrFile

	// Peer/Self Info
	PeerID peer.ID
	Self   *pb.Peer
	Peer   *pb.Peer

	// RPC Structs
	progressClient ProgressClient
	progressServer ProgressServer

	// Stream Info
	id     string
	stream network.Stream
}

// ^ Start New Stream ^ //
func (dsc *DataStreamConn) Transfer(ctx context.Context, sm *sf.SafeMeta) {
	// Create New Auth Stream
	stream, err := dsc.Host.NewStream(ctx, dsc.PeerID, protocol.ID("/sonr/data"))
	if err != nil {
		dsc.Call.Error(err, "Transfer")
	}

	// Set Stream
	dsc.stream = stream
	dsc.id = stream.ID()

	// Set RPC Server
	dsc.progressServer = setSender(dsc.Host, dsc.Call.Progressed)

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go dsc.writeMessages(sm)
}

// ^ Handle Incoming Stream ^ //
func (dsc *DataStreamConn) HandleStream(stream network.Stream) {
	// Set Stream
	dsc.stream = stream
	dsc.id = stream.ID()

	// Set RPC Client
	dsc.progressClient = setReceiver(dsc.Host, dsc.PeerID, dsc.Call.Progressed)

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	reader := msgio.NewReader(dsc.stream)
	go dsc.readBlock(reader)
}

// ^ read Data from Msgio ^ //
func (dsc *DataStreamConn) readBlock(reader msgio.ReadCloser) error {
	for {
		// @ Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			dsc.Call.Error(err, "ReadMsg")
			return err
		}

		// @ Add Buffer Data to File, Check for Completed
		hasCompleted, progress, err := dsc.File.AddBuffer(buffer)
		if err != nil {
			dsc.Call.Error(err, "AddBuffer")
			return err
		}

		// @ Check for Completed
		if hasCompleted {
			// Save The File
			metadata, err := dsc.File.Save(dsc.Peer)
			if err != nil {
				dsc.Call.Error(err, "Save")
			}

			// Convert to Bytes
			bytes, err := proto.Marshal(metadata)
			if err != nil {
				dsc.Call.Error(err, "Completed")
			}

			// Callback Completed
			go dsc.Call.Completed(bytes)
			break
		}

		// @ Send Progress
		go dsc.Call.Progressed(progress)
	}
	return nil
}

func (dsc *DataStreamConn) writeMessages(file *sf.SafeMeta) error {
	// Get Data
	writer := msgio.NewWriter(dsc.stream)
	total := sf.GetSize(file)

	// Check type before splitting
	if file.Metadata().Mime.Type == "image" {
		// Iterate for Entire file as String
		for i, chunk := range sf.ChunkBase64(file) {
			// Create Block Protobuf from Chunk
			chunk := pb.Chunk{
				Size:    int32(len(chunk)),
				B64:     chunk,
				Current: int32(i),
				Total:   total,
			}

			// Convert to bytes
			bytes, err := proto.Marshal(&chunk)
			if err != nil {
				dsc.Call.Error(err, "writeMessages-base64")
			}

			// Write Message Bytes to Stream
			err = writer.WriteMsg(bytes)
			if err != nil {
				dsc.Call.Error(err, "writeMessages-base64")
			}
		}
	} else {
		// Iterate for Entire file as Bytes
		for i, chunk := range sf.ChunkBytes(file) {
			// Create Block Protobuf from Chunk
			chunk := pb.Chunk{
				Size:    int32(len(chunk)),
				Buffer:  chunk,
				Current: int32(i),
				Total:   total,
			}

			// Convert to bytes
			bytes, err := proto.Marshal(&chunk)
			if err != nil {
				dsc.Call.Error(err, "writeMessages-base64")
			}

			// Write Message Bytes to Stream
			err = writer.WriteMsg(bytes)
			if err != nil {
				dsc.Call.Error(err, "writeMessages-base64")
			}
		}
	}

	return nil
}
