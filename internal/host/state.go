package host

import (
	"time"

	"github.com/kataras/golog"
)

// SNRHostStatus is the status of the host
type SNRHostStatus int

// SNRHostStatus Definitions
const (
	Status_IDLE       SNRHostStatus = iota // Host is idle, default state
	Status_CONNECTING                      // Host is connecting
	Status_READY                           // Host is ready
	Status_FAIL                            // Host failed to connect
	Status_CLOSED                          // Host is closed
)

// Equals returns true if given SNRHostStatus matches this one
func (s SNRHostStatus) Equals(other SNRHostStatus) bool {
	return s == other
}

// IsNotIdle returns true if the SNRHostStatus != Status_IDLE
func (s SNRHostStatus) IsNotIdle() bool {
	return s != Status_IDLE
}

// IsReady returns true if the SNRHostStatus == Status_READY
func (s SNRHostStatus) IsReady() bool {
	return s == Status_READY
}

// IsConnecting returns true if the SNRHostStatus == Status_CONNECTING
func (s SNRHostStatus) IsConnecting() bool {
	return s == Status_CONNECTING
}

// IsFail returns true if the SNRHostStatus == Status_FAIL
func (s SNRHostStatus) IsFail() bool {
	return s == Status_FAIL
}

// IsClosed returns true if the SNRHostStatus == Status_CLOSED
func (s SNRHostStatus) IsClosed() bool {
	return s == Status_CLOSED
}

// String returns the string representation of the SNRHostStatus
func (s SNRHostStatus) String() string {
	switch s {
	case Status_IDLE:
		return "IDLE"
	case Status_CONNECTING:
		return "CONNECTING"
	case Status_READY:
		return "READY"
	case Status_FAIL:
		return "FAIL"
	case Status_CLOSED:
		return "CLOSED"
	}
	return "UNKNOWN"
}

// Event Emitter Constants
const (
	Event_STATUS            = "host-status" // Host is ready to accept connections
	DefaultEventLoopTimeout = time.Second * 30
)

// StatusFunc is a function that handles the current status of the host
type StatusFunc func(status SNRHostStatus)

// EventLoopOption is a function that modifies the node options.
type EventLoopOption func(eventLoopOptions)

// WithDoneChannel sets the done channel
func WithMiddlewareFunc(f StatusFunc) EventLoopOption {
	return func(o eventLoopOptions) {
		o.middlewares = append(o.middlewares, f)
	}
}

// WithDoneChannel sets the done channel
func WithDoneChannel(dch chan bool) EventLoopOption {
	return func(o eventLoopOptions) {
		o.doneCh = dch
	}
}

// WithTargetEvent sets the target event
func WithTargetEvent(e SNRHostStatus) EventLoopOption {
	return func(o eventLoopOptions) {
		o.target = e
	}
}

// WithTimeout sets the timeout
func WithTimeout(t time.Duration) EventLoopOption {
	return func(o eventLoopOptions) {
		o.timeout = t
	}
}

// eventLoopOptions is a collection of options for the node.
type eventLoopOptions struct {
	target      SNRHostStatus
	doneCh      chan bool
	host        *SNRHost
	middlewares []StatusFunc
	timeout     time.Duration
}

// defaultEventLoopOptions returns the default node options.
func defaultEventLoopOptions(h *SNRHost) eventLoopOptions {
	return eventLoopOptions{
		host:        h,
		target:      Status_IDLE,
		doneCh:      make(chan bool),
		middlewares: make([]StatusFunc, 0),
		timeout:     DefaultEventLoopTimeout,
	}
}

// createEventLoop is a helper function to handle events
func createEventLoop(h *SNRHost, options ...EventLoopOption) {
	// Create the options
	opts := defaultEventLoopOptions(h)
	for _, o := range options {
		o(opts)
	}

	// Create the event loop
	for {
		select {
		// Handle Timeout
		case <-time.After(opts.timeout):
			logger.Error("Timeout for EventLoop reached \n", golog.Fields{
				"Target":  opts.target.String(),
				"Timeout": opts.timeout,
			})
			return
		}
	}
}

// Handle handles the state.Event and applies the appropriate action
func (eo eventLoopOptions) Handle(s SNRHostStatus) {
	logger.Info("Handling EventLoop Status: ", s)
	// Check if we have a target
	if eo.target.IsNotIdle() {
		// Update channels
		logger.Info("Handling Done Channel and Middlewares for Target: ", eo.target.String())
		if eo.host.status == eo.target {
			eo.doneCh <- true
		} else {
			eo.doneCh <- false
		}

		// Call on target middlewares
		for _, m := range eo.middlewares {
			if eo.host.status == s {
				m(s)
			}
		}
		return
	}
	// Update channels
	logger.Info("Handling Done Channel and Middlewares for 'Status_READY'")
	if s.IsReady() {
		eo.doneCh <- true
	} else if s.IsFail() {
		eo.doneCh <- false
	}

	// Call on ALL middlewares
	for _, m := range eo.middlewares {
		m(s)
	}
	return
}
