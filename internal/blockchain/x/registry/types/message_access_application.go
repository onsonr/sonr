package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAccessApplication = "access_application"

var _ sdk.Msg = &MsgAccessApplication{}

func NewMsgAccessApplication(creator string, appName string) *MsgAccessApplication {
	return &MsgAccessApplication{
		Creator: creator,
		AppName: appName,
	}
}

func (msg *MsgAccessApplication) Route() string {
	return RouterKey
}

func (msg *MsgAccessApplication) Type() string {
	return TypeMsgAccessApplication
}

func (msg *MsgAccessApplication) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAccessApplication) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAccessApplication) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
