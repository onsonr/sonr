package main

import (
	"context"
	"fmt"
	"log"

	md "github.com/sonr-io/core/pkg/models"
)

// Action method handles misceallaneous actions for node
func (s *NodeServer) Action(ctx context.Context, req *md.ActionRequest) (*md.ActionResponse, error) {
	// Check Action
	switch req.Action {
	case md.Action_PING:
		// Ping
		md.LogRPC("action", "ping")
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_PING,
		}, nil
	case md.Action_LOCATION:
		// Location
		md.LogRPC("action", "location")
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_LOCATION,
			Data: &md.ActionResponse_Location{
				Location: s.user.GetLocation(),
			},
		}, nil
	case md.Action_URL_LINK:
		// URL Link
		md.LogRPC("action", "url")
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_URL_LINK,
			Data: &md.ActionResponse_UrlLink{
				UrlLink: md.NewURLLink(req.GetData()),
			},
		}, nil
	case md.Action_PAUSE:
		// Pause
		md.LogRPC("action", "pause")
		s.state = md.Lifecycle_PAUSED
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Pause()
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_PAUSE,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}, nil
	case md.Action_RESUME:
		// Resume
		md.LogRPC("action", "resume")
		s.state = md.Lifecycle_ACTIVE
		s.client.Lifecycle(s.state, s.local)
		md.GetState().Resume()

		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_RESUME,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}, nil
	case md.Action_STOP:
		// Stop
		md.LogRPC("action", "stop")
		s.state = md.Lifecycle_STOPPED
		s.client.Lifecycle(s.state, s.local)
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_STOP,
			Data: &md.ActionResponse_Lifecycle{
				Lifecycle: s.state,
			},
		}, nil
	case md.Action_LIST_LINKERS:
		// List Linkers
		return &md.ActionResponse{
			Success: true,
			Action:  md.Action_LIST_LINKERS,
			Data: &md.ActionResponse_Linkers{
				Linkers: s.local.ListLinkers(),
			},
		}, nil
	default:
		md.LogRPC("action", false)
		return nil, fmt.Errorf("Action: %s not supported", req.Action)
	}
}

// Sign method signs data with user's private key
func (s *NodeServer) Sign(ctx context.Context, req *md.AuthRequest) (*md.AuthResponse, error) {
	log.Println("Sign Called")
	return s.user.Sign(req), nil
}

// Verify validates user Keys
func (s *NodeServer) Verify(ctx context.Context, req *md.VerifyRequest) (*md.VerifyResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Get Key Pair
		kp := s.user.KeyPair()

		// Check Request Type
		if req.GetType() == md.VerifyRequest_VERIFY {
			// Check type and Verify
			if req.IsBuffer() {
				// Verify Result
				result, err := kp.Verify(req.GetBufferValue(), req.GetSignedBuffer())
				if err != nil {
					return &md.VerifyResponse{IsVerified: false}, err
				}

				// Return Result
				return &md.VerifyResponse{IsVerified: result}, nil
			} else if req.IsString() {
				// Verify Result
				result, err := kp.Verify([]byte(req.GetTextValue()), []byte(req.GetSignedText()))
				if err != nil {
					return &md.VerifyResponse{IsVerified: false}, err
				}

				// Return Result
				return &md.VerifyResponse{IsVerified: result}, nil
			}
		} else {
			return s.user.VerifyRead(), nil
		}
	}
	return &md.VerifyResponse{IsVerified: false}, fmt.Errorf("Node is not ready")
}

// Update proximity/direction/contact/properties and notify Lobby
func (s *NodeServer) Update(ctx context.Context, req *md.UpdateRequest) (*md.NoResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Check Update Request Type
		switch req.Data.(type) {
		// Update Position
		case *md.UpdateRequest_Position:
			s.user.UpdatePosition(req.GetPosition().Parameters())

		// Update Contact
		case *md.UpdateRequest_Contact:
			s.user.UpdateContact(req.GetContact())

		// Update Peer Properties
		case *md.UpdateRequest_Properties:
			s.user.UpdateProperties(req.GetProperties())
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
	if s.isReady() {
		// Validate invite
		req = s.user.SignInvite(req)

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
func (s *NodeServer) Respond(ctx context.Context, req *md.InviteResponse) (*md.NoResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Send Response
		s.client.Respond(req)

		// Update Status
		if req.Decision {
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
func (s *NodeServer) Mail(ctx context.Context, req *md.MailboxRequest) (*md.MailboxResponse, error) {
	// Verify Node is Ready
	if s.isReady() {
		// Handle Mail
		resp, serr := s.client.Mail(req)
		if serr != nil {
			s.handleError(serr)
			return nil, serr.Error
		}

		// Return Response
		return resp, nil
	}
	return nil, fmt.Errorf("Node is not ready")
}
