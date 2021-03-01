package transfer

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"

	msgio "github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

// ** Constants for Chunking Data ** //
const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 32000

type OutgoingFile struct {
	onTransmitted md.OnTransmitted
	processedFile *sf.ProcessedFile
}

// ^ Creates New Outgoing ^ //
func NewOutgoingFile(pf *sf.ProcessedFile, oT md.OnTransmitted) *OutgoingFile {
	return &OutgoingFile{
		onTransmitted: oT,
		processedFile: pf,
	}
}

// ^ write file as Base64 in Msgio to Stream ^ //
func (of *OutgoingFile) WriteBase64(writer msgio.WriteCloser, peer *md.Peer) {
	// Initialize Buffer and Encode File
	var base string
	pf := of.processedFile

	if pf.Payload == md.Payload_MEDIA {
		buffer := new(bytes.Buffer)

		if err := pf.EncodeFile(buffer); err != nil {
			log.Fatalln(err)
		}

		// Encode Buffer to base 64
		data := buffer.Bytes()
		base = base64.StdEncoding.EncodeToString(data)
	} else {
		data, err := ioutil.ReadFile(pf.Path)
		if err != nil {
			log.Fatalln(err)
		}
		base = base64.StdEncoding.EncodeToString(data)
	}

	// Set Total
	total := int32(len(base))

	// Iterate for Entire file as String
	for _, chunk := range ChunkBase64(base) {
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
		md.GetState().NeedsWait()
	}

	// Call Completed Sending
	of.onTransmitted(peer)
}

// ^ Helper: Chunks string based on B64ChunkSize ^ //
func ChunkBase64(s string) []string {
	chunkSize := B64ChunkSize
	ss := make([]string, 0, len(s)/chunkSize+1)
	for len(s) > 0 {
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		// Create Current Chunk String
		ss, s = append(ss, s[:chunkSize]), s[chunkSize:]
	}
	return ss
}
