package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnBucket = "burn_bucket"

var _ sdk.Msg = &MsgBurnBucket{}

func NewMsgBurnBucket(creator string, bucketId string) *MsgBurnBucket {
	return &MsgBurnBucket{
		Creator: creator,
		Bucket: &BucketConfig{
			Uuid: bucketId,
		},
	}
}

func (msg *MsgBurnBucket) Route() string {
	return RouterKey
}

func (msg *MsgBurnBucket) Type() string {
	return TypeMsgBurnBucket
}

func (msg *MsgBurnBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
