package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferAppAlias = "sell_alias"

var _ sdk.Msg = &MsgSellAlias{}

func NewMsgSellAlias(creator string, did string, alias string, amount int32) *MsgSellAlias {
	return &MsgSellAlias{
		Creator: creator,
		Alias:   alias,
		Amount:  amount,
	}
}

func (msg *MsgSellAlias) Route() string {
	return RouterKey
}

func (msg *MsgSellAlias) Type() string {
	return TypeMsgTransferAppAlias
}

// GetCreatorDid returns the creator did
func (msg *MsgSellAlias) GetCreatorDid() string {
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

func (msg *MsgSellAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSellAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSellAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
