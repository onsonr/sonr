package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateWhatIs_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateWhatIs
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateWhatIs{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateWhatIs{
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

func TestMsgUpdateWhatIs_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateWhatIs
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateWhatIs{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateWhatIs{
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

func TestMsgDeleteWhatIs_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteWhatIs
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteWhatIs{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteWhatIs{
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
