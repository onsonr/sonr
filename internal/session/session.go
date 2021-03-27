package session

import (
	md "github.com/sonr-io/core/internal/models"
)

type Session struct {
	sender   *md.Peer
	receiver *md.Peer

	incoming *IncomingFile
	outgoing *OutgoingFile
}

func NewOutSession(p *md.Peer, o *OutgoingFile) *Session {
	return &Session{
		sender:   p,
		outgoing: o,
	}
}

func NewInSession(p *md.Peer, s *md.Peer, i *IncomingFile) *Session {
	return &Session{
		sender:   s,
		receiver: p,
		incoming: i,
	}
}
