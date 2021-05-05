package client

import (
	"github.com/libp2p/go-libp2p-core/peer"
	mg "github.com/libp2p/go-msgio"
	se "github.com/sonr-io/core/internal/session"
	md "github.com/sonr-io/core/pkg/models"
	us "github.com/sonr-io/core/pkg/user"
	"google.golang.org/protobuf/proto"
)

// ^ GetContact: Return User Contact Card for FlatContact ^
func (n *Client) GetContact() *md.Contact {
	return n.call.Contact()
}

// ^ OnEvent: Specific Lobby Event ^
func (n *Client) OnEvent(e *md.LobbyEvent) {
	// Convert Message
	bytes, err := proto.Marshal(e)
	if err != nil {
		n.call.Error(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}

	// Call Event
	n.call.Event(bytes)
}

// ^ OnRefresh: Topic has Updated ^
func (n *Client) OnRefresh(l *md.Lobby) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		n.call.Error(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}
	n.call.Refreshed(bytes)
}

// ^ OnInvite: User Received Invite ^
func (n *Client) OnInvite(data []byte) {
	// Send Callback
	n.call.Invited(data)
}

// ^ OnReply: Begins File Transfer when Accepted ^
func (n *Client) OnReply(id peer.ID, reply []byte, session *se.Session) {
	// Call Responded
	n.call.Responded(reply)

	// AuthReply Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(reply, &resp)
	if err != nil {
		n.call.Error(md.NewError(err, md.ErrorMessage_UNMARSHAL))
	}

	// Check if Status is not already transferring
	if n.call.GetStatus() != md.Status_INPROGRESS {
		// Check for File Transfer
		if resp.Decision && resp.Type == md.AuthReply_Transfer {
			// Update Status
			n.call.Status(md.Status_INPROGRESS)

			// Create New Auth Stream
			stream, err := n.Host.StartStream(id, n.router.Transfer(id))
			if err != nil {
				n.call.Error(md.NewError(err, md.ErrorMessage_HOST_STREAM))
				return
			}

			// Write to Stream on Session
			writer := mg.NewWriter(stream)
			go se.WriteToStream(writer, session)
		} else {
			n.session = nil
		}
	} else {
		n.call.Error(md.NewErrorWithType(md.ErrorMessage_TRANSFER_START))
	}
}

// ^ OnResponded: Prepares for Incoming File Transfer when Accepted ^
func (n *Client) OnResponded(inv *md.AuthInvite, p *md.Peer, fs *us.FileSystem) {
	n.session = se.NewInSession(p, inv, n.device, n.call)
	n.Host.HandleStream(n.router.Transfer(n.Host.ID), n.session.ReadFromStream)
}
