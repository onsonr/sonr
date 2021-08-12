package client

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// OnConnected: HostNode Connection Response ^
func (c *client) OnConnected(r *md.ConnectionResponse) {
	// Convert Message
	bytes, err := r.ToGeneric()
	if err != nil {
		c.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
		return
	}
	// Call Event
	c.call.OnResponse(bytes)
}

// OnEvent: Local Lobby Event ^
func (n *client) OnRoomEvent(e *md.RoomEvent) {
	// Only Callback when not in Transfer
	if n.device.IsNotStatus(md.Status_TRANSFER) {
		// Convert Message
		bytes, err := e.ToGeneric()
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(bytes)
	}
}

func (n *client) OnSyncEvent(e *md.SyncEvent) {
			// Convert Message
		bytes, err := e.ToGeneric()
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(bytes)
}

// OnLink: Handle Result of Link Request ^
func (n *client) OnLink(success bool, id peer.ID, data []byte) {
	// Unmarshal Link Response
	resp := md.LinkResponse{}
	err := proto.Unmarshal(data, &resp)
	if err != nil {
		n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
		return
	}

	// Check Success
	if success {
		// Create link Event
		link := &md.LinkEvent{
			Success: success,
			Device:  resp.GetDevice(),
			Contact: resp.GetContact(),
		}

		// Marshal Link Event
		buf, err := link.ToGeneric()
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}
		n.call.OnEvent(buf)
	} else {
		// Unsuccessful Link Request
		link := &md.LinkEvent{
			Success: success,
		}

		// Marshal Link Event
		buf, err := link.ToGeneric()
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}
		n.call.OnEvent(buf)
	}
}

// OnInvite: User Received Invite ^
func (n *client) OnInvite(data []byte) {
	// Update Status
	n.call.SetStatus(md.Status_INVITED)

	// Create Request
	req := md.GenericRequest{
		Type: md.GenericRequest_INVITE,
		Data: data,
	}

	// Marshal Request
	buf, err := proto.Marshal(&req)
	if err != nil {
		n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
		return
	}

	// Send Callback
	n.call.OnRequest(buf)
}

// OnReply: Begins File Transfer when Accepted ^
func (n *client) OnReply(id peer.ID, reply []byte) {
	// Create Response
	req := md.GenericResponse{
		Type: md.GenericResponse_REPLY,
		Data: reply,
	}

	// Marshal Request
	buf, err := proto.Marshal(&req)
	if err != nil {
		n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
		return
	}

	// Call Responded
	n.call.OnResponse(buf)

	// Check Peer ID
	if id != "" {
		// InviteResponse Message
		resp := md.InviteResponse{}
		err := proto.Unmarshal(reply, &resp)
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
		}

		// Check for File Transfer
		if resp.HasAcceptedTransfer() {
			// Update Status
			n.call.SetStatus(md.Status_TRANSFER)

			// Create New Auth Stream
			stream, err := n.Host.StartStream(id, md.SonrProtocol_LocalTransfer.NewIDProtocol(id))
			if err != nil {
				n.call.OnError(md.NewError(err, md.ErrorEvent_HOST_STREAM))
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

// OnResponded: Prepares for Incoming File Transfer when Accepted ^
func (n *client) OnConfirmed(inv *md.InviteRequest) {
	n.session = md.NewInSession(n.device, inv, n)
	n.Host.HandleStream(md.SonrProtocol_LocalTransfer.NewIDProtocol(n.Host.ID()), n.session.ReadFromStream)
}

// OnMail: Callback for Mail Event
func (n *client) OnMail(e *md.MailEvent) {
	// Create Mail and Marshal Data
	buf, err := e.ToGeneric()
	if err != nil {
		md.NewMarshalError(err)
		return
	}
	n.call.OnEvent(buf)
}

// OnProgress: Callback Progress Update
func (n *client) OnProgress(buf []byte) {
	// Marshal and Return
	n.call.OnEvent(buf)
}

// OnMail: Callback for Error Event
func (n *client) OnError(err *md.SonrError) {
	n.call.OnError(err)
}

// OnCompleted: Callback Completed Transfer
func (n *client) OnCompleted(stream network.Stream, pid protocol.ID, completeEvent *md.CompleteEvent) {
	if completeEvent.Direction == md.CompleteEvent_INCOMING {
		// Convert to Generic
		buf, err := completeEvent.ToGeneric()
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(buf)
		n.call.SetStatus(md.Status_AVAILABLE)
		n.Host.CloseStream(pid, stream)
	} else if completeEvent.Direction == md.CompleteEvent_OUTGOING {
		// Convert to Generic
		buf, err := completeEvent.ToGeneric()
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(buf)
		n.call.SetStatus(md.Status_AVAILABLE)
	}
}
