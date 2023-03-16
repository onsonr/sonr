package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDomainRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Service {
	items := make([]types.Service, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetService(ctx, items[i])
	}
	return items
}

func TestDomainRecordGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDomainRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetService(ctx,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestDomainRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDomainRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllServices(ctx)),
	)
}
