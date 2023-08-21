package service

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/service/keeper"
	"github.com/sonrhq/core/x/service/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the serviceRecord
	for _, elem := range genState.ServiceRecordList {
		k.SetServiceRecord(ctx, elem)
	}

	// Set serviceRecord count
	k.SetServiceRecordCount(ctx, genState.ServiceRecordCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ServiceRecordList = k.GetAllServiceRecord(ctx)
	genesis.ServiceRecordCount = k.GetServiceRecordCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
