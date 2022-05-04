package msg_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/msg"
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChannelKeeper(t)
	return msg.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
