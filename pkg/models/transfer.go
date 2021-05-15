package models

import (
	"bufio"
	"io"
	"os"
	sync "sync"

	msg "github.com/libp2p/go-msgio"
)

type ItemReadWriter interface {
	Progress() (bool, float32)
	ReadFrom(reader msg.ReadCloser) error
	WriteTo(writer msg.WriteCloser) error
}

type itemReadWriter struct {
	ItemReadWriter
	mutex  sync.Mutex
	item   *SonrFile_Item
	size   int
}

// Returns Progress of File, Given the written number of bytes
func (p *itemReadWriter) Progress() (bool, float32) {
	// Calculate Tracking
	progress := float32(p.size) / float32(p.item.Size)
	adjusted := int(progress)

	// Check Interval
	if adjusted&5 == 0 {
		return true, progress
	}
	return false, 0
}

// Read Item from Stream with Chunking
func (ir *itemReadWriter) ReadFrom(reader msg.ReadCloser) error {
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
		done, buf, err := DecodeChunk(buffer)
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
			ir.mutex.Unlock()
			return nil
		}
		GetState().NeedsWait()
	}
}

// Write Item to Stream with Chunking
func (iw *itemReadWriter) WriteTo(writer msg.WriteCloser) error {
	//  Open Os File
	f, err := os.Open(iw.item.Path)
	if err != nil {
		return err
	}

	// Initialize Chunk Data
	r := bufio.NewReader(f)
	buf := make([]byte, 0, K_CHUNK_SIZE)

	//  Loop through File
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
		data, err := EncodeChunk(buf, false)
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
	data, err := EncodeChunk(nil, true)
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
