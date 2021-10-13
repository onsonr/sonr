package api

import "github.com/sonr-io/core/internal/common"

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
