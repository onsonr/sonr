package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestWhoIsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RegistryKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateWhoIs{
			Creator: creator,
			Did:     strconv.Itoa(i),
		}
		_, err := srv.CreateWhoIs(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetWhoIs(ctx,
			expected.Did,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestWhoIsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateWhoIs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateWhoIs{Creator: creator,
				Did: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateWhoIs{Creator: "B",
				Did: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateWhoIs{Creator: creator,
				Did: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RegistryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateWhoIs{Creator: creator,
				Did: strconv.Itoa(0),
			}
			_, err := srv.CreateWhoIs(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateWhoIs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetWhoIs(ctx,
					expected.Did,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestWhoIsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteWhoIs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteWhoIs{Creator: creator,
				Did: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteWhoIs{Creator: "B",
				Did: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteWhoIs{Creator: creator,
				Did: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RegistryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateWhoIs(wctx, &types.MsgCreateWhoIs{Creator: creator,
				Did: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteWhoIs(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetWhoIs(ctx,
					tc.request.Did,
				)
				require.False(t, found)
			}
		})
	}
}
