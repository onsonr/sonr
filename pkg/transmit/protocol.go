package transmit

import (
	"container/list"
	"context"
	"fmt"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
)

// TransmitProtocol type
type TransmitProtocol struct {
	node         api.NodeImpl
	ctx          context.Context // Context
	host         *host.SNRHost   // local host
	sessionQueue *SessionQueue   // transfer session queue
	supplyQueue  *list.List      // supply queue
}

// NewProtocol creates a new TransferProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, node api.NodeImpl) (*TransmitProtocol, error) {
	// create a new transfer protocol
	invProtocol := &TransmitProtocol{
		ctx:  ctx,
		host: host,
		sessionQueue: &SessionQueue{
			ctx:   ctx,
			host:  host,
			queue: list.New(),
		},
		supplyQueue: list.New(),
		node:        node,
	}

	// Setup Stream Handlers
	host.SetStreamHandler(RequestPID, invProtocol.onInviteRequest)
	host.SetStreamHandler(ResponsePID, invProtocol.onInviteResponse)
	host.SetStreamHandler(SessionPID, invProtocol.onIncomingTransfer)
	logger.Debug("✅  TransmitProtocol is Activated \n")
	return invProtocol, nil
}

// Request Method sends a request to Transfer Data to a remote peer
func (p *TransmitProtocol) Request(to *common.Peer) error {
	// Create Request
	id, req, err := p.createRequest(to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// Check if the response is valid
	if req == nil {
		return ErrInvalidRequest
	}

	// sign the data
	signature, err := p.host.SignMessage(req)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.host.SendMessage(id, RequestPID, req)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}

	// store the request in the map
	return p.sessionQueue.AddOutgoing(id, req)
}

// Respond Method authenticates or declines a Transfer Request
func (p *TransmitProtocol) Respond(decs bool, to *common.Peer) error {
	// Create Response
	id, resp, err := p.createResponse(decs, to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// Check if the response is valid
	if resp == nil {
		return ErrInvalidResponse
	}

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// Send Response
	err = p.host.SendMessage(id, ResponsePID, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}
	return nil
}

// Supply a transfer item to the queue
func (p *TransmitProtocol) Supply(req *api.SupplyRequest) error {
	// Profile from NodeImpl
	profile, err := p.node.Profile()
	if err != nil {
		logger.Errorf("%s - Failed to Get Profile from Node")
		return err
	}

	// Create Transfer
	payload, err := req.ToPayload(profile)
	if err != nil {
		logger.Errorf("%s - Failed to Supply Paths", err)
		return err
	}

	// Add items to transfer
	p.supplyQueue.PushBack(payload)
	logger.Debug(fmt.Sprintf("Added %v items to supply queue.", req.Count()), golog.Fields{"File Count": payload.FileCount(), "URL Count": payload.URLCount()})

	// Check if Peer is provided
	if req.GetIsPeerSupply() {
		logger.Debug("Peer Supply Request. Sending Invite after supply")
		err = p.Request(req.GetPeer())
		if err != nil {
			logger.Errorf("%s - Failed to Send Request to Peer", err)
			return err
		}
	}
	return nil
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
		direction:   common.Direction_INCOMING,
		payload:     req.GetPayload(),
		from:        req.GetFrom(),
		to:          req.GetTo(),
		lastUpdated: int64(time.Now().Unix()),
		success:     make(map[int32]bool),
		ctx:         sq.ctx,
	}

	// Add to Requests
	sq.queue.PushBack(entry)
	return nil
}

// AddOutgoing adds Outgoing Request to Transfer Queue
func (sq *SessionQueue) AddOutgoing(to peer.ID, req *InviteRequest) error {
	// Create New TransferEntry
	entry := Session{
		direction:   common.Direction_OUTGOING,
		payload:     req.GetPayload(),
		from:        req.GetFrom(),
		to:          req.GetTo(),
		lastUpdated: int64(time.Now().Unix()),
		success:     make(map[int32]bool),
		ctx:         sq.ctx,
	}

	// Add to Requests
	sq.queue.PushBack(entry)
	return nil
}

// Next returns topmost entry in the queue.
func (sq *SessionQueue) Next() (Session, error) {
	// Find Entry for Peer
	entry := sq.queue.Remove(sq.queue.Front()).(Session)
	entry.lastUpdated = int64(time.Now().Unix())
	return entry, nil
}

// Validate takes list of Requests and returns true if Request exists in List and UUID is verified.
// Method also returns the InviteRequest that points to the Response.
func (sq *SessionQueue) Validate(resp *InviteResponse) (Session, error) {
	// Authenticate Message
	valid := sq.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		return Session{}, ErrFailedAuth
	}

	// Check Decision
	if !resp.GetDecision() {
		return Session{}, nil
	}

	// Check if the request is valid
	if sq.queue.Len() == 0 {
		return Session{}, ErrEmptyRequests
	}

	// Get Next Entry
	entry, err := sq.Next()
	if err != nil {
		logger.Errorf("%s - Failed to get Transmit entry", err)
		return Session{}, err
	}

	entry.lastUpdated = int64(time.Now().Unix())
	return entry, nil
}
