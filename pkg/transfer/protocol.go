package transfer

import (
	"container/list"
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/state"
)

// TransferProtocol type
type TransferProtocol struct {
	ctx     context.Context // Context
	host    *host.SNRHost   // local host
	queue   *SessionQueue   // transfer queue
	emitter *state.Emitter  // Handle to signal when done
}

// NewProtocol creates a new TransferProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, em *state.Emitter) (*TransferProtocol, error) {
	// Check parameters
	if err := checkParams(host, em); err != nil {
		logger.Error("Failed to create TransferProtocol", err)
		return nil, err
	}

	// Wait until host is ready
	if err := host.WaitForReady(); err != nil {
		logger.Error("Failed to create TransferProtocol", err)
		return nil, err
	}

	// create a new transfer protocol
	invProtocol := &TransferProtocol{
		ctx:     ctx,
		host:    host,
		emitter: em,
		queue: &SessionQueue{
			ctx:   ctx,
			host:  host,
			queue: list.New(),
		},
	}

	// Setup Stream Handlers
	host.SetStreamHandler(RequestPID, invProtocol.onInviteRequest)
	host.SetStreamHandler(ResponsePID, invProtocol.onInviteResponse)
	host.SetStreamHandler(SessionPID, invProtocol.onIncomingTransfer)
	return invProtocol, nil
}

// Request Method sends a request to Transfer Data to a remote peer
func (p *TransferProtocol) Request(id peer.ID, req *InviteRequest) error {
	// Check if the response is valid
	if req == nil {
		return ErrInvalidRequest
	}

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
		logger.Error("Failed to Send Message to Peer", err)
		return err
	}

	// store the request in the map
	p.queue.AddOutgoing(id, req)
	return nil
}

// Respond Method authenticates or declines a Transfer Request
func (p *TransferProtocol) Respond(id peer.ID, resp *InviteResponse) error {
	// Check if the response is valid
	if resp == nil {
		return ErrInvalidResponse
	}

	// Find Entry
	entry, err := p.queue.Next()
	if err != nil {
		logger.Error("Failed to find transfer entry", err)
		return err
	}

	// Copy UUID
	resp = entry.CopyUUID(resp)

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		logger.Error("Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// Send Response
	err = p.host.SendMessage(id, ResponsePID, resp)
	if err != nil {
		logger.Error("Failed to Send Message to Peer", err)
		return err
	}
	return nil
}
