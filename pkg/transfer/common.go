package transfer

import (
	"errors"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/wallet"
	"github.com/sonr-io/core/pkg/common"
)

// Transfer Emission Events
const (
	Event_INVITED   = "transfer-invited"
	Event_RESPONDED = "transfer-responded"
	Event_PROGRESS  = "transfer-progress"
	Event_COMPLETED = "transfer-completed"
	ITEM_INTERVAL   = 25
)

// Transfer Protocol ID's
const (
	RequestPID  protocol.ID = "/transfer/request/0.0.1"
	ResponsePID protocol.ID = "/transfer/response/0.0.1"
	SessionPID  protocol.ID = "/transfer/session/0.0.1"
)

// Error Definitions
var (
	logger             = golog.Child("protocols/transfer")
	ErrTimeout         = errors.New("Session has Timed out")
	ErrParameters      = errors.New("Failed to create new TransferProtocol, invalid parameters")
	ErrInvalidResponse = errors.New("Invalid Transfer InviteResponse provided to TransferProtocol")
	ErrInvalidRequest  = errors.New("Invalid Transfer InviteRequest provided to TransferProtocol")
	ErrFailedEntry     = errors.New("Failed to get Topmost entry from Queue")
	ErrFailedAuth      = errors.New("Failed to Authenticate message")
	ErrEmptyRequests   = errors.New("Empty Request list provided")
	ErrRequestNotFound = errors.New("Request not found in list")
)

// checkParams Checks if Non-nil Parameters were passed
func checkParams(host *host.SNRHost) error {
	if host == nil {
		logger.Errorf("%s - Host provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.HasRouting()
}

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *api.DecisionEvent {
	return &api.DecisionEvent{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *api.InviteEvent {
	return &api.InviteEvent{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}

// Share a peer to have a transfer
func (p *TransferProtocol) createRequest(to *common.Peer) (peer.ID, *InviteRequest, error) {
	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Fetch Element from Queue
	elem := p.supplyQueue.Front()
	if elem != nil {
		// Get Payload
		payload := p.supplyQueue.Remove(elem).(*common.Payload)

		// Create new Metadata
		meta, err := wallet.Sonr.CreateMetadata(p.host.ID())
		if err != nil {
			logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
			return "", nil, err
		}

		// Create Invite Request
		req := &InviteRequest{
			Payload:  payload,
			Metadata: api.SignedMetadataToProto(meta),
			To:       to,
			From:     from,
		}

		// Fetch Peer ID from Public Key
		toId, err := to.Libp2pID()
		if err != nil {
			logger.Errorf("%s - Failed to fetch peer id from public key", err)
			return "", nil, err
		}
		return toId, req, nil
	}
	logger.Errorf("%s - Failed to get item from Supply Queue.")
	return "", nil, errors.New("No items in Supply Queue.")
}

// Respond to an invite request
func (p *TransferProtocol) createResponse(decs bool, to *common.Peer) (peer.ID, *InviteResponse, error) {
	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Create new Metadata
	meta, err := wallet.Sonr.CreateMetadata(p.host.ID())
	if err != nil {
		logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
		return "", nil, err
	}

	// Create Invite Response
	resp := &InviteResponse{
		Decision: decs,
		Metadata: api.SignedMetadataToProto(meta),
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
}
