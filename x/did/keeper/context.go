package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	nftkeeper "cosmossdk.io/x/nft/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/onsonr/sonr/pkg/vault"
	"github.com/onsonr/sonr/x/did/types"
	"google.golang.org/grpc/peer"
	"gopkg.in/macaroon.v2"
)

func (k Keeper) UnwrapCtx(goCtx context.Context) Context {
	ctx := sdk.UnwrapSDKContext(goCtx)
	peer, _ := peer.FromContext(goCtx)
	return Context{SDKCtx: ctx, Peer: peer, Keeper: k}
}

type Context struct {
	SDKCtx    sdk.Context
	Keeper    Keeper
	Peer      *peer.Peer
	NFTKeeper nftkeeper.Keeper
}

// AssembleVault assembles the initial vault
func (k Keeper) AssembleVault(ctx Context, subject string, origin string) (string, int64, error) {
	v, err := vault.New(subject, origin, "sonr-testnet")
	if err != nil {
		return "", 0, err
	}
	cid, err := k.ipfsClient.Unixfs().Add(context.Background(), v.FS)
	if err != nil {
		return "", 0, err
	}
	return cid.String(), ctx.CalculateExpiration(time.Second * 15), nil
}

// AverageBlockTime returns the average block time in seconds
func (c Context) AverageBlockTime() float64 {
	return float64(c.SDK().BlockTime().Sub(c.SDK().BlockTime()).Seconds())
}

// GetExpirationBlockHeight returns the block height at which the given duration will have passed
func (c Context) CalculateExpiration(duration time.Duration) int64 {
	return c.SDKCtx.BlockHeight() + int64(duration.Seconds()/c.AverageBlockTime())
}

// IPFSConnected returns true if the IPFS client is initialized
func (c Context) IPFSConnected() bool {
	if c.Keeper.ipfsClient == nil {
		ipfsClient, err := rpc.NewLocalApi()
		if err != nil {
			return false
		}
		c.Keeper.ipfsClient = ipfsClient
	}
	return c.Keeper.ipfsClient != nil
}

func (c Context) IsAnonymous() bool {
	if c.Peer == nil {
		return true
	}
	return c.Peer.Addr == nil
}

// IssueMacaroon creates a macaroon with the specified parameters.
func (c Context) IssueMacaroon(sharedMPCPubKey, location, id string, blockExpiry uint64) (*macaroon.Macaroon, error) {
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

func (c Context) Params() *types.Params {
	p, err := c.Keeper.Params.Get(c.SDK())
	if err != nil {
		p = types.DefaultParams()
	}
	params := p.ActiveParams(c.IPFSConnected())
	return &params
}

func (c Context) PeerID() string {
	if c.Peer == nil {
		return ""
	}
	return c.Peer.Addr.String()
}

// PinVaultController pins the initial vault to the local IPFS node
func (k Keeper) PinVaultController(_ sdk.Context, cid string, address string) (bool, error) {
	// Resolve the path
	path, err := path.NewPath(cid)
	if err != nil {
		return false, err
	}

	// 1. Initialize vault.db sqlite database in local IPFS with Mount

	// 2. Insert the InitialWalletAccounts

	// 3. Publish the path to the IPNS
	_, err = k.ipfsClient.Name().Publish(context.Background(), path, options.Name.Key(address))
	if err != nil {
		return false, err
	}

	// 4. Insert the accounts into x/auth

	// 5. Insert the controller into state
	return true, nil
}

func (c Context) SDK() sdk.Context {
	return c.SDKCtx
}

// ValidateOrigin checks if a service origin is valid
func (c Context) ValidateOrigin(origin string) error {
	if origin == "localhost" {
		return nil
	}
	return types.ErrInvalidServiceOrigin
}

// VerifyMinimumStake checks if a validator has a minimum stake
func (c Context) VerifyMinimumStake(addr string) bool {
	address, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return false
	}
	addval, err := sdk.ValAddressFromBech32(addr)
	if err != nil {
		return false
	}
	del, err := c.Keeper.StakingKeeper.GetDelegation(c.SDK(), address, addval)
	if err != nil {
		return false
	}
	if del.Shares.IsZero() {
		return false
	}
	return del.Shares.IsPositive()
}
