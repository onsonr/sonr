package models

import (
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

func (m *SonrFile_Item) NewReader(d *Device) ItemReader {
	return &itemReader{
		item:   m,
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

// ** ─── Session MANAGEMENT ────────────────────────────────────────────────────────
type Session struct {
	// Inherited Properties
	mutex sync.Mutex
	file  *SonrFile
	peer  *Peer
	user  *User

	// Management
	call NodeCallback
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(u *User, req *InviteRequest, tc NodeCallback) *Session {
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

		// Close Stream and Callback
		stream.Close()
		s.call.Received(s.file.CardIn(s.user.GetPeer(), s.peer))
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
		s.call.Transmitted(s.file.CardOut(s.peer, s.user.GetPeer()))
	}(msg.NewWriter(stream))
}
