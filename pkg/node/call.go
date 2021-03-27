package node

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	tr "github.com/sonr-io/core/pkg/transfer"
	"google.golang.org/protobuf/proto"
)

// ^ GetPeer: Returns Peer ^
func (n *Node) GetPeer() *md.Peer {
	return n.call.GetPeer()
}

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

// ^ OnReply: Begins File Transfer when Accepted ^
func (n *Node) OnReply(id peer.ID, p *md.Peer, cf *sf.FileItem, reply []byte) {
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
		stream, err := n.Host.NewStream(n.ctx, id, n.router.Transfer())
		if err != nil {
			n.call.Error(err, "StartOutgoing")
		}

		// Initialize Writer
		writer := msgio.NewWriter(stream)
		// Start Routine
		go tr.WriteBase64(writer, p, cf, n.call)
	}
}

// ^ OnInvite: User Received Invite ^
func (n *Node) OnInvite(invite []byte) {
	// Send Callback
	n.call.Invited(invite)
}
