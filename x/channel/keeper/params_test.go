package keeper_test

import (
	"testing"

	testkeeper "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/x/channel/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChannelKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
