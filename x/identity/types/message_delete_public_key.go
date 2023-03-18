package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeletePublicKey = "delete_public_key"

var _ sdk.Msg = &MsgDeletePublicKey{}

func NewMsgDeletePublicKey(creator string) *MsgDeletePublicKey {
	return &MsgDeletePublicKey{
		Creator: creator,
	}
}

func (msg *MsgDeletePublicKey) Route() string {
	return RouterKey
}

func (msg *MsgDeletePublicKey) Type() string {
	return TypeMsgDeletePublicKey
}

func (msg *MsgDeletePublicKey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePublicKey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePublicKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
