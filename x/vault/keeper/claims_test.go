package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/vault/keeper"
	"github.com/sonrhq/core/x/vault/types"
	"github.com/stretchr/testify/require"
)

func createNClaimableWallet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ClaimableWallet {
	items := make([]types.ClaimableWallet, n)
	for i := range items {
		items[i].Id = keeper.AppendClaimableWallet(ctx, items[i])
	}
	return items
}

func TestClaimableWalletGet(t *testing.T) {
	keeper, ctx := keepertest.VaultKeeper(t)
	items := createNClaimableWallet(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetClaimableWallet(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestClaimableWalletRemove(t *testing.T) {
	keeper, ctx := keepertest.VaultKeeper(t)
	items := createNClaimableWallet(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClaimableWallet(ctx, item.Id)
		_, found := keeper.GetClaimableWallet(ctx, item.Id)
		require.False(t, found)
	}
}

func TestClaimableWalletGetAll(t *testing.T) {
	keeper, ctx := keepertest.VaultKeeper(t)
	items := createNClaimableWallet(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClaimableWallet(ctx)),
	)
}

func TestClaimableWalletCount(t *testing.T) {
	keeper, ctx := keepertest.VaultKeeper(t)
	items := createNClaimableWallet(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetClaimableWalletCount(ctx))
}
