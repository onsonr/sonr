package session

import (
	"bytes"
	"strings"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
)

type Session struct {
	sender   *md.Peer
	receiver *md.Peer

	incoming *incomingFile
	outgoing *outgoingFile

	callback md.NodeCallback
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *md.Peer, req *md.InviteRequest, fs *sf.FileSystem, tc md.NodeCallback) *Session {
	o := newOutgoingFile(req, p)
	return &Session{
		sender:   p,
		receiver: req.To,
		outgoing: o,
		callback: tc,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *md.Peer, inv *md.AuthInvite, fs *sf.FileSystem, tc md.NodeCallback) *Session {
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
