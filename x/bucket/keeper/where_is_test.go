package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/bucket/keeper"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/require"
)

func createNWhereIs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WhereIs {
	items := make([]types.WhereIs, n)
	for i := range items {
		items[i].Did = keeper.AppendWhereIs(ctx, items[i])
	}
	return items
}

func TestWhereIsGet(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	items := createNWhereIs(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetWhereIs(ctx, item.Did)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestWhereIsRemove(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	items := createNWhereIs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhereIs(ctx, item.Did)
		_, found := keeper.GetWhereIs(ctx, item.Did)
		require.False(t, found)
	}
}

func TestWhereIsGetAll(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	items := createNWhereIs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhereIs(ctx)),
	)
}

func TestWhereIsCount(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	items := createNWhereIs(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetWhereIsCount(ctx))
}
