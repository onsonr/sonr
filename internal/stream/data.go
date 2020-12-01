package stream

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"sync"
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
type OnProgressed func(float32)
type OnComplete func(data []byte)

const ChunkSize = 32000

// Struct to Implement Node Callback Methods
type DataCallback struct {
	Progressed OnProgressed
	Completed  OnComplete
	Error      OnError
}

// Struct defines a Chunk of Bytes of File
type Chunk struct {
	Size    int64
	Offset  int64
	Data    []byte
	Current int64
	Total   int64
}

// ^ Struct: Holds/Handles Stream for Authentication  ^ //
type DataStreamConn struct {
	Call     DataCallback
	Self     *pb.Peer
	SavePath string
	Peer     *pb.Peer

	id     string
	data   *pb.Metadata
	stream network.Stream
	writer msgio.Writer
}

// ^ Start New Stream ^ //
func (dsc *DataStreamConn) Transfer(ctx context.Context, h host.Host, id peer.ID, r *pb.Peer, sm *sf.SafeMeta) error {
	// Create New Auth Stream
	stream, err := h.NewStream(ctx, id, protocol.ID("/sonr/data"))
	if err != nil {
		return err
	}

	// Set Stream
	dsc.stream = stream
	dsc.id = stream.ID()
	dsc.Peer = r
	dsc.writer = msgio.NewWriter(dsc.stream)

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	go dsc.writeFile(sm)
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
	go dsc.readMessages(mrw)
}

// ^ read Data from Msgio ^ //
func (dsc *DataStreamConn) readMessages(mrw msgio.ReadCloser) error {
	var builder strings.Builder
	for {
		// Read Length Fixed Bytes
		buffer, err := mrw.ReadMsg()
		if err != nil {
			dsc.Call.Error(err, "read")
			return err
		}
		// Unmarshal Bytes into Proto
		msg := pb.Block{}
		err = proto.Unmarshal(buffer, &msg)
		if err != nil {
			dsc.Call.Error(err, "read")
			return err
		}

		if msg.Current < msg.Total {
			// Add Block to Buffer
			_, err := builder.WriteString(msg.Data)
			if err != nil {
				fmt.Println(err)
			}
			dsc.calcProgress(msg.Current, msg.Total)
		}

		// Save File on Buffer Complete
		if msg.Current == msg.Total {
			// Add Block to Buffer
			_, err := builder.WriteString(msg.Data)
			if err != nil {
				fmt.Println(err)
			}

			// Save The File
			err = dsc.saveFile(builder.String())
			if err != nil {
				dsc.Call.Error(err, "Save")
			}
			break
		}
	}
	return nil
}

func (dsc *DataStreamConn) calcProgress(current int32, total int32) {
	// Adjust Progress to Send on 100 Intervals
	if total > 100 {
		// Check if interval has been met
		updateInterval := total / 100
		remainder := current % updateInterval
		if remainder == 0 {
				// Calculate Progress
	percent := float32(current) / float32(total)

	// Send Callback
	dsc.Call.Progressed(percent)
		}
	} else {
			// Calculate Progress
	percent := float32(current) / float32(total)

	// Send Callback
	dsc.Call.Progressed(percent)
	}
}

func (dsc *DataStreamConn) sendProgress(current int32, total int32) {

}

func (dsc *DataStreamConn) saveFile(data string) error {
	// Get Bytes from base64
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal("error:", err)
	}

	// Create Delay
	time.After(time.Millisecond * 500)

	// Create File at Path
	f, err := os.Create(dsc.SavePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Write Bytes to to file
	if _, err := f.Write(bytes); err != nil {
		log.Fatalln(err)
	}
	if err := f.Sync(); err != nil {
		log.Fatalln(err)
	}

	// Create Delay
	time.After(time.Millisecond * 500)

	// Get Metadata
	metadata, err := sf.GetMetadata(dsc.SavePath, dsc.Peer)
	if err != nil {
		dsc.Call.Error(err, "Save")
	}

	// Create Delay
	time.After(time.Millisecond * 500)

	// Convert to Bytes
	bytes, err = proto.Marshal(metadata)
	if err != nil {
		dsc.Call.Error(err, "Completed")
	}

	// Callback Completed
	dsc.Call.Completed(bytes)

	// Return Block
	return nil
}

func (dsc *DataStreamConn) writeFile(sm *sf.SafeMeta) error {
	// Get Metadata
	var wg sync.WaitGroup
	meta := sm.Metadata()

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

	// Create Delay
	time.After(time.Millisecond * 500)

	// Iterate for Entire file
	total := len(splitString(b64, ChunkSize))
	for i, chunk := range splitString(b64, ChunkSize) {
		wg.Add(1)
		dsc.writeBlock(&wg, chunk, i, total)
	}
	wg.Wait()
	return nil
}

func (dsc *DataStreamConn) writeBlock(wg *sync.WaitGroup, chunk string, curr int, total int) {
	// Create Block Protobuf from Chunk
	block := pb.Block{
		Size:    int32(len(chunk)),
		Data:    chunk,
		Current: int32(curr),
		Total:   int32(total),
	}

	// Convert to bytes
	bytes, err := proto.Marshal(&block)
	if err != nil {
		dsc.Call.Error(err, "writeFileToStream")
	}

	// Write Message Bytes to Stream
	err = dsc.writer.WriteMsg(bytes)
	if err != nil {
		dsc.Call.Error(err, "writeFileToStream")
	}
	wg.Done()
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
