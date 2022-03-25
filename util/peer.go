package util

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
)

// GetProfileFunc returns a function that returns the Profile and error
type GetProfileFunc func() (*types.Profile, error)

// Libp2pID returns the PeerID based on PublicKey from Profile
func Libp2pID(p *types.Peer) (peer.ID, error) {
	// Check if PublicKey is empty
	if len(p.GetPublicKey()) == 0 {
		return "", errors.New("Peer Public Key is not set.")
	}

	pubKey, err := crypto.UnmarshalPublicKey(p.GetPublicKey())
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

// PubKey returns the Public Key from the Peer
func PubKey(p *types.Peer) (crypto.PubKey, error) {
	// Check if PublicKey is empty
	if len(p.GetPublicKey()) == 0 {
		return nil, errors.New("Peer Public Key is not set.")
	}

	// Unmarshal Public Key
	pubKey, err := crypto.UnmarshalPublicKey(p.GetPublicKey())
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to Unmarshal Public Key: %s", p.GetSName()), err)
		return nil, err
	}
	return pubKey, nil
}

// OS returns Peer Device GOOS
func OS(p *types.Peer) string {
	return p.GetDevice().GetOs()
}
