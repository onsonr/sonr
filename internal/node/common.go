package node

import (
	"errors"
	"strings"

	"github.com/libp2p/go-libp2p-core/crypto"
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/tools/logger"
)

// Error Definitions
var (
	ErrEmptyQueue   = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery = errors.New("No SName or PeerID provided.")
)

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	// Get Public Key
	pubKey, err := device.KeyChain.GetPubKey(device.Account)
	if err != nil {
		return nil, logger.Error("Failed to get Public Key", err)
	}

	// Find PublicKey Buffer
	deviceStat, err := device.Stat()
	if err != nil {
		return nil, logger.Error("Failed to get device Stat", err)
	}

	// Marshal Public Key
	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, logger.Error("Failed to marshal public key", err)
	}

	// Get Profile
	profile, err := n.store.GetProfile()
	if err != nil {
		return nil, err
	}

	// Return Peer
	return &common.Peer{
		SName:     strings.ToLower(profile.SName),
		Status:    common.Peer_ONLINE,
		Profile:   profile,
		PublicKey: pubBuf,
		Device: &common.Peer_Device{
			HostName: deviceStat.HostName,
			Os:       deviceStat.Os,
			Id:       deviceStat.Id,
			Arch:     deviceStat.Arch,
		},
	}, nil
}

// ToExchangeQueryRequest converts a query request to an exchange query request.
func (f *SearchRequest) ToExchangeQueryRequest() (*exchange.QueryRequest, error) {
	if f.GetSName() != "" {
		return &exchange.QueryRequest{
			SName: f.GetSName(),
		}, nil
	}

	if f.GetPeerId() != "" {
		return &exchange.QueryRequest{
			PeerId: f.GetPeerId(),
		}, nil
	}
	return nil, logger.Error("Failed to convert FindRequest", ErrInvalidQuery)
}

// ToFindResponse converts PeerInfo to a FindResponse.
func ToFindResponse(p *common.PeerInfo) *SearchResponse {
	return &SearchResponse{
		Success: true,
		Peer:    p.Peer,
		PeerId:  p.PeerID.String(),
		SName:   p.SName,
	}
}
