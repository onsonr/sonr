package client

import (
	"bytes"
	"sync"

	"encoding/base64"
	"log"

	md "github.com/sonr-io/core/pkg/models"
	us "github.com/sonr-io/core/pkg/user"

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
	callback     md.NodeCallback
	filesys      *us.FileSystem
	chunkCh      chan *md.Chunk
	progressCh   chan *md.Progress
	currentIndex int
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *md.Peer, req *md.InviteRequest, fs *us.FileSystem, tc md.NodeCallback) *Session {
	f := req.GetFile()
	return &Session{
		file:         f,
		sender:       p,
		receiver:     req.To,
		callback:     tc,
		filesys:      fs,
		currentIndex: 0,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *md.Peer, inv *md.AuthInvite, fs *us.FileSystem, c md.NodeCallback) *Session {
	s := &Session{
		file:         inv.GetFile(),
		sender:       inv.From,
		receiver:     p,
		callback:     c,
		filesys:      fs,
		chunkCh:      make(chan *md.Chunk),
		progressCh:   make(chan *md.Progress),
		currentIndex: 0,
	}

	// Handle Progress
	go func(pCh chan *md.Progress) {
		for {
			p := <-pCh
			if p.TotalComplete {
				s.callback.Received(s.file.Card(s.receiver, s.sender))
			} else {
				s.callback.Progressed(p.ItemProgress)
			}
		}
	}(s.progressCh)
	return s
}

// ^ Check file type and use corresponding method ^ //
func (s *Session) AddBuffer(curr int, buffer []byte) error {
	// ** Lock/Unlock ** //
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// @ Unmarshal Bytes into Proto
	chunk := &md.Chunk{}
	err := proto.Unmarshal(buffer, chunk)
	if err != nil {
		return err
	}

	// @ Initialize Vars if First Chunk
	if curr == 0 {
		// Retreive Item
		m, err := s.file.ItemAtIndex(s.currentIndex)
		if err != nil {
			return err
		}

		// Begin Item Progress
		s.file.AddItemAtIndex(s.currentIndex, s.filesys.GetPathForMetadata(m), chunk, s.chunkCh, s.progressCh)
	} else {
		s.chunkCh <- chunk
	}
	return nil
}

// ^ Helper: Chunks string based on B64ChunkSize ^ //
func ChunkBase64(s string) []string {
	chunkSize := K_B64_CHUNK
	ss := make([]string, 0, len(s)/chunkSize+1)
	for len(s) > 0 {
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		// Create Current Chunk String
		ss, s = append(ss, s[:chunkSize]), s[chunkSize:]
	}
	return ss
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
			err = s.AddBuffer(i, buffer)
			if err != nil {
				s.callback.Error(md.NewError(err, md.ErrorMessage_TRANSFER_CHUNK))
				break
			}
			md.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), s.file)
}

// ^ write file as Base64 in Msgio to Stream ^ //
func WriteToStream(writer msgio.WriteCloser, s *Session) {
	// Initialize Buffer and Encode File
	buffer := new(bytes.Buffer)
	if err := s.file.Encode(s.currentIndex, buffer); err != nil {
		log.Fatalln(err)
	}

	// Encode Buffer to base 64
	data := buffer.Bytes()
	base := base64.StdEncoding.EncodeToString(data)
	total := int32(len(base))

	// Iterate for Entire file as String
	for _, dat := range ChunkBase64(base) {
		// Create Block Protobuf from Chunk
		chunk := &md.Chunk{
			Size:  int32(len(dat)),
			Base:  dat,
			Total: total,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(chunk)
		if err != nil {
			log.Fatalln(err)
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			log.Fatalln(err)
		}
		md.GetState().NeedsWait()
	}

	// Call Completed Sending
	s.callback.Transmitted(s.file.Card(s.receiver, s.sender))
}
