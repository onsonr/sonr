package transmit

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

func NewOutSession(payload *common.Payload, to *common.Peer, from *common.Peer) *Session {
	// Create Session Items
	sessionPayload := createPayload(payload)
	return &Session{
		Direction:   common.Direction_OUTGOING,
		Payload:     payload,
		To:          to,
		From:        from,
		LastUpdated: int64(time.Now().Unix()),
		Items:       sessionPayload.CreateItems(common.Direction_OUTGOING),
	}
}

func NewInSession(payload *common.Payload, from *common.Peer, to *common.Peer) *Session {
	// Create Session Items
	sessionPayload := createPayload(payload)
	return &Session{
		Direction:   common.Direction_INCOMING,
		Payload:     payload,
		From:        from,
		To:          to,
		LastUpdated: int64(time.Now().Unix()),
		Items:       sessionPayload.CreateItems(common.Direction_INCOMING),
	}
}

// IsIncoming returns true if the session is incoming.
func (s *Session) IsIncoming() bool {
	return s.Direction == common.Direction_INCOMING
}

// IsOutgoing returns true if the session is outgoing.
func (s *Session) IsOutgoing() bool {
	return s.Direction == common.Direction_OUTGOING
}

// ReadFrom reads the next Session from the given stream.
func (s *Session) ReadFrom(stream network.Stream, n api.NodeImpl) (*api.CompleteEvent, error) {
	// Initialize Params
	logger.Debug("Beginning INCOMING Transmit Stream")

	// Handle incoming stream
	rs := msgio.NewReader(stream)
	var wg sync.WaitGroup

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

// WriteTo writes the Session to the given stream.
func (s *Session) WriteTo(stream network.Stream, n api.NodeImpl) (*api.CompleteEvent, error) {
	// Initialize Params
	logger.Debug("Beginning OUTGOING Transmit Stream")
	wc := msgio.NewWriter(stream)
	var wg sync.WaitGroup

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

// Count returns the number of items in Payload
func (s *Session) Count() int {
	return len(s.GetPayload().GetItems())
}
