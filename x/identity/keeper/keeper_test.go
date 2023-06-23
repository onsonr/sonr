package keeper_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	testutil "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/testutil/nullify"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RemoveIndex(s []types.Identification, index int) []types.Identification {
	return append(s[:index], s[index+1:]...)
}

func createDidDocumentsWithPrefix(keeper *keeper.Keeper, ctx sdk.Context, prefix string, n int) []types.DIDDocument {
	items := make([]types.DIDDocument, n)
	for i := range items {
		id := fmt.Sprintf("did:snr:%s%d", prefix, i)
		items[i].Id = id
		items[i].AlsoKnownAs = []string{
			fmt.Sprintf("FirstAka%d", i),
			fmt.Sprintf("SecondAka%d", i),
		}

		keeper.SetIdentity(ctx, items[i])
	}
	return items
}

type KeeperTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	wCtx      context.Context
	keeper    *keeper.Keeper
	docs      []types.Identification
	msgServer types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	k, ctx := testutil.IdentityKeeper(suite.T())
	suite.keeper = k
	suite.ctx = ctx
	suite.msgServer = keeper.NewMsgServerImpl(*k)
	suite.wCtx = sdk.WrapSDKContext(ctx)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func createNDidDocument(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DIDDocument {
	items := make([]types.DIDDocument, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)
		items[i].AlsoKnownAs = []string{strconv.Itoa(i)}

		keeper.SetIdentity(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestDidDocumentGet() {
	keeper := suite.keeper
	ctx := suite.ctx
	items := createNDidDocument(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetIdentity(ctx,
			item.Id,
		)
		suite.Assert().True(found)
		suite.Assert().Equal(
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func (suite *KeeperTestSuite) TestDidDocumentGetAll() {
	keeper := suite.keeper
	ctx := suite.ctx
	items := createNDidDocument(keeper, ctx, 10)
	suite.Assert().ElementsMatch(
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllIdentities(ctx)),
	)
}

func (suite *KeeperTestSuite) TestGetParams() {
	keeper := suite.keeper
	ctx := suite.ctx
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)
	suite.Assert().EqualValues(params, keeper.GetParams(ctx))
}
