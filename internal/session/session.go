package client

import (
	"bytes"
	"sync"

	md "github.com/sonr-io/core/pkg/models"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	msg "github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

const K_BUF_CHUNK = 32000
const K_B64_CHUNK = 31998 // Adjusted for Base64 -- has to be divisible by 3

type Session struct {
	// Inherited Properties
	mutex    sync.Mutex
	sender   *md.Peer
	receiver *md.Peer
	file     *md.SonrFile

	// Management
	callback md.NodeCallback
	device   *md.Device

	// Builders
	bytesBuilder *bytes.Buffer

	// Tracking
	currentIndex int
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *md.Peer, req *md.InviteRequest, tc md.NodeCallback) *Session {
	f := req.GetFile()
	return &Session{
		file:         f,
		sender:       p,
		receiver:     req.To,
		callback:     tc,
		currentIndex: 0,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *md.Peer, inv *md.AuthInvite, d *md.Device, c md.NodeCallback) *Session {
	f := inv.GetFile()
	s := &Session{
		file:         f,
		sender:       inv.From,
		receiver:     p,
		callback:     c,
		device:       d,
		currentIndex: 0,
		bytesBuilder: new(bytes.Buffer),
	}
	return s
}

// ^ Check file type and use corresponding method ^ //
func (s *Session) AddBuffer(curr int, buffer []byte) (bool, error) {
	// ** Lock/Unlock ** //
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Unmarshal Bytes into Proto
	chunk := &md.Chunk{}
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
		if met, p := s.file.ItemAtIndex(s.currentIndex).Progress(curr, n); met {
			s.callback.Progressed(p)
			return false, nil
		}
	}
	return true, nil
}

// ^ read buffers sent on stream and save to file ^ //
func (s *Session) ReadFromStream(stream network.Stream) {
	// Route Data from Stream
	go func(reader msg.ReadCloser, in *md.SonrFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				s.callback.Error(md.NewError(err, md.ErrorMessage_TRANSFER_CHUNK))
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := s.AddBuffer(i, buffer)
			if err != nil {
				s.callback.Error(md.NewError(err, md.ErrorMessage_TRANSFER_CHUNK))
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := s.Save(); err != nil {
					s.callback.Error(md.NewError(err, md.ErrorMessage_TRANSFER_END))
				}
				break
			}
			md.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), s.file)
}

// ^ Check file type and use corresponding method to save to Disk ^ //
func (s *Session) Save() error {
	// Sync file to Disk
	if err := s.device.SaveTransfer(s.file, s.currentIndex, s.bytesBuilder.Bytes()); err != nil {
		s.callback.Error(md.NewError(err, md.ErrorMessage_TRANSFER_END))
	}

	// Send Complete Callback
	s.callback.Received(s.file.CardIn(s.receiver, s.sender))
	return nil
}

// ^ write file as Base64 in Msgio to Stream ^ //
func WriteToStream(writer msgio.WriteCloser, s *Session) {
	// Get Item
	m := s.file.ItemAtIndex(s.currentIndex)

	// Write Item to Stream
	if err := m.WriteTo(writer, s.callback); err != nil {
		s.callback.Error(md.NewError(err, md.ErrorMessage_OUTGOING))
		return
	}

	// Call Completed Sending
	s.callback.Transmitted(s.file.CardOut(s.receiver, s.sender))
}
