package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/schema/keeper"
	"github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/require"
)

func createSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Schema {
	items := make([]types.Schema, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)
		keeper.SetSchema(ctx, items[i])
	}
	return items
}

func TestSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.SchemaKeeper(t)
	items := createSchema(keeper, ctx, 1)
	for _, item := range items {
		schema, found := keeper.GetSchema(ctx, item.Did)
		require.True(t, found)
		require.Equal(t, nullify.Fill(&item), nullify.Fill(&schema))
	}
}

func TestSchemaGetFromID(t *testing.T) {
	keeper, ctx := keepertest.SchemaKeeper(t)
	items := createSchema(keeper, ctx, 1)
	for i, item := range items {
		schema, found := keeper.GetSchemasFromID(ctx, item.Did)
		require.True(t, found)
		// bad, fix tmrrw
		require.Equal(t, nullify.Fill(&items[i]), nullify.Fill(&schema[i]))
	}
}
