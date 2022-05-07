package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAccessName = "access_name"

var _ sdk.Msg = &MsgAccessName{}

func NewMsgAccessName(creator string, name string, c *Credential) *MsgAccessName {
	return &MsgAccessName{
		Creator:    creator,
		Name:       name,
		Credential: c,
	}
}

func (msg *MsgAccessName) Route() string {
	return RouterKey
}

func (msg *MsgAccessName) Type() string {
	return TypeMsgAccessName
}

func (msg *MsgAccessName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAccessName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAccessName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
