package data

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	sync "sync"

	msg "github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/emitter"
	util "github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// Check if Payload is Album
func (p Payload) IsAlbum() bool {
	return p == Payload_ALBUM
}

// Check if Payload is Contact
func (p Payload) IsContact() bool {
	return p == Payload_CONTACT
}

// Check if Payload is File
func (p Payload) IsFile() bool {
	return p == Payload_FILE
}

// Check if Payload is Multiple Files
func (p Payload) IsFiles() bool {
	return p == Payload_FILES
}

// Check if Payload is Media
func (p Payload) IsMedia() bool {
	return p == Payload_MEDIA
}

// Check if Payload is URL
func (p Payload) IsURL() bool {
	return p == Payload_URL
}

// Check if Payload is for any Data Transfer
func (p Payload) IsTransfer() bool {
	return p == Payload_ALBUM || p == Payload_FILE || p == Payload_MEDIA || p == Payload_FILES
}

// Check if Payload is NOT for any Data Transfer
func (p Payload) IsNotTransfer() bool {
	return p == Payload_CONTACT || p == Payload_URL
}

// ** ─── SFile_Item MANAGEMENT ────────────────────────────────────────────────────────
func (i *SFile_Item) NewReader(d *Device, index int, total int, em *emitter.Emitter) ItemReader {
	// Return Reader
	return &itemReader{
		item:    i,
		device:  d,
		size:    0,
		emitter: em,
		index:   index,
		total:   total,
	}
}

func (m *SFile_Item) NewWriter(d *Device, index int, total int, em *emitter.Emitter) ItemWriter {
	return &itemWriter{
		item:    m,
		size:    0,
		device:  d,
		emitter: em,
		index:   index,
		total:   total,
	}
}

func (i *SFile_Item) SetPath(d *Device) string {
	// Set Path
	if i.Mime.IsMedia() {
		// Check for Desktop
		if d.IsDesktop() {
			i.Path = filepath.Join(d.FileSystem.GetDownloads().GetPath(), i.Name)
		} else {
			i.Path = filepath.Join(d.FileSystem.GetTemporary().GetPath(), i.Name)
		}
	} else {
		i.Path = filepath.Join(d.FileSystem.GetDownloads().GetPath(), i.Name)
	}
	return i.Path
}

// ** ─── Transfer (Reader) MANAGEMENT ────────────────────────────────────────────────────────
type ItemReader interface {
	Progress() []byte
	ReadFrom(reader msg.ReadCloser) error
}
type itemReader struct {
	ItemReader
	emitter *emitter.Emitter
	mutex   sync.Mutex
	item    *SFile_Item
	device  *Device
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
	buf, err := update.ToGeneric()
	if err != nil {
		return nil
	}
	return buf
}

func (ir *itemReader) ReadFrom(reader msg.ReadCloser) error {
	// Return Created File
	f, err := os.Create(ir.item.SetPath(ir.device))
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

		done, buf, err := decodeChunk(buffer)
		if err != nil {
			return err
		}

		if !done {
			ir.mutex.Lock()
			n, err := f.Write(buf)
			if err != nil {
				return err
			}
			ir.size = ir.size + n
			ir.mutex.Unlock()

			if (i % 10) == 0 {
				ir.emitter.Emit(emitter.EMIT_PROGRESS_EVENT, ir.Progress())
			}
		} else {
			// Flush File Data
			if err := f.Sync(); err != nil {
				return err
			}

			// Close File
			if err := f.Close(); err != nil {
				return err
			}
			return nil
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
	emitter *emitter.Emitter
	item    *SFile_Item
	device  *Device
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

func (iw *itemWriter) WriteTo(writer msg.WriteCloser) error {
	// Write Item to Stream
	// Open Os File
	f, err := os.Open(iw.item.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("Error to read Item, %s", err.Error()))
	}

	// Initialize Chunk Data
	r := bufio.NewReader(f)
	buf := make([]byte, 0, util.CHUNK_SIZE)

	// Loop through File
	for i := 0; ; i++ {
		// Initialize
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		// Bytes read is zero
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			return err
		}

		// Create Block Protobuf from Chunk
		data, err := encodeChunk(buf, false)
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

		if (i % 10) == 0 {
			iw.emitter.Emit(emitter.EMIT_PROGRESS_EVENT, iw.Progress())
		}
	}

	// Create Block Protobuf from Chunk
	data, err := encodeChunk(nil, true)
	if err != nil {
		return err
	}

	// Write Message Bytes to Stream
	err = writer.WriteMsg(data)
	if err != nil {
		return err
	}
	return nil
}

func decodeChunk(buf []byte) (bool, []byte, error) {
	// Unmarshal Bytes into Proto
	c := &Chunk{}
	err := proto.Unmarshal(buf, c)
	if err != nil {
		return false, nil, err
	}

	if c.IsComplete {
		return true, nil, nil
	} else {
		return false, c.Buffer, nil
	}
}

func encodeChunk(buffer []byte, completed bool) ([]byte, error) {
	if !completed {
		// Create Block Protobuf from Chunk
		chunk := &Chunk{
			Size:       int32(len(buffer)),
			Buffer:     buffer,
			IsComplete: false,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(chunk)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	} else {
		// Create Block Protobuf from Chunk
		chunk := &Chunk{
			IsComplete: true,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(chunk)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	}
}
