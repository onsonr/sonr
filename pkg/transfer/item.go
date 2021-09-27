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
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// ** ────────────────────────────────────────────────────────────────
// ** ─── FileItem Reader ────────────────────────────────────────────
// ** ────────────────────────────────────────────────────────────────
type ItemReader interface {
	Progress() []byte
	ReadFrom(reader msgio.ReadCloser) error
}
type itemReader struct {
	ItemReader
	docsDir string
	emitter *emitter.Emitter
	mutex   sync.Mutex
	item    *common.FileItem
	index   int
	size    int
	total   int
}

func NewReader(pi *common.Payload_Item, index int, total int, docsDir string, em *emitter.Emitter) ItemReader {
	// Return Reader
	return &itemReader{
		item:    pi.GetFile(),
		size:    0,
		emitter: em,
		index:   index,
		total:   total,
		docsDir: docsDir,
	}
}

// Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress() []byte {
	// Calculate Tracking
	currentProgress := float32(p.size) / float32(p.item.Size)

	// Create Update
	update := &common.ProgressEvent{
		Progress: float64(currentProgress),
		Current:  int32(p.index),
		Total:    int32(p.total),
	}

	// Marshal and Return
	buf, err := proto.Marshal(update)
	if err != nil {
		return nil
	}
	return buf
}

func (ir *itemReader) ReadFrom(reader msgio.ReadCloser) error {
	// Return Created File
	f, err := os.Create(filepath.Join(ir.docsDir, ir.item.Name))
	if err != nil {
		return err
	}

	// Route Data from Stream
	for i := 0; ; i++ {
		// Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			return err
		}

		buf, err := decodeChunk(buffer)
		if err != nil {
			return err
		}

		ir.mutex.Lock()
		n, err := f.Write(buf.Data)
		if err != nil {
			return err
		}
		ir.size = ir.size + n
		ir.mutex.Unlock()

		if (i % 10) == 0 {
			//ir.emitter.Emit(emitter.EMIT_PROGRESS_EVENT, ir.Progress())
		}
	}
}

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

// ** ────────────────────────────────────────────────────────────────
// ** ─── FileItem Writer ────────────────────────────────────────────
// ** ────────────────────────────────────────────────────────────────
type ItemWriter interface {
	Progress() []byte
	WriteTo(writer msgio.WriteCloser) error
}

type itemWriter struct {
	ItemWriter
	docsDir string
	emitter *emitter.Emitter
	item    *common.FileItem
	index   int
	size    int
	total   int
}

func NewWriter(pi *common.Payload_Item, index int, total int, docsDir string, em *emitter.Emitter) ItemWriter {
	return &itemWriter{
		item:    pi.GetFile(),
		size:    0,
		emitter: em,
		index:   index,
		total:   total,
		docsDir: docsDir,
	}
}

// Returns Progress of File, Given the written number of bytes
func (p *itemWriter) Progress() []byte {
	// Create Update
	update := &common.ProgressEvent{
		Progress: float64(p.size) / float64(p.item.Size),
		Current:  int32(p.index),
		Total:    int32(p.total),
	}

	// Marshal and Return
	buf, err := proto.Marshal(update)
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
	for {
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
	}
	return nil
}

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
