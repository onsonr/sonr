package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferAppAlias = "transfer_app_alias"

var _ sdk.Msg = &MsgTransferAppAlias{}

func NewMsgTransferAppAlias(creator string, did string, alias string, recipient string) *MsgTransferAppAlias {
	return &MsgTransferAppAlias{
		Creator:   creator,
		Did:       did,
		Alias:     alias,
		Recipient: recipient,
	}
}

func (msg *MsgTransferAppAlias) Route() string {
	return RouterKey
}

func (msg *MsgTransferAppAlias) Type() string {
	return TypeMsgTransferAppAlias
}

func (msg *MsgTransferAppAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferAppAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferAppAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
