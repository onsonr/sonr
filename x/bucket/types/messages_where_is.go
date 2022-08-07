package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWhereIs = "create_where_is"
	TypeMsgUpdateWhereIs = "update_where_is"
	TypeMsgDeleteWhereIs = "delete_where_is"
)

var _ sdk.Msg = &MsgCreateWhereIs{}

func NewMsgCreateWhereIs(creator string) *MsgCreateWhereIs {
	return &MsgCreateWhereIs{
		Creator: creator,
	}
}

func (msg *MsgCreateWhereIs) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhereIs) Type() string {
	return TypeMsgCreateWhereIs
}

func (msg *MsgCreateWhereIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhereIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhereIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhereIs{}

func NewMsgUpdateWhereIs(creator string, id string) *MsgUpdateWhereIs {
	return &MsgUpdateWhereIs{
		Did:      id,
		Creator: creator,
	}
}

func (msg *MsgUpdateWhereIs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhereIs) Type() string {
	return TypeMsgUpdateWhereIs
}

func (msg *MsgUpdateWhereIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWhereIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhereIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWhereIs{}

func NewMsgDeleteWhereIs(creator string, id string) *MsgDeleteWhereIs {
	return &MsgDeleteWhereIs{
		Did:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteWhereIs) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWhereIs) Type() string {
	return TypeMsgDeleteWhereIs
}

func (msg *MsgDeleteWhereIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWhereIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhereIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
