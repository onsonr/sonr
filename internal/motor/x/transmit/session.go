package transmit

import (
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	dv1 "github.com/sonr-io/sonr/internal/motor/x/discover/v1"
	v1 "github.com/sonr-io/sonr/internal/motor/x/transmit/v1"
	"github.com/sonr-io/sonr/pkg/host"
)

// NewInSession creates a new Session from the given payload with Incoming direction.
func NewInSession(payload *v1.Payload, from *dv1.Peer, to *dv1.Peer) *v1.Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &v1.Session{
		Direction: v1.Direction_DIRECTION_INCOMING,
		Payload:   payload,

		LastUpdated:  int64(time.Now().Unix()),
		Items:        CreatePayloadItems(sessionPayload, v1.Direction_DIRECTION_INCOMING),
		CurrentIndex: 0,
		Results:      make(map[int32]bool),
	}
}

// NewOutSession creates a new Session from the given payload with Outgoing direction.
func NewOutSession(payload *v1.Payload, to *dv1.Peer, from *dv1.Peer) *v1.Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &v1.Session{
		Direction: v1.Direction_DIRECTION_OUTGOING,
		Payload:   payload,

		LastUpdated:  int64(time.Now().Unix()),
		Items:        CreatePayloadItems(sessionPayload, v1.Direction_DIRECTION_OUTGOING),
		CurrentIndex: 0,
		Results:      make(map[int32]bool),
	}
}

// FinalIndex returns the final index of the session.
func SessionFinalIndex(s *v1.Session) int {
	return len(s.Items) - 1
}

// HasRead returns true if all files have been read.
func SessionHasRead(s *v1.Session) bool {
	return SessionIsIn(s) && SessionIsDone(s)
}

// HasWrote returns true if all files have been written.
func SessionHasWrote(s *v1.Session) bool {
	return SessionIsOut(s) && SessionIsDone(s)
}

// IsDone returns true if all files have been read or written.
func SessionIsDone(s *v1.Session) bool {
	return int(s.GetCurrentIndex()) >= SessionFinalIndex(s)
}

// IsOut returns true if the session is outgoing.
func SessionIsOut(s *v1.Session) bool {
	return s.Direction == v1.Direction_DIRECTION_OUTGOING
}

// IsIn returns true if the session is incoming.
func SessionIsIn(s *v1.Session) bool {
	return s.Direction == v1.Direction_DIRECTION_INCOMING
}

// Event returns the complete event for the session.
func SessionEvent(s *v1.Session) *v1.OnTransmitCompleteResponse {
	return &v1.OnTransmitCompleteResponse{
		Direction: s.GetDirection(),
		// Payload:    s.GetPayload(),
		CreatedAt:  s.GetPayload().GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
		Results:    s.GetResults(),
	}
}

// RouteStream is used to route the given stream to the given peer.
func RouteSessionStream(s *v1.Session, stream network.Stream, h host.SonrHost) (*v1.OnTransmitCompleteResponse, error) {
	// Initialize Params
	logger.Debugf("Beginning %s Transmit Stream", s.Direction.String())
	doneChan := make(chan bool)

	// Check for Incoming
	if SessionIsIn(s) {
		// Handle incoming stream
		go func(stream network.Stream, dchan chan bool) {
			// Create reader
			rs := msgio.NewReader(stream)

			// Read all items
			for _, v := range s.GetItems() {
				// Read Stream to File
				if err := ReadItemFromStream(v, h, rs); err != nil {
					logger.Errorf("Error reading stream: %v", err)
					dchan <- false
				} else {
					dchan <- true
				}
			}

			// Close Stream on Done Reading
			stream.Close()
		}(stream, doneChan)
	}

	// Check for Outgoing
	if SessionIsOut(s) {
		// Handle outgoing stream
		go func(stream network.Stream, dchan chan bool) {
			// Create writer
			wc := msgio.NewWriter(stream)

			// Write all items
			for _, v := range s.GetItems() {
				// Write File to Stream
				if err := WriteItemToStream(v, h, wc); err != nil {
					logger.Errorf("Error writing file: %v", err)
					dchan <- false
				} else {
					dchan <- true
				}
			}
		}(stream, doneChan)
	}

	// Wait for all files to be written
	for {
		select {
		case r := <-doneChan:
			// Set Result
			if complete := UpdateCurrent(s, r); !complete {
				continue
			} else {
				return SessionEvent(s), nil
			}
		}
	}
}

// UpdateCurrent updates the current index of the session.
func UpdateCurrent(s *v1.Session, result bool) bool {
	logger.Debugf("Item (%v) transmit result: %v", s.GetCurrentIndex(), result)
	s.Results[s.GetCurrentIndex()] = result
	s.CurrentIndex = s.GetCurrentIndex() + 1
	return int(s.GetCurrentIndex()) >= SessionFinalIndex(s)
}
