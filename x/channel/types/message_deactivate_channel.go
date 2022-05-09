package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rt "github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
	ct "go.buf.build/grpc/go/sonr-io/blockchain/channel"
)

const TypeMsgDeactivateChannel = "delete_channel"

var _ sdk.Msg = &MsgDeactivateChannel{}

func NewMsgDeactivateChannel(creator string, did string, s *rt.Session) *MsgDeactivateChannel {
	return &MsgDeactivateChannel{
		Creator: creator,
		Did:     did,
		Session: s,
	}
}

func NewMsgDeactivateChannelFromBuf(msg *ct.MsgDeactivateChannel) *MsgDeactivateChannel {
	return &MsgDeactivateChannel{
		Creator: msg.GetCreator(),
		Did:     msg.GetDid(),
		Session: rt.NewSessionFromBuf(msg.GetSession()),
	}
}

func (msg *MsgDeactivateChannel) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateChannel) Type() string {
	return TypeMsgDeactivateChannel
}

func (msg *MsgDeactivateChannel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeactivateChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeactivateChannel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
