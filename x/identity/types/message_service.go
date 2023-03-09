package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRegisterService   = "register_service"
	TypeMsgUpdateService     = "update_service"
	TypeMsgDeactivateService = "deactivate_service"
)

var _ sdk.Msg = &MsgRegisterService{}

func NewMsgRegisterService(creator string) *MsgRegisterService {
	return &MsgRegisterService{
		Creator: creator,
	}
}

func (msg *MsgRegisterService) Route() string {
	return RouterKey
}

func (msg *MsgRegisterService) Type() string {
	return TypeMsgRegisterService
}

func (msg *MsgRegisterService) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterService) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// NewMsgUpdateService creates a new MsgUpdateService instance
func NewMsgUpdateService(creator string, id string, did string, serviceEndpoint string, description string, properties string, tags string) *MsgUpdateService {
	return &MsgUpdateService{}
}

// Route returns the message route for a MsgUpdateService.
func (msg *MsgUpdateService) Route() string {
	return RouterKey
}

// Type returns the message type for a MsgUpdateService.
func (msg *MsgUpdateService) Type() string {
	return TypeMsgUpdateService
}

// GetSigners returns the addresses of the signers who must sign.
func (msg *MsgUpdateService) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on.
func (msg *MsgUpdateService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg *MsgUpdateService) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// NewMsgDeactivateService creates a new MsgDeactivateService instance
func NewMsgDeactivateService(creator string, id string) *MsgDeactivateService {
	return &MsgDeactivateService{}
}

// Route returns the message route for a MsgDeactivateService.
func (msg *MsgDeactivateService) Route() string {
	return RouterKey
}

// Type returns the message type for a MsgDeactivateService.
func (msg *MsgDeactivateService) Type() string {
	return TypeMsgDeactivateService
}

// GetSigners returns the addresses of the signers who must sign.
func (msg *MsgDeactivateService) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on.
func (msg *MsgDeactivateService) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg *MsgDeactivateService) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
