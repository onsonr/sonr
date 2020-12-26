package lifecycle

import "sync"

const (
	StateRunning = iota
	StatePaused
)

type WContext struct {
	mu    sync.Mutex
	state int
}

func NewContext() *WContext {
	return new(WContext)
}

func (w *WContext) SetState(state int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.state = state
}

func (w *WContext) State() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.state
}
