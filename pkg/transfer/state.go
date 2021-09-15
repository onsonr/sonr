package transfer

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
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
	// UnknownPeer is state for when Peer is not Found
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
	InviteShared state.EventType = "InviteShared"
	PeerAccepted state.EventType = "PeerAccepted"
	PeerRejected state.EventType = "PeerRejected"

	// 2c. Events
	// Common - Sender+Receiver
	TransferSuccess state.EventType = "TransferSuccess"
	TransferFail    state.EventType = "TransferFail"
)

type TransferInviteContext struct {
	Direction TransferDirection
	Invite    *InviteRequest
	To        peer.ID
	From      peer.ID
	Decision  bool
}

type InviteTransferAction struct{}

func (a *InviteTransferAction) Execute(eventCtx state.EventContext) state.EventType {
	invite := eventCtx.(*TransferInviteContext)
	fmt.Println(invite.To.String())
	return InviteShared
}
