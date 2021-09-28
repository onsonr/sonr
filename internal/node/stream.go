package node

// OnNodeStatus method sends a status event to the client.
func (n *NodeRPCService) OnNodeStatus(e *Empty, stream NodeService_OnNodeStatusServer) error {
	for {
		select {
		case m := <-n.statusEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}

// OnLobbyRefresh method sends a lobby refresh event to the client.
func (n *NodeRPCService) OnLobbyRefresh(e *Empty, stream NodeService_OnLobbyRefreshServer) error {
	for {
		select {
		case m := <-n.exchangeEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}

// OnTransferAccepted method sends an accepted event to the client.
func (n *NodeRPCService) OnTransferAccepted(e *Empty, stream NodeService_OnTransferAcceptedServer) error {
	for {
		select {
		case m := <-n.decisionEvents:
			if m != nil {
				if m.Decision {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}

// OnTransferDeclinedmethod sends a decline event to the client.
func (n *NodeRPCService) OnTransferDeclined(e *Empty, stream NodeService_OnTransferDeclinedServer) error {
	for {
		select {
		case m := <-n.decisionEvents:
			if m != nil {
				if !m.Decision {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}

// OnTransferInvite method sends an invite event to the client.
func (n *NodeRPCService) OnTransferInvite(e *Empty, stream NodeService_OnTransferInviteServer) error {
	for {
		select {
		case m := <-n.inviteEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}

// OnTransferProgress method sends a progress event to the client.
func (n *NodeRPCService) OnTransferProgress(e *Empty, stream NodeService_OnTransferProgressServer) error {
	for {
		select {
		case m := <-n.progressEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}

// OnTransferComplete method sends a complete event to the client.
func (n *NodeRPCService) OnTransferComplete(e *Empty, stream NodeService_OnTransferCompleteServer) error {
	for {
		select {
		case m := <-n.completeEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}
