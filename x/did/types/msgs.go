package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
)

// NewMsgUpdateParams creates new instance of MsgUpdateParams
func NewMsgUpdateParams(
	sender sdk.Address,
	someValue bool,
) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: sender.String(),
		Params:    DefaultParams(),
	}
}

// Route returns the name of the module
func (msg MsgUpdateParams) Route() string { return ModuleName }

// Type returns the the action
func (msg MsgUpdateParams) Type() string { return "update_params" }

// GetSignBytes implements the LegacyMsg interface.
func (msg MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgUpdateParams) Validate() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	return msg.Params.Validate()
}

// NewMsgInitializeController creates a new instance of MsgInitializeController
func NewMsgInitializeController(
	sender sdk.Address,
	assertions AssertionList,
	keyshares KeyshareList,
	verifications VerificationList,
) (*MsgInitializeController, error) {

	// Convert assertions to byte arrays
	assertionsBz, err := ConvertAssertionListToByteArray(assertions)
	if err != nil {
		return nil, err
	}

	// Convert keyshares to byte arrays
	keysharesBz, err := ConvertKeyshareListToByteArray(keyshares)
	if err != nil {
		return nil, err
	}

	// Convert verifications to byte arrays
	verificationsBz, err := ConvertVerificationListToByteArray(verifications)
	if err != nil {
		return nil, err
	}

	return &MsgInitializeController{
		Authority:     sender.String(),
		Assertions:    assertionsBz,
		Keyshares:     keysharesBz,
		Verifications: verificationsBz,
	}, nil
}

// Route returns the name of the module
func (msg MsgInitializeController) Route() string { return ModuleName }

// Type returns the the action
func (msg MsgInitializeController) Type() string { return "initialize_controller" }

// GetSignBytes implements the LegacyMsg interface.
func (msg MsgInitializeController) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgInitializeController) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// GetAssertions returns the assertions
func (msg *MsgInitializeController) GetAssertionList() (AssertionList, error) {
	return ConvertByteArrayToAssertionList(msg.Assertions)
}

// GetKeyshares returns the keyshares
func (msg *MsgInitializeController) GetKeyshareList() (KeyshareList, error) {
	return ConvertByteArrayToKeyshareList(msg.Keyshares)
}

// GetVerifications returns the verifications
func (msg *MsgInitializeController) GetVerificationList() (VerificationList, error) {
	return ConvertByteArrayToVerificationList(msg.Verifications)
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgInitializeController) Validate() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}
	_, err := ConvertByteArrayToAssertionList(msg.Assertions)
	if err != nil {
		return errors.Wrap(err, "invalid assertions")
	}

	_, err = ConvertByteArrayToKeyshareList(msg.Keyshares)
	if err != nil {
		return errors.Wrap(err, "invalid keyshares")
	}

	_, err = ConvertByteArrayToVerificationList(msg.Verifications)
	if err != nil {
		return errors.Wrap(err, "invalid verifications")
	}
	return nil
}
