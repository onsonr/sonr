package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/testutil/nullify"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWhoIs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WhoIs {
	items := make([]types.WhoIs, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)

		keeper.SetWhoIs(ctx, items[i])
	}
	return items
}

func TestWhoIsGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNWhoIs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWhoIs(ctx,
			item.Did,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWhoIsRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNWhoIs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhoIs(ctx,
			item.Did,
		)
		_, found := keeper.GetWhoIs(ctx,
			item.Did,
		)
		require.False(t, found)
	}
}

func TestWhoIsGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNWhoIs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhoIs(ctx)),
	)
}
