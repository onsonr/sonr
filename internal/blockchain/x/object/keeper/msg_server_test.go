package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ObjectKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
