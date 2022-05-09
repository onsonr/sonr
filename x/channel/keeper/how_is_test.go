package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/channel/keeper"
	"github.com/sonr-io/sonr/x/channel/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNHowIs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.HowIs {
	items := make([]types.HowIs, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)
		items[i].IsActive = true

		keeper.SetHowIs(ctx, items[i])
	}
	return items
}

func TestHowIsGet(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNHowIs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetHowIs(ctx,
			item.Did,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestHowIsRemove(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNHowIs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHowIs(ctx,
			item.Did,
		)
		_, found := keeper.GetHowIs(ctx,
			item.Did,
		)
		require.False(t, found)
	}
}

func TestHowIsGetAll(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNHowIs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHowIs(ctx)),
	)
}
