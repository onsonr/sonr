package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgImportPublicKey = "import_public_key"

var _ sdk.Msg = &MsgImportPublicKey{}

func NewMsgImportPublicKey(creator string) *MsgImportPublicKey {
	return &MsgImportPublicKey{
		Creator: creator,
	}
}

func (msg *MsgImportPublicKey) Route() string {
	return RouterKey
}

func (msg *MsgImportPublicKey) Type() string {
	return TypeMsgImportPublicKey
}

func (msg *MsgImportPublicKey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgImportPublicKey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgImportPublicKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
