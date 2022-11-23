package keeper_test

import (
	"testing"

	testutil "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/x/identity/keeper"
	"github.com/sonr-io/sonr/x/identity/types"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RemoveIndex(s []types.DidDocument, index int) []types.DidDocument {
	return append(s[:index], s[index+1:]...)
}

type KeeperTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	keeper    *keeper.Keeper
	docs      []types.DidDocument
	msgServer types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	k, ctx := testutil.IdentityKeeper(suite.T())
	suite.keeper = k
	suite.ctx = ctx
	suite.msgServer = keeper.NewMsgServerImpl(*k)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
