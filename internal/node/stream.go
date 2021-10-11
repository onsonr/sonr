package node

import api "github.com/sonr-io/core/internal/api"

// OnLobbyRefresh method sends a lobby refresh event to the client.
func (s *ClientNodeStub) OnLobbyRefresh(e *Empty, stream ClientService_OnLobbyRefreshServer) error {
	for {
		select {
		case m := <-s.node.refreshEvents:
			if m != nil {
				stream.Send(m)
			}
		}
	}
}

// OnMailboxMessage method sends an accepted event to the client.
func (s *ClientNodeStub) OnMailboxMessage(e *Empty, stream ClientService_OnMailboxMessageServer) error {
	for {
		select {
		case m := <-s.node.mailEvents:
			if m != nil {
				stream.Send(m)
			}
		}
	}
}

// OnTransferAccepted method sends an accepted event to the client.
func (s *ClientNodeStub) OnTransferAccepted(e *Empty, stream ClientService_OnTransferAcceptedServer) error {
	for {
		select {
		case m := <-s.node.decisionEvents:
			if m != nil {
				if m.Decision {
					stream.Send(m)
				}
			}
		}
	}
}

// OnTransferDeclinedmethod sends a decline event to the client.
func (s *ClientNodeStub) OnTransferDeclined(e *Empty, stream ClientService_OnTransferDeclinedServer) error {
	for {
		select {
		case m := <-s.node.decisionEvents:
			if m != nil {
				if !m.Decision {
					stream.Send(m)
				}
			}
		}
	}
}

// OnTransferInvite method sends an invite event to the client.
func (s *ClientNodeStub) OnTransferInvite(e *Empty, stream ClientService_OnTransferInviteServer) error {
	for {
		select {
		case m := <-s.node.inviteEvents:
			if m != nil {
				stream.Send(m)
			}
		}
	}
}

// OnTransferProgress method sends a progress event to the client.
func (s *ClientNodeStub) OnTransferProgress(e *Empty, stream ClientService_OnTransferProgressServer) error {
	for {
		select {
		case m := <-s.node.progressEvents:
			if m != nil {
				stream.Send(m)
			}
		}
	}
}

// OnTransferComplete method sends a complete event to the client.
func (s *ClientNodeStub) OnTransferComplete(e *Empty, stream ClientService_OnTransferCompleteServer) error {
	for {
		select {
		case m := <-s.node.completeEvents:
			if m != nil {
				// Check Direction
				stream.Send(m)
				if m.Direction == api.CompleteEvent_INCOMING {
					// Add Sender to Recents
					err := s.node.AddRecent(m.GetFrom().GetProfile())
					if err != nil {
						logger.Child("Client").Error("Failed to add sender's profile to store.", err)
					}
				} else {
					// Add Receiver to Recents
					err := s.node.AddRecent(m.GetTo().GetProfile())
					if err != nil {
						logger.Child("Client").Error("Failed to add receiver's profile to store.", err)
					}
				}
			}
		}
	}
}
