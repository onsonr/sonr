package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rtv1 "github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

const TypeMsgUpdateObject = "update_object"

var _ sdk.Msg = &MsgUpdateObject{}

func NewMsgUpdateObjectFromBuf(ot *ot.MsgUpdateObject) *MsgUpdateObject {
	return &MsgUpdateObject{
		Creator:       ot.GetCreator(),
		Label:         ot.GetLabel(),
		AddedFields:   NewTypeFieldListFromBuf(ot.GetAddedFields()),
		RemovedFields: NewTypeFieldListFromBuf(ot.GetRemovedFields()),
		Session:       rtv1.NewSessionFromBuf(ot.GetSession()),
	}
}

// TODO: Add validation for fields
func NewMsgUpdateObject(creator string, did string) *MsgUpdateObject {
	return &MsgUpdateObject{
		Creator: creator,
	}
}

func (msg *MsgUpdateObject) Route() string {
	return RouterKey
}

func (msg *MsgUpdateObject) Type() string {
	return TypeMsgUpdateObject
}

func (msg *MsgUpdateObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
