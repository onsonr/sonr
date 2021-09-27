package transfer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
)

// TransferEntry is a single entry in the transfer queue.
type TransferEntry struct {
	direction   common.CompleteEvent_Direction
	request     *InviteRequest
	response    *InviteResponse
	fromId      peer.ID
	toId        peer.ID
	lastUpdated int64
	uuid        string
}

// Count returns the number of items in Payload
func (e TransferEntry) Count() int {
	return len(e.request.GetPayload().GetItems())
}

// CopyUUID copies Request UUID to Response
func (e TransferEntry) CopyUUID(resp *InviteResponse) *InviteResponse {
	resp.Uuid = e.uuid
	return resp
}

// Equals checks if given ID is equal to the current UUID.
func (e TransferEntry) Equals(id string) bool {
	return e.uuid == id
}

// transferQueue is a queue for incoming and outgoing requests.
type transferQueue struct {
	ctx      context.Context
	host     *host.SNRHost
	requests map[peer.ID]TransferEntry
}

// AddIncoming adds Incoming Request to Transfer Queue
func (tq *transferQueue) AddIncoming(from peer.ID, req *InviteRequest) error {
	// Authenticate Message
	valid := tq.host.AuthenticateMessage(req, req.Metadata)
	if !valid {
		return fmt.Errorf("Failed to Authorize Invite REQUEST.")
	}

	// Create New TransferEntry
	entry := TransferEntry{
		direction:   common.CompleteEvent_INCOMING,
		request:     req,
		fromId:      from,
		toId:        tq.host.ID(),
		lastUpdated: int64(time.Now().Unix()),
		uuid:        req.GetUuid(),
	}

	// Add to Requests
	tq.requests[from] = entry
	return nil
}

// AddOutgoing adds Outgoing Request to Transfer Queue
func (tq *transferQueue) AddOutgoing(to peer.ID, req *InviteRequest) error {
	// Create New TransferEntry
	entry := TransferEntry{
		direction:   common.CompleteEvent_OUTGOING,
		request:     req,
		fromId:      tq.host.ID(),
		toId:        to,
		lastUpdated: int64(time.Now().Unix()),
		uuid:        req.GetUuid(),
	}

	// Add to Requests
	tq.requests[to] = entry
	return nil
}

// Complete marks the transfer as completed and returns the CompleteEvent.
func (tq *transferQueue) Complete(peer peer.ID) (*common.CompleteEvent, error) {
	// Find Entry for Peer
	entry, ok := tq.requests[peer]
	if !ok {
		return nil, errors.New("Entry for these peer not found in Request Queue")
	}

	// Create CompleteEvent
	event := &common.CompleteEvent{
		Direction: entry.direction,
		Payload:   entry.request.GetPayload(),
		Received:  int64(time.Now().Unix()),
	}

	// Delete Entry
	delete(tq.requests, peer)
	return event, nil
}

// Complete marks the transfer as completed and returns the CompleteEvent.
func (tq *transferQueue) Find(peer peer.ID) (*TransferEntry, error) {
	// Find Entry for Peer
	entry, ok := tq.requests[peer]
	if !ok {
		return nil, errors.New("Entry for these peer not found in Request Queue")
	}
	entry.lastUpdated = int64(time.Now().Unix())
	return &entry, nil
}

// Validate takes list of Requests and returns true if Request exists in List and UUID is verified.
// Method also returns the InviteRequest that points to the Response.
func (tq *transferQueue) Validate(resp *InviteResponse) (*TransferEntry, error) {
	// Authenticate Message
	valid := tq.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		return nil, errors.New("Failed to Authenticate Invite RESPONSE.")
	}

	// Check Decision
	if !resp.GetDecision() {
		return nil, errors.New("Peer rejected Invite")
	}

	// Check if the request is valid
	totalRequests := len(tq.requests)
	if totalRequests == 0 {
		return nil, errors.New("No InviteRequest's provided in list.")
	}

	// Validate UUID
	ok, err := tq.host.AuthenticateId(resp.GetUuid())
	if err != nil {
		return nil, err
	}

	// Check if UUID is valid
	if !ok {
		return nil, errors.New("UUID does not match")
	}

	// Check if Request exists in Map
	for i, entry := range tq.requests {
		if entry.Equals(resp.GetUuid()) {
			logger.Info(fmt.Sprintf("Found matching invite at [%v/%v]", i, totalRequests))
			entry.response = resp
			entry.lastUpdated = int64(time.Now().Unix())
			return &entry, nil
		}
	}
	return nil, errors.New("InviteRequest does not exist in provided list.")
}
