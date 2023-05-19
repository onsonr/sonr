package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the didDocument
	for _, elem := range genState.DidDocuments {
		k.SetDidDocument(ctx, elem)
	}
	for _, elem := range genState.Relationships {
		k.SetAuthentication(ctx, elem)
	}
	// Set all the claimableWallet
	for _, elem := range genState.ClaimableWalletList {
		k.SetClaimableWallet(ctx, elem)
	}

	// Set claimableWallet count
	k.SetClaimableWalletCount(ctx, genState.ClaimableWalletCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DidDocuments = k.GetAllPrimaryIdentities(ctx)
	genesis.Relationships = k.GetAllAuthentication(ctx)
	genesis.ClaimableWalletList = k.GetAllClaimableWallet(ctx)
	genesis.ClaimableWalletCount = k.GetClaimableWalletCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
