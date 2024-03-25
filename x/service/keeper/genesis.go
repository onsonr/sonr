package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/didao-org/sonr/x/service/types"
)

// InitGenesis initializes the middlewares state from a specified GenesisState.
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	for _, record := range state.Records {
		k.RecordsMapping.Set(ctx, record.Origin, record)
	}
}

// ExportGenesis exports the middlewares state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{}
}
