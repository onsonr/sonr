package transfer

import (
	"container/list"
	"context"
	"fmt"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
)

// TransferProtocol type
type TransferProtocol struct {
	api.NodeImpl
	ctx          context.Context // Context
	host         *host.SNRHost   // local host
	sessionQueue *SessionQueue   // transfer session queue
	supplyQueue  *list.List      // supply queue
}

// NewProtocol creates a new TransferProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, node api.NodeImpl) (*TransferProtocol, error) {
	// Check parameters
	if err := checkParams(host); err != nil {
		logger.Error("Failed to create TransferProtocol", err)
		return nil, err
	}

	// create a new transfer protocol
	invProtocol := &TransferProtocol{
		ctx:  ctx,
		host: host,
		sessionQueue: &SessionQueue{
			ctx:   ctx,
			host:  host,
			queue: list.New(),
		},
		supplyQueue: list.New(),
	}

	// Setup Stream Handlers
	host.SetStreamHandler(RequestPID, invProtocol.onInviteRequest)
	host.SetStreamHandler(ResponsePID, invProtocol.onInviteResponse)
	host.SetStreamHandler(SessionPID, invProtocol.onIncomingTransfer)
	logger.Info("✅  TransferProtocol is Activated \n")
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
	p.sessionQueue.AddOutgoing(id, req)
	return nil
}

// Respond Method authenticates or declines a Transfer Request
func (p *TransferProtocol) Respond(id peer.ID, resp *InviteResponse) error {
	// Check if the response is valid
	if resp == nil {
		return ErrInvalidResponse
	}

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

// Supply a transfer item to the queue
func (p *TransferProtocol) Supply(paths []string) error {
	// Profile from NodeImpl
	profile, err := p.GetProfile()
	if err != nil {
		logger.Error("Failed to Get Profile from Node")
		return err
	}

	// Create Transfer
	payload, err := common.NewPayload(profile, paths)
	if err != nil {
		logger.Error("Failed to Supply Paths", err)
		return err
	}

	// Add items to transfer
	p.supplyQueue.PushBack(payload)
	logger.Info(fmt.Sprintf("Added %v items to supply queue.", len(paths)), golog.Fields{"File Count": payload.FileCount(), "URL Count": payload.URLCount()})
	return nil
}
