package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateWhichIs_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateWhichIs
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateWhichIs{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateWhichIs{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateWhichIs_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateWhichIs
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateWhichIs{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateWhichIs{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteWhichIs_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteWhichIs
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteWhichIs{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteWhichIs{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
