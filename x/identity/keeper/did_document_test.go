package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/identity/keeper"
	"github.com/sonr-io/sonr/x/identity/types"
)

func createNDidDocument(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DidDocument {
	items := make([]types.DidDocument, n)
	for i := range items {
		items[i].ID = strconv.Itoa(i)
		items[i].AlsoKnownAs = []string{strconv.Itoa(i)}

		keeper.SetDidDocument(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestDidDocumentGet() {
	keeper := suite.keeper
	ctx := suite.ctx
	items := createNDidDocument(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDidDocument(ctx,
			item.ID,
		)
		suite.Assert().True(found)
		suite.Assert().Equal(
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func (suite *KeeperTestSuite) TestDidDocumentRemove() {
	keeper := suite.keeper
	ctx := suite.ctx
	items := createNDidDocument(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDidDocument(ctx,
			item.ID,
		)
		_, found := keeper.GetDidDocument(ctx,
			item.ID,
		)
		suite.Assert().False(found)
	}
}

func (suite *KeeperTestSuite) TestDidDocumentGetAll() {
	keeper := suite.keeper
	ctx := suite.ctx
	items := createNDidDocument(keeper, ctx, 10)
	suite.Assert().ElementsMatch(
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDidDocument(ctx)),
	)
}

func (suite *KeeperTestSuite) TestGetDidDocumentByAKA() {
	keeper := suite.keeper
	ctx := suite.ctx
	items := createDidDocumentsWithPrefix(keeper, ctx, "AKA", 10)
	for _, item := range items {
		rst, found := keeper.GetDidDocumentByAKA(ctx,
			item.AlsoKnownAs[0],
		)
		suite.Assert().True(found)
		suite.Assert().Equal(
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
		keeper.RemoveDidDocument(ctx,
			item.ID,
		)
	}
}
