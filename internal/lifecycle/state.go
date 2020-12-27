package lifecycle

import "sync"

const (
	StateRunning = iota
	StatePaused
)

type Worker struct {
	mu    sync.Mutex
	state int
}

func (w *Worker) SetState(state int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.state = state
}

func (w *Worker) State() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.state
}
