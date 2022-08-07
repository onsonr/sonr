package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeprecateSchema = "deprecate_schema"

var _ sdk.Msg = &MsgDeprecateSchema{}

func NewMsgDeprecateSchema(creator, did string) *MsgDeprecateSchema {
	return &MsgDeprecateSchema{
		Creator: creator,
		Did:     did,
	}
}

func (msg *MsgDeprecateSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeprecateSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
