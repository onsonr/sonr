package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/sonr-io/core/pkg/common"
)

// IsIncoming returns true if the event is incoming
func (e *CompleteEvent) IsIncoming() bool {
	return e.GetDirection() == common.Direction_INCOMING
}

// IsOutgoing returns true if the event is outgoing
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

// Title returns the title of the event
func (d *CompleteEvent) Title() string {
	return fmt.Sprintf("[Transfer-Complete] from %s at %v", d.GetFrom().GetProfile().GetSName(), time.Now())
}

// Message returns the message of the event
func (d *CompleteEvent) Message() string {
	paths := ""
	for _, v := range d.GetPayload().GetItems() {
		paths += fmt.Sprintf("\n-\t%s", v.GetFile().GetPath())
	}
	return fmt.Sprintf("Size: %v \n Paths: %s", d.GetPayload().GetSize(), paths)
}

// Title returns the title of the event
func (d *DecisionEvent) Title() string {
	return fmt.Sprintf("[Transfer-Decision] from %s", d.GetFrom().GetProfile().GetSName())
}

// Message returns the message of the event
func (d *DecisionEvent) Message() string {
	return fmt.Sprintf("Result: %v", d.GetDecision())
}

// Title returns the title of the event
func (d *InviteEvent) Title() string {
	fname := d.GetFrom().GetProfile().GetFirstName()
	lname := d.GetFrom().GetProfile().GetLastName()
	sname := d.GetFrom().GetProfile().GetSName()
	platform := strings.ToUpper(d.GetFrom().GetDevice().GetOs())
	return fmt.Sprintf("[Transfer-Invite] from %s %s (%s) on (%s)", fname, lname, sname, platform)
}

// Message returns the message of the event
func (d *InviteEvent) Message() string {
	fcount := d.GetPayload().FileCount()
	ucount := d.GetPayload().URLCount()
	tcount := fcount + ucount
	countStr := fmt.Sprintf("%d, (Files) %d, (Urls) %d", tcount, fcount, ucount)
	mimes := ""
	for _, v := range d.GetPayload().GetItems() {
		mimes += fmt.Sprintf("%s, ", v.GetMime().GetValue())
	}
	return fmt.Sprintf("Count: %s \nMimes: %s \nSize: %v", countStr, mimes, d.GetPayload().GetSize())
}
