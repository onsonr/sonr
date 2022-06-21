package keeper_test

import (
	"testing"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestGetWhatIsFromCreator(t *testing.T) {
	k, ctx := keepertest.SchemaKeeper(t)
	items := keepertest.CreateWhatIsWithDID(k, ctx, 2)

	for _, item := range items {
		_, found := k.GetWhatIsFromCreator(ctx, item.Creator)
		require.True(t, found)
	}
}

func TestWhatIsGet(t *testing.T) {
	keeper, ctx := keepertest.SchemaKeeper(t)
	items := keepertest.CreateWhatIsWithDID(keeper, ctx, 1)
	for _, item := range items {
		_, found := keeper.GetWhatIs(ctx, item.Did)
		require.True(t, found)
	}
}

func TestWhatIsGetFromLabel(t *testing.T) {
	keeper, ctx := keepertest.SchemaKeeper(t)
	items := keepertest.CreateWhatIsWithDID(keeper, ctx, 1)
	for _, item := range items {
		_, found := keeper.GetWhatIsFromLabel(ctx, item.Schema.Label)
		require.True(t, found)
	}
}
