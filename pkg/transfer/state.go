package transfer

import (
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/state"
)

type TransferDirection int

const (
	DirectionInbound TransferDirection = iota
	DirectionOutbound
)

const (
	// ### - Invite Method Types -
	// 1. States
	Available state.StateType = "Available"
	// Pending is state for when Peer is not Found
	Pending state.StateType = "Pending"
	// Pending is state for when Invite Succeeds
	InProgress state.StateType = "InProgress"

	// 2a. Events
	// Receiver
	InviteReceived state.EventType = "InviteReceived"
	DecisionAccept state.EventType = "DecisionAccept"
	DecisionReject state.EventType = "DecisionReject"

	// 2b. Events
	// Sender
	InviteFailed state.EventType = "InviteFailed"
	InviteShared state.EventType = "InviteShared"
	PeerAccepted state.EventType = "PeerAccepted"
	PeerRejected state.EventType = "PeerRejected"

	// 2c. Events
	// Common - Sender+Receiver
	TransferSuccess state.EventType = "TransferSuccess"
	TransferFail    state.EventType = "TransferFail"
)

type TransferSessionContext struct {
	To        peer.ID
	From      peer.ID
	Direction TransferDirection
	Decision  bool
	Invite    *InviteRequest
	Transfer  *common.Payload
	LastEvent state.EventType
}

// initStateMachine initializes the state machine
func (p *TransferProtocol) initStateMachine() {
	p.state = state.StateMachine{
		States: state.States{
			state.Default: state.State{
				Events: state.Events{
					InviteReceived: Pending,
					InviteFailed:   Available,
					InviteShared:   Pending,
				},
			},
			Available: state.State{
				Action: &TransferInviteAction{},
				Events: state.Events{
					InviteReceived: Pending,
					InviteFailed:   Available,
					InviteShared:   Pending,
				},
			},
			Pending: state.State{
				Action: &TransferPendingAction{},
				Events: state.Events{
					PeerAccepted:   InProgress,
					PeerRejected:   Available,
					DecisionAccept: InProgress,
					DecisionReject: Available,
				},
			},
			InProgress: state.State{
				Action: &TransferInProgressAction{},
				Events: state.Events{
					TransferSuccess: Available,
					TransferFail:    Available,
				},
			},
		},
	}
}
