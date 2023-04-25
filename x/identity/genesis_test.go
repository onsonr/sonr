package identity_test

import (
	"testing"

	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/identity"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PrimaryIdentities: []types.DidDocument{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
		ClaimableWalletList: []types.ClaimableWallet{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ClaimableWalletCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PrimaryIdentities, got.PrimaryIdentities)
	require.ElementsMatch(t, genesisState.ClaimableWalletList, got.ClaimableWalletList)
	require.Equal(t, genesisState.ClaimableWalletCount, got.ClaimableWalletCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
