package models

import (
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	msg "github.com/libp2p/go-msgio"
)

// ** ─── CALLBACK MANAGEMENT ────────────────────────────────────────────────────────
type OnConnected func(r *ConnectionResponse)
type HTTPHandler func(http.ResponseWriter, *http.Request)
type SetStatus func(s Status)
type OnProtobuf func([]byte)
type OnError func(err *SonrError)
type Callback struct {
	OnEvent    OnProtobuf
	OnRequest  OnProtobuf
	OnResponse OnProtobuf
	OnError    OnError
	SetStatus  SetStatus
}

// ** ─── State MANAGEMENT ────────────────────────────────────────────────────────
type state struct {
	flag uint64
	chn  chan bool
}

var (
	instance *state
	once     sync.Once
)

func GetState() *state {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &state{chn: chn}
	})

	return instance
}

// Checks rather to wait or does not need
func (c *state) NeedsWait() {
	<-c.chn
}

// Says all of goroutines to resume execution
func (c *state) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Says all of goroutines to pause execution
func (c *state) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.chn = make(chan bool)
	}
}

// ** ─── Transfer MANAGEMENT ────────────────────────────────────────────────────────
// Returns Transfer for URLLink
func (u *URLLink) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_Url{
			Url: u,
		},
	}
}

// Returns URLLink as Transfer_Url Data
func (u *URLLink) ToData() *Transfer_Url {
	return &Transfer_Url{
		Url: u,
	}
}

// Returns Transfer for SFile
func (f *SFile) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_File{
			File: f,
		},
	}
}

// Returns SFile as Transfer_File Data
func (f *SFile) ToData() *Transfer_File {
	return &Transfer_File{
		File: f,
	}
}

// Returns Transfer for Contact
func (c *Contact) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_Contact{
			Contact: c,
		},
	}
}

// Returns Contact as Transfer_Contact Data
func (c *Contact) ToData() *Transfer_Contact {
	return &Transfer_Contact{
		Contact: c,
	}
}

// ** ─── MIME MANAGEMENT ────────────────────────────────────────────────────────
// Method adjusts extension for JPEG
func (m *MIME) Ext() string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// Checks if Mime is Audio
func (m *MIME) IsAudio() bool {
	return m.Type == MIME_AUDIO
}

// Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.Type == MIME_AUDIO || m.Type == MIME_IMAGE || m.Type == MIME_VIDEO
}

// Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// ** ─── SFile MANAGEMENT ────────────────────────────────────────────────────────
// Checks if File contains single item
func (f *SFile) IsSingle() bool {
	return len(f.Items) == 1
}

// Checks if Single File is Media
func (f *SFile) IsMedia() bool {
	return f.Payload == Payload_MEDIA || (f.IsSingle() && f.Single().Mime.IsMedia())
}

// Checks if File contains multiple items
func (f *SFile) IsMultiple() bool {
	return len(f.Items) > 1
}

// Returns Index of Item from File
func (f *SFile) IndexOf(item *SFile_Item) int {
	for i, v := range f.Items {
		if v == item {
			return i
		}
	}
	return -1
}

// Method Returns Single if Applicable
func (f *SFile) Single() *SFile_Item {
	if f.IsSingle() {
		return f.Items[0]
	} else {
		return nil
	}
}

// ** ─── Session MANAGEMENT ────────────────────────────────────────────────────────
type Session struct {
	// Inherited Properties
	file      *SFile
	owner     *Peer
	receiver  *Peer
	pid       protocol.ID
	user      *User
	direction CompleteEvent_Direction

	// Management
	call SessionHandler
}

type SessionHandler interface {
	OnCompleted(stream network.Stream, pid protocol.ID, completeEvent *CompleteEvent)
	OnProgress(buf []byte)
	OnError(err *SonrError)
}

// Prepare for Outgoing Session ^ //
func NewOutSession(u *User, req *InviteRequest, sh SessionHandler) *Session {
	return &Session{
		file:      req.GetFile(),
		owner:     req.GetFrom(),
		receiver:  req.GetTo(),
		pid:       req.ProtocolID(),
		user:      u,
		direction: CompleteEvent_OUTGOING,
		call:      sh,
	}
}

// Prepare for Incoming Session ^ //
func NewInSession(u *User, inv *InviteRequest, sh SessionHandler) *Session {
	// Return Session
	return &Session{
		file:      inv.GetFile(),
		owner:     inv.GetFrom(),
		receiver:  inv.GetTo(),
		pid:       inv.ProtocolID(),
		user:      u,
		direction: CompleteEvent_INCOMING,
		call:      sh,
	}
}

// Returns SFile as TransferCard given Receiver and Owner
func (s *Session) Event() *CompleteEvent {
	return &CompleteEvent{
		Direction: s.direction,
		Transfer: &Transfer{
			// SQL Properties
			Payload:  s.file.Payload,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    s.owner.GetProfile(),
			Receiver: s.receiver.GetProfile(),

			// Data Properties
			Data: s.file.ToData(),
		},
	}
}

// read buffers sent on stream and save to file ^ //
func (s *Session) ReadFromStream(stream network.Stream) {
	// Concurrent Function
	go func(rs msg.ReadCloser) {
		// Read All Files
		for i, m := range s.file.Items {
			r := m.NewReader(s.user.Device, i, int(s.file.GetCount()), s.call)
			err := r.ReadFrom(rs)
			if err != nil {
				s.call.OnError(NewError(err, ErrorEvent_INCOMING))
			}
		}
		// Set Status
		s.call.OnCompleted(stream, s.pid, s.Event())
	}(msg.NewReader(stream))
}

// write file as Base64 in Msgio to Stream ^ //
func (s *Session) WriteToStream(stream network.Stream) {
	// Concurrent Function
	go func(ws msg.WriteCloser) {
		// Write All Files
		for i, m := range s.file.Items {
			w := m.NewWriter(s.user.Device, i, int(s.file.GetCount()), s.call)
			err := w.WriteTo(ws)
			if err != nil {
				s.call.OnError(NewError(err, ErrorEvent_OUTGOING))
			}
		}
		// Handle Complete
		s.call.OnCompleted(stream, s.pid, s.Event())
	}(msg.NewWriter(stream))
}
