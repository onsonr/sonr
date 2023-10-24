package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/core/x/identity/keeper"
	"github.com/sonr-io/core/x/identity/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the dIDDocument
	for _, elem := range genState.DIDDocumentList {
		k.SetDIDDocument(ctx, elem)
	}
	// Set all the controllerAccount
	for _, elem := range genState.ControllerAccountList {
		k.SetControllerAccount(ctx, elem)
	}

	// Set controllerAccount count
	k.SetControllerAccountCount(ctx, genState.ControllerAccountCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.DIDDocumentList = k.GetAllDIDDocument(ctx)
	genesis.ControllerAccountList = k.GetAllControllerAccount(ctx)
	genesis.ControllerAccountCount = k.GetControllerAccountCount(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
