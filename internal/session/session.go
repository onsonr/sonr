package client

import (
	"bytes"
	"strings"
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
	payload  md.Payload
	preview  []byte

	// Management
	callback md.NodeCallback
	filesys  *us.FileSystem

	// Builders
	stringsBuilder *strings.Builder
	bytesBuilder   *bytes.Buffer

	// Tracking
	currentSize int
	interval    int
	totalChunks int
	totalSize   int
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *md.Peer, req *md.InviteRequest, fs *us.FileSystem, tc md.NodeCallback) *Session {
	f := req.GetFile()
	prev := f.Preview()
	return &Session{
		file:     f,
		sender:   p,
		receiver: req.To,
		callback: tc,
		filesys:  fs,
		preview:  prev,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *md.Peer, inv *md.AuthInvite, fs *us.FileSystem, tc md.NodeCallback) *Session {
	return &Session{
		file:           inv.Card.GetFile(),
		sender:         inv.From,
		receiver:       p,
		callback:       tc,
		filesys:        fs,
		payload:        inv.Payload,
		preview:        inv.Card.Preview,
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),
	}
}

// ^ Check file type and use corresponding method ^ //
func (s *Session) AddBuffer(curr int, buffer []byte) (bool, error) {
	// ** Lock/Unlock ** //
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// @ Unmarshal Bytes into Proto
	chunk := &md.Chunk64{}
	err := proto.Unmarshal(buffer, chunk)
	if err != nil {
		return true, err
	}

	// @ Initialize Vars if First Chunk
	if curr == 0 {
		// Calculate Tracking Data
		totalChunks := int(chunk.Total) / K_B64_CHUNK
		interval := totalChunks / 100

		// Set Data
		s.totalSize = int(chunk.Total)
		s.totalChunks = totalChunks
		s.interval = interval
	}

	// @ Add Buffer by File Type
	// Add Base64 Chunk to Buffer
	n, err := s.stringsBuilder.WriteString(chunk.Data)
	if err != nil {
		return true, err
	}

	// Update Tracking
	s.currentSize = s.currentSize + n

	// @ Check Completed
	if s.currentSize < s.totalSize {
		// Validate Interval
		if s.interval != 0 {
			// Check for Interval
			if curr%s.interval == 0 {
				// Send Callback
				s.callback.Progressed(float32(s.currentSize) / float32(s.totalSize))
			}
		}
		return false, nil
	} else {
		return true, nil
	}
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
				if err := s.Save(0); err != nil {
					s.callback.Error(md.NewError(err, md.ErrorMessage_TRANSFER_END))
				}
				break
			}
			md.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), s.file)
}

func (s *Session) OutgoingCard() *md.TransferCard {
	return s.file.ToCard(s.receiver, s.sender, s.preview)
}

// ^ Check file type and use corresponding method to save to Disk ^ //
func (s *Session) Save(index int) error {
	// Retreive Item
	meta, err := s.file.ItemAtIndex(index)
	if err != nil {
		return err
	}

	// Get Path
	path := s.filesys.GetPathForMetadata(meta)

	// Get Bytes from base64
	data, err := base64.StdEncoding.DecodeString(s.stringsBuilder.String())
	if err != nil {
		return err
	}

	// Sync file
	if err := s.file.SaveItem(path, data, index); err != nil {
		s.callback.Error(md.NewError(err, md.ErrorMessage_TRANSFER_END))
	}

	// Send Complete Callback
	s.callback.Received(s.file.ToCard(s.receiver, s.sender, s.preview))
	return nil
}

// ^ write file as Base64 in Msgio to Stream ^ //
func WriteToStream(writer msgio.WriteCloser, s *Session) {
	// Initialize Buffer and Encode File
	buffer := new(bytes.Buffer)
	if err := s.file.EncodeSingle(buffer); err != nil {
		log.Fatalln(err)
	}

	// Encode Buffer to base 64
	data := buffer.Bytes()
	base := base64.StdEncoding.EncodeToString(data)
	total := int32(len(base))

	// Iterate for Entire file as String
	for _, dat := range ChunkBase64(base) {
		// Create Block Protobuf from Chunk
		chunk := &md.Chunk64{
			Size:  int32(len(dat)),
			Data:  dat,
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
	s.callback.Transmitted(s.file.ToCard(s.receiver, s.sender, s.preview))
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
