package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

const TypeMsgDeactivateObject = "delete_object"

var _ sdk.Msg = &MsgDeactivateObject{}

func NewMsgDeactivateObject(creator string, did string) *MsgDeactivateObject {
	return &MsgDeactivateObject{
		Creator: creator,
		Did:     did,
	}
}

func NewMsgDeactivateObjectFromBuf(msg *ot.MsgDeactivateObject) *MsgDeactivateObject {
	return NewMsgDeactivateObject(msg.GetCreator(), msg.GetDid())
}

func (msg *MsgDeactivateObject) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateObject) Type() string {
	return TypeMsgDeactivateObject
}

func (msg *MsgDeactivateObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeactivateObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeactivateObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
