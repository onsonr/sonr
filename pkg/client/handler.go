package client

import (
	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ OnEvent: Local Lobby Event ^
func (n *Client) OnEvent(e *md.LobbyEvent) {
	// Convert Message
	bytes, err := proto.Marshal(e)
	if err != nil {
		n.call.OnError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}

	// Call Event
	n.call.OnEvent(bytes)
}

// ^ OnRefresh: Topic has Updated ^
func (n *Client) OnRefresh(l *md.Lobby) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		n.call.OnError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}
	n.call.OnRefresh(bytes)
}

// ^ OnInvite: User Received Invite ^
func (n *Client) OnInvite(data []byte) {
	// Update Status
	n.call.SetStatus(md.Status_INVITED)

	// Send Callback
	n.call.OnInvite(data)
}

// ^ OnReply: Begins File Transfer when Accepted ^
func (n *Client) OnReply(id peer.ID, reply []byte) {

	// Call Responded
	n.call.OnReply(reply)

	// InviteResponse Message
	resp := md.InviteResponse{}
	err := proto.Unmarshal(reply, &resp)
	if err != nil {
		n.call.OnError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
	}

	// Check for File Transfer
	if resp.HasAcceptedTransfer() {
		// Update Status
		n.call.SetStatus(md.Status_TRANSFER)

		// Create New Auth Stream
		stream, err := n.Host.StartStream(id, n.user.GetRouter().LocalTransferProtocol(id))
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorMessage_HOST_STREAM))
			return
		}

		// Write to Stream on Session
		n.session.WriteToStream(stream)
	} else {
		n.call.SetStatus(md.Status_AVAILABLE)
	}
}

// ^ OnResponded: Prepares for Incoming File Transfer when Accepted ^
func (n *Client) OnResponded(inv *md.InviteRequest) {
	n.session = md.NewInSession(n.user, inv, n.call)
	n.Host.HandleStream(n.user.GetRouter().LocalTransferProtocol(n.Host.ID()), n.session.ReadFromStream)
}
