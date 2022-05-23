package types

import (
	"fmt"
	"strings"

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
		Creator:     owner,
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
	owner, err := sdk.AccAddressFromBech32(msg.Creator)
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// GetCreatorDid returns the creator did
func (msg *MsgCreateWhoIs) GetCreatorDid() string {
	rawCreator := msg.GetCreator()

	// Trim snr account prefix
	if strings.HasPrefix(rawCreator, "snr") {
		rawCreator = strings.TrimLeft(rawCreator, "snr")
	}

	// Trim cosmos account prefix
	if strings.HasPrefix(rawCreator, "cosmos") {
		rawCreator = strings.TrimLeft(rawCreator, "cosmos")
	}
	return fmt.Sprintf("did:snr:%s", rawCreator)
}

var _ sdk.Msg = &MsgUpdateWhoIs{}

func NewMsgUpdateWhoIs(owner string, id string, doc []byte) *MsgUpdateWhoIs {
	return &MsgUpdateWhoIs{
		Creator:     owner,
		DidDocument: doc,
	}
}

func (msg *MsgUpdateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhoIs) Type() string {
	return TypeMsgUpdateWhoIs
}

func (msg *MsgUpdateWhoIs) GetCreatorDid() string {
	rawCreator := msg.GetCreator()
	var identifier string

	// Trim snr account prefix
	if strings.HasPrefix(rawCreator, "snr") {
		identifier = strings.TrimLeft(rawCreator, "snr")
	}

	// Trim cosmos account prefix
	if strings.HasPrefix(identifier, "cosmos") {
		identifier = strings.TrimLeft(identifier, "cosmos")
	}
	return fmt.Sprintf("did:snr:%s", identifier)
}

func (msg *MsgUpdateWhoIs) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Creator)
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeactivateWhoIs{}

func NewMsgDeactivateWhoIs(owner string, id string) *MsgDeactivateWhoIs {
	return &MsgDeactivateWhoIs{
		Creator: owner,
	}
}
func (msg *MsgDeactivateWhoIs) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateWhoIs) Type() string {
	return TypeMsgDeactivateWhoIs
}

func (msg *MsgDeactivateWhoIs) GetCreatorDid() string {
	rawCreator := msg.GetCreator()
	var identifier string

	// Trim snr account prefix
	if strings.HasPrefix(rawCreator, "snr") {
		identifier = strings.TrimLeft(rawCreator, "snr")
	}

	// Trim cosmos account prefix
	if strings.HasPrefix(identifier, "cosmos") {
		identifier = strings.TrimLeft(identifier, "cosmos")
	}
	return fmt.Sprintf("did:snr:%s", identifier)
}

func (msg *MsgDeactivateWhoIs) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Creator)
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
