package common

import "github.com/sonr-io/core/tools/state"

// Transfer Emission Events
const (
	Event_INVITED      = "transfer-invited"
	Event_RESPONDED    = "transfer-responded"
	Event_PROGRESS     = "transfer-progress"
	Event_COMPLETED    = "transfer-completed"
	Event_LIST_REFRESH = "lobby-list-refresh"
)

type EventChannel struct {
	// Emitter receiving events
	emitter *state.Emitter

	// TransferProtocol - decisionEvents
	decisionEvents chan *DecisionEvent

	// LobbyProtocol - refreshEvents
	refreshEvents chan *RefreshEvent

	// MailboxProtocol - mailEvents
	mailEvents chan *MailboxEvent

	// TransferProtocol - inviteEvents
	inviteEvents chan *InviteEvent

	// TransferProtocol - progressEvents
	progressEvents chan *ProgressEvent

	// TransferProtocol - completeEvents
	completeEvents chan *CompleteEvent
}
