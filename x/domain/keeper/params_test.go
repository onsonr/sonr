package keeper_test

import (
	"testing"

	testkeeper "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/x/domain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DomainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
