package client

import (
	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ OnConnected: HostNode Connection Response ^
func (c *client) OnConnected(r *md.ConnectionResponse) {
	// Convert Message
	bytes, err := proto.Marshal(r)
	if err != nil {
		c.call.OnError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}
	// Call Event
	c.call.OnConnected(bytes)
}

// ^ OnEvent: Local Lobby Event ^
func (n *client) OnEvent(e *md.LobbyEvent) {
	// Only Callback when not in Transfer
	if n.user.IsNotStatus(md.Status_TRANSFER) {
		// Convert Message
		bytes, err := proto.Marshal(e)
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(bytes)
	}
}

// ^ OnInvite: User Received Invite ^
func (n *client) OnInvite(data []byte) {
	// Update Status
	n.call.SetStatus(md.Status_INVITED)

	// Send Callback
	n.call.OnInvite(data)
}

// ^ OnReply: Begins File Transfer when Accepted ^
func (n *client) OnReply(id peer.ID, reply []byte) {
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
func (n *client) OnConfirmed(inv *md.InviteRequest) {
	n.session = md.NewInSession(n.user, inv, n.call)
	n.Host.HandleStream(n.user.GetRouter().LocalTransferProtocol(n.Host.ID()), n.session.ReadFromStream)
}

// ^ OnMail: Callback for Mail Event
func (n *client) OnMail(mail *md.MailEvent) {
	buf, err := proto.Marshal(mail)
	if err != nil {
		n.call.OnError(md.NewUnmarshalError(err))
	}
	n.call.OnMail(buf)
}
