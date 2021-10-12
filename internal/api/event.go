package api

import "github.com/sonr-io/core/internal/common"

func (e *CompleteEvent) IsIncoming() bool {
	return e.GetDirection() == CompleteEvent_INCOMING
}

func (e *CompleteEvent) IsOutgoing() bool {
	return e.GetDirection() == CompleteEvent_OUTGOING
}

// Recent returns the profile of CompleteEvent by Direction
func (e *CompleteEvent) Recent() *common.Profile {
	if e.Direction == CompleteEvent_INCOMING {
		return e.GetFrom().GetProfile()
	}
	return e.GetTo().GetProfile()
}
