package node

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	mg "github.com/libp2p/go-msgio"
	md "github.com/sonr-io/core/internal/models"
	se "github.com/sonr-io/core/internal/session"
	us "github.com/sonr-io/core/internal/user"
	pn "github.com/sonr-io/core/pkg/peer"
	"google.golang.org/protobuf/proto"
)

// ^ OnEvent: Specific Lobby Event ^
func (n *Node) OnEvent(e *md.LobbyEvent) {
	// Convert Message
	bytes, err := proto.Marshal(e)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Call Event
	n.call.Event(bytes)
}

// ^ OnRefresh: Topic has Updated ^
func (n *Node) OnRefresh(l *md.Lobby) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
		return
	}
	n.call.Refreshed(bytes)
}

// ^ OnInvite: User Received Invite ^
func (n *Node) OnInvite(invite []byte) {
	// Send Callback
	n.call.Invited(invite)
}

// ^ OnReply: Begins File Transfer when Accepted ^
func (n *Node) OnReply(id peer.ID, reply []byte, session *se.Session) {
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
func (n *Node) OnResponded(inv *md.AuthInvite, p *pn.PeerNode, fs *us.FileSystem) {
	n.session = se.NewInSession(p.Get(), inv, fs, n.call)
	n.Host.HandleStream(n.router.Transfer(), n.session.ReadFromStream)
}
