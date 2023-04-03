package keeper_test


import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sonrhq/core/testutil/keeper"
	"github.com/sonrhq/core/x/service/keeper"
	"github.com/sonrhq/core/x/service/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestServiceRecordMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ServiceKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	Controller := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateServiceRecord{Controller: Controller,
			Id: strconv.Itoa(i),
		}
		_, err := srv.CreateServiceRecord(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetServiceRecord(ctx,
			expected.Id,
		)
		require.True(t, found)
		require.Equal(t, expected.Controller, rst.Controller)
	}
}

func TestServiceRecordMsgServerUpdate(t *testing.T) {
	Controller := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateServiceRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateServiceRecord{Controller: Controller,
				Id: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateServiceRecord{Controller: "B",
				Id: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateServiceRecord{Controller: Controller,
				Id: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ServiceKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateServiceRecord{Controller: Controller,
				Id: strconv.Itoa(0),
			}
			_, err := srv.CreateServiceRecord(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateServiceRecord(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetServiceRecord(ctx,
					expected.Id,
				)
				require.True(t, found)
				require.Equal(t, expected.Controller, rst.Controller)
			}
		})
	}
}

func TestServiceRecordMsgServerDelete(t *testing.T) {
	Controller := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteServiceRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteServiceRecord{Controller: Controller,
				Id: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteServiceRecord{Controller: "B",
				Id: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteServiceRecord{Controller: Controller,
				Id: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ServiceKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateServiceRecord(wctx, &types.MsgCreateServiceRecord{Controller: Controller,
				Id: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteServiceRecord(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetServiceRecord(ctx,
					tc.request.Id,
				)
				require.False(t, found)
			}
		})
	}
}
