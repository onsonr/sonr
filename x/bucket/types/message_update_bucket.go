package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rt "github.com/sonr-io/sonr/x/registry/types"
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

const TypeMsgUpdateBucket = "update_bucket"

var _ sdk.Msg = &MsgUpdateBucket{}

func NewMsgUpdateBucket(creator string, label string, description string, session *rt.Session, addObjs []string, removObjs []string) *MsgUpdateBucket {
	return &MsgUpdateBucket{
		Creator:           creator,
		Label:             label,
		Description:       description,
		Session:           session,
		AddedObjectDids:   addObjs,
		RemovedObjectDids: removObjs,
	}
}

func NewMsgUpdateBucketFromBuf(msg *bt.MsgUpdateBucket) *MsgUpdateBucket {
	return NewMsgUpdateBucket(msg.GetCreator(), msg.GetLabel(), msg.GetDescription(), rt.NewSessionFromBuf(msg.GetSession()), msg.GetAddedObjectDids(), msg.GetRemovedObjectDids())
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
