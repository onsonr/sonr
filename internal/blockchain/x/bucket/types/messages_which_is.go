package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

const (
	TypeMsgCreateWhichIs = "create_which_is"
	TypeMsgUpdateWhichIs = "update_which_is"
	TypeMsgDeleteWhichIs = "delete_which_is"
)

var _ sdk.Msg = &MsgCreateWhichIs{}

func NewMsgCreateWhichIs(
	creator string,
	did string,
	bucketDoc *BucketDoc,
) *MsgCreateWhichIs {
	return &MsgCreateWhichIs{
		Creator: creator,
		Did:     did,
		Bucket:  bucketDoc,
	}
}

func NewMsgCreateWhichIsFromBuf(msg *bt.MsgCreateWhichIs) *MsgCreateWhichIs {
	return &MsgCreateWhichIs{
		Creator: msg.GetCreator(),
		Did:     msg.GetDid(),
		Bucket:  NewBucketDocFromBuf(msg.GetBucket()),
	}
}

func (msg *MsgCreateWhichIs) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhichIs) Type() string {
	return TypeMsgCreateWhichIs
}

func (msg *MsgCreateWhichIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhichIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhichIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhichIs{}

func NewMsgUpdateWhichIs(
	creator string,
	did string,
	bucketDoc *BucketDoc,
) *MsgUpdateWhichIs {
	return &MsgUpdateWhichIs{
		Creator: creator,
		Did:     did,
		Bucket:  bucketDoc,
	}
}

func (msg *MsgUpdateWhichIs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhichIs) Type() string {
	return TypeMsgUpdateWhichIs
}

func (msg *MsgUpdateWhichIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWhichIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhichIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWhichIs{}

func NewMsgDeleteWhichIs(
	creator string,
	did string,
) *MsgDeleteWhichIs {
	return &MsgDeleteWhichIs{
		Creator: creator,
		Did:     did,
	}
}
func (msg *MsgDeleteWhichIs) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWhichIs) Type() string {
	return TypeMsgDeleteWhichIs
}

func (msg *MsgDeleteWhichIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWhichIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhichIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
