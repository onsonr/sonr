package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRegisterServiceRecord = "register_service_record"
	TypeMsgUpdateServiceRecord   = "update_service_record"
	TypeMsgBurnServiceRecord     = "burn_service_record"
)

var _ sdk.Msg = &MsgRegisterServiceRecord{}

func NewMsgRegisterServiceRecord(
	Controller string,
	record ServiceRecord,
) *MsgRegisterServiceRecord {
	return &MsgRegisterServiceRecord{
		Controller: Controller,
		Record:     &record,
	}
}

func (msg *MsgRegisterServiceRecord) Route() string {
	return RouterKey
}

func (msg *MsgRegisterServiceRecord) Type() string {
	return TypeMsgRegisterServiceRecord
}

func (msg *MsgRegisterServiceRecord) GetSigners() []sdk.AccAddress {
	Controller, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{Controller}
}

func (msg *MsgRegisterServiceRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterServiceRecord) ValidateBasic() error {
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
		Record:     &record,
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

var _ sdk.Msg = &MsgBurnServiceRecord{}

func NewMsgBurnServiceRecord(
	Controller string,
	index string,

) *MsgBurnServiceRecord {
	return &MsgBurnServiceRecord{
		Controller: Controller,
		Id:         index,
	}
}
func (msg *MsgBurnServiceRecord) Route() string {
	return RouterKey
}

func (msg *MsgBurnServiceRecord) Type() string {
	return TypeMsgBurnServiceRecord
}

func (msg *MsgBurnServiceRecord) GetSigners() []sdk.AccAddress {
	Controller, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{Controller}
}

func (msg *MsgBurnServiceRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnServiceRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Controller)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Controller address (%s)", err)
	}
	return nil
}
