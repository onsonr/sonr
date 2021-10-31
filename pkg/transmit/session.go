package transmit

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

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

	// Write All Files
	for i, v := range s.Items() {
		// Write to File
		wg.Add(1)
		go ReadItem(i, s.Count(), v, &wg, n, rs)
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
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)
		go WriteItem(i, s.Count(), v, &wg, n, wc)
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

// MapItems performs PayloadItemFunc on each item in the Payload.
func (s *Session) Items() []*common.Payload_Item {
	return s.GetPayload().GetItems()
}
