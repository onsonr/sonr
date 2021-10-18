package node

import "github.com/pterm/pterm"

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
					if s.isTerminal {
						PrintTerminal(m.Title(), m.Message())
					}
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
					if s.isTerminal {
						PrintTerminal(m.Title(), m.Message())
					}
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
				if s.isTerminal {
					PrintTerminal(m.Title(), m.Message())
				}
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
				if s.isTerminal {
					PrintTerminal(m.Title(), m.Message())
				}
				// Check Direction
				stream.Send(m)
				// Add Receiver to Recents
				err := s.node.AddRecent(m.Recent())
				if err != nil {
					logger.Error("Failed to add receiver's profile to store.", err)
				}

				// Add Payload to History
				if m.IsIncoming() {
					err = s.node.AddHistory(m.GetPayload())
					if err != nil {
						logger.Error("Failed to add payload to store.", err)
					}
				}
			}
		}
	}
}

func PrintTerminal(title string, msg string) {
	// Print a section with level one.
	pterm.DefaultSection.Println(title)
	// Print placeholder.
	pterm.Info.Println(msg)

}
