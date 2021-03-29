package client

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	mg "github.com/libp2p/go-msgio"
	se "github.com/sonr-io/core/internal/session"
	md "github.com/sonr-io/core/pkg/models"
	us "github.com/sonr-io/core/pkg/user"
	"google.golang.org/protobuf/proto"
)

// ^ OnEvent: Specific Lobby Event ^
func (n *Client) OnEvent(e *md.LobbyEvent) {
	// Convert Message
	bytes, err := proto.Marshal(e)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Call Event
	n.call.Event(bytes)
}

// ^ OnRefresh: Topic has Updated ^
func (n *Client) OnRefresh(l *md.Lobby) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
		return
	}
	n.call.Refreshed(bytes)
}

// ^ OnInvite: User Received Invite ^
func (n *Client) OnInvite(invite []byte) {
	// Send Callback
	n.call.Invited(invite)
}

// ^ OnReply: Begins File Transfer when Accepted ^
func (n *Client) OnReply(id peer.ID, reply []byte, session *se.Session) {
	// Call Responded
	n.call.Responded(reply)

	// AuthReply Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(reply, &resp)
	if err != nil {
		n.call.Error(err, "handleReply")
	}

	// Check for File Transfer
	if resp.Decision && resp.Type == md.AuthReply_Transfer {
		// Create New Auth Stream
		stream, err := n.Host.StartStream(id, n.router.Transfer())
		if err != nil {
			n.call.Error(err, "StartOutgoing")
			return
		}

		// Write to Stream on Session
		writer := mg.NewWriter(stream)
		go se.WriteToStream(writer, session)
	} else {
		n.session = nil
	}
}

// ^ OnResponded: Prepares for Incoming File Transfer when Accepted ^
func (n *Client) OnResponded(inv *md.AuthInvite, p *md.Peer, fs *us.FileSystem) {
	n.session = se.NewInSession(p, inv, fs, n.call)
	n.Host.HandleStream(n.router.Transfer(), n.session.ReadFromStream)
}
