package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateBucket = "create_bucket"

var _ sdk.Msg = &MsgCreateBucket{}

func NewMsgCreateBucket(creator string, label string, description string, kind string, ibDids []string) *MsgCreateBucket {
	return &MsgCreateBucket{
		Creator:           creator,
		Label:             label,
		Description:       description,
		Kind:              kind,
		InitialObjectDids: ibDids,
	}
}


func (msg *MsgCreateBucket) Route() string {
	return RouterKey
}

func (msg *MsgCreateBucket) Type() string {
	return TypeMsgCreateBucket
}

func (msg *MsgCreateBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
