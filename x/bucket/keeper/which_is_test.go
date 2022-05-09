package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/bucket/keeper"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWhichIs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WhichIs {
	items := make([]types.WhichIs, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)
		items[i].IsActive = true

		keeper.SetWhichIs(ctx, items[i])
	}
	return items
}

func TestWhichIsGet(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	items := createNWhichIs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWhichIs(ctx,
			item.Did,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWhichIsRemove(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	items := createNWhichIs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhichIs(ctx,
			item.Did,
		)
		_, found := keeper.GetWhichIs(ctx,
			item.Did,
		)
		require.False(t, found)
	}
}

func TestWhichIsGetAll(t *testing.T) {
	keeper, ctx := keepertest.BucketKeeper(t)
	items := createNWhichIs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhichIs(ctx)),
	)
}
