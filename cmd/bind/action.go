package bind

import (
	"github.com/sonr-io/core/internal/logger"
	"github.com/sonr-io/core/pkg/data"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// ** ─── Node Binded Actions ────────────────────────────────────────────────────────
// Signing Request for Data

func (n *Node) Sign(buf []byte) []byte {
	// Unmarshal Data to Request
	request := &data.AuthRequest{}
	err := proto.Unmarshal(buf, request)
	if err != nil {
		// Handle Error
		n.handleError(data.NewUnmarshalError(err))

		// Initialize invalid Response
		invalidResp := data.AuthResponse{
			IsSigned: false,
		}

		// Send Invalid Response
		buf, err := proto.Marshal(&invalidResp)
		if err != nil {
			n.handleError(data.NewMarshalError(err))
			return nil
		}
		return buf
	}

	// Sign Buffer
	result := n.account.SignAuth(request)
	res, err := proto.Marshal(result)
	if err != nil {
		n.handleError(data.NewMarshalError(err))
		return nil
	}
	return res
}

// Verification Request for Signed Data
func (n *Node) Verify(buf []byte) []byte {
	// Check Read
	// Get Key Pair
	kp := n.account.AccountKeys()

	// Unmarshal Data to Request
	request := &data.VerifyRequest{}
	if err := proto.Unmarshal(buf, request); err != nil {
		// Handle Error
		n.handleError(data.NewUnmarshalError(err))

		// Send Invalid Response
		return data.NewInvalidVerifyResponseBuf()
	}

	// Check Request Type
	if request.GetType() == data.VerifyRequest_VERIFY {
		// Check type and Verify
		if request.IsBuffer() {
			// Verify Result
			result, err := kp.Verify(request.GetBufferValue(), request.GetSignedBuffer())
			if err != nil {
				return data.NewInvalidVerifyResponseBuf()
			}

			// Return Result
			return data.NewVerifyResponseBuf(result)
		} else if request.IsString() {
			// Verify Result
			result, err := kp.Verify([]byte(request.GetTextValue()), []byte(request.GetSignedText()))
			if err != nil {
				return data.NewInvalidVerifyResponseBuf()
			}

			// Return Result
			return data.NewVerifyResponseBuf(result)
		}
	} else {
		resp := n.account.VerifyRead()
		buf, err := proto.Marshal(resp)
		if err != nil {
			n.handleError(data.NewMarshalError(err))
		}
		return buf
	}

	// Send Invalid Response
	return data.NewInvalidVerifyResponseBuf()
}

// Update proximity/direction and Notify Lobby
func (n *Node) Update(buf []byte) {
	// Unmarshal Data to Request
	update := &data.UpdateRequest{}
	if err := proto.Unmarshal(buf, update); err != nil {
		n.handleError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}

	// Check Update Request Type
	switch update.Data.(type) {
	// Update Position
	case *data.UpdateRequest_Position:
		n.account.CurrentDevice().UpdatePosition(update.GetPosition().Parameters())

	// Update Contact
	case *data.UpdateRequest_Contact:
		n.account.UpdateContact(update.GetContact())

	// Update Peer Properties
	case *data.UpdateRequest_Properties:
		n.account.CurrentDevice().UpdateProperties(update.GetProperties())
	}

	// Notify Local Lobby
	err := n.client.Update(n.local)
	if err != nil {
		n.handleError(err)
		return
	}
}

// Invite Processes Data and Sends Invite to Peer
func (n *Node) Invite(buf []byte) {
	// Unmarshal Data to Request
	req := &data.InviteRequest{}
	if err := proto.Unmarshal(buf, req); err != nil {
		n.handleError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}

	// Validate invite
	req = n.account.SignInvite(req)
	logger.Info("Signed Invite.", zap.String("invite", req.String()))

	// Send Invite
	err := n.client.Invite(req, n.local)
	if err != nil {
		n.handleError(err)
		return
	}
}

// Link method starts device linking channel
func (n *Node) Link(buf []byte) []byte {
	// Unmarshal Data to Request
	req := &data.LinkRequest{}
	if err := proto.Unmarshal(buf, req); err != nil {
		n.handleError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return nil
	}
	req = n.account.CurrentDevice().SignLink(req)

	// Send to Client
	resp, serr := n.client.Link(req, n.local)
	if serr != nil {
		n.handleError(serr)
		return nil
	}

	// Marshal Response
	buf, err := proto.Marshal(resp)
	if err != nil {
		n.handleError(data.NewMarshalError(err))
		return nil
	}
	return buf
}

// Mail handles request for a message in Mailbox
func (n *Node) Mail(buf []byte) []byte {
	// Unmarshal Data to Request
	req := &data.MailboxRequest{}
	if err := proto.Unmarshal(buf, req); err != nil {
		n.handleError(data.NewUnmarshalError(err))
		return nil
	}

	// Handle Mail
	resp, serr := n.client.Mail(req)
	if serr != nil {
		n.handleError(serr)
		return nil
	}

	// Marshal Response
	buf, err := proto.Marshal(resp)
	if err != nil {
		n.handleError(data.NewMarshalError(err))
		return nil
	}

	// Return Response
	return buf
}

// Respond to an Invite with Decision
func (n *Node) Respond(buf []byte) {
	// Unmarshal Data to Request
	resp := &data.DecisionRequest{}
	if err := proto.Unmarshal(buf, resp); err != nil {
		n.handleError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return
	}

	// Send Response
	n.client.Respond(resp.ToResponse())
}

// ** ─── Misc Methods ────────────────────────────────────────────────────────
// Action method handles misceallaneous actions for node
func (s *Node) Action(buf []byte) []byte {
	// Unmarshal Data to Request
	req := &data.ActionRequest{}
	if err := proto.Unmarshal(buf, req); err != nil {
		s.handleError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
		return nil
	}

	// Check Action
	switch req.Action {
	case data.Action_PING:
		// Ping
		resp := &data.ActionResponse{
			Success: true,
			Action:  data.Action_PING,
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(data.NewMarshalError(err))
			return nil
		}
		return bytes
	case data.Action_LOCATION:
		// Location
		resp := &data.ActionResponse{
			Success: true,
			Action:  data.Action_LOCATION,
			Data: &data.ActionResponse_Location{
				Location: s.account.CurrentDevice().GetLocation(),
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(data.NewMarshalError(err))
			return nil
		}
		return bytes
	case data.Action_URL_LINK:
		// URL Link
		resp := &data.ActionResponse{
			Success: true,
			Action:  data.Action_URL_LINK,
			Data: &data.ActionResponse_UrlLink{
				UrlLink: data.NewURLLink(req.GetData()),
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(data.NewMarshalError(err))
			return nil
		}
		return bytes
	case data.Action_LIST_LINKERS:
		// List Linkers
		resp := &data.ActionResponse{
			Success: true,
			Action:  data.Action_LIST_LINKERS,
			Data: &data.ActionResponse_Linkers{
				Linkers: s.local.ListLinkers(),
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(data.NewMarshalError(err))
			return nil
		}
		return bytes
	case data.Action_PAUSE:
		// Pause
		logger.Info("Lifecycle Update", zap.String("value", data.Action_name[int32(data.Action_PAUSE)]))
		s.state = data.Lifecycle_PAUSED
		s.client.Lifecycle(s.state, s.local)
		data.GetState().Pause()
		resp := &data.ActionResponse{
			Success: true,
			Action:  data.Action_PAUSE,
			Data: &data.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(data.NewMarshalError(err))
			return nil
		}
		return bytes
	case data.Action_RESUME:
		// Resume
		logger.Info("Lifecycle Update", zap.String("value", data.Action_name[int32(data.Action_RESUME)]))
		s.state = data.Lifecycle_ACTIVE
		s.client.Lifecycle(s.state, s.local)
		data.GetState().Resume()

		// Create Response
		resp := &data.ActionResponse{
			Success: true,
			Action:  data.Action_RESUME,
			Data: &data.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(data.NewMarshalError(err))
			return nil
		}
		return bytes
	case data.Action_STOP:
		// Stop
		logger.Info("Lifecycle Update", zap.String("value", data.Action_name[int32(data.Action_STOP)]))
		s.state = data.Lifecycle_STOPPED
		s.client.Lifecycle(s.state, s.local)

		// Create Response
		resp := &data.ActionResponse{
			Success: true,
			Action:  data.Action_STOP,
			Data: &data.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(data.NewMarshalError(err))
			return nil
		}
		return bytes
	}
	return nil
}
