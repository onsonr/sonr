package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/object/keeper"
	"github.com/sonr-io/sonr/x/object/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWhatIs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WhatIs {
	items := make([]types.WhatIs, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)
		items[i].IsActive = true

		keeper.SetWhatIs(ctx, items[i])
	}
	return items
}

func TestWhatIsGet(t *testing.T) {
	keeper, ctx := keepertest.ObjectKeeper(t)
	items := createNWhatIs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWhatIs(ctx,
			item.Did,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWhatIsRemove(t *testing.T) {
	keeper, ctx := keepertest.ObjectKeeper(t)
	items := createNWhatIs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhatIs(ctx,
			item.Did,
		)
		_, found := keeper.GetWhatIs(ctx,
			item.Did,
		)
		require.False(t, found)
	}
}

func TestWhatIsGetAll(t *testing.T) {
	keeper, ctx := keepertest.ObjectKeeper(t)
	items := createNWhatIs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhatIs(ctx)),
	)
}
