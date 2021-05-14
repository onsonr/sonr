package models

import (
	"bytes"
	"net/http"
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

// ** ─── URLLink MANAGEMENT ────────────────────────────────────────────────────────
// Creates New Link
func NewURLLink(url string) *URLLink {
	link := &URLLink{
		Url:         url,
		Initialized: false,
	}
	link.SetData()
	return link
}

// Sets URLLink Data
func (u *URLLink) SetData() {
	if !u.Initialized {
		// Create Request
		resp, err := http.Get(u.Url)
		if err != nil {
			return
		}

		// Get Info
		info, err := getPageInfoFromResponse(resp)
		if err != nil {
			return
		}

		// Set Link
		u.Initialized = true
		u.Title = info.Title
		u.Type = info.Type
		u.Site = info.Site
		u.SiteName = info.SiteName
		u.Description = info.Description
		u.Locale = info.Locale

		// Get Images
		if info.Images != nil {
			for _, v := range info.Images {
				u.Images = append(u.Images, &URLLink_OpenGraphImage{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Width:     int32(v.Width),
					Height:    int32(v.Height),
					Type:      v.Type,
				})
			}
		}

		// Get Videos
		if info.Videos != nil {
			for _, v := range info.Videos {
				u.Videos = append(u.Videos, &URLLink_OpenGraphVideo{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Width:     int32(v.Width),
					Height:    int32(v.Height),
					Type:      v.Type,
				})
			}
		}

		// Get Audios
		if info.Audios != nil {
			for _, v := range info.Videos {
				u.Audios = append(u.Audios, &URLLink_OpenGraphAudio{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Type:      v.Type,
				})
			}
		}

		// Get Twitter
		if info.Twitter != nil {
			u.Twitter = &URLLink_TwitterCard{
				Card:        info.Twitter.Card,
				Site:        info.Twitter.Site,
				SiteId:      info.Twitter.SiteId,
				Creator:     info.Twitter.Creator,
				CreatorId:   info.Twitter.CreatorId,
				Description: info.Twitter.Description,
				Title:       info.Twitter.Title,
				Image:       info.Twitter.Image,
				ImageAlt:    info.Twitter.ImageAlt,
				Url:         info.Twitter.Url,
				Player: &URLLink_TwitterCard_Player{
					Url:    info.Twitter.Player.Url,
					Width:  int32(info.Twitter.Player.Width),
					Height: int32(info.Twitter.Player.Height),
					Stream: info.Twitter.Player.Stream,
				},
				Iphone: &URLLink_TwitterCard_IPhone{
					Name: info.Twitter.IPhone.Name,
					Id:   info.Twitter.IPhone.Id,
					Url:  info.Twitter.IPhone.Url,
				},
				Ipad: &URLLink_TwitterCard_IPad{
					Name: info.Twitter.IPad.Name,
					Id:   info.Twitter.IPad.Id,
					Url:  info.Twitter.IPad.Url,
				},
				GooglePlay: &URLLink_TwitterCard_GooglePlay{
					Name: info.Twitter.Googleplay.Name,
					Id:   info.Twitter.Googleplay.Id,
					Url:  info.Twitter.Googleplay.Url,
				},
			}
		}
	}
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
	call    NodeCallback
	builder *bytes.Buffer

	// Tracking
	direction Direction
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(u *User, req *InviteRequest, tc NodeCallback) *Session {
	return &Session{
		file:      req.GetFile(),
		peer:      req.GetTo(),
		user:      u,
		call:      tc,
		direction: Direction_Outgoing,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(u *User, inv *AuthInvite, c NodeCallback) *Session {
	return &Session{
		file:      inv.GetFile(),
		peer:      inv.GetFrom(),
		user:      u,
		call:      c,
		builder:   new(bytes.Buffer),
		direction: Direction_Incoming,
	}
}

// ^ Check file type and use corresponding method ^ //
func (s *Session) AddBuffer(curr int, buffer []byte, index int) (bool, error) {
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
		_, err := s.builder.Write(chunk.Buffer)
		if err != nil {
			return true, err
		}

		// Check for Interval and Send Callback
		if met, p := s.file.ItemAtIndex(index).Progress(s.builder.Len()); met {
			s.call.Progressed(p)
		}

		// Not Complete
		return false, nil
	}
	return true, nil
}

// ^ read buffers sent on stream and save to file ^ //
func (s *Session) ReadFromStream(stream network.Stream) {
	// Route Data from Stream
	go func(reader msg.ReadCloser, index int) {
		for i := 0; ; i++ {
			// Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				s.call.Error(NewError(err, ErrorMessage_TRANSFER_CHUNK))
				break
			}

			// Unmarshal Bytes into Proto
			hasCompleted, err := s.AddBuffer(i, buffer, index)
			if err != nil {
				s.call.Error(NewError(err, ErrorMessage_TRANSFER_CHUNK))
				break
			}

			// Check if All Buffer Received to Save
			if hasCompleted {
				// Save file
				if done, newIndex := s.Save(index); done {
					break
				} else {
					index = newIndex
				}
			}
			GetState().NeedsWait()
		}
	}(msg.NewReader(stream), 0)
}

// ^ Check file type and use corresponding method to save to Disk ^ //
func (s *Session) Save(index int) (bool, int) {
	// Sync file to Disk
	if err := s.user.Device.SaveTransfer(s.file.ItemAtIndex(index), s.builder.Bytes()); err != nil {
		s.call.Error(NewError(err, ErrorMessage_TRANSFER_END))
	}

	// Reset Builder
	s.builder.Reset()

	// Check Completion
	if s.file.IsFinalIndex(index) {
		// Completed
		s.call.Received(s.file.CardIn(s.user.GetPeer(), s.peer))
		return true, index
	} else {
		// Next Item
		index += 1
		return false, index
	}
}

// ^ write file as Base64 in Msgio to Stream ^ //
func (s *Session) WriteToStream(writer msgio.WriteCloser) {
	// Write All Files
	for _, m := range s.file.Files {
		// Write Item to Stream
		if err := m.WriteTo(writer, s.call); err != nil {
			s.call.Error(NewError(err, ErrorMessage_OUTGOING))
			return
		}
		GetState().NeedsWait()
	}

	// Callback
	s.call.Transmitted(s.file.CardOut(s.peer, s.user.GetPeer()))
}
