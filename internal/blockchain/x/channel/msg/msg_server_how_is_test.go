package msg_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/msg"
	"github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestHowIsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ChannelKeeper(t)
	srv := msg.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateHowIs{Creator: creator,
			Did: strconv.Itoa(i),
		}
		_, err := srv.CreateHowIs(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetHowIs(ctx,
			expected.Did,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestHowIsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateHowIs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateHowIs{Creator: creator,
				Did: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateHowIs{Creator: "B",
				Did: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateHowIs{Creator: creator,
				Did: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ChannelKeeper(t)
			srv := msg.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateHowIs{Creator: creator,
				Did: strconv.Itoa(0),
			}
			_, err := srv.CreateHowIs(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateHowIs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetHowIs(ctx,
					expected.Did,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestHowIsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteHowIs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteHowIs{Creator: creator,
				Did: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteHowIs{Creator: "B",
				Did: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteHowIs{Creator: creator,
				Did: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ChannelKeeper(t)
			srv := msg.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateHowIs(wctx, &types.MsgCreateHowIs{Creator: creator,
				Did: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteHowIs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetHowIs(ctx,
					tc.request.Did,
				)
				require.False(t, found)
			}
		})
	}
}
