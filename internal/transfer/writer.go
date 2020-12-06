package transfer

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"log"
	"os"

	md "github.com/sonr-io/core/internal/models"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	msgio "github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

// ** Constants for Chunking Data ** //
const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 32000

// ^ User has accepted, Begin Sending Transfer ^ //
func (pc *PeerConnection) SendFile(h host.Host, pid peer.ID) {
	// Create New Auth Stream
	stream, err := h.NewStream(context.Background(), pid, protocol.ID("/sonr/data/transfer"))
	if err != nil {
		onError(err, "Transfer")
		log.Fatalln(err)
	}

	// Initialize Writer
	writer := msgio.NewWriter(stream)
	meta := pc.SafeFile.GetMetadata()

	// @ Check Type
	if pc.SafeFile.Mime.Type == md.MIME_image {
		// Start Routine
		log.Println("Starting Base64 Write Routine")
		go writeBase64ToStream(writer, meta)
	} else {
		total := meta.Size

		// Start Routine
		log.Println("Starting Bytes Write Routine")
		go writeBytesToStream(writer, meta, total)
	}
}

// ^ write file as Base64 in Msgio to Stream ^ //
func writeBase64ToStream(writer msgio.WriteCloser, meta *md.Metadata) {
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
	for i, chunk := range ChunkBase64(data) {
		log.Println("Chunk Number: ", i)
		// Create Block Protobuf from Chunk
		chunk := md.Chunk{
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
func writeBytesToStream(writer msgio.WriteCloser, meta *md.Metadata, total int32) {
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
				log.Println(err)
			}
			// File Complete
			break
		}

		// Create Block Protobuf from Chunk
		chunk := md.Chunk{
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
