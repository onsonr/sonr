package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBucket = "update_bucket"

var _ sdk.Msg = &MsgUpdateBucket{}

func NewMsgUpdateBucket(creator string, label string, description string, addObjs []string, removObjs []string) *MsgUpdateBucket {
	return &MsgUpdateBucket{
		Creator:           creator,
		Label:             label,
		Description:       description,
		AddedObjectDids:   addObjs,
		RemovedObjectDids: removObjs,
	}
}

func (msg *MsgUpdateBucket) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBucket) Type() string {
	return TypeMsgUpdateBucket
}

func (msg *MsgUpdateBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
