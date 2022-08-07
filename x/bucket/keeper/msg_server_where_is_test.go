package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sonr-io/sonr/x/bucket/types"
)

func TestWhereIsMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateWhereIs(ctx, &types.MsgCreateWhereIs{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, resp.Did)
	}
}

func TestWhereIsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateWhereIs
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateWhereIs{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateWhereIs{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateWhereIs{Creator: creator, Did: "did:sonr:B"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateWhereIs(ctx, &types.MsgCreateWhereIs{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateWhereIs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWhereIsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteWhereIs
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteWhereIs{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteWhereIs{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteWhereIs{Creator: creator, Did: "did:snr:10"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateWhereIs(ctx, &types.MsgCreateWhereIs{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteWhereIs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
