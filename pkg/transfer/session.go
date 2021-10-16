package transfer

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
)

// Session is a single entry in the transfer queue.
type Session struct {
	direction   common.Direction
	from        *common.Peer
	to          *common.Peer
	payload     *common.Payload
	lastUpdated int64
}

// IsIncoming returns true if the session is incoming.
func (s *Session) IsIncoming() bool {
	return s.direction == common.Direction_INCOMING
}

// IsOutgoing returns true if the session is outgoing.
func (s *Session) IsOutgoing() bool {
	return s.direction == common.Direction_OUTGOING
}

// ReadFrom reads the next Session from the given stream.
func (s *Session) ReadFrom(stream network.Stream) *api.CompleteEvent {
	// Initialize Params
	logger.Info("Beginning INCOMING Transfer Stream")

	// Handle incoming stream
	rs := msgio.NewReader(stream)
	var wg sync.WaitGroup

	// Write All Files
	for i, v := range s.Items() {
		// Create Reader
		r := NewReader(i, s.Count(), v)

		// Write to File
		wg.Add(1)
		go func(idx, total int) {
			defer wg.Done()
			r.ReadFrom(rs)
			logger.Info(fmt.Sprintf("Finished RECEIVING File (%v/%v)", idx, total))
		}(i, s.Count())
	}
	wg.Wait()

	// Return Complete Event
	return &api.CompleteEvent{
		From:       s.from,
		To:         s.to,
		Direction:  s.direction,
		Payload:    s.SetPayload(),
		CreatedAt:  s.SetPayload().GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
	}
}

// WriteTo writes the Session to the given stream.
func (s *Session) WriteTo(stream network.Stream) *api.CompleteEvent {
	// Initialize Params
	logger.Info("Beginning OUTGOING Transfer Stream")
	wc := msgio.NewWriter(stream)
	var wg sync.WaitGroup

	// Create New Writer
	for i, v := range s.Items() {
		// Create New Writer
		w, err := NewWriter(i, s.Count(), v)
		if err != nil {
			logger.Error("Failed to create new writer.", err)
			wc.Close()
			return nil
		}

		// Write File to Stream
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.WriteTo(wc)
		}()
	}

	// Wait for all writes to finish
	wg.Wait()

	// Return Complete Event
	return &api.CompleteEvent{
		From:       s.from,
		To:         s.to,
		Direction:  s.direction,
		Payload:    s.SetPayload(),
		CreatedAt:  s.SetPayload().GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
	}
}

// Count returns the number of items in Payload
func (s *Session) Count() int {
	return len(s.payload.GetItems())
}

// MapItems performs PayloadItemFunc on each item in the Payload.
func (s *Session) Items() []*common.Payload_Item {
	return s.payload.GetItems()
}

// SetPayload sets the Payload for the Session.
func (s *Session) SetPayload() *common.Payload {
	if s.IsIncoming() {
		s.payload = s.payload.ReplaceItemsDir(device.DownloadsPath)
		s.lastUpdated = common.NewLastUpdated()
	}
	return s.payload
}

// SessionQueue is a queue for incoming and outgoing requests.
type SessionQueue struct {
	ctx   context.Context
	host  *host.SNRHost
	queue *list.List
}

// AddIncoming adds Incoming Request to Transfer Queue
func (sq *SessionQueue) AddIncoming(from peer.ID, req *InviteRequest) error {
	// Authenticate Message
	valid := sq.host.AuthenticateMessage(req, req.Metadata)
	if !valid {
		return ErrFailedAuth
	}

	// Create New TransferEntry
	entry := &Session{
		direction:   common.Direction_INCOMING,
		payload:     req.GetPayload(),
		from:        req.GetFrom(),
		to:          req.GetTo(),
		lastUpdated: int64(time.Now().Unix()),
	}

	// Add to Requests
	sq.queue.PushBack(entry)
	return nil
}

// AddOutgoing adds Outgoing Request to Transfer Queue
func (sq *SessionQueue) AddOutgoing(to peer.ID, req *InviteRequest) error {
	// Create New TransferEntry
	entry := &Session{
		direction:   common.Direction_OUTGOING,
		payload:     req.GetPayload(),
		from:        req.GetFrom(),
		to:          req.GetTo(),
		lastUpdated: int64(time.Now().Unix()),
	}

	// Add to Requests
	sq.queue.PushBack(entry)
	return nil
}

// Next returns topmost entry in the queue.
func (sq *SessionQueue) Next() (*Session, error) {
	// Find Entry for Peer
	entry := sq.queue.Front()
	if entry == nil {
		return nil, ErrFailedEntry
	}

	val := entry.Value.(*Session)
	val.lastUpdated = int64(time.Now().Unix())
	return val, nil
}

// Validate takes list of Requests and returns true if Request exists in List and UUID is verified.
// Method also returns the InviteRequest that points to the Response.
func (sq *SessionQueue) Validate(resp *InviteResponse) (*Session, error) {
	// Authenticate Message
	valid := sq.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		return nil, ErrFailedAuth
	}

	// Check Decision
	if !resp.GetDecision() {
		return nil, nil
	}

	// Check if the request is valid
	if sq.queue.Len() == 0 {
		return nil, ErrEmptyRequests
	}

	// Get Next Entry
	entry, err := sq.Next()
	if err != nil {
		logger.Error("Failed to get Transfer entry", err)
		return nil, err
	}

	// Check if Request exists in Map
	if entry != nil {
		entry.lastUpdated = int64(time.Now().Unix())
		return entry, nil
	}
	return nil, ErrRequestNotFound
}
