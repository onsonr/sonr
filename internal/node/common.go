package node

import (
	"errors"
	"strings"

	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/internal/store"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/tools/logger"
)

// NodeType is the type of the node (Client, Highway)
type NodeType int

const (
	// NodeType_CLIENT is the Node utilized by Desktop, Mobile and Web Clients
	NodeType_CLIENT NodeType = iota

	// NodeType_HIGHWAY is the Node utilized by long running Server processes
	NodeType_HIGHWAY
)

// Error Definitions
var (
	ErrEmptyQueue   = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery = errors.New("No SName or PeerID provided.")
)

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	// Get Public Key
	pubKey, err := device.KeyChain.GetSnrPubKey(keychain.Account)
	if err != nil {
		return nil, logger.Error("Failed to get Public Key", err)
	}

	// Find PublicKey Buffer
	deviceStat, err := device.Stat()
	if err != nil {
		return nil, logger.Error("Failed to get device Stat", err)
	}

	// Marshal Public Key
	pubBuf, err := pubKey.Buffer()
	if err != nil {
		return nil, logger.Error("Failed to marshal public key", err)
	}

	// Get Profile
	profile, err := n.Profile()
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

// Profile method returns the profile of the node
func (n *Node) Profile() (*common.Profile, error) {
	pro, err := n.store.GetProfile()
	if err != nil {
		return nil, logger.Error("Failed to retreive Profile", err)
	}
	return pro, nil
}

// Recents method returns the recent peers of the node
func (n *Node) Recents() (store.RecentsHistory, error) {
	rec, err := n.store.GetRecents()
	if err != nil {
		return nil, logger.Error("Failed to get recents", err)
	}
	return rec, nil
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
