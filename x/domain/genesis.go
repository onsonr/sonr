package domain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/domain/keeper"
	"github.com/sonrhq/core/x/domain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the tLDRecord
	for _, elem := range genState.TLDRecordList {
		k.SetTLDRecord(ctx, elem)
	}
	// Set all the sLDRecord
	for _, elem := range genState.SLDRecordList {
		k.SetSLDRecord(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TLDRecordList = k.GetAllTLDRecord(ctx)
	genesis.SLDRecordList = k.GetAllSLDRecord(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
