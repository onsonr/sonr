package exchange

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/libp2p/go-libp2p-core/crypto"
	dv1 "github.com/sonr-io/sonr/internal/motor/x/discover/v1"
	v1 "github.com/sonr-io/sonr/internal/motor/x/exchange/v1"
	tv1 "github.com/sonr-io/sonr/internal/motor/x/transmit/v1"
)

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func RequestToEvent(ir *v1.InviteRequest) *tv1.OnTransmitInviteResponse {
	return &tv1.OnTransmitInviteResponse{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
	}
}

// createRequest creates a new InviteRequest
func (p *ExchangeProtocol) createRequest(to *dv1.Peer, payload *tv1.Payload) (peer.ID, *v1.InviteRequest, error) {
	// Call Peer from Node
	// from, err := p.node.Peer()
	// if err != nil {
	// 	logger.Errorf("%s - Failed to Get Peer from Node", err)
	// 	return "", nil, err
	// }

	// Fetch Peer ID from Public Key
	toId, err := Libp2pID(to)
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
		To: to,
		// From: from,
	}
	return toId, req, nil
}

// createResponse creates a new InviteResponse
func (p *ExchangeProtocol) createResponse(decs bool, to *dv1.Peer) (peer.ID, *v1.InviteResponse, error) {

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
	resp := &v1.InviteResponse{
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
func Libp2pID(p *dv1.Peer) (peer.ID, error) {
	pubKey, err := crypto.UnmarshalPublicKey(nil)
	if err != nil {
		return "", err
	}

	// Return Peer ID
	id, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	return id, nil
}
