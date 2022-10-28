package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sonr-io/sonr/x/bucket/types"
)

func TestWhereIsMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "cosmos1pvnkmcpmtsxjuprqvu5qsdn2rnlenwnqsh276f"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateWhereIs(ctx, &types.MsgCreateWhereIs{Creator: creator, Label: strconv.Itoa(i)})
		require.NoError(t, err)
		require.Equal(t, strconv.Itoa(i), resp.WhereIs.Label)
		require.Equal(t, creator, resp.WhereIs.Creator)
	}
}

func TestWhereIsMsgServerUpdate(t *testing.T) {
	creator := "cosmos1pvnkmcpmtsxjuprqvu5qsdn2rnlenwnqsh276f"
	intruder := "cosmos1pvnkmcpmtsexuprqvu5qsdn2rnlenwnqkx66ky"

	srv, ctx := setupMsgServer(t)
	resp, err := srv.CreateWhereIs(ctx, &types.MsgCreateWhereIs{Creator: creator, Visibility: types.BucketVisibility_PUBLIC})
	require.NoError(t, err)
	did1 := resp.WhereIs.Did
	resp, err = srv.CreateWhereIs(ctx, &types.MsgCreateWhereIs{Creator: creator, Visibility: types.BucketVisibility_PUBLIC})
	require.NoError(t, err)
	did2 := resp.WhereIs.Did

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateWhereIs
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateWhereIs{Creator: creator, Did: did1},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateWhereIs{Creator: intruder, Did: did2},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateWhereIs{Creator: creator, Did: "did:sonr:B"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
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
	creator := "cosmos1pvnkmcpmtsxjuprqvu5qsdn2rnlenwnqsh276f"
	intruder := "cosmos1pvnkmcpmtsexuprqvu5qsdn2rnlenwnqkx66ky"

	srv, ctx := setupMsgServer(t)
	resp, err := srv.CreateWhereIs(ctx, &types.MsgCreateWhereIs{Creator: creator, Visibility: types.BucketVisibility_PUBLIC})
	require.NoError(t, err)
	did1 := resp.WhereIs.Did
	resp, err = srv.CreateWhereIs(ctx, &types.MsgCreateWhereIs{Creator: creator, Visibility: types.BucketVisibility_PUBLIC})
	require.NoError(t, err)
	did2 := resp.WhereIs.Did

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteWhereIs
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteWhereIs{Creator: creator, Did: did1},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteWhereIs{Creator: intruder, Did: did2},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteWhereIs{Creator: creator, Did: "did:snr:10"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DeleteWhereIs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
