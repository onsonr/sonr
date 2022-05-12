package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyAppAlias = "buy_app_alias"

var _ sdk.Msg = &MsgBuyAppAlias{}

func NewMsgBuyAppAlias(creator string, did string, name string) *MsgBuyAppAlias {
	return &MsgBuyAppAlias{
		Creator: creator,
		Did:     did,
		Name:    name,
	}
}

func (msg *MsgBuyAppAlias) Route() string {
	return RouterKey
}

func (msg *MsgBuyAppAlias) Type() string {
	return TypeMsgBuyAppAlias
}

func (msg *MsgBuyAppAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyAppAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyAppAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
