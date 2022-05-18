package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferAlias = "transfer_alias"

var _ sdk.Msg = &MsgTransferAlias{}

func NewMsgTransferAlias(creator string, did string, alias string, recipient string) *MsgTransferAlias {
	return &MsgTransferAlias{
		Creator:   creator,
		Did:       did,
		Alias:     alias,
		Recipient: recipient,
	}
}

func (msg *MsgTransferAlias) Route() string {
	return RouterKey
}

func (msg *MsgTransferAlias) Type() string {
	return TypeMsgTransferAlias
}

func (msg *MsgTransferAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
