package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWhoIs     = "create_who_is"
	TypeMsgUpdateWhoIs     = "update_who_is"
	TypeMsgDeactivateWhoIs = "delete_who_is"
)

var _ sdk.Msg = &MsgCreateWhoIs{}

func NewMsgCreateWhoIs(owner string, didDoc []byte, t WhoIsType) *MsgCreateWhoIs {
	return &MsgCreateWhoIs{
		Owner:       owner,
		DidDocument: didDoc,
		WhoisType:   t,
	}
}

func (msg *MsgCreateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhoIs) Type() string {
	return TypeMsgCreateWhoIs
}

func (msg *MsgCreateWhoIs) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgCreateWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhoIs{}

func NewMsgUpdateWhoIs(owner string, id string, doc []byte) *MsgUpdateWhoIs {
	return &MsgUpdateWhoIs{
		Did:   id,
		Owner: owner,
		DidDocument: doc,
	}
}

func (msg *MsgUpdateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhoIs) Type() string {
	return TypeMsgUpdateWhoIs
}

func (msg *MsgUpdateWhoIs) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgUpdateWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeactivateWhoIs{}

func NewMsgDeactivateWhoIs(owner string, id string) *MsgDeactivateWhoIs {
	return &MsgDeactivateWhoIs{
		Did:   id,
		Owner: owner,
	}
}
func (msg *MsgDeactivateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateWhoIs) Type() string {
	return TypeMsgDeactivateWhoIs
}

func (msg *MsgDeactivateWhoIs) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgDeactivateWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeactivateWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
