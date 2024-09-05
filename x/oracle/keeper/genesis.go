package keeper

import (
	"github.com/onsonr/sonr/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the middlewares state from a specified GenesisState.
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
}

// ExportGenesis exports the middlewares state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{}
}
