package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyNameAlias = "buy_name_alias"

var _ sdk.Msg = &MsgBuyNameAlias{}

func NewMsgBuyNameAlias(creator string, did string, amount int32, name string) *MsgBuyNameAlias {
	return &MsgBuyNameAlias{
		Creator: creator,
		Did:     did,
		Amount:  amount,
		Name:    name,
	}
}

func (msg *MsgBuyNameAlias) Route() string {
	return RouterKey
}

func (msg *MsgBuyNameAlias) Type() string {
	return TypeMsgBuyNameAlias
}

func (msg *MsgBuyNameAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyNameAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyNameAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
