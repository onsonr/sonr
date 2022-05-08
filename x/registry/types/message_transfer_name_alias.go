package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferNameAlias = "transfer_name_alias"

var _ sdk.Msg = &MsgTransferNameAlias{}

func NewMsgTransferNameAlias(creator string, did string, alias string, recipient string) *MsgTransferNameAlias {
	return &MsgTransferNameAlias{
		Creator:   creator,
		Did:       did,
		Alias:     alias,
		Recipient: recipient,
	}
}

func (msg *MsgTransferNameAlias) Route() string {
	return RouterKey
}

func (msg *MsgTransferNameAlias) Type() string {
	return TypeMsgTransferNameAlias
}

func (msg *MsgTransferNameAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferNameAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferNameAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
