package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-hq/sonr/testutil/keeper"
	"github.com/sonr-hq/sonr/testutil/nullify"
	"github.com/sonr-hq/sonr/x/identity/keeper"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDomainRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DomainRecord {
	items := make([]types.DomainRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetDomainRecord(ctx, items[i])
	}
	return items
}

func TestDomainRecordGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDomainRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDomainRecord(ctx,
			item.Index,
			item.Domain,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDomainRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDomainRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDomainRecord(ctx,
			item.Index,
			item.Domain,
		)
		_, found := keeper.GetDomainRecord(ctx,
			item.Index,
			item.Domain,
		)
		require.False(t, found)
	}
}

func TestDomainRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDomainRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDomainRecord(ctx)),
	)
}
