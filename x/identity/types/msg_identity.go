package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterIdentity = "register_identity"

var _ sdk.Msg = &MsgRegisterIdentity{}

func NewMsgRegisterIdentity(creator string, doc *DIDDocument) *MsgRegisterIdentity {
	msg := &MsgRegisterIdentity{
		Creator:  creator,
		DidDocument: doc,
	}
	return msg
}

func (msg *MsgRegisterIdentity) Route() string {
	return RouterKey
}

func (msg *MsgRegisterIdentity) Type() string {
	return TypeMsgRegisterIdentity
}

func (msg *MsgRegisterIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
