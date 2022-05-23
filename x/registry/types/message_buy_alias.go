package types

import (
	"fmt"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyAlias = "buy_alias"

var _ sdk.Msg = &MsgBuyAlias{}

func NewMsgBuyAlias(creator string, did string, name string) *MsgBuyAlias {
	return &MsgBuyAlias{
		Creator: creator,
		Did:     did,
		Name:    name,
	}
}

func (msg *MsgBuyAlias) Route() string {
	return RouterKey
}

func (msg *MsgBuyAlias) Type() string {
	return TypeMsgBuyAlias
}

// GetCreatorDid returns the creator did
func (msg *MsgBuyAlias) GetCreatorDid() string {
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

func (msg *MsgBuyAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgBuyAlias) ValidateAlias() error {
	name := msg.GetName()
	// Check if Alias length is valid
	if len(name) < 3 {
		return sdkerrors.Wrap(ErrAliasUnavailable, "Alias must be at least 3 characters long")
	}
	// Check if alias is only alpha-numeric
	regexp := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !regexp.MatchString(name) {
		return sdkerrors.Wrap(ErrAliasUnavailable, "Alias must be alphanumeric")
	}
	return nil
}
