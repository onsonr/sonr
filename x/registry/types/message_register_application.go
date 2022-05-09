package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterApplication = "register_application"

var _ sdk.Msg = &MsgRegisterApplication{}

func NewMsgRegisterApplication(creator string, serviceName string, c *Credential) *MsgRegisterApplication {
	return &MsgRegisterApplication{
		Creator:     creator,
		ApplicationName: serviceName,
		Credential: c,
	}
}

func (msg *MsgRegisterApplication) Route() string {
	return RouterKey
}

func (msg *MsgRegisterApplication) Type() string {
	return TypeMsgRegisterApplication
}

func (msg *MsgRegisterApplication) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterApplication) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterApplication) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
