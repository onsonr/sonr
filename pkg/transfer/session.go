package transfer

import (
	"container/list"
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
)

// Session is a single entry in the transfer queue.
type Session struct {
	direction   common.CompleteEvent_Direction
	request     *InviteRequest
	response    *InviteResponse
	fromId      peer.ID
	toId        peer.ID
	lastUpdated int64
	uuid        *common.UUID
}

// Count returns the number of items in Payload
func (s Session) Count() int {
	return len(s.request.GetPayload().GetItems())
}

// CopyUUID copies Request UUID to Response
func (s Session) CopyUUID(resp *InviteResponse) *InviteResponse {
	resp.Uuid = s.uuid
	return resp
}

// Equals checks if given ID is equal to the current UUID.
func (s Session) Equals(id *common.UUID) bool {
	return s.uuid.GetValue() == id.GetValue()
}

// MapItems performs PayloadItemFunc on each item in the Payload.
func (s Session) MapItems(f common.PayloadItemFunc) error {
	return s.request.GetPayload().MapItems(f)
}

func (s Session) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(10 * time.Minute):
		return ErrTimeout
	}
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
	entry := Session{
		direction:   common.CompleteEvent_INCOMING,
		request:     req,
		fromId:      from,
		toId:        sq.host.ID(),
		lastUpdated: int64(time.Now().Unix()),
		uuid:        req.GetUuid(),
	}

	// Add to Requests
	sq.queue.PushBack(entry)
	return nil
}

// AddOutgoing adds Outgoing Request to Transfer Queue
func (sq *SessionQueue) AddOutgoing(to peer.ID, req *InviteRequest) error {
	// Create New TransferEntry
	entry := Session{
		direction:   common.CompleteEvent_OUTGOING,
		request:     req,
		fromId:      sq.host.ID(),
		toId:        to,
		lastUpdated: int64(time.Now().Unix()),
		uuid:        req.GetUuid(),
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

	val := entry.Value.(Session)
	val.lastUpdated = int64(time.Now().Unix())
	return &val, nil
}

// Done marks the transfer as completed and returns the CompleteEvent.
func (sq *SessionQueue) Done() (*common.CompleteEvent, error) {
	// Find Entry for Peer
	entry := sq.queue.Front()
	if entry == nil {
		return nil, ErrFailedEntry
	}

	// Pop Value of Entry from Queue
	val := sq.queue.Remove(entry).(Session)
	rawPayload := val.request.GetPayload()

	// Adjust Payload item paths
	adjPayload, err := rawPayload.ReplaceItemsDir(device.DownloadsPath)
	if err != nil {
		return nil, err
	}

	// Create CompleteEvent
	event := &common.CompleteEvent{
		From:       val.request.GetFrom(),
		To:         val.request.GetTo(),
		Direction:  val.direction,
		Payload:    adjPayload,
		CreatedAt:  adjPayload.GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
	}
	return event, nil
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

	// Validate UUID
	ok, err := sq.host.AuthenticateId(resp.GetUuid())
	if err != nil {
		return nil, err
	}

	// Check if UUID is valid
	if !ok {
		return nil, ErrMismatchUUID
	}

	// Get Next Entry
	entry, err := sq.Next()
	if err != nil {
		 logger.Error("Failed to get Transfer entry", err)
		return nil, err
	}

	// Check if Request exists in Map
	if entry.Equals(resp.GetUuid()) {
		entry.response = resp
		entry.lastUpdated = int64(time.Now().Unix())
		return entry, nil
	}
	return nil, ErrRequestNotFound
}
