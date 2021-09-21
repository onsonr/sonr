package node

import (
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/state"
)

// OnDecision method sends a decision event to the client.
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
		state.GetState().NeedsWait()
	}
}

// OnLocalJoin method sends a join event to the client.
func (n *NodeRPCService) OnLocalJoin(e *Empty, stream NodeService_OnLocalJoinServer) error {
	for {
		select {
		case m := <-n.exchangeEvents:
			if m != nil {
				if m.GetType() == common.LobbyEvent_JOIN {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnLocalJoin method sends a join event to the client.
func (n *NodeRPCService) OnLocalUpdate(e *Empty, stream NodeService_OnLocalUpdateServer) error {
	for {
		select {
		case m := <-n.exchangeEvents:
			if m != nil {
				if m.GetType() == common.LobbyEvent_UPDATE {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnLocalExit method sends a join event to the client.
func (n *NodeRPCService) OnLocalExit(e *Empty, stream NodeService_OnLocalExitServer) error {
	for {
		select {
		case m := <-n.exchangeEvents:
			if m != nil {
				if m.GetType() == common.LobbyEvent_EXIT {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnDecision-Accepted method sends a decision event to the client.
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
		state.GetState().NeedsWait()
	}
}

// OnDecision-Declined method sends a decision event to the client.
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
		state.GetState().NeedsWait()
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
		state.GetState().NeedsWait()
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
		state.GetState().NeedsWait()
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
		state.GetState().NeedsWait()
	}
}
