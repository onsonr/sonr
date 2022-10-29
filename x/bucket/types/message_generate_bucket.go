package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGenerateBucket = "generate_bucket"

var _ sdk.Msg = &MsgAllocateBucket{}

func NewMsgGenerateBucket(creator string) *MsgAllocateBucket {
	return &MsgAllocateBucket{
		Creator: creator,
	}
}

func (msg *MsgAllocateBucket) Route() string {
	return RouterKey
}

func (msg *MsgAllocateBucket) Type() string {
	return TypeMsgGenerateBucket
}

func (msg *MsgAllocateBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAllocateBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAllocateBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
