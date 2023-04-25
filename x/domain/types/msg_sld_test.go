package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateSLDRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSLDRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSLDRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSLDRecord{
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

func TestMsgUpdateSLDRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSLDRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSLDRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSLDRecord{
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

func TestMsgDeleteSLDRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSLDRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSLDRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSLDRecord{
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
