package keeper_test

import (
	"context"
	"fmt"
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

func createDidDocumentsWithPrefix(keeper *keeper.Keeper, ctx sdk.Context, prefix string, n int) []types.DidDocument {
	items := make([]types.DidDocument, n)
	for i := range items {
		id := fmt.Sprintf("did:snr:%s%d", prefix, i)
		items[i].ID = id
		items[i].AlsoKnownAs = []string{
			fmt.Sprintf("FirstAka%d", i),
			fmt.Sprintf("SecondAka%d", i),
		}
		items[i].Service = &types.Services{
			Data: []*types.Service{
				{
					ID: fmt.Sprintf("%s#FirstSvc", id),
				},
				{
					ID: fmt.Sprintf("%s#SecondSvc", id),
				},
			},
		}
		items[i].VerificationMethod = &types.VerificationMethods{
			Data: []*types.VerificationMethod{
				{
					ID: fmt.Sprintf("%s#Key", id),
				},
			},
		}
		keeper.SetDidDocument(ctx, items[i])
	}
	return items
}

type KeeperTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	wCtx      context.Context
	keeper    *keeper.Keeper
	docs      []types.DidDocument
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
