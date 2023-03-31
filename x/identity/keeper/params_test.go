package keeper_test

import (
	"github.com/sonr-io/sonr/x/identity/types"
)

func (suite *KeeperTestSuite) TestGetParams() {
	keeper := suite.keeper
	ctx := suite.ctx
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)
	suite.Assert().EqualValues(params, keeper.GetParams(ctx))
}
