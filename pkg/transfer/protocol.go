package transfer

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
)

// TransferProtocol type
type TransferProtocol struct {
	ctx     context.Context // Context
	host    *host.SNRHost   // local host
	queue   *transferQueue  // transfer queue
	emitter *state.Emitter  // Handle to signal when done
}

// NewProtocol creates a new TransferProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, em *state.Emitter) *TransferProtocol {
	// create a new transfer protocol
	invProtocol := &TransferProtocol{
		ctx:     ctx,
		host:    host,
		emitter: em,
		queue: &transferQueue{
			ctx:      ctx,
			host:     host,
			requests: make(map[peer.ID]TransferEntry),
		},
	}

	// Setup Stream Handlers
	host.SetStreamHandler(RequestPID, invProtocol.onInviteRequest)
	host.SetStreamHandler(ResponsePID, invProtocol.onInviteResponse)
	host.SetStreamHandler(SessionPID, invProtocol.onIncomingTransfer)
	return invProtocol
}

// Request Method sends a request to Transfer Data to a remote peer
func (p *TransferProtocol) Request(id peer.ID, req *InviteRequest) error {
	// sign the data
	signature, err := p.host.SignMessage(req)
	if err != nil {
		logger.Error("Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.host.SendMessage(id, RequestPID, req)
	if err != nil {
		return logger.Error("Failed to Send Message to Peer", err)
	}

	// store the request in the map
	p.queue.AddOutgoing(id, req)
	return nil
}

// Respond Method authenticates or declines a Transfer Request
func (p *TransferProtocol) Respond(id peer.ID, resp *InviteResponse) error {
	// Find Entry
	entry, err := p.queue.Find(id)
	if err != nil {
		return logger.Error("Failed to find transfer entry", err)
	}

	// Copy UUID
	resp = entry.CopyUUID(resp)

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		return logger.Error("Failed to Sign Response Message", err)
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// Send Response
	err = p.host.SendMessage(id, ResponsePID, resp)
	if err != nil {
		return logger.Error("Failed to Send Message to Peer", err)
	}
	return nil
}

// TransferEntry is a single entry in the transfer queue.
type TransferEntry struct {
	direction   common.CompleteEvent_Direction
	request     *InviteRequest
	response    *InviteResponse
	fromId      peer.ID
	toId        peer.ID
	lastUpdated int64
	uuid        *common.UUID
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
func (e TransferEntry) Equals(id *common.UUID) bool {
	return e.uuid.GetValue() == id.GetValue()
}

// MapItems performs PayloadItemFunc on each item in the Payload.
func (e TransferEntry) MapItems(f common.PayloadItemFunc) error {
	return e.request.GetPayload().MapItems(f)
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
		return ErrFailedAuth
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
		return nil, logger.Error("Failed to find Entry", ErrInvalidEntry)
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
		return nil, logger.Error("Failed to find Entry", ErrInvalidEntry)
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
		return nil, ErrFailedAuth
	}

	// Check Decision
	if !resp.GetDecision() {
		return nil, nil
	}

	// Check if the request is valid
	totalRequests := len(tq.requests)
	if totalRequests == 0 {
		return nil, ErrEmptyRequests
	}

	// Validate UUID
	ok, err := tq.host.AuthenticateId(resp.GetUuid())
	if err != nil {
		return nil, err
	}

	// Check if UUID is valid
	if !ok {
		return nil, ErrMismatchUUID
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
	return nil, ErrRequestNotFound
}
