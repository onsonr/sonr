package transmit

import (
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

// NewInSession creates a new Session from the given payload with Incoming direction.
func NewInSession(payload *common.Payload, from *common.Peer, to *common.Peer) *Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &Session{
		Direction:    common.Direction_INCOMING,
		Payload:      payload,
		From:         from,
		To:           to,
		LastUpdated:  int64(time.Now().Unix()),
		Items:        sessionPayload.CreateItems(common.Direction_INCOMING),
		CurrentIndex: 0,
		Results:      make(map[int32]bool),
	}
}

// NewOutSession creates a new Session from the given payload with Outgoing direction.
func NewOutSession(payload *common.Payload, to *common.Peer, from *common.Peer) *Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &Session{
		Direction:    common.Direction_OUTGOING,
		Payload:      payload,
		To:           to,
		From:         from,
		LastUpdated:  int64(time.Now().Unix()),
		Items:        sessionPayload.CreateItems(common.Direction_OUTGOING),
		CurrentIndex: 0,
		Results:      make(map[int32]bool),
	}
}

// HasRead returns true if all files have been read.
func (s *Session) HasRead() bool {
	return s.IsIn() && s.IsDone()
}

// HasWrote returns true if all files have been written.
func (s *Session) HasWrote() bool {
	return s.IsOut() && s.IsDone()
}

// IsDone returns true if all files have been read or written.
func (s *Session) IsDone() bool {
	return int(s.CurrentIndex) >= len(s.GetItems())
}

// IsOut returns true if the session is outgoing.
func (s *Session) IsOut() bool {
	return s.Direction == common.Direction_OUTGOING
}

// IsIn returns true if the session is incoming.
func (s *Session) IsIn() bool {
	return s.Direction == common.Direction_INCOMING
}

// Event returns the complete event for the session.
func (s *Session) Event() *api.CompleteEvent {
	return &api.CompleteEvent{
		From:       s.GetFrom(),
		To:         s.GetTo(),
		Direction:  s.GetDirection(),
		Payload:    s.GetPayload(),
		CreatedAt:  s.GetPayload().GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
		Results:    s.GetResults(),
	}
}

// RouteStream is used to route the given stream to the given peer.
func (s *Session) RouteStream(stream network.Stream, n api.NodeImpl) (*api.CompleteEvent, error) {
	// Initialize Params
	logger.Debugf("Beginning %s Transmit Stream", s.Direction.String())
	doneChan := make(chan bool)

	// Check for Incoming
	if s.IsIn() {
		// Handle incoming stream
		rs := msgio.NewReader(stream)
		for _, v := range s.GetItems() {
			// Read Stream to File
			go v.Read(doneChan, n, rs)
		}
	}

	// Check for Outgoing
	if s.IsOut() {
		// Handle outgoing stream
		wc := msgio.NewWriter(stream)
		for _, v := range s.GetItems() {
			// Write File to Stream
			go v.Write(doneChan, n, wc)
		}
	}

	// Wait for all files to be written
	for {
		select {
		case r := <-doneChan:
			// Set Result
			s.UpdateCurrent(r)

			// Close Stream on Done Reading
			if s.HasRead() {
				stream.Close()
			}

			// Return Event
			return s.Event(), nil
		}
	}
}

// UpdateCurrent updates the current index of the session.
func (s *Session) UpdateCurrent(result bool) {
	logger.Debugf("Item (%v) transmit result: %v", s.CurrentIndex, result)
	s.Results[s.CurrentIndex] = result
	s.CurrentIndex = s.CurrentIndex + 1
}
