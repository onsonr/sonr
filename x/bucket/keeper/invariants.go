package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/bucket/types"
)

const TOTAL_WHERE_IS = "total-where-is"

func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, TOTAL_WHERE_IS,
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
				TOTAL_WHERE_IS,
				fmt.Sprintf("Count of WhereIs: %v, whereIsCount in store: %v", len(allWhereIs), whereIsCount)),
			broken
	}
}
