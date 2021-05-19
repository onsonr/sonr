package client

import (
	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

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
func (n *Client) refreshed(l *md.Lobby) {
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
func (n *Client) OnReply(id peer.ID, reply []byte) {
	// Call Responded
	n.call.Responded(reply)

	// AuthReply Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(reply, &resp)
	if err != nil {
		n.call.Error(md.NewError(err, md.ErrorMessage_UNMARSHAL))
	}

	// Check if Status is not already transferring
	if n.user.IsNotStatus(md.Status_INPROGRESS) {
		// Check for File Transfer
		if resp.HasAcceptedTransfer() {
			// Update Status
			n.call.Status(md.Status_INPROGRESS)

			// Create New Auth Stream
			stream, err := n.Host.StartStream(id, n.user.GetRouter().Transfer(id))
			if err != nil {
				n.call.Error(md.NewError(err, md.ErrorMessage_HOST_STREAM))
				return
			}

			// Write to Stream on Session
			n.session.WriteToStream(stream)
		}
	} else {
		n.call.Error(md.NewErrorWithType(md.ErrorMessage_TRANSFER_START))
	}
}

// ^ OnResponded: Prepares for Incoming File Transfer when Accepted ^
func (n *Client) OnResponded(inv *md.AuthInvite) {
	n.session = md.NewInSession(n.user, inv, n.call)
	n.Host.HandleStream(n.user.GetRouter().Transfer(n.Host.ID), n.session.ReadFromStream)
}
