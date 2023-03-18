package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/core/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgImportPublicKey_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgImportPublicKey
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgImportPublicKey{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgImportPublicKey{
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
