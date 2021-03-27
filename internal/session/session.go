package session

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"

	"github.com/libp2p/go-libp2p-core/network"
	msg "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

type Session struct {
	sender   *md.Peer
	receiver *md.Peer

	incoming *incomingFile
	outgoing *outgoingFile

	callback dt.NodeCallback
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *md.Peer, req *md.InviteRequest, fs *sf.FileSystem, tc dt.NodeCallback) *Session {
	o := newOutgoingFile(req, p)
	return &Session{
		sender:   p,
		receiver: req.To,
		outgoing: o,
		callback: tc,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *md.Peer, inv *md.AuthInvite, fs *sf.FileSystem, tc dt.NodeCallback) *Session {
	return &Session{
		sender:   inv.From,
		receiver: p,
		callback: tc,
		incoming: &incomingFile{
			// Inherited Properties
			properties: inv.Card.Properties,
			payload:    inv.Payload,
			owner:      inv.From.Profile,
			preview:    inv.Card.Preview,
			fs:         fs,
			call:       tc,

			// Builders
			stringsBuilder: new(strings.Builder),
			bytesBuilder:   new(bytes.Buffer),
		},
	}
}

func (s *Session) OutgoingCard() *md.TransferCard {
	return s.outgoing.Card()
}

// ^ read buffers sent on stream and save to file ^ //
func (s *Session) ReadFromStream(stream network.Stream) {
	// Route Data from Stream
	go func(reader msg.ReadCloser, t *incomingFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				s.callback.Error(err, "HandleIncoming:ReadMsg")
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(i, buffer)
			if err != nil {
				s.callback.Error(err, "HandleIncoming:AddBuffer")
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := s.incoming.Save(); err != nil {
					s.callback.Error(err, "HandleIncoming:Save")
				}
				break
			}
			dt.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), s.incoming)
}

// ^ write file as Base64 in Msgio to Stream ^ //
func (s *Session) WriteToStream(stream network.Stream) {
	// Initialize Writer
	wr := msg.NewWriter(stream)

	// Begin Routine
	go func(writer msg.WriteCloser) {
		// Initialize Buffer and Encode File
		var base string
		if s.outgoing.Payload == md.Payload_MEDIA {
			buffer := new(bytes.Buffer)

			if err := s.outgoing.EncodeFile(buffer); err != nil {
				log.Fatalln(err)
			}

			// Encode Buffer to base 64
			data := buffer.Bytes()
			base = base64.StdEncoding.EncodeToString(data)
		} else {
			data, err := ioutil.ReadFile(s.outgoing.Path)
			if err != nil {
				log.Fatalln(err)
			}
			base = base64.StdEncoding.EncodeToString(data)
		}

		// Set Total
		total := int32(len(base))

		// Iterate for Entire file as String
		for _, dat := range ChunkBase64(base) {
			// Create Block Protobuf from Chunk
			chunk := md.Chunk64{
				Size:  int32(len(dat)),
				Data:  dat,
				Total: total,
			}

			// Convert to bytes
			bytes, err := proto.Marshal(&chunk)
			if err != nil {
				log.Fatalln(err)
			}

			// Write Message Bytes to Stream
			err = writer.WriteMsg(bytes)
			if err != nil {
				log.Fatalln(err)
			}
			dt.GetState().NeedsWait()
		}

		// Call Completed Sending
		s.callback.Transmitted(s.receiver)
	}(wr)
}
