package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

// GetExpirationBlockHeight returns the block height at which the given duration will have passed
func (k Keeper) GetExpirationBlockHeight(ctx sdk.Context, duration time.Duration) int64 {
	return ctx.BlockHeight() + int64(duration.Seconds()/k.GetAverageBlockTime(ctx))
}

// ValidServiceOrigin checks if a service origin is valid
func (k Keeper) ValidServiceOrigin(ctx sdk.Context, origin string) bool {
	rec, err := k.OrmDB.ServiceRecordTable().GetByOriginUri(ctx, origin)
	if err != nil {
		return false
	}
	if rec == nil {
		return false
	}
	return true
}

// VerifyMinimumStake checks if a validator has a minimum stake
func (k Keeper) VerifyMinimumStake(ctx sdk.Context, addr string) bool {
	address, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return false
	}
	addval, err := sdk.ValAddressFromBech32(addr)
	if err != nil {
		return false
	}
	del, err := k.StakingKeeper.GetDelegation(ctx, address, addval)
	if err != nil {
		return false
	}
	if del.Shares.IsZero() {
		return false
	}
	return del.Shares.IsPositive()
}

// VerifyServicePermissions checks if a service has permission
func (k Keeper) VerifyServicePermissions(ctx sdk.Context, addr string, service string, permissions string) bool {
	return false
}
