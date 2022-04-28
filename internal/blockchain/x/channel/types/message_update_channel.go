package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ot "github.com/sonr-io/sonr/internal/blockchain/x/object/types"
	rt "github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
	ct "go.buf.build/grpc/go/sonr-io/blockchain/channel"
)

const TypeMsgUpdateChannel = "update_channel"

var _ sdk.Msg = &MsgUpdateChannel{}

func NewMsgUpdateChannel(creator string, did string, label string, description string, objectToRegister *ot.ObjectDoc, session *rt.Session) *MsgUpdateChannel {
	return &MsgUpdateChannel{
		Creator:          creator,
		Did:              did,
		Label:            label,
		Session:          session,
		ObjectToRegister: objectToRegister,
		Description:      description,
	}
}

func NewMsgUpdateChannelFromBuf(msg *ct.MsgUpdateChannel) *MsgUpdateChannel {
	return &MsgUpdateChannel{
		Creator:          msg.GetCreator(),
		Did:              msg.GetDid(),
		Label:            msg.GetLabel(),
		Description:      msg.GetDescription(),
		ObjectToRegister: ot.NewObjectDocFromBuf(msg.GetObjectToRegister()),
		Session:          rt.NewSessionFromBuf(msg.GetSession()),
	}
}
func (msg *MsgUpdateChannel) Route() string {
	return RouterKey
}

func (msg *MsgUpdateChannel) Type() string {
	return TypeMsgUpdateChannel
}

func (msg *MsgUpdateChannel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChannel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
