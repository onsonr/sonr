package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.T) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdentityKeeper(&t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDidDocumentMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.IdentityKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDidDocument{Creator: creator,
			Primary: types.NewSonrIdentity(creator),
		}
		_, _ = srv.CreateDidDocument(wctx, expected)
		rst, found := k.GetDidDocument(ctx,
			expected.Primary.Id,
		)
		accAddr, err := rst.AccAddress()
		require.Error(t, err)
		require.True(t, found)
		require.NotEqual(t, expected.Creator, accAddr)
	}
}

func TestDidDocumentMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDidDocument
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDidDocument{Creator: creator,
				Primary: types.NewSonrIdentity(creator),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDidDocument{Creator: "B",
				Primary: types.NewSonrIdentity(creator),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDidDocument{Creator: creator,
				Primary: types.NewSonrIdentity(creator),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IdentityKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateDidDocument{Creator: creator,
				Primary: types.NewSonrIdentity(creator),
			}
			_, err := srv.CreateDidDocument(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDidDocument(wctx, tc.request)
			if tc.err != nil {
				require.Error(t, err, tc.err)
			} else {
				//require.NoError(t, err)
				rst, found := k.GetDidDocument(ctx,
					expected.Primary.Id,
				)
				require.True(t, found)
				accAddr, err := rst.AccAddress()
				require.Error(t, err)
				require.NotEqual(t, expected.Creator, accAddr)
			}
		})
	}
}
