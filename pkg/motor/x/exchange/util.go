package exchange

import (
	"github.com/libp2p/go-libp2p-core/peer"

	ct "github.com/sonr-io/sonr/third_party/types/common"
	st "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
)

// // ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
// func ResponseToEvent(ir *v1.InviteResponse) *motor.OnTransmitDecisionResponse {
// 	return &motor.OnTransmitDecisionResponse{
// 		From:     ir.GetFrom(),
// 		Received: int64(time.Now().Unix()),
// 		Decision: ir.GetDecision(),
// 	}
// }

// // ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
// func RequestToEvent(ir *v1.InviteRequest) *motor.OnTransmitInviteResponse {
// 	return &motor.OnTransmitInviteResponse{
// 		Received: int64(time.Now().Unix()),
// 		From:     ir.GetFrom(),
// 	}
// }

// // createRequest creates a new InviteRequest
// func (p *ExchangeProtocol) createRequest(to *motor.Peer, payload *motor.Payload) (peer.ID, *v1.InviteRequest, error) {
// 	// // Call Peer from Node
// 	// from, err := p.node.Peer()
// 	// if err != nil {
// 	// 	logger.Errorf("%s - Failed to Get Peer from Node", err)
// 	// 	return "", nil, err
// 	// }

// 	// Fetch Peer ID from Public Key
// 	toId, err := Libp2pID(to)
// 	if err != nil {
// 		logger.Errorf("%s - Failed to fetch peer id from public key", err)
// 		return "", nil, err
// 	}

// 	// Create new Metadata
// 	// meta, err := wallet.CreateMetadata(p.host.ID())
// 	// if err != nil {
// 	// 	logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
// 	// 	return "", nil, err
// 	// }

// 	// Create Invite Request
// 	req := &v1.InviteRequest{
// 		// TODO: Implement Signed Meta to Proto Method
// 		// Metadata: api.SignedMetadataToProto(meta),
// 		To:   to,
// 		// From: from,
// 	}
// 	return toId, req, nil
// }

// createResponse creates a new InviteResponse
func (p *ExchangeProtocol) createResponse(decs bool, to *ct.Peer) (peer.ID, *st.InviteResponse, error) {

	// // Call Peer from Node
	// from, err := p.node.Peer()
	// if err != nil {
	// 	logger.Errorf("%s - Failed to Get Peer from Node", err)
	// 	return "", nil, err
	// }

	// Create new Metadata
	// meta, err := wallet.CreateMetadata(p.host.ID())
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
	// 	return "", nil, err
	// }

	// Create Invite Response
	resp := &st.InviteResponse{
		Decision: decs,
		// TODO: Implement Signed Meta to Proto Method
		//Metadata: api.SignedMetadataToProto(meta),
		// From: from,
		To: to,
	}

	// Fetch Peer ID from Public Key
	toId, err := Libp2pID(to)
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
}

// Libp2pID returns the PeerID based on PublicKey from Profile
func Libp2pID(p *ct.Peer) (peer.ID, error) {
	// Return Peer ID
	id, err := peer.IDFromString(p.GetPeerId())
	if err != nil {
		return "", err
	}
	return id, nil
}
