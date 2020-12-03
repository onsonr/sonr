package stream

import (
	"bytes"
	"context"
	"encoding/base64"
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
	Call         DataCallback
	Host         host.Host
	TransferFile sonrFile.TransferFile

	// Peer/Self Info
	PeerID peer.ID
	Self   *pb.Peer
	Peer   *pb.Peer

	// Stream Info
	id     string
	stream network.Stream
}

// ^ Start New Stream ^ //
func (dsc *DataStreamConn) Transfer(ctx context.Context, sf *sonrFile.SafeFile) {
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

	// Initialize Writer
	writer := msgio.NewWriter(stream)

	// @ Check Type
	if sf.Mime.Type == pb.MIME_image {
		// Get Metadata
		meta := sf.Metadata()

		// Start Routine
		log.Println("Starting Base64 Write Routine")
		go writeBase64ToStream(writer, meta)
	} else {
		// Get Metadata and Size
		meta := sf.Metadata()
		total := meta.Size

		// Start Routine
		log.Println("Starting Bytes Write Routine")
		go writeBytesToStream(writer, meta, total)
	}
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
	reader := msgio.NewReader(stream)
	go dsc.readFile(reader)
}

// ^ read Data from Msgio ^ //
func (dsc *DataStreamConn) readFile(reader msgio.ReadCloser) {
	for i := 0; ; i++ {
		// @ Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			log.Fatalln(err)
			break
		}

		// @ Add Buffer Data to File, Check for Completed
		hasCompleted, progress, err := dsc.TransferFile.AddBuffer(buffer)
		if err != nil {
			log.Fatalln(err)
			break
		}

		// @ Check for Completed
		if hasCompleted {
			// Save The File
			metadata, err := dsc.TransferFile.Save(dsc.Peer)
			if err != nil {
				log.Fatalln(err)
				break
			}

			// Convert to Bytes
			bytes, err := proto.Marshal(metadata)
			if err != nil {
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

// ^ write file as Base64 in Msgio to Stream ^ //
func writeBase64ToStream(writer msgio.WriteCloser, meta *pb.Metadata) {
	// Initialize Buffer
	imgBuffer := new(bytes.Buffer)

	// @ Check Image type
	if meta.Mime.Subtype == "jpeg" {
		// Get JPEG Encoded Buffer
		err := sonrFile.EncodeJpegBuffer(imgBuffer, meta)
		if err != nil {
			log.Fatalln(err)
		}
	} else if meta.Mime.Subtype == "png" {
		// Get PNG Encoded Buffer
		err := sonrFile.EncodePngBuffer(imgBuffer, meta)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Encode Buffer to base 64
	imgBytes := imgBuffer.Bytes()
	data := base64.StdEncoding.EncodeToString(imgBytes)
	total := int32(len(data))

	// Iterate for Entire file as String
	for i, chunk := range sonrFile.ChunkBase64(data, B64ChunkSize) {
		log.Println("Chunk Number: ", i)
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
			log.Fatalln(err)
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// ^ write file as Bytes in Msgio to Stream ^ //
func writeBytesToStream(writer msgio.WriteCloser, meta *pb.Metadata, total int32) {
	// Open File
	file, err := os.Open(meta.Path)
	if err != nil {
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
			log.Fatalln(err)
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
