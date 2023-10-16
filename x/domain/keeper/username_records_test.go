package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/core/testutil/keeper"
	"github.com/sonr-io/core/testutil/nullify"
	"github.com/sonr-io/core/x/domain/keeper"
	"github.com/sonr-io/core/x/domain/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUsernameRecords(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UsernameRecord {
	items := make([]types.UsernameRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUsernameRecords(ctx, items[i])
	}
	return items
}

func TestUsernameRecordsGet(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNUsernameRecords(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUsernameRecords(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUsernameRecordsRemove(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNUsernameRecords(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUsernameRecords(ctx,
			item.Index,
		)
		_, found := keeper.GetUsernameRecords(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUsernameRecordsGetAll(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNUsernameRecords(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUsernameRecords(ctx)),
	)
}
