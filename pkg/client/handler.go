package client

import (
	"net/http"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ OnLocalEvent: Local Lobby Event ^
func (n *Client) OnLocalEvent(e *md.LocalEvent) {
	// Convert Message
	bytes, err := proto.Marshal(e)
	if err != nil {
		n.call.Error(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}

	// Call Event
	n.call.LocalEvent(bytes)
}

// ^ OnRemoteEvent: Remote Lobby Event ^
func (n *Client) OnRemoteEvent(e *md.RemoteEvent) {
	// Convert Message
	bytes, err := proto.Marshal(e)
	if err != nil {
		n.call.Error(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}

	// Call Event
	n.call.RemoteEvent(bytes)
}

// ^ OnLinkRequest: When Device is Initiating Link
func (n *Client) OnLinkRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi!"))
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
			stream, err := n.Host.StartStream(id, n.user.GetRouter().LocalTransfer(id))
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
	n.Host.HandleStream(n.user.GetRouter().LocalTransfer(n.Host.ID), n.session.ReadFromStream)
}
