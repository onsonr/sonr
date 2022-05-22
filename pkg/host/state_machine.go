package host

import (
	"context"

	"github.com/sonr-io/sonr/pkg/config"
)

// HostStatus is the status of the host
type HostStatus string

// Equals returns true if given SNRHostStatus matches this one
func (s HostStatus) Equals(other HostStatus) bool {
	return s == other
}

// IsNotIdle returns true if the SNRHostStatus != Status_IDLE
func (s HostStatus) IsNotIdle() bool {
	return s != Status_IDLE
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

// String returns the string representation of the SNRHostStatus
func (s HostStatus) String() string {
	return s.String()
}

func (sm *SFSM) GetCurrent() string {
	return string(sm.Current)
}

type SFSM struct {
	States       *[]HostStatus
	StateMapping *map[HostStatus][]HostStatus
	Current      HostStatus
}

// SNRHostStatus Definitions
const (
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
	}
)

func New(ctx context.Context, c *config.Config) *SFSM {
	states := []HostStatus{
		Status_IDLE,
		Status_READY,
		Status_CONNECTING,
		Status_FAIL,
		Status_STANDBY,
	}

	return &SFSM{
		States:       &states,
		StateMapping: &STATE_MAPPINGS,
		Current:      Status_IDLE,
	}

}
