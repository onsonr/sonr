package domain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sonr.io/core/x/domain/keeper"
	"sonr.io/core/x/domain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the UsernameRecord
	for _, elem := range genState.UsernameRecordsList {
		k.SetUsernameRecords(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.UsernameRecordsList = k.GetAllUsernameRecords(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
