package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
)

func setupMsgServer(t testing.T) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdentityKeeper(&t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
