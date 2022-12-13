package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-hq/sonr/testutil/keeper"
	"github.com/sonr-hq/sonr/x/identity/keeper"
	"github.com/sonr-hq/sonr/x/identity/types"
)

func setupMsgServer(t *testing.T) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdentityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
