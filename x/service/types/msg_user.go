package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterUserEntity = "register_user_entity"

var _ sdk.Msg = &MsgRegisterUserEntity{}

func NewMsgRegisterUserEntity(creator string) *MsgRegisterUserEntity {
	return &MsgRegisterUserEntity{
		Creator: creator,
	}
}

func (msg *MsgRegisterUserEntity) Route() string {
	return RouterKey
}

func (msg *MsgRegisterUserEntity) Type() string {
	return TypeMsgRegisterUserEntity
}

func (msg *MsgRegisterUserEntity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterUserEntity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterUserEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgAuthenticateUserEntity = "authenticate_user_entity"

var _ sdk.Msg = &MsgAuthenticateUserEntity{}

func NewMsgAuthenticateUserEntity(creator string) *MsgAuthenticateUserEntity {
	return &MsgAuthenticateUserEntity{
		Creator: creator,
	}
}

func (msg *MsgAuthenticateUserEntity) Route() string {
	return RouterKey
}

func (msg *MsgAuthenticateUserEntity) Type() string {
	return TypeMsgAuthenticateUserEntity
}

func (msg *MsgAuthenticateUserEntity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAuthenticateUserEntity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAuthenticateUserEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
