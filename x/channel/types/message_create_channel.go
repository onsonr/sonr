package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ot "github.com/sonr-io/sonr/x/object/types"
)

const TypeMsgCreateChannel = "create_channel"

var _ sdk.Msg = &MsgCreateChannel{}

func NewMsgCreateChannel(creator string, name string, description string, object *ot.ObjectDoc, ttl int64, maxSize int64) *MsgCreateChannel {
	return &MsgCreateChannel{
		Creator:          creator,
		Label:            name,
		Description:      description,
		ObjectToRegister: object,
	}
}



func (msg *MsgCreateChannel) Route() string {
	return RouterKey
}

func (msg *MsgCreateChannel) Type() string {
	return TypeMsgCreateChannel
}

func (msg *MsgCreateChannel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChannel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
