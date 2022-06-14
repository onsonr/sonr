package keeper_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/did"
	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/x/schema/keeper"
	"github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/require"
)

func createSchemaWithDID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.WhatIs {
	items := make([]types.WhatIs, n)
	for i := range items {
		id := fmt.Sprintf("did:snr:%s", strconv.Itoa(i))
		items[i].Did = id
		doc, _ := did.NewDocument(id)
		var whatIs = types.WhatIs{
			Did: doc.GetID().String(),
			Schema: &types.SchemaReference{
				Did:   doc.GetID().String(),
				Label: "test",
				Cid:   strconv.Itoa(i),
			},
			Timestamp: time.Now().Unix(),
		}
		keeper.SetWhatIs(ctx, whatIs)
		items[i] = whatIs
	}
	return items
}

func TestWhatIsGet(t *testing.T) {
	keeper, ctx := keepertest.SchemaKeeper(t)
	items := createSchemaWithDID(keeper, ctx, 1)
	for _, item := range items {
		_, found := keeper.GetWhatIs(ctx, item.Did)
		require.True(t, found)
	}
}

func TestWhatIsGetFromLabel(t *testing.T) {
	keeper, ctx := keepertest.SchemaKeeper(t)
	items := createSchemaWithDID(keeper, ctx, 1)
	for _, item := range items {
		_, found := keeper.GetWhatIsFromLabel(ctx, item.Schema.Label)
		require.True(t, found)
	}
}
