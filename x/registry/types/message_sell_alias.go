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
		Did:     did,
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
