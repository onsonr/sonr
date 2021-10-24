package app

import (
	"sync"
	"sync/atomic"
)

var (
	instance *state
	once     sync.Once
)

// state is the internal state of the API
type state struct {
	flag uint64
	chn  chan bool
}

// GetState returns the current state of the API
func GetState() *state {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &state{chn: chn}
	})
	return instance
}

// HasStarted returns true if the Node has been started
func HasStarted() bool {
	return Node != nil
}

// NeedsWait Checks rather to wait or does not need
func (c *state) NeedsWait() {
	<-c.chn
}

// Resume tells all of goroutines to resume execution
func (c *state) Resume() {
	if HasStarted() {
		if atomic.LoadUint64(&c.flag) == 1 {
			close(c.chn)
			atomic.StoreUint64(&c.flag, 0)
		}
	}
}

// Pause tells all of goroutines to pause execution
func (c *state) Pause() {
	if HasStarted() {
		if atomic.LoadUint64(&c.flag) == 0 {
			atomic.StoreUint64(&c.flag, 1)
			c.chn = make(chan bool)
		}
	}
}
