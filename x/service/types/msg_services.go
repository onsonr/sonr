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

func NewMsgCreateServiceRecord(creator string) *MsgCreateServiceRecord {
	return &MsgCreateServiceRecord{
		Creator: creator,
	}
}

func (msg *MsgCreateServiceRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateServiceRecord) Type() string {
	return TypeMsgCreateServiceRecord
}

func (msg *MsgCreateServiceRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateServiceRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateServiceRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateServiceRecord{}

func NewMsgUpdateServiceRecord(creator string, id uint64) *MsgUpdateServiceRecord {
	return &MsgUpdateServiceRecord{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgUpdateServiceRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateServiceRecord) Type() string {
	return TypeMsgUpdateServiceRecord
}

func (msg *MsgUpdateServiceRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateServiceRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateServiceRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteServiceRecord{}

func NewMsgDeleteServiceRecord(creator string, id uint64) *MsgDeleteServiceRecord {
	return &MsgDeleteServiceRecord{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteServiceRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteServiceRecord) Type() string {
	return TypeMsgDeleteServiceRecord
}

func (msg *MsgDeleteServiceRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteServiceRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteServiceRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
