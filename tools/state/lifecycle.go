package state

import (
	"sync"
	"sync/atomic"
)

// ** ─── lifecycle MANAGEMENT ────────────────────────────────────────────────────────
type lifecycle struct {
	flag uint64
	chn  chan bool
}

var (
	instance *lifecycle
	once     sync.Once
)

func GetState() *lifecycle {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &lifecycle{chn: chn}
	})

	return instance
}

// Checks rather to wait or does not need
func (c *lifecycle) NeedsWait() {
	<-c.chn
}

// Says all of goroutines to resume execution
func (c *lifecycle) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Says all of goroutines to pause execution
func (c *lifecycle) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.chn = make(chan bool)
	}
}
