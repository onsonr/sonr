package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDomainRecord = "create_domain_registry"
	TypeMsgUpdateDomainRecord = "update_domain_registry"
	TypeMsgDeleteDomainRecord = "delete_domain_registry"
)

var _ sdk.Msg = &MsgCreateDomainRecord{}

func NewMsgCreateDomainRecord(
	creator string,
	index string,

) *MsgCreateDomainRecord {
	return &MsgCreateDomainRecord{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateDomainRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateDomainRecord) Type() string {
	return TypeMsgCreateDomainRecord
}

func (msg *MsgCreateDomainRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDomainRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDomainRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDomainRecord{}

func NewMsgUpdateDomainRecord(
	creator string,
	index string,

) *MsgUpdateDomainRecord {
	return &MsgUpdateDomainRecord{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateDomainRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDomainRecord) Type() string {
	return TypeMsgUpdateDomainRecord
}

func (msg *MsgUpdateDomainRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDomainRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDomainRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDomainRecord{}

func NewMsgDeleteDomainRecord(
	creator string,
	index string,

) *MsgDeleteDomainRecord {
	return &MsgDeleteDomainRecord{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteDomainRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDomainRecord) Type() string {
	return TypeMsgDeleteDomainRecord
}

func (msg *MsgDeleteDomainRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDomainRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDomainRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
