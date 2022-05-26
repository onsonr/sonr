package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferAlias = "transfer_alias"

var _ sdk.Msg = &MsgTransferAlias{}

func NewMsgTransferAlias(creator string, did string, alias string, recipient string) *MsgTransferAlias {
	return &MsgTransferAlias{
		Creator:   creator,
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

// GetCreatorDid returns the creator did
func (msg *MsgTransferAlias) GetCreatorDid() string {
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
