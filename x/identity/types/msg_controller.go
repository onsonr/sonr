package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateControllerAccount = "create_controller_account"
	TypeMsgUpdateControllerAccount = "update_controller_account"
	TypeMsgDeleteControllerAccount = "delete_controller_account"
)

var _ sdk.Msg = &MsgCreateControllerAccount{}

func NewMsgCreateControllerAccount(address string, publicKey string, auths ...string) *MsgCreateControllerAccount {
	return &MsgCreateControllerAccount{
		Address:        address,
		PublicKey:      publicKey,
		Authenticators: auths,
	}
}

func (msg *MsgCreateControllerAccount) Route() string {
	return RouterKey
}

func (msg *MsgCreateControllerAccount) Type() string {
	return TypeMsgCreateControllerAccount
}

func (msg *MsgCreateControllerAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateControllerAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateControllerAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateControllerAccount{}

func NewMsgUpdateControllerAccount(creator string, id uint64, address string, publicKey string) *MsgUpdateControllerAccount {
	return &MsgUpdateControllerAccount{
		Id:      id,
		Address: creator,
	}
}

func (msg *MsgUpdateControllerAccount) Route() string {
	return RouterKey
}

func (msg *MsgUpdateControllerAccount) Type() string {
	return TypeMsgUpdateControllerAccount
}

func (msg *MsgUpdateControllerAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateControllerAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateControllerAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteControllerAccount{}

func NewMsgDeleteControllerAccount(creator string, address string) *MsgDeleteControllerAccount {
	return &MsgDeleteControllerAccount{
		Address: address,
		Creator: creator,
	}
}
func (msg *MsgDeleteControllerAccount) Route() string {
	return RouterKey
}

func (msg *MsgDeleteControllerAccount) Type() string {
	return TypeMsgDeleteControllerAccount
}

func (msg *MsgDeleteControllerAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteControllerAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteControllerAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
