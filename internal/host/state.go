package host

// SNRHostStatus is the status of the host
type SNRHostStatus int

// SNRHostStatus Definitions
const (
	Status_IDLE       SNRHostStatus = iota // Host is idle, default state
	Status_STANDBY                         // Host is standby, waiting for connection
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

// IsStandby returns true if the SNRHostStatus == Status_STANDBY
func (s SNRHostStatus) IsStandby() bool {
	return s == Status_STANDBY
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
	case Status_STANDBY:
		return "STANDBY"
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

// SetStatus sets the host status and emits the event
func (h *SNRHost) SetStatus(s SNRHostStatus) {
	// Check if status is changed
	if h.status == s {
		logger.Info("SetStatus: Same status provided, " + s.String())
		return
	}

	// Update Status
	h.status = s
	h.emitter.Emit(Event_STATUS, s)
}
