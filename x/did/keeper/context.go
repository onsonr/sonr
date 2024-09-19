package keeper

import (
	"time"

	"github.com/ipfs/kubo/client/rpc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onsonr/sonr/x/did/types"
	"google.golang.org/grpc/peer"
)

type Context struct {
	SDKCtx sdk.Context
	Keeper Keeper
	Peer   *peer.Peer
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
