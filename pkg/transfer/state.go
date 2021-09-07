package transfer

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/tools/state"
)

const (
	// ### - Invite Method Types -
	// 1. States
	FindingPeer state.StateType = "FindingPeer"
	// UnknownPeer is state for when Peer is not Found
	UnknownPeer state.StateType = "UnknownPeer"

	// Pending is state for when Invite Succeeds
	Pending state.StateType = "Pending"

	// 2. Events
	// FindPeer Method Events
	FailFindPeer    state.EventType = "FailFindPeer"
	SucceedFindPeer state.EventType = "SucceedFindPeer"

	// SendInvite Method Events
	FailSendInvite    state.EventType = "FailSendInvite"
	SucceedSendInvite state.EventType = "SucceedSendInvite"
)

type TransferInviteContext struct {
	invite   *InviteRequest
	to       peer.ID
	decision bool
}

type InviteTransferAction struct{}

func (a *InviteTransferAction) Execute(eventCtx state.EventContext) state.EventType {
	invite := eventCtx.(*TransferInviteContext)
	fmt.Println(invite.to.String())
	return SucceedFindPeer
}
