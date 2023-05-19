package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDidDocument = "create_did_document"
	TypeMsgUpdateDidDocument = "update_did_document"
)

var _ sdk.Msg = &MsgCreateDidDocument{}

func NewMsgCreateDidDocument(creator string, wallet_id uint32, alias string, didDoc *Identity, blockDocs ...*Identity) *MsgCreateDidDocument {
	return &MsgCreateDidDocument{
		Alias:       alias,
		Creator:     creator,
		Primary:     didDoc,
		Blockchains: blockDocs,
		WalletId:    wallet_id,
	}
}

func (msg *MsgCreateDidDocument) Route() string {
	return RouterKey
}

func (msg *MsgCreateDidDocument) Type() string {
	return TypeMsgCreateDidDocument
}

func (msg *MsgCreateDidDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDidDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDidDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDidDocument{}

func NewMsgUpdateDidDocument(
	creator string,
	primary *Identity,
	blockDocs ...*Identity,
) *MsgUpdateDidDocument {
	return &MsgUpdateDidDocument{
		Creator:     creator,
		Primary:     primary,
		Blockchains: blockDocs,
	}
}

func (msg *MsgUpdateDidDocument) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDidDocument) Type() string {
	return TypeMsgUpdateDidDocument
}

func (msg *MsgUpdateDidDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDidDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDidDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
