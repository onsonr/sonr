package common

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	sync "sync"

	"github.com/gabriel-vasile/mimetype"
	"go.uber.org/zap"

	msg "github.com/libp2p/go-msgio"

	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"google.golang.org/protobuf/proto"
)

// Ext Method adjusts extension for JPEG
func (m *MIME) Ext() string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// IsAudio Checks if Mime is Audio
func (m *MIME) IsAudio() bool {
	return m.Type == MIME_AUDIO
}

// IsMedia Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.Type == MIME_AUDIO || m.Type == MIME_IMAGE || m.Type == MIME_VIDEO
}

// IsImage Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// IsVideo Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// ToTransferItem Returns Transfer for FileItem
func (f *FileItem) ToTransferItem() *Payload_Item {
	return &Payload_Item{
		Data: &Payload_Item_File{
			File: f,
		},
		Type: Payload_Item_FILE,
	}
}

// NewTransferFileItem creates a new transfer file item
func NewTransferFileItem(path string) (*Payload_Item, error) {
	// Extracts File Infrom from path
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Get MIME Type and Set Proto Enum
	mtype, err := mimetype.DetectFile(path)
	if err != nil {
		return nil, err
	}

	// Create File Item
	fileItem := &FileItem{
		Path: path,
		Size: int32(fi.Size()),
		Name: fi.Name(),
		Mime: &MIME{
			Type:    MIME_Type(MIME_Type_value[strings.ToUpper(mtype.Parent().String())]),
			Value:   mtype.String(),
			Subtype: mtype.Extension(),
		},
	}

	// // // Check if File is Image
	// if fileItem.Mime.IsImage() {
	// 	// Create Thumbnail
	// 	name := filepath.Base(path)
	// 	dir := filepath.Dir(path)
	// 	outPath := filepath.Join(dir, name+"_thumb")
	// 	err := thumbnail.NewImage(path, outPath)
	// 	if err != nil {
	// 		return nil, err
	// 	} else {
	// 		fileItem.Thumb = &Thumbnail{
	// 			Path: outPath,
	// 		}
	// 	}
	// }

	// Returns transfer item
	return &Payload_Item{
		Type: Payload_Item_FILE,
		Data: &Payload_Item_File{
			File: fileItem,
		},
	}, nil
}

// ** ─── SFile_Item MANAGEMENT ────────────────────────────────────────────────────────
func NewReader(pi *Payload_Item, index int, total int, docsDir string, em *emitter.Emitter) ItemReader {
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

func NewWriter(pi *Payload_Item, index int, total int, docsDir string, em *emitter.Emitter) ItemWriter {
	return &itemWriter{
		item:    pi.GetFile(),
		size:    0,
		emitter: em,
		index:   index,
		total:   total,
		docsDir: docsDir,
	}
}

// ** ─── Transfer (Reader) MANAGEMENT ────────────────────────────────────────────────────────
type ItemReader interface {
	Progress() []byte
	ReadFrom(reader msg.ReadCloser) error
}
type itemReader struct {
	ItemReader
	docsDir string
	emitter *emitter.Emitter
	mutex   sync.Mutex
	item    *FileItem
	index   int
	size    int
	total   int
}

// Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress() []byte {
	// Calculate Tracking
	currentProgress := float32(p.size) / float32(p.item.Size)

	// Create Update
	update := &ProgressEvent{
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

func (ir *itemReader) ReadFrom(reader msg.ReadCloser) error {
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

// ** ─── Transfer (Writer) MANAGEMENT ────────────────────────────────────────────────────────
type ItemWriter interface {
	Progress() []byte
	WriteTo(writer msg.WriteCloser) error
}

type itemWriter struct {
	ItemWriter
	docsDir string
	emitter *emitter.Emitter
	item    *FileItem
	index   int
	size    int
	total   int
}

// Returns Progress of File, Given the written number of bytes
func (p *itemWriter) Progress() []byte {
	// Create Update
	update := &ProgressEvent{
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
func (iw *itemWriter) WriteTo(writer msg.WriteCloser) error {
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
	data, err := proto.Marshal(&Chunk{
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

func decodeChunk(buf []byte) (config.Chunk, error) {
	// Decode Chunk
	chunk := &Chunk{}
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
