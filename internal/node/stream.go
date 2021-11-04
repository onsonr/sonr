package node

import api "github.com/sonr-io/core/internal/api"

// OnDecision is callback for NodeImpl for decisionEvents
func (n *Node) OnDecision(event *api.DecisionEvent, invite *api.InviteEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnDecision")
		return
	}
	n.decisionEvents <- event
	n.motor.TransmitProtocol.Outgoing(invite.GetPayload(), event.GetFrom())
}

// OnInvite is callback for NodeImpl for inviteEvents
func (n *Node) OnInvite(event *api.InviteEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnInvite")
		return
	}
	n.inviteEvents <- event
}

// OnMailbox is callback for NodeImpl for mailEvents
func (n *Node) OnMailbox(event *api.MailboxEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnMailbox")
		return
	}
	n.mailEvents <- event
}

// OnRefresh is callback for NodeImpl for refreshEvents
func (n *Node) OnRefresh(event *api.RefreshEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnRefresh")
		return
	}
	n.refreshEvents <- event
}

// OnProgress is callback for NodeImpl for progressEvents
func (n *Node) OnProgress(event *api.ProgressEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnProgress")
		return
	}
	n.progressEvents <- event
}

// OnComplete is callback for NodeImpl for completeEvents
func (n *Node) OnComplete(event *api.CompleteEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnComplete")
		return
	}
	n.completeEvents <- event
}

// OnLobbyRefresh method sends a lobby refresh event to the client.
func (s *NodeMotorStub) OnLobbyRefresh(e *Empty, stream MotorStub_OnLobbyRefreshServer) error {
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
func (s *NodeMotorStub) OnMailboxMessage(e *Empty, stream MotorStub_OnMailboxMessageServer) error {
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
func (s *NodeMotorStub) OnTransmitAccepted(e *Empty, stream MotorStub_OnTransmitAcceptedServer) error {
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
func (s *NodeMotorStub) OnTransmitDeclined(e *Empty, stream MotorStub_OnTransmitDeclinedServer) error {
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
func (s *NodeMotorStub) OnTransmitInvite(e *Empty, stream MotorStub_OnTransmitInviteServer) error {
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
func (s *NodeMotorStub) OnTransmitProgress(e *Empty, stream MotorStub_OnTransmitProgressServer) error {
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
func (s *NodeMotorStub) OnTransmitComplete(e *Empty, stream MotorStub_OnTransmitCompleteServer) error {
	for {
		select {
		case m := <-s.node.completeEvents:
			if m != nil {
				// Check Direction
				stream.Send(m)
				// Add Receiver to Recents
				err := s.node.identity.AddRecent(m.Recent())
				if err != nil {
					logger.Errorf("%s - Failed to add receiver's profile to store.", err)
					continue
				}
			}
		}
	}
}
