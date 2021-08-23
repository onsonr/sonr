package main

import (
	"context"
	"fmt"

	md "github.com/sonr-io/core/pkg/models"
)

// Action method handles miscellaneous actions for node
func (s *NodeServer) Action(ctx context.Context, req *md.ActionRequest) (*md.NoResponse, error) {
	// Check Action
	switch req.Action {
	case md.Action_PING:
		// Ping
		md.LogRPC("action", "ping")
		s.actionResponses <- &md.ActionResponse{
			Success: true,
			Action:  md.Action_PING,
		}
	case md.Action_LOCATION:
		// Location
		md.LogRPC("action", "location")
		s.actionResponses <- &md.ActionResponse{
			Success: true,
			Action:  md.Action_LOCATION,
			Data: &md.ActionResponse_Location{
				Location: s.account.CurrentDevice().GetLocation(),
			},
		}
	case md.Action_URL_LINK:
		// URL Link
		md.LogRPC("action", "url")
		s.actionResponses <- &md.ActionResponse{
			Success: true,
			Action:  md.Action_URL_LINK,
			Data: &md.ActionResponse_UrlLink{
				UrlLink: md.NewURLLink(req.GetData()),
			},
		}
	case md.Action_PAUSE:
		// Pause
		md.LogRPC("action", "pause")
		s.state = md.Lifecycle_PAUSED
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Pause()
		s.actionResponses <- &md.ActionResponse{
			Success: true,
			Action:  md.Action_PAUSE,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}
	case md.Action_RESUME:
		// Resume
		md.LogRPC("action", "resume")
		s.state = md.Lifecycle_ACTIVE
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Resume()

		s.actionResponses <- &md.ActionResponse{
			Success: true,
			Action:  md.Action_RESUME,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}
	case md.Action_STOP:
		// Stop
		md.LogRPC("action", "stop")
		s.state = md.Lifecycle_STOPPED
		s.client.Lifecycle(s.state, s.local)
		s.actionResponses <- &md.ActionResponse{
			Success: true,
			Action:  md.Action_STOP,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}
	case md.Action_LIST_LINKERS:
		// List Linkers
		s.actionResponses <- &md.ActionResponse{
			Success: true,
			Action:  md.Action_LIST_LINKERS,
			Data: &md.ActionResponse_Linkers{
				Linkers: s.local.ListLinkers(),
			},
		}
	default:
		md.LogRPC("action", false)
		return nil, fmt.Errorf("Action: %s not supported", req.Action)
	}
	return &md.NoResponse{}, nil
}

// Sign method signs data with device's private key
func (s *NodeServer) Sign(ctx context.Context, req *md.AuthRequest) (*md.NoResponse, error) {
	md.LogRPC("Sign", req)
	s.authResponses <- s.account.SignAuth(req)
	return &md.NoResponse{}, nil
}

// Link method starts device linking channel
func (s *NodeServer) Link(ctx context.Context, req *md.LinkRequest) (*md.NoResponse, error) {
	md.LogRPC("Link", req)
	req = s.account.CurrentDevice().SignLink(req)

	// Check Link Request Type
	resp, err := s.client.Link(req, s.local)
	if err != nil {
		s.handleError(err)
		return nil, err.Error
	}

	// Return Link Response
	s.linkResponses <- resp
	return &md.NoResponse{}, nil
}

// Verify validates device Keys
func (s *NodeServer) Verify(ctx context.Context, req *md.VerifyRequest) (*md.NoResponse, error) {
	md.LogRPC("Verify", req)
	// Get Key Pair
	kp := s.account.AccountKeys()

	// Check Request Type
	if req.GetType() == md.VerifyRequest_VERIFY {
		// Check type and Verify
		if req.IsBuffer() {
			// Verify Result
			result, err := kp.Verify(req.GetBufferValue(), req.GetSignedBuffer())
			if err != nil {
				s.verifyResponses <- &md.VerifyResponse{Success: false}
			}

			// Return Result
			s.verifyResponses <- &md.VerifyResponse{Success: result}
			return &md.NoResponse{}, nil
		} else if req.IsString() {
			// Verify Result
			result, err := kp.Verify([]byte(req.GetTextValue()), []byte(req.GetSignedText()))
			if err != nil {
				s.verifyResponses <- &md.VerifyResponse{Success: false}
				return &md.NoResponse{}, nil
			}

			// Return Result
			s.verifyResponses <- &md.VerifyResponse{Success: result}
			return &md.NoResponse{}, nil
		}
	}
	s.verifyResponses <- s.account.VerifyRead()
	return &md.NoResponse{}, nil
}

// Update proximity/direction/contact/properties and notify Lobby
func (s *NodeServer) Update(ctx context.Context, req *md.UpdateRequest) (*md.NoResponse, error) {
	// Verify Node is Ready
	if s.account.IsReady() {
		// Check Update Request Type
		switch req.Data.(type) {
		// Update Position
		case *md.UpdateRequest_Position:
			s.account.CurrentDevice().UpdatePosition(req.GetPosition().Parameters())

		// Update Contact
		case *md.UpdateRequest_Contact:
			s.account.UpdateContact(req.GetContact())

		// Update Peer Properties
		case *md.UpdateRequest_Properties:
			s.account.CurrentDevice().UpdateProperties(req.GetProperties())
		}

		// Notify Local Lobby
		err := s.client.Update(s.local)
		if err != nil {
			s.handleError(err)
			return nil, err.Error
		}
	}

	// Return Blank Response
	return &md.NoResponse{}, nil
}

// Invite pushes Invite request to Peer
func (s *NodeServer) Invite(ctx context.Context, req *md.InviteRequest) (*md.NoResponse, error) {
	// Verify Node is Ready
	if s.account.IsReady() {
		// Validate invite
		req = s.account.CurrentDevice().SignInvite(req)

		// Send Invite
		err := s.client.Invite(req, s.local)
		if err != nil {
			s.handleError(err)
			return nil, err.Error
		}
	}
	// Return Blank Response
	return &md.NoResponse{}, nil
}

// Respond handles a respond request
func (s *NodeServer) Respond(ctx context.Context, req *md.DecisionRequest) (*md.NoResponse, error) {
	// Verify Node is Ready
	if s.account.IsReady() {
		// Send Response
		s.client.Respond(req.ToResponse())

		// Update Status
		if req.Decision.Accepted() {
			s.setStatus(md.Status_TRANSFER)
		} else {
			s.setStatus(md.Status_AVAILABLE)
		}

		// Return Blank Response
		return nil, nil
	}
	return nil, fmt.Errorf("Node is not ready")
}

// Mail method handles a mail request
func (s *NodeServer) Mail(ctx context.Context, req *md.MailboxRequest) (*md.NoResponse, error) {
	// Verify Node is Ready
	if s.account.IsReady() {
		// Handle Mail
		resp, serr := s.client.Mail(req)
		if serr != nil {
			s.handleError(serr)
			return nil, serr.Error
		}

		s.mailboxResponses <- resp
		// Return Response
		return &md.NoResponse{}, nil
	}
	return nil, fmt.Errorf("Node is not ready")
}
