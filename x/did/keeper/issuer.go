package keeper

import (
	"crypto/sha256"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	didtypes "github.com/onsonr/sonr/x/did/types"
	"gopkg.in/macaroon.v2"
)

var fourYears = time.Hour * 24 * 365 * 4

// IssueAdminMacaroon creates a macaroon with the specified parameters.
func (k Keeper) IssueAdminMacaroon(sdkctx sdk.Context, controller didtypes.ControllerI) (*macaroon.Macaroon, error) {
	// Derive the root key by hashing the shared MPC public key
	rootKey := sha256.Sum256([]byte(controller.RawPublicKey()))
	// Create the macaroon
	m, err := macaroon.New(rootKey[:], []byte(controller.SonrAddress()), controller.ChainID(), macaroon.LatestVersion)
	if err != nil {
		return nil, err
	}

	// Add the block expiry caveat
	caveat := fmt.Sprintf("block-expiry=%d", calculateBlockExpiry(sdkctx, fourYears))
	err = m.AddFirstPartyCaveat([]byte(caveat))
	if err != nil {
		return nil, err
	}

	return m, nil
}

// IssueServiceMacaroon creates a macaroon with the specified parameters.
func (k Keeper) IssueServiceMacaroon(sdkctx sdk.Context, sharedMPCPubKey, location, id string, blockExpiry uint64) (*macaroon.Macaroon, error) {
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

func calculateBlockExpiry(sdkctx sdk.Context, duration time.Duration) uint64 {
	blockTime := sdkctx.BlockTime()
	avgBlockTime := float64(blockTime.Sub(blockTime).Seconds())
	return uint64(duration.Seconds() / avgBlockTime)
}
