package host

import (
	"context"

	"github.com/looplab/fsm"
	"github.com/sonr-io/sonr/pkg/config"
)

type SFSM struct {
	start     string
	states    *[]fsm.EventDesc
	callbacks *map[string]fsm.Callback
	fsm       *fsm.FSM
}

func New(ctx context.Context, c *config.Config) *SFSM {
	initial_state := ""
	events := c.States
	callbacks := c.StateCallbacks

	fsm.NewFSM(
		initial_state,
		events,
		callbacks,
	)

	return &SFSM{
		start:     initial_state,
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
