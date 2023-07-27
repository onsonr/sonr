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

		DIDDocumentList: []types.DIDDocument{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},

		ControllerAccountList: []types.ControllerAccount{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		ControllerAccountCount: 2,
		EscrowAccountList: []types.EscrowAccount{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EscrowAccountCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DIDDocumentList, got.DIDDocumentList)
	require.ElementsMatch(t, genesisState.ControllerAccountList, got.ControllerAccountList)
	require.Equal(t, genesisState.ControllerAccountCount, got.ControllerAccountCount)
	require.ElementsMatch(t, genesisState.EscrowAccountList, got.EscrowAccountList)
	require.Equal(t, genesisState.EscrowAccountCount, got.EscrowAccountCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
