package motor

import api "github.com/sonr-io/core/node/api"

// OnDecision is callback for NodeImpl for decisionEvents
func (n *MotorStub) OnDecision(event *api.DecisionEvent, invite *api.InviteEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnDecision")
		return
	}
	n.decisionEvents <- event
	n.TransmitProtocol.Outgoing(invite.GetPayload(), event.GetFrom())
}

// OnInvite is callback for NodeImpl for inviteEvents
func (n *MotorStub) OnInvite(event *api.InviteEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnInvite")
		return
	}
	n.inviteEvents <- event
}

// OnMailbox is callback for NodeImpl for mailEvents
func (n *MotorStub) OnMailbox(event *api.MailboxEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnMailbox")
		return
	}
	n.mailEvents <- event
}

// OnRefresh is callback for NodeImpl for refreshEvents
func (n *MotorStub) OnRefresh(event *api.RefreshEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnRefresh")
		return
	}
	n.refreshEvents <- event
}

// OnProgress is callback for NodeImpl for progressEvents
func (n *MotorStub) OnProgress(event *api.ProgressEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnProgress")
		return
	}
	n.progressEvents <- event
}

// OnComplete is callback for NodeImpl for completeEvents
func (n *MotorStub) OnComplete(event *api.CompleteEvent) {
	if event == nil {
		logger.Warn("Received nil event: OnComplete")
		return
	}
	n.completeEvents <- event
}
