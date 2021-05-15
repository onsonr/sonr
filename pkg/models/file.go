package models

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	msg "github.com/libp2p/go-msgio"
)

const K_CHUNK_SIZE = 4 * 1024

// ** ─── SONRFILE MANAGEMENT ────────────────────────────────────────────────────────
// Checks if File contains single item
func (f *SonrFile) IsSingle() bool {
	return len(f.Items) == 1
}

// Checks if Single File is Media
func (f *SonrFile) IsMedia() bool {
	return f.Payload == Payload_MEDIA
}

// Checks if File contains multiple items
func (f *SonrFile) IsMultiple() bool {
	return len(f.Items) > 1
}

// Returns All Item Readers for This File
func (f *SonrFile) Readers(d *Device) []ItemReadWriter {
	readers := make([]ItemReadWriter, len(f.Items))
	for i, m := range f.Items {
		// Set Item Path
		m.SetPath(d)

		// Create Reader
		readers[i] = &itemReadWriter{
			item: m,
			size: 0,
		}
	}
	return readers
}

// Returns All Item Writers for This File
func (f *SonrFile) Writers() []ItemReadWriter {
	writers := make([]ItemReadWriter, len(f.Items))
	for i, m := range f.Items {
		writers[i] = &itemReadWriter{
			item: m,
			size: 0,
		}
	}
	return writers
}

// ** ─── SonrFile_Item MANAGEMENT ────────────────────────────────────────────────────────
// Sets New Path for Item Provided Device
func (i *SonrFile_Item) SetPath(d *Device) {
	// Check for Media
	if i.Mime.IsMedia() {
		// Check for Desktop
		if d.IsDesktop() {
			i.Path = filepath.Join(d.FileSystem.GetDownloads(), i.Name)
		} else {
			i.Path = filepath.Join(d.FileSystem.GetTemporary(), i.Name)
		}
	} else {
		// Check for Desktop
		if d.IsDesktop() {
			i.Path = filepath.Join(d.FileSystem.GetDownloads(), i.Name)
		} else {
			i.Path = filepath.Join(d.FileSystem.GetDocuments(), i.Name)
		}
	}
}

// ** ─── Session MANAGEMENT ────────────────────────────────────────────────────────
// Interface for Transfer Session
type Session interface {
	ReadFromStream(stream network.Stream)
	WriteToStream(stream network.Stream)
}

// Underlying Implementation
type session struct {
	Session
	// Inherited Properties
	mutex sync.Mutex
	file  *SonrFile
	peer  *Peer
	user  *User

	// Management
	call NodeCallback
}

// Returns SonrFile as TransferCard given Receiver and Owner
func (s *session) Card() *TransferCard {
	return &TransferCard{
		// SQL Properties
		Payload:  s.file.Payload,
		Received: int32(time.Now().Unix()),

		// Owner Properties
		Owner:    s.user.Peer.GetProfile(),
		Receiver: s.peer.GetProfile(),

		// Data Properties
		File: s.file,
	}
}

// ^ read buffers sent on stream and save to file ^ //
func (s *session) ReadFromStream(stream network.Stream) {
	// Concurrent Function
	go func(rs msg.ReadCloser) {
		// Read All Files
		for _, r := range s.file.Readers(s.user.Device) {
			err := r.ReadFrom(rs)
			if err != nil {
				s.call.Error(NewError(err, ErrorMessage_INCOMING))
			}
		}

		// Close Stream and Callback
		stream.Close()
		s.call.Received(s.Card())
	}(msg.NewReader(stream))
}

// ^ Writes Files to Stream ^ //
func (s *session) WriteToStream(stream network.Stream) {
	// Concurrent Function
	go func(ws msg.WriteCloser) {
		// Write All Files
		for _, w := range s.file.Writers() {
			err := w.WriteTo(ws)
			if err != nil {
				s.call.Error(NewError(err, ErrorMessage_OUTGOING))
			}
			GetState().NeedsWait()
		}

		// Callback
		s.call.Transmitted(s.Card())
	}(msg.NewWriter(stream))
}
