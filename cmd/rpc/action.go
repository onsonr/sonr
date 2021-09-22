package main

// // Action method handles miscellaneous actions for node
// func (s *NodeServer) Action(ctx context.Context, req *data.ActionRequest) (*data.NoResponse, error) {
// 	// Check Action
// 	switch req.Action {
// 	case data.Action_PING:
// 		// Ping
// 		s.actionResponses <- &data.ActionResponse{
// 			Success: true,
// 			Action:  data.Action_PING,
// 		}
// 	case data.Action_LOCATION:
// 		// Location
// 		s.actionResponses <- &data.ActionResponse{
// 			Success: true,
// 			Action:  data.Action_LOCATION,
// 			Data: &data.ActionResponse_Location{
// 				Location: s.account.CurrentDevice().GetLocation(),
// 			},
// 		}
// 	case data.Action_URL_LINK:
// 		// URL Link
// 		s.actionResponses <- &data.ActionResponse{
// 			Success: true,
// 			Action:  data.Action_URL_LINK,
// 			Data: &data.ActionResponse_UrlLink{
// 				UrlLink: data.NewURLLink(req.GetData()),
// 			},
// 		}
// 	case data.Action_PAUSE:
// 		// Pause
// 		s.state = data.Lifecycle_PAUSED
// 		s.client.Lifecycle(s.state, s.local)
// 		data.GetState().Pause()
// 		s.actionResponses <- &data.ActionResponse{
// 			Success: true,
// 			Action:  data.Action_PAUSE,
// 			Data: &data.ActionResponse_Lifecycle{
// 				Lifecycle: s.state,
// 			},
// 		}
// 	case data.Action_RESUME:
// 		// Resume
// 		s.state = data.Lifecycle_ACTIVE
// 		s.client.Lifecycle(s.state, s.local)
// 		data.GetState().Resume()

// 		s.actionResponses <- &data.ActionResponse{
// 			Success: true,
// 			Action:  data.Action_RESUME,
// 			Data: &data.ActionResponse_Lifecycle{
// 				Lifecycle: s.state,
// 			},
// 		}
// 	case data.Action_STOP:
// 		// Stop
// 		s.state = data.Lifecycle_STOPPED
// 		s.client.Lifecycle(s.state, s.local)
// 		s.actionResponses <- &data.ActionResponse{
// 			Success: true,
// 			Action:  data.Action_STOP,
// 			Data: &data.ActionResponse_Lifecycle{
// 				Lifecycle: s.state,
// 			},
// 		}
// 	case data.Action_LIST_LINKERS:
// 		// List Linkers
// 		s.actionResponses <- &data.ActionResponse{
// 			Success: true,
// 			Action:  data.Action_LIST_LINKERS,
// 			Data: &data.ActionResponse_Linkers{
// 				Linkers: s.local.ListLinkers(),
// 			},
// 		}
// 	default:
// 		return nil, fmt.Errorf("Action: %s not supported", req.Action)
// 	}
// 	return &data.NoResponse{}, nil
// }

// // Sign method signs data with device's private key
// func (s *NodeServer) Sign(ctx context.Context, req *data.AuthRequest) (*data.NoResponse, error) {
// 	s.authResponses <- s.account.SignAuth(req)
// 	return &data.NoResponse{}, nil
// }

// // Link method starts device linking channel
// func (s *NodeServer) Link(ctx context.Context, req *data.LinkRequest) (*data.NoResponse, error) {
// 	req = s.account.CurrentDevice().SignLink(req)

// 	// Check Link Request Type
// 	resp, err := s.client.Link(req, s.local)
// 	if err != nil {
// 		s.handleError(err)
// 		return nil, err.Error
// 	}

// 	// Return Link Response
// 	s.linkResponses <- resp
// 	return &data.NoResponse{}, nil
// }

// // Verify validates device Keys
// func (s *NodeServer) Verify(ctx context.Context, req *data.VerifyRequest) (*data.NoResponse, error) {
// 	// Get Key Pair
// 	kp := s.account.AccountKeys()

// 	// Check Request Type
// 	if req.GetType() == data.VerifyRequest_VERIFY {
// 		// Check type and Verify
// 		if req.IsBuffer() {
// 			// Verify Result
// 			result, err := kp.Verify(req.GetBufferValue(), req.GetSignedBuffer())
// 			if err != nil {
// 				s.verifyResponses <- &data.VerifyResponse{Success: false}
// 			}

// 			// Return Result
// 			s.verifyResponses <- &data.VerifyResponse{Success: result}
// 			return &data.NoResponse{}, nil
// 		} else if req.IsString() {
// 			// Verify Result
// 			result, err := kp.Verify([]byte(req.GetTextValue()), []byte(req.GetSignedText()))
// 			if err != nil {
// 				s.verifyResponses <- &data.VerifyResponse{Success: false}
// 				return &data.NoResponse{}, nil
// 			}

// 			// Return Result
// 			s.verifyResponses <- &data.VerifyResponse{Success: result}
// 			return &data.NoResponse{}, nil
// 		}
// 	}
// 	s.verifyResponses <- s.account.VerifyRead()
// 	return &data.NoResponse{}, nil
// }

// // Update proximity/direction/contact/properties and notify Lobby
// func (s *NodeServer) Update(ctx context.Context, req *data.UpdateRequest) (*data.NoResponse, error) {
// 	// Check Update Request Type
// 	switch req.Data.(type) {
// 	// Update Position
// 	case *data.UpdateRequest_Position:
// 		s.account.CurrentDevice().UpdatePosition(req.GetPosition().Parameters())

// 	// Update Contact
// 	case *data.UpdateRequest_Contact:
// 		s.account.UpdateContact(req.GetContact())

// 	// Update Peer Properties
// 	case *data.UpdateRequest_Properties:
// 		s.account.CurrentDevice().UpdateProperties(req.GetProperties())
// 	}

// 	// Notify Local Lobby
// 	err := s.client.Update(s.local)
// 	if err != nil {
// 		s.handleError(err)
// 		return nil, err.Error
// 	}

// 	// Return Blank Response
// 	return &data.NoResponse{}, nil
// }

// // Invite pushes Invite request to Peer
// func (s *NodeServer) Invite(ctx context.Context, req *data.InviteRequest) (*data.NoResponse, error) {
// 	// Validate invite
// 	req = s.account.SignInvite(req)

// 	// Send Invite
// 	err := s.client.Invite(req, s.local)
// 	if err != nil {
// 		s.handleError(err)
// 		return nil, err.Error
// 	}

// 	// Return Blank Response
// 	return &data.NoResponse{}, nil
// }

// // Respond handles a respond request
// func (s *NodeServer) Respond(ctx context.Context, req *data.DecisionRequest) (*data.NoResponse, error) {
// 	// Send Response
// 	s.client.Respond(req.ToResponse())

// 	// Return Blank Response
// 	return nil, nil
// }

// // Mail method handles a mail request
// func (s *NodeServer) Mail(ctx context.Context, req *data.MailboxRequest) (*data.NoResponse, error) {
// 	// Handle Mail
// 	resp, serr := s.client.Mail(req)
// 	if serr != nil {
// 		s.handleError(serr)
// 		return nil, serr.Error
// 	}

// 	s.mailboxResponses <- resp
// 	// Return Response
// 	return &data.NoResponse{}, nil
// }
