package transfer

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"

	msgio "github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

// ** Constants for Chunking Data ** //
const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 32000

// ^ write file as Base64 in Msgio to Stream ^ //
func writeBase64ToStream(writer msgio.WriteCloser, onCompleted OnProtobuf, meta *md.Metadata, peer []byte) {
	// Initialize Buffer
	imgBuffer := new(bytes.Buffer)

	// @ Check Image type
	if meta.Mime.Subtype == "jpeg" {
		// Get JPEG Encoded Buffer
		err := EncodeJpegBuffer(imgBuffer, meta)
		if err != nil {
			log.Fatalln(err)
		}
	} else if meta.Mime.Subtype == "png" {
		// Get PNG Encoded Buffer
		err := EncodePngBuffer(imgBuffer, meta)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Encode Buffer to base 64
	imgBytes := imgBuffer.Bytes()
	data := base64.StdEncoding.EncodeToString(imgBytes)
	total := int32(len(data))

	// Iterate for Entire file as String
	for _, chunk := range ChunkBase64(data) {
		// Create Block Protobuf from Chunk
		chunk := md.Chunk{
			Size:  int32(len(chunk)),
			B64:   chunk,
			Total: total,
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
		lifecycle.GetState().NeedsWait()
	}

	// Call Completed Sending
	onCompleted(peer)
}

// ^ write file as Bytes in Msgio to Stream ^ //
func writeBytesToStream(writer msgio.WriteCloser, onCompleted OnProtobuf, meta *md.Metadata, peer []byte) {
	// Open File
	file, err := os.Open(meta.Path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Set Chunk Variables
	ps := make([]byte, BufferChunkSize)
	total := meta.Size

	// Iterate file
	for {
		// Read Bytes
		bytesread, err := file.Read(ps)

		// Check for Error
		if err != nil {
			// Non EOF Error
			if err != io.EOF {
				log.Println(err)
			}
			// File Complete
			break
		}

		// Create Block Protobuf from Chunk
		chunk := md.Chunk{
			Size:   int32(len(ps[:bytesread])),
			Buffer: ps[:bytesread],
			Total:  total,
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
		lifecycle.GetState().NeedsWait()
	}

	// Call Completed Sending
	onCompleted(peer)
}
