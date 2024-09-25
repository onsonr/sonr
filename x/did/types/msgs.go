package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgUpdateParams{}

//
// [UpdateParams]
//

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

//
// [RegisterService]
//

// NewMsgRegisterController creates a new instance of MsgRegisterController
func NewMsgRegisterService(
	sender sdk.Address,
) (*MsgRegisterService, error) {
	return &MsgRegisterService{
		Controller: sender.String(),
	}, nil
}

// Route returns the name of the module
func (msg MsgRegisterService) Route() string { return ModuleName }

// Type returns the the action
func (msg MsgRegisterService) Type() string { return "register_service" }

// GetSignBytes implements the LegacyMsg interface.
func (msg MsgRegisterService) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgRegisterService) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Controller)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgRegisterService) Validate() error {
	return nil
}

//
// [RegisterController]
//

// NewMsgRegisterController creates a new instance of MsgRegisterController
func NewMsgRegisterController(
	sender sdk.Address,
) (*MsgRegisterController, error) {
	return &MsgRegisterController{
		Authority: sender.String(),
	}, nil
}

// Route returns the name of the module
func (msg MsgRegisterController) Route() string { return ModuleName }

// Type returns the the action
func (msg MsgRegisterController) Type() string { return "register_controller" }

// GetSignBytes implements the LegacyMsg interface.
func (msg MsgRegisterController) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgRegisterController) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgRegisterController) Validate() error {
	return nil
}
