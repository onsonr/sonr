package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateUsernameRecords = "create_username_records"
	TypeMsgUpdateUsernameRecords = "update_username_records"
	TypeMsgDeleteUsernameRecords = "delete_username_records"
)

var _ sdk.Msg = &MsgCreateUsernameRecords{}

func NewMsgCreateEmailUsernameRecord(creator string, email string) *MsgCreateUsernameRecords {
	return &MsgCreateUsernameRecords{
		Creator: creator,
		Index:  EmailIndex(email),
		Method: EmailMethod,
	}
}

func NewMsgCreateUsernameRecords(
	creator string,
	index string,
	method string,
) *MsgCreateUsernameRecords {
	return &MsgCreateUsernameRecords{
		Creator: creator,
		Index:   index,
		Method:  method,
	}
}

func (msg *MsgCreateUsernameRecords) Route() string {
	return RouterKey
}

func (msg *MsgCreateUsernameRecords) Type() string {
	return TypeMsgCreateUsernameRecords
}

func (msg *MsgCreateUsernameRecords) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateUsernameRecords) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUsernameRecords) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUsernameRecords{}

func NewMsgUpdateUsernameRecords(
	creator string,
	index string,

) *MsgUpdateUsernameRecords {
	return &MsgUpdateUsernameRecords{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateUsernameRecords) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUsernameRecords) Type() string {
	return TypeMsgUpdateUsernameRecords
}

func (msg *MsgUpdateUsernameRecords) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUsernameRecords) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUsernameRecords) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUsernameRecords{}

func NewMsgDeleteUsernameRecords(
	creator string,
	index string,

) *MsgDeleteUsernameRecords {
	return &MsgDeleteUsernameRecords{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteUsernameRecords) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUsernameRecords) Type() string {
	return TypeMsgDeleteUsernameRecords
}

func (msg *MsgDeleteUsernameRecords) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUsernameRecords) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUsernameRecords) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
