package transfer

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
)

// Transfer Emission Events
const (
	Event_INVITED   = "transfer-invited"
	Event_RESPONDED = "transfer-responded"
	Event_PROGRESS  = "transfer-progress"
	Event_COMPLETED = "transfer-completed"
)

// Transfer Protocol ID's
const (
	RequestPID  protocol.ID = "/transfer/request/0.0.1"
	ResponsePID protocol.ID = "/transfer/response/0.0.1"
	SessionPID  protocol.ID = "/transfer/session/0.0.1"
)

// TransferProtocol type
type TransferProtocol struct {
	ctx     context.Context  // Context
	host    *host.SNRHost    // local host
	queue   *transferQueue   // transfer queue
	emitter *emitter.Emitter // Handle to signal when done
}

// NewProtocol creates a new TransferProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, em *emitter.Emitter) *TransferProtocol {
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
		logger.Error("Failed to Sign Response Message", zap.Error(err))
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.host.SendMessage(id, RequestPID, req)
	if err != nil {
		logger.Error("Failed to Send Message to Peer", zap.Error(err))
		return err
	}

	// store the request in the map
	p.queue.AddOutgoing(id, req)
	return nil
}

// Respond Method authenticates or declines a Transfer Request
func (p *TransferProtocol) Respond(resp *InviteResponse) error {
	// Get ToID
	toID, err := resp.GetTo().PeerID()
	if err != nil {
		logger.Error("Failed to Get PeerID", zap.Error(err))
		return err
	}

	// Find Entry
	entry, err := p.queue.Find(toID)
	if err != nil {
		logger.Error("Failed to find transfer entry", zap.Error(err))
		return err
	}

	// Copy UUID
	resp = entry.CopyUUID(resp)

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		logger.Error("Failed to Sign Response Message", zap.Error(err))
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// Send Response
	err = p.host.SendMessage(toID, ResponsePID, resp)
	if err != nil {
		logger.Error("Failed to Send Message to Peer", zap.Error(err))
		return err
	}
	return nil
}
