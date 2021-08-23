package bind

import (
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ** ─── Node Binded Actions ────────────────────────────────────────────────────────
// Signing Request for Data
func (n *Node) Sign(data []byte) []byte {
	// Unmarshal Data to Request
	request := &md.AuthRequest{}
	err := proto.Unmarshal(data, request)
	if err != nil {
		// Handle Error
		n.handleError(md.NewUnmarshalError(err))

		// Initialize invalid Response
		invalidResp := md.AuthResponse{
			IsSigned: false,
		}

		// Send Invalid Response
		buf, err := proto.Marshal(&invalidResp)
		if err != nil {
			n.handleError(md.NewMarshalError(err))
			return nil
		}
		return buf
	}

	// Sign Buffer
	result := n.account.SignAuth(request)
	buf, err := proto.Marshal(result)
	if err != nil {
		n.handleError(md.NewMarshalError(err))
		return nil
	}
	return buf
}

// Verification Request for Signed Data
func (n *Node) Verify(data []byte) []byte {
	// Check Ready
	if n.account.IsReady() {
		// Get Key Pair
		kp := n.account.AccountKeys()

		// Unmarshal Data to Request
		request := &md.VerifyRequest{}
		if err := proto.Unmarshal(data, request); err != nil {
			// Handle Error
			n.handleError(md.NewUnmarshalError(err))

			// Send Invalid Response
			return md.NewInvalidVerifyResponseBuf()
		}

		// Check Request Type
		if request.GetType() == md.VerifyRequest_VERIFY {
			// Check type and Verify
			if request.IsBuffer() {
				// Verify Result
				result, err := kp.Verify(request.GetBufferValue(), request.GetSignedBuffer())
				if err != nil {
					return md.NewInvalidVerifyResponseBuf()
				}

				// Return Result
				return md.NewVerifyResponseBuf(result)
			} else if request.IsString() {
				// Verify Result
				result, err := kp.Verify([]byte(request.GetTextValue()), []byte(request.GetSignedText()))
				if err != nil {
					return md.NewInvalidVerifyResponseBuf()
				}

				// Return Result
				return md.NewVerifyResponseBuf(result)
			}
		} else {
			resp := n.account.VerifyRead()
			buf, err := proto.Marshal(resp)
			if err != nil {
				n.handleError(md.NewMarshalError(err))
			}
			return buf
		}

	}

	// Send Invalid Response
	return md.NewInvalidVerifyResponseBuf()
}

// Update proximity/direction and Notify Lobby
func (n *Node) Update(data []byte) {
	if n.account.IsReady() {
		// Unmarshal Data to Request
		update := &md.UpdateRequest{}
		if err := proto.Unmarshal(data, update); err != nil {
			n.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Check Update Request Type
		switch update.Data.(type) {
		// Update Position
		case *md.UpdateRequest_Position:
			n.account.CurrentDevice().UpdatePosition(update.GetPosition().Parameters())

		// Update Contact
		case *md.UpdateRequest_Contact:
			n.account.UpdateContact(update.GetContact())

		// Update Peer Properties
		case *md.UpdateRequest_Properties:
			n.account.CurrentDevice().UpdateProperties(update.GetProperties())
		}

		// Notify Local Lobby
		err := n.client.Update(n.local)
		if err != nil {
			n.handleError(err)
			return
		}
	}
}

// Invite Processes Data and Sends Invite to Peer
func (n *Node) Invite(data []byte) {
	if n.account.IsReady() {
		// Unmarshal Data to Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			n.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Validate invite
		req = n.account.CurrentDevice().SignInvite(req)

		// Send Invite
		err := n.client.Invite(req, n.local)
		if err != nil {
			n.handleError(err)
			return
		}
	}
}

// Link method starts device linking channel
func (n *Node) Link(data []byte) []byte {
	// Unmarshal Data to Request
	req := &md.LinkRequest{}
	if err := proto.Unmarshal(data, req); err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
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
		n.handleError(md.NewMarshalError(err))
		return nil
	}
	return buf
}

// Mail handles request for a message in Mailbox
func (n *Node) Mail(data []byte) []byte {
	// Check Ready
	if n.account.IsReady() {
		// Unmarshal Data to Request
		req := &md.MailboxRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			n.handleError(md.NewUnmarshalError(err))
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
			n.handleError(md.NewMarshalError(err))
			return nil
		}

		// Return Response
		return buf
	}
	return nil
}

// Respond to an Invite with Decision
func (n *Node) Respond(data []byte) {
	if n.account.IsReady() {
		// Unmarshal Data to Request
		resp := &md.DecisionRequest{}
		if err := proto.Unmarshal(data, resp); err != nil {
			n.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
			return
		}

		// Send Response
		n.client.Respond(resp.ToResponse())

		// Update Status
		if resp.Decision.Accepted() {
			n.setStatus(md.Status_TRANSFER)
		} else {
			n.setStatus(md.Status_AVAILABLE)
		}
	}
}

// ** ─── Misc Methods ────────────────────────────────────────────────────────
// Action method handles misceallaneous actions for node
func (s *Node) Action(buf []byte) []byte {
	// Unmarshal Data to Request
	req := &md.ActionRequest{}
	if err := proto.Unmarshal(buf, req); err != nil {
		s.handleError(md.NewError(err, md.ErrorEvent_UNMARSHAL))
		return nil
	}

	// Check Action
	switch req.Action {
	case md.Action_PING:
		// Ping
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_PING,
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_LOCATION:
		// Location
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_LOCATION,
			Data: &md.ActionResponse_Location{
				Location: s.account.CurrentDevice().GetLocation(),
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_URL_LINK:
		// URL Link
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_URL_LINK,
			Data: &md.ActionResponse_UrlLink{
				UrlLink: md.NewURLLink(req.GetData()),
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_LIST_LINKERS:
		// List Linkers
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_LIST_LINKERS,
			Data: &md.ActionResponse_Linkers{
				Linkers: s.local.ListLinkers(),
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_PAUSE:
		// Pause
		md.LogInfo("Lifecycle Pause Called")
		s.state = md.Lifecycle_PAUSED
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Pause()
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_PAUSE,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_RESUME:
		// Resume
		md.LogInfo("Lifecycle Resume Called")
		s.state = md.Lifecycle_ACTIVE
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Resume()

		// Create Response
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_RESUME,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	case md.Action_STOP:
		// Stop
		md.LogInfo("Lifecycle Stop Called")
		s.state = md.Lifecycle_STOPPED
		s.client.Lifecycle(s.state, s.local)

		// Create Response
		resp := &md.ActionResponse{
			Success: true,
			Action:  md.Action_STOP,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}

		// Marshal Response
		bytes, err := proto.Marshal(resp)
		if err != nil {
			s.handleError(md.NewMarshalError(err))
			return nil
		}
		return bytes
	}
	return nil
}
