package node

import (
	"context"
	"errors"
	"sync/atomic"
)

// HostStatus is the status of the host
type HostStatus string

// Equals returns true if given SNRHostStatus matches this one
func (s HostStatus) Equals(other HostStatus) bool {
	return s == other
}

// IsIdle returns true if the SNRHostStatus != Status_IDLE
func (s HostStatus) IsIdle() bool {
	return s == Status_IDLE
}

// IsStandby returns true if the SNRHostStatus == Status_STANDBY
func (s HostStatus) IsStandby() bool {
	return s == Status_STANDBY
}

// IsReady returns true if the SNRHostStatus == Status_READY
func (s HostStatus) IsReady() bool {
	return s == Status_READY
}

// IsConnecting returns true if the SNRHostStatus == Status_CONNECTING
func (s HostStatus) IsConnecting() bool {
	return s == Status_CONNECTING
}

// IsFail returns true if the SNRHostStatus == Status_FAIL
func (s HostStatus) IsFail() bool {
	return s == Status_FAIL
}

// IsClosed returns true if the SNRHostStatus == Status_CLOSED
func (s HostStatus) IsClosed() bool {
	return s == Status_CLOSED
}

func (sm SFSM) GetCurrent() string {
	return string(sm.CurrentStatus)
}

type SFSM struct {
	CurrentStatus HostStatus
	Chn           chan bool
	flag          uint64
	States        *[]HostStatus
	StateMapping  *map[HostStatus][]HostStatus
}

// SNRHostStatus Definitions
const (
	// States
	Status_IDLE       HostStatus = "IDLE"
	Status_STANDBY    HostStatus = "STANDBY"    // Host is standby, waiting for connection
	Status_CONNECTING HostStatus = "CONNECTING" // Host is connecting
	Status_READY      HostStatus = "READY"      // Host is ready
	Status_FAIL       HostStatus = "FAILURE"    // Host failed to connect
	Status_CLOSED     HostStatus = "CLOSED"     // Host is closed
)

var (
	//state mapings
	STATE_MAPPINGS = map[HostStatus][]HostStatus{
		Status_IDLE:       {Status_STANDBY, Status_CLOSED},
		Status_STANDBY:    {Status_READY, Status_CLOSED},
		Status_CONNECTING: {Status_READY, Status_FAIL, Status_CLOSED},
		Status_READY:      {Status_STANDBY, Status_CLOSED},
		Status_CLOSED:     {Status_IDLE, Status_STANDBY},
		Status_FAIL:       {Status_CONNECTING, Status_STANDBY},
	}

	// Errors
	ErrTRANSITION = errors.New("cannot transition to state")
	ErrOPERATION  = errors.New("cannot perform operation, already active")
)

func NewFSM(ctx context.Context) *SFSM {
	states := []HostStatus{
		Status_IDLE,
		Status_READY,
		Status_CONNECTING,
		Status_FAIL,
		Status_STANDBY,
		Status_CLOSED,
	}

	return &SFSM{
		States:        &states,
		StateMapping:  &STATE_MAPPINGS,
		CurrentStatus: Status_IDLE,
	}

}

// SetState sets the host status and emits the event
func (fsm *SFSM) SetState(s HostStatus) error {
	// Check if status is changed
	if fsm.CurrentStatus == s {
		return nil
	}
	status_bucket := STATE_MAPPINGS[fsm.CurrentStatus]
	for _, status := range status_bucket {
		if status == s {
			fsm.CurrentStatus = s
			break
		}
	}

	return ErrTRANSITION
}

// NeedsWait checks if state is Resumed or Paused
// 0 -> Running
// 1 -> Paused
func (sfm *SFSM) NeedsWait() {
	<-sfm.Chn
}

// Resume tells all of goroutines to resume execution
func (fsm *SFSM) ResumeOperation() error {
	if atomic.LoadUint64(&fsm.flag) == 1 {
		close(fsm.Chn)
		atomic.StoreUint64(&fsm.flag, 0)
		return nil
	}

	return ErrOPERATION
}

// Pause tells all of goroutines to pause execution
func (fsm *SFSM) PauseOperation() error {
	if atomic.LoadUint64(&fsm.flag) == 0 {
		atomic.StoreUint64(&fsm.flag, 1)
		fsm.Chn = make(chan bool)
		return nil
	}

	return ErrOPERATION
}
