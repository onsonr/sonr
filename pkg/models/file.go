package models

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	sync "sync"
	"time"

	"github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

const K_CHUNK_SIZE = 4 * 1024

// ** ─── SONRFILE MANAGEMENT ────────────────────────────────────────────────────────
// Checks if File contains single item
func (f *SonrFile) IsSingle() bool {
	return len(f.Items) == 1
}

// Checks if Single File is Media
func (f *SonrFile) IsMedia() bool {
	return f.Payload == Payload_MEDIA || (f.IsSingle() && f.Single().Mime.IsMedia())
}

// Checks if File contains multiple items
func (f *SonrFile) IsMultiple() bool {
	return len(f.Items) > 1
}

// Checks if Given Index is Final Item
func (f *SonrFile) IsFinalIndex(i int) bool {
	return i == f.FinalIndex()
}

// Checks if Given Index is Final Item
func (f *SonrFile) IsFinalItem(i *SonrFile_Item) bool {
	return f.IndexOf(i) == f.FinalIndex()
}

// Returns SonrFile as TransferCard given Receiver and Owner
func (f *SonrFile) CardIn(receiver *Peer, owner *Peer) *TransferCard {
	// Create Card
	return &TransferCard{
		// SQL Properties
		Payload:  f.Payload,
		Received: int32(time.Now().Unix()),

		// Owner Properties
		Owner:    owner.GetProfile(),
		Receiver: receiver.GetProfile(),

		// Data Properties
		File: f,
	}
}

// Returns SonrFile as TransferCard given Receiver and Owner
func (f *SonrFile) CardOut(receiver *Peer, owner *Peer) *TransferCard {
	// Create Card
	return &TransferCard{
		// SQL Properties
		Payload: f.Payload,

		// Owner Properties
		Receiver: receiver.GetProfile(),
		Owner:    owner.GetProfile(),
		File:     f,
	}
}

// Method Returns Final Index of Metadata
func (f *SonrFile) FinalIndex() int {
	return len(f.Items) - 1
}

func (f *SonrFile) IndexOf(element *SonrFile_Item) int {
	for k, v := range f.Items {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func (f *SonrFile) NextItem(i *SonrFile_Item) *SonrFile_Item {
	idx := f.IndexOf(i)
	if f.IsFinalIndex(idx) {
		return nil
	} else {
		return f.Items[idx+1]
	}
}

// Method Returns Single if Applicable
func (f *SonrFile) Single() *SonrFile_Item {
	if f.IsSingle() {
		return f.Items[0]
	} else {
		return nil
	}
}

// ** ─── SONRFILE_Item MANAGEMENT ────────────────────────────────────────────────────────
type ItemFileReader struct {
	mutex sync.Mutex
	item  *SonrFile_Item
	file  *os.File
	size  int
}

func (m *SonrFile_Item) NewReader(d *Device) ItemReader {
	return &itemReader{
		item:   m,
		device: d,
		size:   0,
	}
}

func (m *SonrFile_Item) NewWriter(d *Device) ItemWriter {
	return &itemWriter{
		item:   m,
		size:   0,
	}
}

func (m *SonrFile_Item) Create(d *Device) (*ItemFileReader, error) {
	// Check for Media
	if m.Mime.IsMedia() {
		// Check for Desktop
		if d.IsDesktop() {
			m.Path = filepath.Join(d.FileSystem.GetDownloads(), m.Name)
		} else {
			m.Path = filepath.Join(d.FileSystem.GetTemporary(), m.Name)
		}
	} else {
		// Check for Desktop
		if d.IsDesktop() {
			m.Path = filepath.Join(d.FileSystem.GetDownloads(), m.Name)
		} else {
			m.Path = filepath.Join(d.FileSystem.GetDocuments(), m.Name)
		}
	}

	// Return Created File
	f, err := os.Create(m.Path)
	if err != nil {
		return nil, err
	}

	return &ItemFileReader{
		file: f,
		size: 0,
		item: m,
	}, nil
}

func (iw *ItemFileReader) Next(f *SonrFile, d *Device) (*ItemFileReader, error) {
	idx := f.IndexOf(iw.item)
	// Check if Last Item
	if f.IsFinalIndex(idx) {
		return nil, nil
	} else {
		// Get Next item
		newItem := f.NextItem(iw.item)

		// Create New Writer
		iw, err := newItem.Create(d)
		if err != nil {
			return nil, err
		}
		return iw, nil
	}
}

// Returns Progress of File, Given the written number of bytes
func (p *ItemFileReader) Progress() (bool, float32) {
	// Calculate Tracking
	progress := float32(p.size) / float32(p.item.Size)
	adjusted := int(progress)

	// Check Interval
	if adjusted&5 == 0 {
		return true, progress
	}
	return false, 0
}

func (p *ItemFileReader) Write(data []byte) (n int, err error) {
	// Unmarshal Bytes into Proto
	c := &Chunk{}
	err = proto.Unmarshal(data, c)
	if err != nil {
		return -1, err
	}

	// Check for Complete
	if !c.IsComplete {
		n, err = p.file.Write(c.GetBuffer())
		if err != nil {
			return -1, err
		}
		p.size += n
		return p.size, nil
	} else {
		err := p.file.Sync()
		if err != nil {
			return -1, err
		}
		err = p.file.Close()
		if err != nil {
			return -1, err
		}
		return 0, nil
	}
}

// Metadata Info returns: Total Bytes, Total Chunks, error
func (m *SonrFile_Item) WriteTo(writer msgio.WriteCloser, call NodeCallback) error {
	// @ Open Os File
	f, err := os.Open(m.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("Error to read [file=%v]: %v", m.Name, err.Error()))
	}

	// @ Initialize Chunk Data
	nBytes, nChunks := int32(0), int32(0)
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

		// Increment
		nChunks++
		nBytes += int32(len(buf))

		// Write to Stream
		streamBuffer(buf, false, writer, call)

		// Unexpected Error
		if err != nil && err != io.EOF {
			return err
		}
	}

	// Send Completed
	streamBuffer(nil, true, writer, call)
	f.Close()
	return nil
}

// Writes data to provided writer until completed is called
func streamBuffer(buf []byte, completed bool, writer msgio.WriteCloser, call NodeCallback) {
	if !completed {
		// Create Block Protobuf from Chunk
		chunk := &Chunk{
			Size:       int32(len(buf)),
			Buffer:     buf,
			IsComplete: false,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(chunk)
		if err != nil {
			call.Error(NewError(err, ErrorMessage_OUTGOING))
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			call.Error(NewError(err, ErrorMessage_OUTGOING))
		}
	} else {
		// Create Block Protobuf from Chunk
		chunk := &Chunk{
			IsComplete: true,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(chunk)
		if err != nil {
			call.Error(NewError(err, ErrorMessage_OUTGOING))
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			call.Error(NewError(err, ErrorMessage_OUTGOING))
		}
	}
}
