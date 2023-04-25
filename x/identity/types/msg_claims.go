package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateClaimableWallet = "create_claimable_wallet"
	TypeMsgUpdateClaimableWallet = "update_claimable_wallet"
	TypeMsgDeleteClaimableWallet = "delete_claimable_wallet"
)

var _ sdk.Msg = &MsgCreateClaimableWallet{}

func NewMsgCreateClaimableWallet(creator string, claimableWallet *ClaimableWallet) *MsgCreateClaimableWallet {
	return &MsgCreateClaimableWallet{
		Creator: creator,
		ClaimableWallet: claimableWallet,
	}
}

func (msg *MsgCreateClaimableWallet) Route() string {
	return RouterKey
}

func (msg *MsgCreateClaimableWallet) Type() string {
	return TypeMsgCreateClaimableWallet
}

func (msg *MsgCreateClaimableWallet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClaimableWallet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClaimableWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateClaimableWallet{}

func NewMsgUpdateClaimableWallet(creator string, id uint64) *MsgUpdateClaimableWallet {
	return &MsgUpdateClaimableWallet{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgUpdateClaimableWallet) Route() string {
	return RouterKey
}

func (msg *MsgUpdateClaimableWallet) Type() string {
	return TypeMsgUpdateClaimableWallet
}

func (msg *MsgUpdateClaimableWallet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateClaimableWallet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateClaimableWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteClaimableWallet{}

func NewMsgDeleteClaimableWallet(creator string, id uint64) *MsgDeleteClaimableWallet {
	return &MsgDeleteClaimableWallet{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteClaimableWallet) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClaimableWallet) Type() string {
	return TypeMsgDeleteClaimableWallet
}

func (msg *MsgDeleteClaimableWallet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClaimableWallet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClaimableWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
