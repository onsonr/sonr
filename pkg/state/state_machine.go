package state

import (
	"context"

	"github.com/looplab/fsm"
	"github.com/sonr-io/sonr/pkg/config"
)

// HostStatus is the status of the host
type HostStatus int

// SNRHostStatus Definitions
const (
	Status_IDLE       HostStatus = iota // Host is idle, default state
	Status_STANDBY                      // Host is standby, waiting for connection
	Status_CONNECTING                   // Host is connecting
	Status_READY                        // Host is ready
	Status_FAIL                         // Host failed to connect
	Status_CLOSED                       // Host is closed
)

type SFSM struct {
	name      string
	states    *[]fsm.EventDesc
	callbacks *map[string]fsm.Callback
	fsm       *fsm.FSM
}

func New(ctx context.Context, c *config.Config) *SFSM {
	name := ""
	if c.Role == config.Role_HIGHWAY {
		name = "highway"
	} else {
		name = "motor"
	}
	events := c.States
	callbacks := c.StateCallbacks

	fsm.NewFSM(
		name,
		events,
		callbacks,
	)

	return &SFSM{
		name:      name,
		states:    &events,
		callbacks: &callbacks,
	}
}

func (sm *SFSM) Transition(state string) *error {
	err := sm.fsm.Event(state)

	return &err
}

func (sm *SFSM) GetCurrent() string {
	return sm.fsm.Current()
}
