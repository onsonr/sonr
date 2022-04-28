package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rt "github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

const TypeMsgCreateBucket = "create_bucket"

var _ sdk.Msg = &MsgCreateBucket{}

func NewMsgCreateBucket(creator string, label string, description string, kind string, session *rt.Session, ibDids []string) *MsgCreateBucket {
	return &MsgCreateBucket{
		Creator:           creator,
		Label:             label,
		Description:       description,
		Kind:              kind,
		Session:           session,
		InitialObjectDids: ibDids,
	}
}

func NewMsgCreateBucketFromBuf(msg *bt.MsgCreateBucket) *MsgCreateBucket {
	return NewMsgCreateBucket(msg.GetCreator(), msg.GetLabel(), msg.GetDescription(), msg.GetKind(), rt.NewSessionFromBuf(msg.GetSession()), msg.GetInitialObjectDids())
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
