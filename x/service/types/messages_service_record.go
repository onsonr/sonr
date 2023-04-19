package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateServiceRecord = "create_service_record"
	TypeMsgUpdateServiceRecord = "update_service_record"
	TypeMsgDeleteServiceRecord = "delete_service_record"
)

var _ sdk.Msg = &MsgCreateServiceRecord{}

func NewMsgCreateServiceRecord(
	Controller string,
	record ServiceRecord,
) *MsgCreateServiceRecord {
	return &MsgCreateServiceRecord{
		Controller: Controller,
		Record:         &record,
	}
}

func (msg *MsgCreateServiceRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateServiceRecord) Type() string {
	return TypeMsgCreateServiceRecord
}

func (msg *MsgCreateServiceRecord) GetSigners() []sdk.AccAddress {
	Controller, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{Controller}
}

func (msg *MsgCreateServiceRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateServiceRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Controller address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateServiceRecord{}

func NewMsgUpdateServiceRecord(
	Controller string,
record ServiceRecord,

) *MsgUpdateServiceRecord {
	return &MsgUpdateServiceRecord{
		Controller: Controller,
		Record:         &record,
	}
}

func (msg *MsgUpdateServiceRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateServiceRecord) Type() string {
	return TypeMsgUpdateServiceRecord
}

func (msg *MsgUpdateServiceRecord) GetSigners() []sdk.AccAddress {
	Controller, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{Controller}
}

func (msg *MsgUpdateServiceRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateServiceRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Controller address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteServiceRecord{}

func NewMsgDeleteServiceRecord(
	Controller string,
	index string,

) *MsgDeleteServiceRecord {
	return &MsgDeleteServiceRecord{
		Controller: Controller,
		Id:         index,
	}
}
func (msg *MsgDeleteServiceRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteServiceRecord) Type() string {
	return TypeMsgDeleteServiceRecord
}

func (msg *MsgDeleteServiceRecord) GetSigners() []sdk.AccAddress {
	Controller, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{Controller}
}

func (msg *MsgDeleteServiceRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteServiceRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Controller address (%s)", err)
	}
	return nil
}
