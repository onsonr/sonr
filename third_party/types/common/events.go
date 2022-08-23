package common

import "github.com/kataras/go-events"

const (
	// ON_REFRESH is an event type for when the LobbyProtocol is refreshed
	ON_REFRESH = events.EventName("on_refresh")

	// ON_MAILBOX is an event type for when the MailboxProtocol receives a MailboxEvent
	ON_MAILBOX = events.EventName("on_mailbox")

	// ON_INVITE is an event type for when the TransferProtocol receives InviteEvent
	ON_INVITE = events.EventName("on_invite")

	// ON_DECISION is an event type for when the TransferProtocol receives a DecisionEvent
	ON_DECISION = events.EventName("on_decision")

	// ON_PROGRESS is an event type for when the TransferProtocol sends or receives a ProgressEvent
	ON_PROGRESS = events.EventName("on_progress")

	// ON_COMPLETE is an event type for when the TransferProtocol completes a transfer and pushes a CompleteEvent
	ON_COMPLETE = events.EventName("on_complete")
)
