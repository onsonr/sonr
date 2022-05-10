package types

import (
	"fmt"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyNameAlias = "buy_name_alias"

var _ sdk.Msg = &MsgBuyNameAlias{}

func NewMsgBuyNameAlias(creator string, did string, name string) *MsgBuyNameAlias {
	return &MsgBuyNameAlias{
		Creator: creator,
		Did:     did,
		Name:    name,
	}
}

func (msg *MsgBuyNameAlias) Route() string {
	return RouterKey
}

func (msg *MsgBuyNameAlias) Type() string {
	return TypeMsgBuyNameAlias
}

func (msg *MsgBuyNameAlias) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyNameAlias) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyNameAlias) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func ValidateAlias(name string) error {
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

func FormatNameAlias(alias string) string {
	if strings.Contains(alias, ".snr") {
		return alias
	}
	return fmt.Sprintf("%s.snr", alias)
}

func FormatAppAlias(alias string) string {
	if strings.Contains(alias, "/") {
		return alias
	}
	return fmt.Sprintf("/%s", alias)
}
