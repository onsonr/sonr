package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonrhq/core/testutil/keeper"
	testkeeper "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/service/keeper"
	"github.com/sonrhq/core/x/service/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNServiceRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ServiceRecord {
	items := make([]types.ServiceRecord, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetServiceRecord(ctx, items[i])
	}
	return items
}

func TestServiceRecordGet(t *testing.T) {
	keeper, ctx := keepertest.ServiceKeeper(t)
	items := createNServiceRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetServiceRecord(ctx,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestServiceRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.ServiceKeeper(t)
	items := createNServiceRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveServiceRecord(ctx,
			item.Id,
		)
		_, found := keeper.GetServiceRecord(ctx,
			item.Id,
		)
		require.False(t, found)
	}
}

func TestServiceRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.ServiceKeeper(t)
	items := createNServiceRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllServiceRecord(ctx)),
	)
}


func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ServiceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
