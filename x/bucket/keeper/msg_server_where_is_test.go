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
		resp, err := srv.DefineBucket(ctx, &types.MsgDefineBucket{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, resp.WhereIs)
	}
}

func TestWhereIsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateBucket
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateBucket{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBucket{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBucket{Creator: creator, Did: "did:sonr:B"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.DefineBucket(ctx, &types.MsgDefineBucket{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateBucket(ctx, tc.request)
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
		request *types.MsgDeleteBucket
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteBucket{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteBucket{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteBucket{Creator: creator, Did: "did:snr:10"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.DefineBucket(ctx, &types.MsgDefineBucket{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteBucket(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
