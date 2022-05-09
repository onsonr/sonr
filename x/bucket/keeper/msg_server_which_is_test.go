package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/x/bucket/keeper"
	"github.com/sonr-io/sonr/x/bucket/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestWhichIsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BucketKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateWhichIs{
			Creator: creator,
			Did:     strconv.Itoa(i),
		}
		_, err := srv.CreateWhichIs(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetWhichIs(ctx,
			expected.Did,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestWhichIsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateWhichIs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateWhichIs{
				Creator: creator,
				Did:     strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateWhichIs{Creator: "B",
				Did: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateWhichIs{Creator: creator,
				Did: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BucketKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateWhichIs{Creator: creator,
				Did: strconv.Itoa(0),
			}
			_, err := srv.CreateWhichIs(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateWhichIs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetWhichIs(ctx,
					expected.Did,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestWhichIsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteWhichIs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteWhichIs{
				Creator: creator,
				Did:     strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteWhichIs{
				Creator: "B",
				Did:     strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteWhichIs{
				Creator: creator,
				Did:     strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BucketKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateWhichIs(wctx, &types.MsgCreateWhichIs{Creator: creator,
				Did: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteWhichIs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetWhichIs(ctx,
					tc.request.Did,
				)
				require.False(t, found)
			}
		})
	}
}
