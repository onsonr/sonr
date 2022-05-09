package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rt "go.buf.build/grpc/go/sonr-io/blockchain/registry"
)

const TypeMsgUpdateApplication = "update_application"

var _ sdk.Msg = &MsgUpdateApplication{}

func NewMsgUpdateApplication(creator string, did string, m map[string]string, s *Session) *MsgUpdateApplication {
	return &MsgUpdateApplication{
		Creator:  creator,
		Did:      did,
		Metadata: m,
		Session:  s,
	}
}

func NewMsgUpdateApplicationFromBuf(msg *rt.MsgUpdateApplication) *MsgUpdateApplication {
	return NewMsgUpdateApplication(msg.GetCreator(), msg.GetDid(), msg.GetMetadata(), NewSessionFromBuf(msg.GetSession()))
}

func (msg *MsgUpdateApplication) Route() string {
	return RouterKey
}

func (msg *MsgUpdateApplication) Type() string {
	return TypeMsgUpdateApplication
}

func (msg *MsgUpdateApplication) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateApplication) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateApplication) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
