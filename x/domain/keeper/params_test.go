package keeper_test

import (
	"testing"

	testkeeper "sonr.io/core/testutil/keeper"
	"sonr.io/core/x/domain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DomainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
