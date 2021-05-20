package models

import (
	"path/filepath"
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
	return f.Payload == Payload_MEDIA || (f.IsSingle() && f.Single().Mime.IsMedia())
}

// Checks if File contains multiple items
func (f *SonrFile) IsMultiple() bool {
	return len(f.Items) > 1
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

func (i *SonrFile_Item) NewReader(d *Device) ItemReader {
	// Return Reader
	return &itemReader{
		item:   i,
		device: d,
		size:   0,
	}
}

func (m *SonrFile_Item) NewWriter(d *Device) ItemWriter {
	return &itemWriter{
		item: m,
		size: 0,
	}
}

func (i *SonrFile_Item) SetPath(d *Device) string {
	// Set Path
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
	return i.Path
}

// ** ─── Session MANAGEMENT ────────────────────────────────────────────────────────
type Session struct {
	// Inherited Properties
	file *SonrFile
	peer *Peer
	user *User

	// Management
	call NodeCallback
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(u *User, req *AuthInvite, tc NodeCallback) *Session {
	return &Session{
		file: req.GetFile(),
		peer: req.GetTo(),
		user: u,
		call: tc,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(u *User, inv *AuthInvite, c NodeCallback) *Session {
	// Return Session
	return &Session{
		file: inv.GetFile(),
		peer: inv.GetFrom(),
		user: u,
		call: c,
	}
}

// Returns SonrFile as TransferCard given Receiver and Owner
func (s *Session) Card() *Transfer {
	return &Transfer{
		// SQL Properties
		Payload:  s.file.Payload,
		Received: int32(time.Now().Unix()),

		// Owner Properties
		Owner:    s.user.Peer.GetProfile(),
		Receiver: s.peer.GetProfile(),

		// Data Properties
		Data: s.file.GetTransfer().Data,
	}
}

// ^ read buffers sent on stream and save to file ^ //
func (s *Session) ReadFromStream(stream network.Stream) {
	// Concurrent Function
	go func(rs msg.ReadCloser) {
		// Read All Files
		for _, m := range s.file.Items {
			r := m.NewReader(s.user.Device)
			err := r.ReadFrom(rs)
			if err != nil {
				s.call.Error(NewError(err, ErrorMessage_INCOMING))
			}
		}
		s.call.Received(s.Card())
	}(msg.NewReader(stream))
}

// ^ write file as Base64 in Msgio to Stream ^ //
func (s *Session) WriteToStream(stream network.Stream) {
	// Concurrent Function
	go func(ws msg.WriteCloser) {
		// Write All Files
		for _, m := range s.file.Items {
			w := m.NewWriter(s.user.Device)
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
