package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEscrowAccount = "create_escrow_account"
	TypeMsgUpdateEscrowAccount = "update_escrow_account"
	TypeMsgDeleteEscrowAccount = "delete_escrow_account"
)

var _ sdk.Msg = &MsgCreateEscrowAccount{}

func NewMsgCreateEscrowAccount(creator string, address string, publicKey string, lockupUsdBalance string) *MsgCreateEscrowAccount {
	return &MsgCreateEscrowAccount{
		Creator:          creator,
		Address:          address,
		PublicKey:        publicKey,
		LockupUsdBalance: lockupUsdBalance,
	}
}

func (msg *MsgCreateEscrowAccount) Route() string {
	return RouterKey
}

func (msg *MsgCreateEscrowAccount) Type() string {
	return TypeMsgCreateEscrowAccount
}

func (msg *MsgCreateEscrowAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEscrowAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEscrowAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateEscrowAccount{}

func NewMsgUpdateEscrowAccount(creator string, id uint64, address string, publicKey string, lockupUsdBalance string) *MsgUpdateEscrowAccount {
	return &MsgUpdateEscrowAccount{
		Id:               id,
		Creator:          creator,
		Address:          address,
		PublicKey:        publicKey,
		LockupUsdBalance: lockupUsdBalance,
	}
}

func (msg *MsgUpdateEscrowAccount) Route() string {
	return RouterKey
}

func (msg *MsgUpdateEscrowAccount) Type() string {
	return TypeMsgUpdateEscrowAccount
}

func (msg *MsgUpdateEscrowAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateEscrowAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateEscrowAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEscrowAccount{}

func NewMsgDeleteEscrowAccount(creator string, address string) *MsgDeleteEscrowAccount {
	return &MsgDeleteEscrowAccount{
		Address:      address,
		Creator: creator,
	}
}
func (msg *MsgDeleteEscrowAccount) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEscrowAccount) Type() string {
	return TypeMsgDeleteEscrowAccount
}

func (msg *MsgDeleteEscrowAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEscrowAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEscrowAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
