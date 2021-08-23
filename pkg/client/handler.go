package client

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	data "github.com/sonr-io/core/pkg/data"

	"google.golang.org/protobuf/proto"
)

// OnConnected: HostNode Connection Response ^
func (c *client) OnConnected(r *data.ConnectionResponse) {
	// Convert Message
	bytes, err := r.ToGeneric()
	if err != nil {
		c.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}
	// Call Event
	c.call.OnResponse(bytes)
}

// OnEvent: Local Lobby Event ^
func (n *client) OnRoomEvent(e *data.RoomEvent) {
	// Only Callback when not in Transfer
	if n.account.CurrentDevice().IsNotStatus(data.Status_TRANSFER) {
		// Convert Message
		bytes, err := e.ToGeneric()
		if err != nil {
			n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(bytes)
	}
}

func (n *client) OnSyncEvent(e *data.SyncEvent) {
	// Convert Message
	bytes, err := e.ToGeneric()
	if err != nil {
		n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}

	// Call Event
	n.call.OnEvent(bytes)
}

// OnLink: Handle Result of Link Request ^
func (n *client) OnLink(success bool, incoming bool, id peer.ID, buf []byte) {
	// Unmarshal Link Response
	resp := data.LinkResponse{}
	err := proto.Unmarshal(buf, &resp)
	if err != nil {
		n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}

	// Check Success
	if success {
		// Create link Event
		link := &data.LinkEvent{
			Success: success,
			Device:  resp.GetDevice(),
			Contact: resp.GetContact(),
		}

		// Marshal Link Event
		buf, err := link.ToGeneric()
		if err != nil {
			n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
			return
		}
		n.call.OnEvent(buf)

		// Pass Keys with Linker
		if incoming {
			// Open Stream if Incoming
			pid := data.SonrProtocol_Linker.NewIDProtocol(n.Host.ID())
			n.Host.HandleStream(pid, n.account.ReadFromLink)
		} else {
			// Create Stream if Outgoing
			pid := data.SonrProtocol_Linker.NewIDProtocol(id)
			// Write Stream
			stream, err := n.Host.StartStream(id, pid)
			if err != nil {
				n.call.OnError(data.NewError(err, data.ErrorEvent_HOST_STREAM))
				return
			}
			n.account.WriteToLink(stream, &resp)
		}
	} else {
		// Unsuccessful Link Request
		link := &data.LinkEvent{
			Success: success,
		}

		// Marshal Link Event
		buf, err := link.ToGeneric()
		if err != nil {
			n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
			return
		}
		n.call.OnEvent(buf)
	}
}

// OnInvite: User Received Invite ^
func (n *client) OnInvite(buf []byte) {
	// Update Status
	n.call.SetStatus(data.Status_INVITED)

	// Create Request
	req := data.GenericRequest{
		Type: data.GenericRequest_INVITE,
		Data: data,
	}

	// Marshal Request
	buf, err := proto.Marshal(&req)
	if err != nil {
		n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}

	// Send Callback
	n.call.OnRequest(buf)
}

// OnReply: Begins File Transfer when Accepted ^
func (n *client) OnReply(id peer.ID, reply []byte) {
	// Create Response
	req := data.GenericResponse{
		Type: data.GenericResponse_REPLY,
		Data: reply,
	}

	// Marshal Request
	buf, err := proto.Marshal(&req)
	if err != nil {
		n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}

	// Call Responded
	n.call.OnResponse(buf)

	// Check Peer ID
	if id != "" {
		// InviteResponse Message
		resp := data.InviteResponse{}
		err := proto.Unmarshal(reply, &resp)
		if err != nil {
			n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		}

		// Check for File Transfer
		if resp.HasAcceptedTransfer() {
			// Update Status
			n.call.SetStatus(data.Status_TRANSFER)

			// Create New Auth Stream
			stream, err := n.Host.StartStream(id, data.SonrProtocol_LocalTransfer.NewIDProtocol(id))
			if err != nil {
				n.call.OnError(data.NewError(err, data.ErrorEvent_HOST_STREAM))
				return
			}

			// Write to Stream on Session
			n.session.WriteToStream(stream)
		} else {
			n.call.SetStatus(data.Status_AVAILABLE)
		}
	} else {
		n.call.SetStatus(data.Status_AVAILABLE)
	}
}

// OnResponded: Prepares for Incoming File Transfer when Accepted ^
func (n *client) OnConfirmed(inv *data.InviteRequest) {
	n.session = data.NewInSession(n.account.CurrentDevice(), inv, n)
	n.Host.HandleStream(data.SonrProtocol_LocalTransfer.NewIDProtocol(n.Host.ID()), n.session.ReadFromStream)
}

// OnMail: Callback for Mail Event
func (n *client) OnMail(e *data.MailEvent) {
	// Create Mail and Marshal Data
	buf, err := e.ToGeneric()
	if err != nil {
		data.NewMarshalError(err)
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
func (n *client) OnError(err *data.SonrError) {
	n.call.OnError(err)
}

// OnCompleted: Callback Completed Transfer
func (n *client) OnCompleted(stream network.Stream, pid protocol.ID, completeEvent *data.CompleteEvent) {
	if completeEvent.Direction == data.CompleteEvent_INCOMING {
		// Convert to Generic
		buf, err := completeEvent.ToGeneric()
		if err != nil {
			n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(buf)
		n.call.SetStatus(data.Status_AVAILABLE)
		n.Host.CloseStream(pid, stream)
	} else if completeEvent.Direction == data.CompleteEvent_OUTGOING {
		// Convert to Generic
		buf, err := completeEvent.ToGeneric()
		if err != nil {
			n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(buf)
		n.call.SetStatus(data.Status_AVAILABLE)
	}
}
