package node

import (
	"github.com/sonr-io/core/types/go/node/motor/v1"
)

// CallbackImpl is the implementation of Callback interface
type CallbackImpl interface {
	// OnRefresh is called when the LobbyProtocol is refreshed and pushes a RefreshEvent
	OnRefresh(event *motor.OnLobbyRefreshResponse)

	// OnMailbox is called when the MailboxProtocol receives a MailboxEvent
	OnMailbox(event *motor.OnMailboxMessageResponse)

	// OnInvite is called when the TransferProtocol receives InviteEvent
	OnInvite(event *motor.OnTransmitInviteResponse)

	// OnDecision is called when the TransferProtocol receives a DecisionEvent
	OnDecision(event *motor.OnTransmitDecisionResponse, invite *motor.OnTransmitInviteResponse)

	// OnProgress is called when the TransferProtocol sends or receives a ProgressEvent
	OnProgress(event *motor.OnTransmitProgressResponse)

	// OnTransfer is called when the TransferProtocol completes a transfer and pushes a CompleteEvent
	OnComplete(event *motor.OnTransmitCompleteResponse)
}
