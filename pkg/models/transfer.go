package models

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	sync "sync"

	msg "github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

type ItemReader interface {
	Progress() (bool, float32)
	ReadFrom(reader msg.ReadCloser) error
}
type itemReader struct {
	ItemReader
	mutex  sync.Mutex
	item   *SonrFile_Item
	device *Device
	size   int
}

// Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress() (bool, float32) {
	// Calculate Tracking
	progress := float32(p.size) / float32(p.item.Size)
	adjusted := int(progress)

	// Check Interval
	if adjusted&5 == 0 {
		return true, progress
	}
	return false, 0
}

func (ir *itemReader) ReadFrom(reader msg.ReadCloser) error {
	// Set Item Path
	ir.item.SetPath(ir.device)

	// Return Created File
	f, err := os.Create(ir.item.Path)
	if err != nil {
		return err
	}

	// Route Data from Stream
	for {
		// Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			return err
		}

		ir.mutex.Lock()
		done, buf, err := decodeChunk(buffer)
		if err != nil {
			return err
		}

		if !done {
			n, err := f.Write(buf)
			if err != nil {
				return err
			}
			ir.size = ir.size + n
			ir.mutex.Unlock()
		} else {
			// Flush File Data
			if err := f.Sync(); err != nil {
				return err
			}

			// Close File
			if err := f.Close(); err != nil {
				return err
			}
			ir.mutex.Unlock()
			return nil
		}
		GetState().NeedsWait()
	}
}

type ItemWriter interface {
	Progress() (bool, float32)
	WriteTo(writer msg.WriteCloser) error
}

type itemWriter struct {
	ItemWriter
	item *SonrFile_Item
	size int
}

// Returns Progress of File, Given the written number of bytes
func (p *itemWriter) Progress() (bool, float32) {
	// Calculate Tracking
	progress := float32(p.size) / float32(p.item.Size)
	adjusted := int(progress)

	// Check Interval
	if adjusted&5 == 0 {
		return true, progress
	}
	return false, 0
}

func (iw *itemWriter) WriteTo(writer msg.WriteCloser) error {
	// Write Item to Stream
	// @ Open Os File
	f, err := os.Open(iw.item.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("Error to read Item, %s", err.Error()))
	}

	// @ Initialize Chunk Data
	r := bufio.NewReader(f)
	buf := make([]byte, 0, K_CHUNK_SIZE)

	// @ Loop through File
	for {
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

	// Callback
	return nil
}

func decodeChunk(data []byte) (bool, []byte, error) {
	// Unmarshal Bytes into Proto
	c := &Chunk{}
	err := proto.Unmarshal(data, c)
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
