package models

import (
	"bytes"
	"sync"
	"sync/atomic"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	msg "github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

const K_BUF_CHUNK = 32000
const K_B64_CHUNK = 31998 // Adjusted for Base64 -- has to be divisible by 3

// ** ─── CALLBACK MANAGEMENT ────────────────────────────────────────────────────────
// Define Function Types
var callback NodeCallback
type SetStatus func(s Status)
type OnProtobuf func([]byte)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *TransferCard)
type OnTransmitted func(data *TransferCard)
type OnError func(err *SonrError)
type NodeCallback struct {
	Invited     OnInvite
	Refreshed   OnProtobuf
	Event       OnProtobuf
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	Status      SetStatus
	Transmitted OnTransmitted
	Error       OnError
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

// Returns Transfer for SonrFile
func (f *SonrFile) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_File{
			File: f,
		},
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

// ** ─── Session MANAGEMENT ────────────────────────────────────────────────────────
type Session struct {
	// Inherited Properties
	mutex sync.Mutex
	file  *SonrFile
	peer  *Peer
	user  *User

	// Management
	callback NodeCallback

	// Builders
	bytesBuilder *bytes.Buffer

	// Tracking
	index     int
	direction Direction
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *User, req *InviteRequest, tc NodeCallback) *Session {
	return &Session{
		file:      req.GetFile(),
		peer:      req.GetTo(),
		user:      p,
		callback:  tc,
		index:     0,
		direction: Direction_Outgoing,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *User, inv *AuthInvite, c NodeCallback) *Session {
	return &Session{
		file:         inv.GetFile(),
		peer:         inv.GetFrom(),
		user:         p,
		callback:     c,
		index:        0,
		bytesBuilder: new(bytes.Buffer),
		direction:    Direction_Incoming,
	}
}

// ^ Check file type and use corresponding method ^ //
func (s *Session) AddBuffer(curr int, buffer []byte) (bool, error) {
	// ** Lock/Unlock ** //
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Unmarshal Bytes into Proto
	chunk := &Chunk{}
	err := proto.Unmarshal(buffer, chunk)
	if err != nil {
		return true, err
	}

	// Check for Complete
	if !chunk.IsComplete {
		// Add Buffer by File Type
		n, err := s.bytesBuilder.Write(chunk.Buffer)
		if err != nil {
			return true, err
		}

		// Check for Interval and Send Callback
		if met, p := s.file.ItemAtIndex(s.index).Progress(curr, n); met {
			s.callback.Progressed(p)
		}

		// Not Complete
		return false, nil
	}
	return true, nil
}

// ^ read buffers sent on stream and save to file ^ //
func (s *Session) ReadFromStream(stream network.Stream) {
	// Route Data from Stream
	go func(reader msg.ReadCloser) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				s.callback.Error(NewError(err, ErrorMessage_TRANSFER_CHUNK))
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := s.AddBuffer(i, buffer)
			if err != nil {
				s.callback.Error(NewError(err, ErrorMessage_TRANSFER_CHUNK))
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if done := s.Save(); done {
					break
				}
			}
			GetState().NeedsWait()
		}
	}(msg.NewReader(stream))
}

// ^ Check file type and use corresponding method to save to Disk ^ //
func (s *Session) Save() bool {
	// Sync file to Disk
	if err := s.user.Device.SaveTransfer(s.file, s.index, s.bytesBuilder.Bytes()); err != nil {
		s.callback.Error(NewError(err, ErrorMessage_TRANSFER_END))
	}

	// Send Complete Callback
	if s.index+1 == int(s.file.GetCount()) {
		s.callback.Received(s.file.CardIn(s.user.GetPeer(), s.peer))
		return true
	} else {
		s.index = s.index + 1
		return false
	}
}

// ^ write file as Base64 in Msgio to Stream ^ //
func WriteToStream(writer msgio.WriteCloser, s *Session) {
	// Write All Files
	for i := 0; i < int(s.file.GetCount()); i++ {
		// Get Item
		m := s.file.ItemAtIndex(i)

		// Write Item to Stream
		if err := m.WriteTo(writer, s.callback); err != nil {
			s.callback.Error(NewError(err, ErrorMessage_OUTGOING))
			return
		}
		GetState().NeedsWait()
	}

	// Callback
	s.callback.Transmitted(s.file.CardOut(s.peer, s.user.GetPeer()))
}
