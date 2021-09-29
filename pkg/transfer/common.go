package transfer

import (
	"time"

	"github.com/sonr-io/core/internal/common"
)

// ITEM_INTERVAL is the interval in which progress events are emitted
const ITEM_INTERVAL = 25

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *common.DecisionEvent {
	return &common.DecisionEvent{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *common.InviteEvent {
	return &common.InviteEvent{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}
