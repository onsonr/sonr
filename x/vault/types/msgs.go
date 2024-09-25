package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgUpdateParams{}

// NewMsgUpdateParams creates new instance of MsgUpdateParams
func NewMsgUpdateParams(
	sender sdk.Address,
	someValue bool,
) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: sender.String(),
		Params: Params{
			IpfsActive: someValue,
		},
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
// [AllocateVault]
//

// NewMsgAllocateVault creates a new instance of MsgAllocateVault
func NewMsgAllocateVault(
	sender sdk.Address,
) (*MsgAllocateVault, error) {
	return &MsgAllocateVault{}, nil
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *MsgAllocateVault) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// Route returns the name of the module
func (msg MsgAllocateVault) Route() string { return ModuleName }

// Type returns the the action
func (msg MsgAllocateVault) Type() string { return "allocate_vault" }

// GetSignBytes implements the LegacyMsg interface.
func (msg MsgAllocateVault) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// Vaalidate does a sanity check on the provided data.
func (msg *MsgAllocateVault) Validate() error {
	return nil
}
