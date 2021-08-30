package client

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/internal/emitter"
	data "github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"

	"google.golang.org/protobuf/proto"
)

func (c *client) initEmitter() {
	c.emitter = emitter.New(util.MAX_EMITTER_CAP)
	c.emitter.On(emitter.EMIT_COMPLETED, c.onCompleted)
	c.emitter.On(emitter.EMIT_CONFIRMED, c.onConfirmed)
	c.emitter.Once(emitter.EMIT_CONNECTED, c.onConnected)
	c.emitter.On(emitter.EMIT_REPLY, c.onReply)
	c.emitter.On(emitter.EMIT_ERROR, c.onError)
	c.emitter.On(emitter.EMIT_INVITE, c.onInvite)
	c.emitter.On(emitter.EMIT_LINK, c.onLink)
	c.emitter.On(emitter.EMIT_MAIL_EVENT, c.onMailEvent)
	c.emitter.On(emitter.EMIT_PROGRESS_EVENT, c.onProgressEvent)
	c.emitter.On(emitter.EMIT_ROOM_EVENT, c.onRoomEvent)
	c.emitter.On(emitter.EMIT_SYNC_EVENT, c.onSyncEvent)
}

// onConnected: HostNode Connection Response ^
// Params: resp *data.ConnectionResponse
func (c *client) onConnected(e *emitter.Event) {
	// Fetch Response
	resp := e.Args[0].(*data.ConnectionResponse)
	// Convert Message
	bytes, err := resp.ToGeneric()
	if err != nil {
		c.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}
	// Call Event
	c.call.OnResponse(bytes)
}

// onMailEvent: Callback for Mail Event
// Params: e *data.MailEvent
func (n *client) onMailEvent(e *emitter.Event) {
	// Fetch Mail Event
	mail := e.Args[0].(*data.MailEvent)

	// Create Mail and Marshal Data
	buf, err := mail.ToGeneric()
	if err != nil {
		data.NewMarshalError(err)
		return
	}
	n.call.OnEvent(buf)
}

// onProgressEvent: Callback Progress Update
// Params: buf []byte
func (n *client) onProgressEvent(e *emitter.Event) {
	// Fetch Variables
	buf := e.Args[0].([]byte)
	n.call.OnEvent(buf)
}

// onRoomEvent handles a Local Lobby Event
// Params: e *data.RoomEvent
func (n *client) onRoomEvent(e *emitter.Event) {
	// Only Callback when not in Transfer
	if n.account.CurrentDevice().IsNotStatus(data.Status_TRANSFER) && len(e.Args) > int(0) {
		if i, okay := e.Args[0].(*data.RoomEvent); okay {
			// Convert Message
			bytes, err := i.ToGeneric()
			if err != nil {
				n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
				return
			}

			// Call Event
			n.call.OnEvent(bytes)
		}
	}
}

// onSyncEvent handles a Local Sync Event
// Params: e *data.SyncEvent
func (n *client) onSyncEvent(e *emitter.Event) {
	if i, okay := e.Args[0].(*data.SyncEvent); okay {
		// Convert Message
		bytes, err := i.ToGeneric()
		if err != nil {
			n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(bytes)
	}
}

// onLink: Handle Result of Link Request ^
// Params: success bool, incoming bool, id peer.ID, buf []byte
func (n *client) onLink(e *emitter.Event) {
	// Fetch Variables
	success := e.Args[0].(bool)
	incoming := e.Args[1].(bool)
	id := e.Args[2].(peer.ID)
	buf := e.Args[3].([]byte)

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

// onInvite: User Received Invite ^
// Params: buf []byte
func (n *client) onInvite(e *emitter.Event) {
	// Fetch Invite
	buf := e.Args[0].([]byte)

	// Update Status
	data.LogInfo("Received Invite")

	// Create Request
	req := data.GenericRequest{
		Type: data.GenericRequest_INVITE,
		Data: buf,
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

// onReply: Begins File Transfer when Accepted ^
// Params: id peer.ID, reply []byte
func (n *client) onReply(e *emitter.Event) {
	// Fetch Variables
	id := e.Args[0].(peer.ID)
	reply := e.Args[1].([]byte)

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
		if resp.HasAcceptedTransfer() && n.session != nil {
			data.LogInfo("Beginning Transfer")
			pid := data.SonrProtocol_LocalTransfer.NewIDProtocol(id)

			// Create New Auth Stream
			stream, err := n.Host.StartStream(id, pid)
			if err != nil {
				n.call.OnError(data.NewError(err, data.ErrorEvent_HOST_STREAM))
				return
			}

			// Write to Stream on Session
			n.session.WriteToStream(stream)
		}
	}
}

// OnResponded: Prepares for Incoming File Transfer when Accepted ^
// Params: inv *data.InviteRequest
func (c *client) onConfirmed(e *emitter.Event) {
	// Fetch Invite
	inv := e.Args[0].(*data.InviteRequest)
	c.session = data.NewInSession(c.account.CurrentDevice(), inv, c.emitter)
	c.Host.HandleStream(data.SonrProtocol_LocalTransfer.NewIDProtocol(c.Host.ID()), c.session.ReadFromStream)
}

// OnMail: Callback for Error Event
// Params: e *data.SonrError
func (n *client) onError(e *emitter.Event) {
	// Fetch Variables
	err := e.Args[0].(*data.SonrError)
	n.call.OnError(err)
}

// onCompleted: Callback Completed Transfer
// Params: stream network.Stream, pid protocol.ID, completeEvent *data.CompleteEvent
func (n *client) onCompleted(e *emitter.Event) {
	// Fetch Variables
	stream := e.Args[0].(network.Stream)
	pid := e.Args[1].(protocol.ID)
	completeEvent := e.Args[2].(*data.CompleteEvent)

	// Check Status
	if completeEvent.Direction == data.CompleteEvent_INCOMING {
		// Convert to Generic
		buf, err := completeEvent.ToGeneric()
		if err != nil {
			n.call.OnError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
			return
		}

		// Call Event
		n.call.OnEvent(buf)
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
	}
}
