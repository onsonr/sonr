package stream

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

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

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	reader := msgio.NewReader(dsc.stream)
	go dsc.readBlock(reader)
}

// ^ read Data from Msgio ^ //
func (dsc *DataStreamConn) readBlock(reader msgio.ReadCloser) {
	for {
		// @ Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			dsc.Call.Error(err, "ReadMsg")
			log.Fatalln(err)
			break
		}

		// @ Add Buffer Data to File, Check for Completed
		hasCompleted, progress, err := dsc.File.AddBuffer(buffer)
		if err != nil {
			dsc.Call.Error(err, "AddBuffer")
			log.Fatalln(err)
			break
		}

		// @ Check for Completed
		if hasCompleted {
			// Save The File
			metadata, err := dsc.File.Save(dsc.Peer)
			if err != nil {
				dsc.Call.Error(err, "Save")
				log.Fatalln(err)
				break
			}

			// Convert to Bytes
			bytes, err := proto.Marshal(metadata)
			if err != nil {
				dsc.Call.Error(err, "Completed")
				log.Fatalln(err)
				break
			}

			// Callback Completed
			dsc.Call.Completed(bytes)
			break
		} else {
			// @ Send Progress
			// Only 20 Callbacks per transfer to limit UI thread
			rounded := int(progress) * 100
			if rounded%5 == 0 {
				dsc.Call.Progressed(progress)
			}
		}
	}
}

// ^ write file to Msgio ^ //
func (dsc *DataStreamConn) writeMessages(file *sf.SafeMeta) {
	// Get Data
	writer := msgio.NewWriter(dsc.stream)
	meta := file.Metadata()
	imgBuffer := new(bytes.Buffer)

	// Check Type for image
	if meta.Mime.Type == "image" {
		// New File for ThumbNail
		file, err := os.Open(meta.Path)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		// Convert to Image Object
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatalln(err)
		}

		// Encode as Jpeg into buffer
		err = jpeg.Encode(imgBuffer, img, &jpeg.Options{Quality: 100})
		if err != nil {
			log.Fatalln(err)
		}

		// Return Adjusted Size
		b64, length := Base64(imgBuffer.Bytes())
		total := int32(length)

		// Iterate for Entire file as String
		for i, chunk := range ChunkBase64(b64) {
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
				log.Fatalln(err)
			}

			// Write Message Bytes to Stream
			err = writer.WriteMsg(bytes)
			if err != nil {
				dsc.Call.Error(err, "writeMessages-base64")
				log.Fatalln(err)
			}
		}
	} else {
		// Return Given Size
		total := meta.Size

		// Iterate for Entire file as Bytes
		for i, chunk := range ChunkBytes(meta.Path, int(total)) {
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
}
