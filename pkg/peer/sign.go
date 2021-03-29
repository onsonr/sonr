package peer

import (
	"time"

	md "github.com/sonr-io/core/internal/models"
)

// ^ SignMessage Creates Lobby Event with Message ^
func (pn *PeerNode) SignMessage(m string, to string) *md.LobbyEvent {
	return &md.LobbyEvent{
		Event:   md.LobbyEvent_MESSAGE,
		From:    pn.Get(),
		Id:      pn.peer.Id.Peer,
		Message: m,
		To:      to,
	}
}

// ^ SignUpdate Creates Lobby Event with Peer Data ^
func (pn *PeerNode) SignUpdate() *md.LobbyEvent {
	return &md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		From:  pn.Get(),
		Id:    pn.peer.Id.Peer,
	}
}

// ^ SignReply Creates AuthReply ^
func (pn *PeerNode) SignReply() *md.AuthReply {
	p := pn.Get()
	return &md.AuthReply{
		From: p,
		Type: md.AuthReply_Contact,
		Card: &md.TransferCard{
			// SQL Properties
			Payload:  md.Payload_CONTACT,
			Received: int32(time.Now().Unix()),
			Preview:  p.Profile.Picture,
			Platform: p.Platform,

			// Transfer Properties
			Status: md.TransferCard_REPLY,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,
		},
	}
}

// ^ SignReply Creates AuthReply with Contact  ^
func (pn *PeerNode) SignReplyWithContact(c *md.Contact) *md.AuthReply {
	p := pn.Get()
	return &md.AuthReply{
		From: p,
		Type: md.AuthReply_Contact,
		Card: &md.TransferCard{
			// SQL Properties
			Payload:  md.Payload_CONTACT,
			Received: int32(time.Now().Unix()),
			Preview:  p.Profile.Picture,
			Platform: p.Platform,

			// Transfer Properties
			Status: md.TransferCard_REPLY,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,

			// Data Properties
			Contact: c,
		},
	}
}
