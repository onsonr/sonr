package stream

import (
	"context"
	"fmt"
	"io"
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

const BlockSize = 64000

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
	Call DataCallback
	Self *pb.Peer
	File sf.SonrFile
	Peer *pb.Peer

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
	go dsc.readBlock(mrw)
}

// ^ read Data from Msgio ^ //
func (dsc *DataStreamConn) readBlock(mrw msgio.ReadCloser) error {
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
			dsc.File.AddBlock(msg.Data)
			dsc.sendProgress(msg.Current, msg.Total)
		}

		// Save File on Buffer Complete
		if msg.Current == msg.Total {
			// Add Block to Buffer
			fmt.Println("Completed All Blocks, Save the File")
			dsc.File.AddBlock(msg.Data)

			// Save The File
			savePath := dsc.File.Save()

			// Create Completed Protobuf
			completedMessage := pb.CompletedMessage{
				From:     dsc.Peer,
				Metadata: dsc.File.Metadata,
				Path:     savePath,
				Received: int64(time.Now().Unix()),
			}

			// Convert to Bytes
			bytes, err := proto.Marshal(&completedMessage)
			if err != nil {
				dsc.Call.Error(err, "Completed")
			}

			// Callback Completed
			dsc.Call.Completed(bytes)
			break
		}
	}
	return nil
}

func (dsc *DataStreamConn) sendProgress(current int64, total int64) {
	// Calculate Progress
	progress := float32(current) / float32(total)

	// Create Message
	progressMessage := pb.ProgressMessage{
		Current:  current,
		Total:    total,
		Progress: progress,
	}

	// Convert to bytes
	bytes, err := proto.Marshal(&progressMessage)
	if err != nil {
		dsc.Call.Error(err, "SendProgress")
	}

	// Send Callback
	dsc.Call.Progressed(bytes)
}

func (dsc *DataStreamConn) writeFile(sm *sf.SafeMeta) error {
	// Get Metadata
	meta := sm.Metadata()

	// Open File
	file, err := os.Open(meta.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Initialize Chunk Buffer
	buffer := make([]byte, BlockSize)
	i := 0

	// Iterate for Entire file
	for {
		i++
		// Read bytes at file
		bytesread, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		// Create Block Protobuf from Chunk
		block := pb.Block{
			Size:    int64(len(buffer)),
			Data:    buffer[:bytesread],
			Current: int64(i),
			Total:   meta.Blocks,
		}
		fmt.Println("Block: ", block.String())

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

		fmt.Println("bytes read: ", bytesread)
	}
	return nil
}
