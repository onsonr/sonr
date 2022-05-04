package msg_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/msg"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RegistryKeeper(t)
	return msg.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
