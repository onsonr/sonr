package transfer

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	sync "sync"

	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/config"

	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// itemReader is a Reader for a FileItem
type itemReader struct {
	docsDir string
	emitter *state.Emitter
	mutex   sync.Mutex
	item    *common.FileItem
	index   int
	count   int
	size    int64
}

// NewReader Returns a new Reader for the given FileItem
func NewReader(pi *common.Payload_Item, index int, count int, docsDir string, em *state.Emitter) *itemReader {
	// Return Reader
	return &itemReader{
		item:    pi.GetFile(),
		size:    pi.GetSize(),
		emitter: em,
		index:   index,
		count:   count,
		docsDir: docsDir,
	}
}

// Progress Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress(i int64) []byte {
	// Marshal and Return
	buf, err := proto.Marshal(&common.ProgressEvent{
		Progress: (float64(i) / float64(p.size)),
		Current:  int32(p.index),
		Total:    int32(p.count),
	})
	if err != nil {
		return nil
	}
	return buf
}

// ReadFrom Reads from the given Reader and writes to the given File
func (ir *itemReader) ReadFrom(reader msgio.ReadCloser) error {
	// Return Created File
	f, err := os.Create(filepath.Join(ir.docsDir, ir.item.Name))
	if err != nil {
		return err
	}

	// Route Data from Stream
	for i := int64(0); i <= ir.size; {
		// Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			return err
		}

		// Decode Chunk
		buf, err := decodeChunk(buffer)
		if err != nil {
			return err
		}

		// Write to File, and Update Progress
		ir.mutex.Lock()
		n, err := f.Write(buf.Data)
		if err != nil {
			return err
		}
		i += int64(n)
		ir.mutex.Unlock()

		// Emit Progress
		if (i % 10) == 0 {
			ir.emitter.Emit(Event_PROGRESS, ir.Progress(i))
		}
	}
	return nil
}

// decodeChunk Decodes a Chunk from a Message
func decodeChunk(buf []byte) (config.Chunk, error) {
	// Decode Chunk
	chunk := &common.Chunk{}
	err := proto.Unmarshal(buf, chunk)
	if err != nil {
		return config.Chunk{}, err
	}

	// Convert to Chunk
	c := config.Chunk{
		Offset:      int(chunk.Offset),
		Length:      int(chunk.Length),
		Data:        chunk.Data,
		Fingerprint: uint64(chunk.Fingerprint),
	}
	return c, nil
}

// itemWriter is a Writer for FileItems
type itemWriter struct {
	docsDir string
	emitter *state.Emitter
	item    *common.FileItem
	index   int
	count   int
	size    int64
}

// NewWriter Returns a new Writer for the given FileItem
func NewWriter(pi *common.Payload_Item, index int, count int, docsDir string, em *state.Emitter) *itemWriter {
	return &itemWriter{
		item:    pi.GetFile(),
		size:    pi.GetSize(),
		emitter: em,
		index:   index,
		count:   count,
		docsDir: docsDir,
	}
}

// Returns Progress of File, Given the written number of bytes
func (p *itemWriter) Progress(i int64) []byte {
	// Marshal and Return
	buf, err := proto.Marshal(&common.ProgressEvent{
		Progress: (float64(i) / float64(p.size)),
		Current:  int32(p.index),
		Total:    int32(p.count),
	})
	if err != nil {
		return nil
	}
	return buf
}

// Write Item to Stream
func (iw *itemWriter) WriteTo(writer msgio.WriteCloser) error {
	// Print Item Info
	logger.Info("Current Item Info: ", zap.String("Path", iw.item.Path), zap.String("Name", iw.item.Name), zap.Int("Size", int(iw.item.Size)), zap.String("Mime", iw.item.GetMime().String()))

	// Open Os File
	f, err := os.Open(iw.item.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("Error to read Item, %s", err.Error()))
	}

	chunker, err := config.NewChunker(f, config.ChunkerOptions{})
	if err != nil {
		return err
	}

	// Loop through File
	for i := int64(0); i <= iw.size; {
		c, err := chunker.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Create Block Protobuf from Chunk
		data, err := encodeChunk(c)
		if err != nil {
			return err
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(data)
		if err != nil {
			return err
		}

		// Unexpected Error
		if err != nil && err != io.EOF {
			return err
		}

		// Update Progress
		i += int64(c.Length)

		// Emit Progress
		if (i % 10) == 0 {
			iw.emitter.Emit(Event_PROGRESS, iw.Progress(i))
		}
	}
	return nil
}

// encodeChunk Encodes a Chunk into a Protobuf
func encodeChunk(c config.Chunk) ([]byte, error) {
	// Create Block Protobuf from Chunk
	data, err := proto.Marshal(&common.Chunk{
		Offset:      int32(c.Offset),
		Length:      int32(c.Length),
		Data:        c.Data,
		Fingerprint: int64(c.Fingerprint),
	})

	if err != nil {
		return nil, err
	}
	return data, nil
}
