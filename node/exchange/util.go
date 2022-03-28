package exchange

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/util"

	v1 "go.buf.build/grpc/go/sonr-io/core/host/exchange/v1"
	motor "go.buf.build/grpc/go/sonr-io/core/motor/v1"
	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
)

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func ResponseToEvent(ir *v1.InviteResponse) *motor.OnTransmitDecisionResponse {
	return &motor.OnTransmitDecisionResponse{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func RequestToEvent(ir *v1.InviteRequest) *motor.OnTransmitInviteResponse {
	return &motor.OnTransmitInviteResponse{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}

// createRequest creates a new InviteRequest
func (p *ExchangeProtocol) createRequest(to *types.Peer, payload *types.Payload) (peer.ID, *v1.InviteRequest, error) {
	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Fetch Peer ID from Public Key
	toId, err := util.Libp2pID(to)
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}

	// Create new Metadata
	// meta, err := wallet.CreateMetadata(p.host.ID())
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
	// 	return "", nil, err
	// }

	// Create Invite Request
	req := &v1.InviteRequest{
		Payload: payload,
		// TODO: Implement Signed Meta to Proto Method
		// Metadata: api.SignedMetadataToProto(meta),
		To:   to,
		From: from,
	}
	return toId, req, nil
}

// createResponse creates a new InviteResponse
func (p *ExchangeProtocol) createResponse(decs bool, to *types.Peer) (peer.ID, *v1.InviteResponse, error) {

	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Create new Metadata
	// meta, err := wallet.CreateMetadata(p.host.ID())
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
	// 	return "", nil, err
	// }

	// Create Invite Response
	resp := &v1.InviteResponse{
		Decision: decs,
		// TODO: Implement Signed Meta to Proto Method
		//Metadata: api.SignedMetadataToProto(meta),
		From: from,
		To:   to,
	}

	// Fetch Peer ID from Public Key
	toId, err := util.Libp2pID(to)
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
}
