package motor

import (
	mv1 "go.buf.build/grpc/go/sonr-io/core/motor/v1"
)

// MotorCallback is the implementation of Callback interface
type MotorCallback interface {
	// OnRefresh is called when the LobbyProtocol is refreshed and pushes a RefreshEvent
	OnRefresh(event *mv1.OnLobbyRefreshResponse)

	// OnMailbox is called when the MailboxProtocol receives a MailboxEvent
	OnMailbox(event *mv1.OnMailboxMessageResponse)

	// OnInvite is called when the TransferProtocol receives InviteEvent
	OnInvite(event *mv1.OnTransmitInviteResponse)

	// OnDecision is called when the TransferProtocol receives a DecisionEvent
	OnDecision(event *mv1.OnTransmitDecisionResponse, invite *mv1.OnTransmitInviteResponse)

	// OnProgress is called when the TransferProtocol sends or receives a ProgressEvent
	OnProgress(event *mv1.OnTransmitProgressResponse)

	// OnTransfer is called when the TransferProtocol completes a transfer and pushes a CompleteEvent
	OnComplete(event *mv1.OnTransmitCompleteResponse)
}
