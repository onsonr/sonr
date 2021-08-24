package main

import (
	"github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

// CallAuthResponse is called when Auth Response is received
func (s *NodeServer) CallAuthResponse(rsp *data.NoRequest, stream data.NodeService_CallAuthResponseServer) error {
	for {
		select {
		case m := <-s.authResponses:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// CallActionResponse is called when Action Response is received
func (s *NodeServer) CallActionResponse(rsp *data.NoRequest, stream data.NodeService_CallActionResponseServer) error {
	for {
		select {
		case m := <-s.actionResponses:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// CallDecisionResponse is called when Action Response is received
func (s *NodeServer) CallDecisionResponse(rsp *data.NoRequest, stream data.NodeService_CallDecisionResponseServer) error {
	for {
		select {
		case m := <-s.decisionResponses:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// CallLinkResponse is called when Link Response is received
func (s *NodeServer) CallLinkResponse(rsp *data.NoRequest, stream data.NodeService_CallLinkResponseServer) error {
	for {
		select {
		case m := <-s.linkResponses:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// CallMailboxResponse is called when Action Response is received
func (s *NodeServer) CallMailboxResponse(rsp *data.NoRequest, stream data.NodeService_CallMailboxResponseServer) error {
	for {
		select {
		case m := <-s.mailboxResponses:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// CallActionResponse is called when Auth Response is received
func (s *NodeServer) CallVerifyResponse(rsp *data.NoRequest, stream data.NodeService_CallVerifyResponseServer) error {
	for {
		select {
		case m := <-s.verifyResponses:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// OnComplete is called when a complete event is received
func (s *NodeServer) OnComplete(req *data.NoRequest, stream data.NodeService_OnCompleteServer) error {
	for {
		select {
		case m := <-s.completeEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// OnInvite is called when device is invited by a Peer
func (s *NodeServer) OnInvite(req *data.NoRequest, stream data.NodeService_OnInviteServer) error {
	for {
		select {
		case m := <-s.inviteRequests:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// OnReply is called when a peer responds to invite
func (s *NodeServer) OnReply(req *data.NoRequest, stream data.NodeService_OnReplyServer) error {
	for {
		select {
		case m := <-s.inviteResponses:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// OnLink is called when a Link Event has completed for User
func (s *NodeServer) OnLink(req *data.NoRequest, stream data.NodeService_OnLinkServer) error {
	for {
		select {
		case m := <-s.linkEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// OnMail is called when a new mail is received from User
func (s *NodeServer) OnMail(req *data.NoRequest, stream data.NodeService_OnMailServer) error {
	for {
		select {
		case m := <-s.mailEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// OnProgress is called when a file is being transferred
func (s *NodeServer) OnProgress(req *data.NoRequest, stream data.NodeService_OnProgressServer) error {
	for {
		select {
		case m := <-s.progressEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// OnStatus is called when the node receives a status event
func (s *NodeServer) OnStatus(req *data.NoRequest, stream data.NodeService_OnStatusServer) error {
	for {
		select {
		case m := <-s.statusEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// OnRoom is called when Room Event is received
func (s *NodeServer) OnRoom(req *data.NoRequest, stream data.NodeService_OnRoomServer) error {
	for {
		select {
		case m := <-s.RoomEvents:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}

		data.GetState().NeedsWait()
	}
}

// OnError is called when Internal Node Error occurs
func (s *NodeServer) OnError(req *data.NoRequest, stream data.NodeService_OnErrorServer) error {
	for {
		select {
		case m := <-s.errorEvents:
			stream.Send(m)
		case <-s.ctx.Done():
			return nil
		}
		data.GetState().NeedsWait()
	}
}

// Passes binded Methods to Node
func (s *NodeServer) callback() data.Callback {
	return data.Callback{
		OnEvent:    s.handleEvent,
		OnRequest:  s.handleRequest,
		OnResponse: s.handleResponse,
		OnError:    s.handleError,
		SetStatus:  s.setStatus,
	}
}

// Handle Event and Send to Channel after unmarshal
func (s *NodeServer) handleEvent(buf []byte) {
	// Unmarshal Generic Event
	event := &data.GenericEvent{}
	err := proto.Unmarshal(buf, event)
	if err != nil {
		data.LogFatal(err)
		return
	}

	// Switch case event type
	eventType := event.GetType()
	switch eventType {
	case data.GenericEvent_COMPLETE:
		// Unmarshal Complete Event
		ce := &data.CompleteEvent{}
		err = proto.Unmarshal(event.GetData(), ce)
		if err != nil {
			data.LogFatal(err)
			return
		}

		// Logging
		eventType.Log(ce.String())

		// Send Event to Channel
		s.completeEvents <- ce
	case data.GenericEvent_PROGRESS:
		// Unmarshal Progress Event
		pe := &data.ProgressEvent{}
		err = proto.Unmarshal(event.GetData(), pe)
		if err != nil {
			data.LogFatal(err)
			return
		}

		// Logging
		eventType.Log(pe.String())

		// Send Event to Channel
		s.progressEvents <- pe
	case data.GenericEvent_ROOM:
		// Unmarshal Room Event
		te := &data.RoomEvent{}
		err = proto.Unmarshal(event.GetData(), te)
		if err != nil {
			data.LogFatal(err)
			return
		}

		// Logging
		eventType.Log(te.String())

		// Send Event to Channel
		s.RoomEvents <- te

	case data.GenericEvent_MAIL:
		// Unmarshal Mail Event
		me := &data.MailEvent{}
		err = proto.Unmarshal(event.GetData(), me)
		if err != nil {
			data.LogFatal(err)
			return
		}

		// Logging
		eventType.Log(me.String())

		// Send Event to Channel
		s.mailEvents <- me

	case data.GenericEvent_LINK:
		// Unmarshal Link Event
		le := &data.LinkEvent{}
		err = proto.Unmarshal(event.GetData(), le)
		if err != nil {
			data.LogFatal(err)
			return
		}

		// Logging
		eventType.Log(le.String())

		// Send Event to Channel
		s.linkEvents <- le
	}
}

// Handle Request and Send to Channel after unmarshal
func (s *NodeServer) handleRequest(buf []byte) {
	// Unmarshal Generic Request
	request := &data.GenericRequest{}
	err := proto.Unmarshal(buf, request)
	if err != nil {
		data.LogFatal(err)
		return
	}

	// Get Type
	requestType := request.GetType()

	// Switch case request type
	switch requestType {
	case data.GenericRequest_INVITE:
		// Unmarshal Invite Request
		ir := &data.InviteRequest{}
		err = proto.Unmarshal(request.GetData(), ir)
		if err != nil {
			data.LogFatal(err)
			return
		}

		// Logging
		requestType.Log(ir.String())

		// Send Request to Channel
		s.inviteRequests <- ir
	}
}

// Handle Request and Send to Channel after unmarshal
func (s *NodeServer) handleResponse(buf []byte) {
	// Unmarshal Generic Response
	response := &data.GenericResponse{}
	err := proto.Unmarshal(buf, response)
	if err != nil {
		data.LogFatal(err)
		return
	}
	// Get Type
	requestType := response.GetType()

	// Switch case response type
	switch requestType {
	case data.GenericResponse_CONNECTION:
		// Unmarshal Connection Response
		cr := &data.ConnectionResponse{}
		err = proto.Unmarshal(response.GetData(), cr)
		if err != nil {
			data.LogFatal(err)
			return
		}

		// Logging
		requestType.Log(cr.String())

		// Send Response to Channel
		s.connectionResponses <- cr
	case data.GenericResponse_REPLY:
		// Unmarshal Reply Response
		rr := &data.InviteResponse{}
		err = proto.Unmarshal(response.GetData(), rr)
		if err != nil {
			data.LogFatal(err)
			return
		}

		// Logging
		requestType.Log(rr.String())

		// Send Response to Channel
		s.inviteResponses <- rr
	}
}

// handleError Callback with handleError instance, and method
func (s *NodeServer) handleError(errMsg *data.SonrError) {
	// Check for Error
	if errMsg.HasError {
		// Send Callback
		s.errorEvents <- errMsg.Message()
	}
}
