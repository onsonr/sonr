package exchange

import (
	"errors"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/wallet"
	common "github.com/sonr-io/core/pkg/common"
)

// Textile API definitions
const (
	// Textile Client API URL
	TextileClientURL = "https://api.textile.io"

	// Textile Miner Index Target
	TextileMinerIdx = "api.minerindex.hub.textile.io:443"

	// Textile Mailbox Directory
	TextileMailboxDirName = "mailbox"

	RequestPID protocol.ID = "/exchange/request/0.0.1"

	ResponsePID protocol.ID = "/exchange/response/0.0.1"
)

var (
	logger              = golog.Default.Child("protocols/mailbox")
	ErrMailboxDisabled  = errors.New("Mailbox not enabled, cannot perform request.")
	ErrMissingAPIKey    = errors.New("Missing Textile API Key in env")
	ErrMissingAPISecret = errors.New("Missing Textile API Secret in env")
	ErrRequestNotFound  = errors.New("Request not found in protocol cache")
)

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

// createRequest creates a new InviteRequest
func (p *ExchangeProtocol) createRequest(to *common.Peer, payload *common.Payload) (peer.ID, *InviteRequest, error) {
	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Fetch Peer ID from Public Key
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}

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
	return toId, req, nil
}

// createResponse creates a new InviteResponse
func (p *ExchangeProtocol) createResponse(decs bool, to *common.Peer) (peer.ID, *InviteResponse, error) {

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
