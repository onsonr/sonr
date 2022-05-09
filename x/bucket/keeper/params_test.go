package keeper_test

import (
	"testing"

	testkeeper "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BucketKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
