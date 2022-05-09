package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rt "github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

const TypeMsgDeactivateBucket = "delete_bucket"

var _ sdk.Msg = &MsgDeactivateBucket{}

func NewMsgDeactivateBucket(creator string, did string, session *rt.Session) *MsgDeactivateBucket {
	return &MsgDeactivateBucket{
		Creator: creator,
		Did:     did,
		Session: session,
	}
}

func NewMsgDeactivateBucketFromBuf(msg *bt.MsgDeactivateBucket) *MsgDeactivateBucket {
	return NewMsgDeactivateBucket(msg.GetCreator(), msg.GetDid(), rt.NewSessionFromBuf(msg.GetSession()))
}

func (msg *MsgDeactivateBucket) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateBucket) Type() string {
	return TypeMsgDeactivateBucket
}

func (msg *MsgDeactivateBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeactivateBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeactivateBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
