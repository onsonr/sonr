package keeper_test

import (
	"testing"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/stretchr/testify/require"
)

// func ToString(w types.WhatIs) string {
// 	return fmt.Sprintf("WhatIs: (Schema '%s')\nDID:\t%s\nCID:\t%s\nIsActive:\t%d\nCreator:\t%s\nTimestamp:\t%d\n",
// 		w.Schema.Label, w.Did, w.Schema.Cid, w.IsActive, w.Creator, w.Timestamp)
// }

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
