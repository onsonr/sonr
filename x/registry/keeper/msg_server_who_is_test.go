package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sonr-io/sonr/x/registry/types"
)

func TestWhoIsMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	owner := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateWhoIs(ctx, &types.MsgCreateWhoIs{Creator: owner})
		require.NoError(t, err)
		require.Equal(t, i, resp.Did)
	}
}

func TestWhoIsMsgServerUpdate(t *testing.T) {
	owner := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateWhoIs
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateWhoIs{Creator: owner},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateWhoIs{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateWhoIs{Creator: owner, Did: "10"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateWhoIs(ctx, &types.MsgCreateWhoIs{Creator: owner})
			require.NoError(t, err)

			_, err = srv.UpdateWhoIs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWhoIsMsgServerDelete(t *testing.T) {
	owner := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeactivateWhoIs
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeactivateWhoIs{Creator: owner},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeactivateWhoIs{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeactivateWhoIs{Creator: owner, Did: "10"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateWhoIs(ctx, &types.MsgCreateWhoIs{Creator: owner})
			require.NoError(t, err)
			_, err = srv.DeleteWhoIs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
