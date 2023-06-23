package vault_test

import (
	"testing"

	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/vault"
	"github.com/sonrhq/core/x/vault/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClaimableWalletList: []types.ClaimableWallet{
			{
				Index: 0,
			},
			{
				Index: 1,
			},
		},
		ClaimableWalletCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VaultKeeper(t)
	vault.InitGenesis(ctx, *k, genesisState)
	got := vault.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClaimableWalletList, got.ClaimableWalletList)
	require.Equal(t, genesisState.ClaimableWalletCount, got.ClaimableWalletCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
