package transmit

import (
	"sync"
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
		Direction:   common.Direction_INCOMING,
		Payload:     payload,
		From:        from,
		To:          to,
		LastUpdated: int64(time.Now().Unix()),
		Items:       sessionPayload.CreateItems(common.Direction_INCOMING),
	}
}

// NewOutSession creates a new Session from the given payload with Outgoing direction.
func NewOutSession(payload *common.Payload, to *common.Peer, from *common.Peer) *Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &Session{
		Direction:   common.Direction_OUTGOING,
		Payload:     payload,
		To:          to,
		From:        from,
		LastUpdated: int64(time.Now().Unix()),
		Items:       sessionPayload.CreateItems(common.Direction_OUTGOING),
	}
}

// WriteTo writes the Session to the given stream.
func (s *Session) Handle(stream network.Stream, n api.NodeImpl) (*api.CompleteEvent, error) {
	// Initialize Params
	logger.Debugf("Beginning %s Transmit Stream", s.Direction.String())
	var wg sync.WaitGroup

	// Check for Incoming
	if s.Direction == common.Direction_INCOMING {
		// Handle incoming stream
		rs := msgio.NewReader(stream)

		// Read All Files
		for _, v := range s.GetItems() {
			// Write to File
			wg.Add(1)
			go v.Read(&wg, n, rs)
		}
		wg.Wait()
		stream.Close()

		// Return Complete Event
		return &api.CompleteEvent{
			From:       s.GetFrom(),
			To:         s.GetTo(),
			Direction:  s.GetDirection(),
			Payload:    s.GetPayload(),
			CreatedAt:  s.GetPayload().GetCreatedAt(),
			ReceivedAt: int64(time.Now().Unix()),
		}, nil
	}

	// Check for Outgoing
	if s.Direction == common.Direction_OUTGOING {
		// Initialize Params
		wc := msgio.NewWriter(stream)

		// Create New Writer
		for _, v := range s.GetItems() {
			// Write File to Stream
			wg.Add(1)
			go v.Write(&wg, n, wc)
		}

		// Wait for all writes to finish
		wg.Wait()

		// Return Complete Event
		return &api.CompleteEvent{
			From:       s.GetFrom(),
			To:         s.GetTo(),
			Direction:  s.GetDirection(),
			Payload:    s.GetPayload(),
			CreatedAt:  s.GetPayload().GetCreatedAt(),
			ReceivedAt: int64(time.Now().Unix()),
		}, nil
	}

	// Return Error
	return nil, ErrInvalidDirection
}
