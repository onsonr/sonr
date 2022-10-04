package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "test",
		TotalWhereIsCountInvariant(keeper))
}

func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, broken := TotalWhereIsCountInvariant(k)(ctx); broken {
			return res, broken
		}

		return "Every invariant condition is fulfilled correctly", true
	}
}

func TotalWhereIsCountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		whereIsCount := k.GetWhereIsCount(ctx)
		allWhereIs := k.GetAllWhereIs(ctx)
		broken := len(allWhereIs) != int(whereIsCount)
		return sdk.FormatInvariant(types.ModuleName,
				"total-where-is",
				fmt.Sprintf("Total WhereIs: %v, whereIsCount: %v", len(allWhereIs), whereIsCount)),
			broken
	}
}
