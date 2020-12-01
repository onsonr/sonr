package stream

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"os"
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

const ChunkSize = 16002 // Adjusted for Base64

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
func (dsc *DataStreamConn) Transfer(ctx context.Context, peerId peer.ID, peer *pb.Peer, sm *sf.SafeMeta) {
	// Create New Auth Stream
	stream, err := dsc.Host.NewStream(ctx, peerId, protocol.ID("/sonr/data"))
	if err != nil {
		dsc.Call.Error(err, "Transfer")
	}

	// Set Stream
	dsc.stream = stream
	dsc.id = stream.ID()
	dsc.Peer = peer
	dsc.PeerID = peerId

	writer := msgio.NewWriter(dsc.stream)

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go dsc.writeMessages(writer, sm)
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
func (dsc *DataStreamConn) readBlock(reader msgio.ReadCloser) error {
	for {
		// Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			dsc.Call.Error(err, "read")
			return err
		}
		// Unmarshal Bytes into Proto
		msg := pb.Chunk{}
		err = proto.Unmarshal(buffer, &msg)
		if err != nil {
			dsc.Call.Error(err, "read")
			return err
		}

		if msg.Current < msg.Total {
			// Add Block to Buffer
			dsc.File.AddBlock(msg.Data)

			// Send Receiver Progress Update
			go dsc.sendProgress(msg.Current, msg.Total)
		}

		// Save File on Buffer Complete
		if msg.Current == msg.Total {
			// Add Block to Buffer
			dsc.File.AddBlock(msg.Data)

			// Send Receiver Progress Update
			dsc.sendProgress(msg.Current, msg.Total)

			// Save The File
			savePath, err := dsc.File.Save()
			if err != nil {
				dsc.Call.Error(err, "Save")
			}

			// Get Metadata
			metadata, err := sf.GetMetadata(savePath, dsc.Peer)
			if err != nil {
				dsc.Call.Error(err, "Save")
			}

			// Create Delay
			time.After(time.Millisecond * 500)

			// Convert to Bytes
			bytes, err := proto.Marshal(metadata)
			if err != nil {
				dsc.Call.Error(err, "Completed")
			}

			// Callback Completed
			go dsc.Call.Completed(bytes)
			break
		}
	}
	return nil
}

func (dsc *DataStreamConn) sendProgress(current int32, total int32) {
	// Create Message
	progressMessage := pb.ProgressUpdate{
		Current: current,
		Total:   total,
		Percent: float32(current) / float32(total),
	}

	// Convert to bytes
	bytes, err := proto.Marshal(&progressMessage)
	if err != nil {
		dsc.Call.Error(err, "SendProgress")
	}

	// Send Callback
	dsc.Call.Progressed(bytes)
}

func (dsc *DataStreamConn) writeMessages(writer msgio.Writer, sf *sf.SafeMeta) error {
	// Get Metadata
	meta := sf.Metadata()
	imgBuffer := new(bytes.Buffer)

	// New File for ThumbNail
	file, err := os.Open(meta.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Encode as Jpeg into buffer
	err = jpeg.Encode(imgBuffer, img, nil)
	if err != nil {
		return err
	}

	b64 := base64.StdEncoding.EncodeToString(imgBuffer.Bytes())

	// Iterate for Entire file
	for i, chunk := range splitString(b64, ChunkSize) {
		// Create Block Protobuf from Chunk
		chunk := pb.Chunk{
			Size:    int32(len(chunk)),
			Data:    chunk,
			Current: int32(i),
			Total:   int32(len(b64)),
		}

		// Convert to bytes
		bytes, err := proto.Marshal(&chunk)
		if err != nil {
			dsc.Call.Error(err, "writeMessages")
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			dsc.Call.Error(err, "writeMessages")
		}
	}
	return nil
}

func splitString(s string, size int) []string {
	ss := make([]string, 0, len(s)/size+1)
	for len(s) > 0 {
		if len(s) < size {
			size = len(s)
		}
		ss, s = append(ss, s[:size]), s[size:]

	}
	return ss
}
