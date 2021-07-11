package client

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
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

	// Check Peer ID
	if id != "" {
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
			stream, err := n.Host.StartStream(id, md.SonrProtocol_LocalTransfer.NewIDProtocol(id))
			if err != nil {
				n.call.OnError(md.NewError(err, md.ErrorMessage_HOST_STREAM))
				return
			}

			// Write to Stream on Session
			n.session.WriteToStream(stream)
		} else {
			n.call.SetStatus(md.Status_AVAILABLE)
		}
	} else {
		n.call.SetStatus(md.Status_AVAILABLE)
	}
}

// ^ OnResponded: Prepares for Incoming File Transfer when Accepted ^
func (n *client) OnConfirmed(inv *md.InviteRequest) {
	pid := md.SonrProtocol_LocalTransfer.NewIDProtocol(n.Host.ID())
	n.session = md.NewInSession(n.user, inv, pid, n.call)
	n.Host.HandleStream(pid, n.session.ReadFromStream)
}

// ^ OnMail: Callback for Mail Event
func (n *client) OnMail(buf []byte) {
	n.call.OnMail(buf)
}

func (n *client) OnCompleted(stream network.Stream, pid protocol.ID, buf []byte) {

}
