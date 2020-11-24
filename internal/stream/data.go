package stream

import (
	"context"
	"fmt"
	"os"
	"sync"

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

const BlockSize = 16000

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

	id     string
	data   *pb.Metadata
	remote *pb.Peer
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
	dsc.remote = r
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
	go dsc.read(mrw)
}

// ^ read Data from Msgio ^ //
func (dsc *DataStreamConn) read(mrw msgio.ReadCloser) error {
	// Read Length Fixed Bytes
	lengthBytes, err := mrw.ReadMsg()
	if err != nil {
		dsc.Call.Error(err, "read")
	}

	// Unmarshal Bytes into Proto
	protoMsg := &pb.Block{}
	err = proto.Unmarshal(lengthBytes, protoMsg)
	if err != nil {
		dsc.Call.Error(err, "read")
	}

	dsc.handleBlock(protoMsg)
	mrw.ReleaseMsg(lengthBytes)
	return nil
}

// ^ Handle Received Message ^ //
func (dsc *DataStreamConn) handleBlock(msg *pb.Block) {
	// Verify Bytes Remaining
	fmt.Println("Current ", msg.Current, "Total ", msg.Total)

	if msg.Current < msg.Total {
		// Add Block to Buffer
		dsc.File.AddBlock(msg.Data)
		go dsc.sendProgress(msg.Current, msg.Total)
	}

	// Save File on Buffer Complete
	if msg.Current == msg.Total {
		// Add Block to Buffer
		dsc.File.AddBlock(msg.Data)
		fmt.Println("Completed All Blocks, Save the File")
	}
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

	// Number of go routines we need to spawn.
	concurrency := int(meta.Blocks)
	// buffer sizes that each of the go routine below should use. ReadAt
	// returns an error if the buffer size is larger than the bytes returned
	// from the file.
	chunksizes := make([]Chunk, concurrency)

	// All buffer sizes are the same in the normal case. Offsets depend on the
	// index. Second go routine should start at 100, for example, given our
	// buffer size of 100.
	for i := 0; i < concurrency; i++ {
		chunksizes[i].Size = BlockSize
		chunksizes[i].Offset = int64(BlockSize * i)
	}

	// check for any left over bytes. Add the residual number of bytes as the
	// the last chunk size.
	if remainder := meta.Size % BlockSize; remainder != 0 {
		c := Chunk{Size: remainder, Offset: int64(concurrency * BlockSize)}
		concurrency++
		chunksizes = append(chunksizes, c)
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(chunksizes []Chunk, i int) {
			defer wg.Done()

			// Create Chunk Data
			chunk := chunksizes[i]
			chunk.Current = int64(i)
			chunk.Total = int64(concurrency)
			chunk.Data = make([]byte, chunk.Size)
			bytesread, err := file.ReadAt(chunk.Data, chunk.Offset)
			if err != nil {
				fmt.Println(err)
			}

			// Create Block Protobuf from Chunk
			block := &pb.Block{
				Size:    chunk.Size,
				Offset:  chunk.Offset,
				Data:    chunk.Data,
				Current: chunk.Current,
				Total:   chunk.Total,
			}
			fmt.Println("Block: ", block.String())

			// Convert to bytes
			bytes, err := proto.Marshal(block)
			if err != nil {
				dsc.Call.Error(err, "writeFileToStream")
			}

			// Write Message Bytes to Stream
			if err := dsc.writer.WriteMsg(bytes); err != nil {
				dsc.Call.Error(err, "writeFileToStream")
			}

			fmt.Println("bytes read, string(bytestream): ", bytesread)
			fmt.Println("bytestream to string: ", string(chunk.Data))
		}(chunksizes, i)
	}

	wg.Wait()
	return nil
}