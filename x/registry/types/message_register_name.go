package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	"github.com/duo-labs/webauthn/webauthn"
)

const TypeMsgRegisterName = "register_name"

var _ sdk.Msg = &MsgRegisterName{}

// NewMsgRegisterName creates a new MsgRegisterName object
func NewMsgRegisterName(creator string, nameToRegister string, cred webauthn.Credential) *MsgRegisterName {
	return &MsgRegisterName{
		Creator:        creator,
		Credential:     ConvertToSonrCredential(cred),
		NameToRegister: CleanNameForSuffix(nameToRegister),
	}
}

// CleanName checks if the username is available
func CleanNameForSuffix(name string) string {
	if strings.Contains(name, ".snr") {
		return name
	}
	return name + ".snr"
}

// Route returns the message type used for routing the message.
func (msg *MsgRegisterName) Route() string {
	return RouterKey
}

// Type returns the action type
func (msg *MsgRegisterName) Type() string {
	return TypeMsgRegisterName
}

// GetSigners returns the creator of the message
func (msg *MsgRegisterName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes encodes the message for signing
func (msg *MsgRegisterName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// Validate verifies the message details
func (msg *MsgRegisterName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
