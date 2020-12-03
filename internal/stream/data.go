package stream

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	msgio "github.com/libp2p/go-msgio"
	sonrFile "github.com/sonr-io/core/internal/file"
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

const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 32000

// ^ Struct: Holds/Handles Stream for Authentication  ^ //
type DataStreamConn struct {
	// Properties
	Call DataCallback
	Host host.Host
	File sonrFile.TransferFile

	// Peer/Self Info
	PeerID peer.ID
	Self   *pb.Peer
	Peer   *pb.Peer

	// Stream Info
	id     string
	stream network.Stream
}

// ^ Start New Stream ^ //
func (dsc *DataStreamConn) Transfer(ctx context.Context, sm *sonrFile.SafeFile) {
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
	dsc.writeMessages(sm)
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
	for i := 0; ; i++ {
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
			// @ Send Progress every ~0.5MB
			if i%16 == 0 {
				dsc.Call.Progressed(progress)
			}
		}
	}
}

// ^ write file to Msgio ^ //
func (dsc *DataStreamConn) writeMessages(sf *sonrFile.SafeFile) {
	// Initialize Writer
	writer := msgio.NewWriter(dsc.stream)

	// @ Check Type
	if sf.Mime.Type == pb.MIME_image {
		// Retreive Base64 Value for Image File
		b64, total := sf.Base64()
		// Chunk in Goroutine
		go func(writer msgio.WriteCloser, data string, total int32) {
			// Iterate for Entire file as String
			for i, chunk := range sonrFile.ChunkBase64(data, B64ChunkSize) {
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
		}(writer, b64, total)
	} else {
		meta := sf.Metadata()

		// Return Given Size
		total := meta.Size

		// Open File
		file, err := os.Open(meta.Path)
		if err != nil {
			dsc.Call.Error(err, "writeMessages-buffer")
			log.Fatalln(err)
		}
		defer file.Close()

		// Set Chunk Variables
		ps := make([]byte, BufferChunkSize)

		// Iterate file
		for i := 0; ; i++ {
			// Read Bytes
			bytesread, err := file.Read(ps)

			// Check for Error
			if err != nil {
				// Non EOF Error
				if err != io.EOF {
					fmt.Println(err)
				}
				// File Complete
				break
			}

			// Create Block Protobuf from Chunk
			chunk := pb.Chunk{
				Size:    int32(len(ps[:bytesread])),
				Buffer:  ps[:bytesread],
				Current: int32(i),
				Total:   total,
			}

			// Convert to bytes
			bytes, err := proto.Marshal(&chunk)
			if err != nil {
				dsc.Call.Error(err, "writeMessages-buffer")
				log.Fatalln(err)
			}

			// Write Message Bytes to Stream
			err = writer.WriteMsg(bytes)
			if err != nil {
				dsc.Call.Error(err, "writeMessages-buffer")
				log.Fatalln(err)
			}
		}
	}
}
