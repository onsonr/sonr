package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDidDocument = "create_did_document"
	TypeMsgUpdateDidDocument = "update_did_document"
	TypeMsgDeleteDidDocument = "delete_did_document"
)

var _ sdk.Msg = &MsgCreateDidDocument{}

func NewMsgCreateDidDocument(creator string) *MsgCreateDidDocument {
	return &MsgCreateDidDocument{
		Creator: creator,
		Did:     "",
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
	did string,
	context string,
	controller string,
	verificationMethod string,
	authentication string,
	assertionMethod string,
	capibilityInvocation string,
	capabilityDelegation string,
	keyAgreement string,
	service string,
	alsoKnownAs string,

) *MsgUpdateDidDocument {
	return &MsgUpdateDidDocument{
		Creator:              creator,
		Did:                  did,
		Context:              context,
		Controller:           controller,
		VerificationMethod:   verificationMethod,
		Authentication:       authentication,
		AssertionMethod:      assertionMethod,
		CapibilityInvocation: capibilityInvocation,
		CapabilityDelegation: capabilityDelegation,
		KeyAgreement:         keyAgreement,
		Service:              service,
		AlsoKnownAs:          alsoKnownAs,
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

var _ sdk.Msg = &MsgDeleteDidDocument{}

func NewMsgDeleteDidDocument(
	creator string,
	did string,

) *MsgDeleteDidDocument {
	return &MsgDeleteDidDocument{
		Creator: creator,
		Did:     did,
	}
}
func (msg *MsgDeleteDidDocument) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDidDocument) Type() string {
	return TypeMsgDeleteDidDocument
}

func (msg *MsgDeleteDidDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDidDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDidDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
