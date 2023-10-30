package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateUsernameRecords_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateUsernameRecords
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateUsernameRecords{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateUsernameRecords{
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

func TestMsgUpdateUsernameRecords_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateUsernameRecords
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateUsernameRecords{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateUsernameRecords{
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

func TestMsgDeleteUsernameRecords_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteUsernameRecords
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteUsernameRecords{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteUsernameRecords{
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
