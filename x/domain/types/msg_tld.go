package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTLDRecord = "create_tld_record"
	TypeMsgUpdateTLDRecord = "update_tld_record"
	TypeMsgDeleteTLDRecord = "delete_tld_record"
)

var _ sdk.Msg = &MsgCreateTLDRecord{}

func NewMsgCreateTLDRecord(
	creator string,
	tldRecord *TLDRecord,
) *MsgCreateTLDRecord {
	return &MsgCreateTLDRecord{
		Creator:   creator,
		TldRecord: tldRecord,
	}
}

func (msg *MsgCreateTLDRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateTLDRecord) Type() string {
	return TypeMsgCreateTLDRecord
}

func (msg *MsgCreateTLDRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTLDRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTLDRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTLDRecord{}

func NewMsgUpdateTLDRecord(
	creator string,
	tldRecord *TLDRecord,
) *MsgUpdateTLDRecord {
	return &MsgUpdateTLDRecord{
		Creator:   creator,
		TldRecord: tldRecord,
	}
}

func (msg *MsgUpdateTLDRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTLDRecord) Type() string {
	return TypeMsgUpdateTLDRecord
}

func (msg *MsgUpdateTLDRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTLDRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTLDRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTLDRecord{}

func NewMsgDeleteTLDRecord(
	creator string,
	index string,

) *MsgDeleteTLDRecord {
	return &MsgDeleteTLDRecord{
		Creator: creator,
		Name:    index,
	}
}
func (msg *MsgDeleteTLDRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTLDRecord) Type() string {
	return TypeMsgDeleteTLDRecord
}

func (msg *MsgDeleteTLDRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTLDRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTLDRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
