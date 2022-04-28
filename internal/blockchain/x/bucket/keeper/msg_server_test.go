package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/bucket/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/bucket/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BucketKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
