package models

import (
	"sync"
	"sync/atomic"
)

// ************************* //
// ** Callback Management ** //
// ************************* //

// Define Function Types
type OnProtobuf func([]byte)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *TransferCard)
type OnTransmitted func(data *Peer)
type OnError func(err error, method string)
type NodeCallback struct {
	Invited     OnInvite
	Refreshed   OnProtobuf
	Event       OnProtobuf
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	Transmitted OnTransmitted
	Error       OnError
}

// ***************************** //
// ** Binary State Management ** //
// ***************************** //
type state struct {
	flag uint64
	chn  chan bool
}

var (
	instance *state
	once     sync.Once
)

func GetState() *state {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &state{chn: chn}
	})

	return instance
}

// Checks rather to wait or does not need
func (c *state) NeedsWait() {
	<-c.chn
}

// Says all of goroutines to resume execution
func (c *state) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Says all of goroutines to pause execution
func (c *state) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.chn = make(chan bool)
	}
}
