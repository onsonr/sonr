package session

import (
	"bytes"
	"strings"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	st "github.com/sonr-io/core/pkg/state"
)

type Session struct {
	sender   *md.Peer
	receiver *md.Peer

	incoming *incomingFile
	outgoing *outgoingFile

	callback st.NodeCallback
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *md.Peer, req *md.InviteRequest, fs *sf.FileSystem, tc st.NodeCallback) *Session {
	o := newOutgoingFile(req, p)
	return &Session{
		sender:   p,
		receiver: req.To,
		outgoing: o,
		callback: tc,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *md.Peer, inv *md.AuthInvite, fs *sf.FileSystem, tc st.NodeCallback) *Session {
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
