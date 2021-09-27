package node

import (
	"errors"

	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/pkg/exchange"
)

// ToExchangeQueryRequest converts a query request to an exchange query request.
func (f *FindRequest) ToExchangeQueryRequest() (*exchange.QueryExchangeRequest, error) {
	if f.GetSName() != "" {
		return &exchange.QueryExchangeRequest{
			SName: f.GetSName(),
		}, nil
	}

	if f.GetPeerId() != "" {
		return &exchange.QueryExchangeRequest{
			PeerId: f.GetPeerId(),
		}, nil
	}
	return nil, errors.New("No SName or PeerID provided.")
}

// ToFindResponse converts PeerInfo to a FindResponse.
func ToFindResponse(p *common.PeerInfo) *FindResponse {
	return &FindResponse{
		Success: true,
		Peer:    p.Peer,
		PeerId:  p.PeerID.String(),
		SName:   p.SName,
	}
}
