package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonrhq/core/testutil/keeper"
	testkeeper "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/domain/keeper"
	"github.com/sonrhq/core/x/domain/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DomainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

func createNSLDRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SLDRecord {
	items := make([]types.SLDRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSLDRecord(ctx, items[i])
	}
	return items
}

func TestSLDRecordGet(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNSLDRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSLDRecord(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSLDRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNSLDRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSLDRecord(ctx,
			item.Index,
		)
		_, found := keeper.GetSLDRecord(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSLDRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNSLDRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSLDRecord(ctx)),
	)
}

func createNTLDRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TLDRecord {
	items := make([]types.TLDRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTLDRecord(ctx, items[i])
	}
	return items
}

func TestTLDRecordGet(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNTLDRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTLDRecord(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTLDRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNTLDRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTLDRecord(ctx,
			item.Index,
		)
		_, found := keeper.GetTLDRecord(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTLDRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	items := createNTLDRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTLDRecord(ctx)),
	)
}

func TestDNSRecords(t *testing.T) {
	records, err := keeper.ResolveHNSTLD(types.WithDomains("sonr", "welcome.nb"))
	if err != nil {
		t.Fatal(err)
	}
	if len(records) == 0 {
		t.Fatal("no records found")
	}
	for _, record := range records {
		t.Logf("%+v", record)
	}
}
