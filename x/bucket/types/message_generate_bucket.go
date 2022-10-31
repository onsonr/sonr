package types

import (

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

const TypeMsgGenerateBucket = "generate_bucket"

var _ sdk.Msg = &MsgGenerateBucket{}

func NewMsgGenerateBucket(creator string, bucketId string) *MsgGenerateBucket {
	return &MsgGenerateBucket{
		Creator:  creator,
		BucketId: bucketId,
	}
}

func (msg *MsgGenerateBucket) Route() string {
	return RouterKey
}

func (msg *MsgGenerateBucket) Type() string {
	return TypeMsgGenerateBucket
}

func (msg *MsgGenerateBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGenerateBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGenerateBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return msg.ValidateBucketId()
}

func (msg *MsgGenerateBucket) ValidateBucketId() error {
	_, err := uuid.Parse(msg.BucketId)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bucket id (%s)", err)
	}
	return nil
}
