package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/identity/keeper"
	"github.com/sonr-io/sonr/x/identity/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDidDocument(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DidDocument {
	items := make([]types.DidDocument, n)
	for i := range items {
		items[i].ID = strconv.Itoa(i)

		keeper.SetDidDocument(ctx, items[i])
	}
	return items
}

func TestDidDocumentGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDidDocument(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDidDocument(ctx,
			item.ID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDidDocumentRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDidDocument(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDidDocument(ctx,
			item.ID,
		)
		_, found := keeper.GetDidDocument(ctx,
			item.ID,
		)
		require.False(t, found)
	}
}

func TestDidDocumentGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNDidDocument(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDidDocument(ctx)),
	)
}
