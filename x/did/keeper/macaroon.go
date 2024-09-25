package keeper

import (
	"crypto/sha256"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/macaroon.v2"
)

// IssueMacaroon creates a macaroon with the specified parameters.
func (k Keeper) IssueMacaroon(ctx sdk.Context, sharedMPCPubKey, location, id string, blockExpiry uint64) (*macaroon.Macaroon, error) {
	// Derive the root key by hashing the shared MPC public key
	rootKey := sha256.Sum256([]byte(sharedMPCPubKey))
	// Create the macaroon
	m, err := macaroon.New(rootKey[:], []byte(id), location, macaroon.LatestVersion)
	if err != nil {
		return nil, err
	}

	// Add the block expiry caveat
	caveat := fmt.Sprintf("block-expiry=%d", blockExpiry)
	err = m.AddFirstPartyCaveat([]byte(caveat))
	if err != nil {
		return nil, err
	}

	return m, nil
}
