package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rt "go.buf.build/grpc/go/sonr-io/blockchain/registry"
)

const (
	TypeMsgCreateWhoIs = "create_who_is"
	TypeMsgUpdateWhoIs = "update_who_is"
	TypeMsgDeleteWhoIs = "delete_who_is"
)

var _ sdk.Msg = &MsgCreateWhoIs{}

func NewMsgCreateWhoIs(
	creator string,
	did string,
	doc []byte,
	c []*Credential,
	n string,

) *MsgCreateWhoIs {
	return &MsgCreateWhoIs{
		Creator:     creator,
		Did:         did,
		Document:    doc,
		Credentials: c,
		Name:        n,
	}
}

func NewMsgCreateWhoIsFromBuf(msg *rt.MsgCreateWhoIs) *MsgCreateWhoIs {
	return NewMsgCreateWhoIs(msg.GetCreator(), msg.GetDid(), msg.GetDocument(), NewCredentialListFromBuf(msg.GetCredentials()), msg.GetName())
}

func (msg *MsgCreateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhoIs) Type() string {
	return TypeMsgCreateWhoIs
}

func (msg *MsgCreateWhoIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhoIs{}

func NewMsgUpdateWhoIs(
	creator string,
	did string,
	doc []byte,
	c []*Credential,
) *MsgUpdateWhoIs {
	return &MsgUpdateWhoIs{
		Creator:     creator,
		Did:         did,
		Document:    doc,
		Credentials: c,
	}
}

func (msg *MsgUpdateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhoIs) Type() string {
	return TypeMsgUpdateWhoIs
}

func (msg *MsgUpdateWhoIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWhoIs{}

func NewMsgDeleteWhoIs(
	creator string,
	did string,

) *MsgDeleteWhoIs {
	return &MsgDeleteWhoIs{
		Creator: creator,
		Did:     did,
	}
}
func (msg *MsgDeleteWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWhoIs) Type() string {
	return TypeMsgDeleteWhoIs
}

func (msg *MsgDeleteWhoIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWhoIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhoIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
