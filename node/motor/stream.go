package motor

// OnLobbyRefresh method sends a lobby refresh event to the client.
func (s *MotorStub) OnLobbyRefresh(e *Empty, stream MotorStub_OnLobbyRefreshServer) error {
	for {
		select {
		case m := <-s.refreshEvents:
			if m != nil {
				stream.Send(m)
			}
		}
	}
}

// OnMailboxMessage method sends an accepted event to the client.
func (s *MotorStub) OnMailboxMessage(e *Empty, stream MotorStub_OnMailboxMessageServer) error {
	for {
		select {
		case m := <-s.mailEvents:
			if m != nil {
				stream.Send(m)
			}
		}
	}
}

// OnTransferAccepted method sends an accepted event to the client.
func (s *MotorStub) OnTransmitAccepted(e *Empty, stream MotorStub_OnTransmitAcceptedServer) error {
	for {
		select {
		case m := <-s.decisionEvents:
			if m != nil {
				if m.Decision {
					stream.Send(m)
				}
			}
		}
	}
}

// OnTransferDeclinedmethod sends a decline event to the client.
func (s *MotorStub) OnTransmitDeclined(e *Empty, stream MotorStub_OnTransmitDeclinedServer) error {
	for {
		select {
		case m := <-s.decisionEvents:
			if m != nil {
				if !m.Decision {
					stream.Send(m)
				}
			}
		}
	}
}

// OnTransferInvite method sends an invite event to the client.
func (s *MotorStub) OnTransmitInvite(e *Empty, stream MotorStub_OnTransmitInviteServer) error {
	for {
		select {
		case m := <-s.inviteEvents:
			if m != nil {
				stream.Send(m)
			}
		}
	}
}

// OnTransferProgress method sends a progress event to the client.
func (s *MotorStub) OnTransmitProgress(e *Empty, stream MotorStub_OnTransmitProgressServer) error {
	for {
		select {
		case m := <-s.progressEvents:
			if m != nil {
				stream.Send(m)
			}
		}
	}
}

// OnTransferComplete method sends a complete event to the client.
func (s *MotorStub) OnTransmitComplete(e *Empty, stream MotorStub_OnTransmitCompleteServer) error {
	for {
		select {
		case m := <-s.completeEvents:
			if m != nil {
				// Check Direction
				stream.Send(m)
			}
		}
	}
}
