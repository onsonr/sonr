package api

import (
	"fmt"

	"github.com/sonr-io/core/internal/common"
)

func (e *CompleteEvent) IsIncoming() bool {
	return e.GetDirection() == common.Direction_INCOMING
}

func (e *CompleteEvent) IsOutgoing() bool {
	return e.GetDirection() == common.Direction_OUTGOING
}

// Recent returns the profile of CompleteEvent by Direction
func (e *CompleteEvent) Recent() *common.Profile {
	if e.Direction == common.Direction_INCOMING {
		return e.GetFrom().GetProfile()
	}
	return e.GetTo().GetProfile()
}

func (d *CompleteEvent) Title() string {
	return fmt.Sprintf("Completed Transfer from %s", d.GetFrom().GetProfile().GetSName())
}

func (d *CompleteEvent) Message() string {
	return fmt.Sprintf("Size: %v", d.GetPayload().GetSize())
}

func (d *DecisionEvent) Title() string {
	return fmt.Sprintf("Got Decision from %s", d.GetFrom().GetProfile().GetSName())
}

func (d *DecisionEvent) Message() string {
	return fmt.Sprintf("Result: %v", d.GetDecision())
}

func (d *InviteEvent) Title() string {
	return fmt.Sprintf("Got Invite from %s", d.GetFrom().GetProfile().GetSName())
}

func (d *InviteEvent) Message() string {
	return fmt.Sprintf("Payload: %v", d.GetPayload().String())
}
