package file

import (
	"fmt"
	"os"
	"sync"

	pb "github.com/sonr-io/core/internal/models"
)

// Struct defines a File Transferring
type TransferFile struct {
	Call   FileCallback
	Meta   *pb.Metadata
	mutex  sync.Mutex
	blocks []*pb.Block
}

// Struct defines a Chunk of Bytes of File
type Chunk struct {
	bufsize int64
	offset  int64
}

// ^ Create generates file metadata ^ //
func (tf *TransferFile) Generate() {
	// ** Lock ** //
	tf.mutex.Lock()

	// Initialize
	tf.blocks = make([]*pb.Block, tf.Meta.Blocks)

	// Open File
	file, err := os.Open(tf.Meta.Path)
	if err != nil {
		tf.Call.Error(err, "Generate")
	}
	defer file.Close()

	// Number of go routines we need to spawn.
	concurrency := int(tf.Meta.Blocks)
	// buffer sizes that each of the go routine below should use. ReadAt
	// returns an error if the buffer size is larger than the bytes returned
	// from the file.
	chunksizes := make([]Chunk, concurrency)

	// All buffer sizes are the same in the normal case. Offsets depend on the
	// index. Second go routine should start at 100, for example, given our
	// buffer size of 100.
	for i := 0; i < concurrency; i++ {
		chunksizes[i].bufsize = BlockSize
		chunksizes[i].offset = int64(BlockSize * i)
	}

	// check for any left over bytes. Add the residual number of bytes as the
	// the last chunk size.
	if remainder := tf.Meta.Size % BlockSize; remainder != 0 {
		c := Chunk{bufsize: remainder, offset: int64(concurrency * BlockSize)}
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
			buffer := make([]byte, chunk.bufsize)
			bytesread, err := file.ReadAt(buffer, chunk.offset)

			// Create Block
			block := &pb.Block{
				Data:    buffer,
				Size:    chunk.bufsize,
				Offset:  chunk.offset,
				Current: int64(i),
				Total:   int64(concurrency),
			}
			tf.blocks = append(tf.blocks, block)

			if err != nil {
				tf.Call.Error(err, "Generate")
			}

			fmt.Println("bytes read, string(bytestream): ", bytesread)
			fmt.Println("bytestream to string: ", string(buffer))
		}(chunksizes, i)
	}

	wg.Wait()

	fmt.Println("Blocks for file generated")
	tf.mutex.Unlock()
}

// ^ Safely returns Blocks depending on lock ^ //
func (tf *TransferFile) Blocks() []*pb.Block {
	// ** Lock File wait for access ** //
	tf.mutex.Lock()
	defer tf.mutex.Unlock()

	// @ 2. Return Value
	return tf.blocks
}
