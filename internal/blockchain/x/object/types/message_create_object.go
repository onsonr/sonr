package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rtv1 "github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

const TypeMsgCreateObject = "create_object"

var _ sdk.Msg = &MsgCreateObject{}

func NewMsgCreateObjectFromBuf(ot *ot.MsgCreateObject) *MsgCreateObject {
	ot.GetInitialFields()
	return &MsgCreateObject{
		Creator:       ot.GetCreator(),
		Label:         ot.GetLabel(),
		Description:   ot.GetDescription(),
		InitialFields: NewTypeFieldListFromBuf(ot.GetInitialFields()),
		Session:       rtv1.NewSessionFromBuf(ot.GetSession()),
	}
}

func NewMsgCreateObject(creator string, label string, description string) *MsgCreateObject {
	return &MsgCreateObject{
		Creator:     creator,
		Label:       label,
		Description: description,
	}
}

func (msg *MsgCreateObject) Route() string {
	return RouterKey
}

func (msg *MsgCreateObject) Type() string {
	return TypeMsgCreateObject
}

func (msg *MsgCreateObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
