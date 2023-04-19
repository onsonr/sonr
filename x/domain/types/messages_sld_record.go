package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSLDRecord = "create_sld_record"
	TypeMsgUpdateSLDRecord = "update_sld_record"
	TypeMsgDeleteSLDRecord = "delete_sld_record"
)

var _ sdk.Msg = &MsgCreateSLDRecord{}

func NewMsgCreateSLDRecord(
	creator string,
	sldRecord *SLDRecord,
) *MsgCreateSLDRecord {
	return &MsgCreateSLDRecord{
		Creator: creator,
		SldRecord: sldRecord,
	}
}

func (msg *MsgCreateSLDRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateSLDRecord) Type() string {
	return TypeMsgCreateSLDRecord
}

func (msg *MsgCreateSLDRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSLDRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSLDRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSLDRecord{}

func NewMsgUpdateSLDRecord(
	creator string,
	sldRecord *SLDRecord,

) *MsgUpdateSLDRecord {
	return &MsgUpdateSLDRecord{
		Creator: creator,
		SldRecord:  sldRecord,
	}
}

func (msg *MsgUpdateSLDRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSLDRecord) Type() string {
	return TypeMsgUpdateSLDRecord
}

func (msg *MsgUpdateSLDRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSLDRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSLDRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSLDRecord{}

func NewMsgDeleteSLDRecord(
	creator string,
	index string,

) *MsgDeleteSLDRecord {
	return &MsgDeleteSLDRecord{
		Creator: creator,
		Name:   index,
	}
}
func (msg *MsgDeleteSLDRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSLDRecord) Type() string {
	return TypeMsgDeleteSLDRecord
}

func (msg *MsgDeleteSLDRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSLDRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSLDRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
