package keeper

import (
	"context"
	"time"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onsonr/sonr/x/did/types"
)

// Logger returns the logger
func (k Keeper) Logger() log.Logger {
	return k.logger
}

// InitGenesis initializes the module's state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {
	// this line is used by starport scaffolding # genesis/module/init
	if err := data.Params.Validate(); err != nil {
		return err
	}

	return k.Params.Set(ctx, data.Params)
}

// ExportGenesis exports the module's state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	// this line is used by starport scaffolding # genesis/module/export

	return &types.GenesisState{
		Params: params,
	}
}

// CheckValidatorExists checks if a validator exists
func (k Keeper) CheckValidatorExists(ctx sdk.Context, addr string) bool {
	address, err := sdk.ValAddressFromBech32(addr)
	if err != nil {
		return false
	}
	ok, err := k.StakingKeeper.Validator(ctx, address)
	if err != nil {
		return false
	}
	if ok != nil {
		return true
	}
	return false
}

// GetAverageBlockTime returns the average block time in seconds
func (k Keeper) GetAverageBlockTime(ctx sdk.Context) float64 {
	return float64(ctx.BlockTime().Sub(ctx.BlockTime()).Seconds())
}

// GetParams returns the module parameters.
func (k Keeper) GetParams(ctx sdk.Context) *types.Params {
	p, err := k.Params.Get(ctx)
	if err != nil {
		p = types.DefaultParams()
	}
	return &p
}

// GetExpirationBlockHeight returns the block height at which the given duration will have passed
func (k Keeper) GetExpirationBlockHeight(ctx sdk.Context, duration time.Duration) int64 {
	return ctx.BlockHeight() + int64(duration.Seconds()/k.GetAverageBlockTime(ctx))
}
