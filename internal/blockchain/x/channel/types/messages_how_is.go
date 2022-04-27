package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ct "go.buf.build/grpc/go/sonr-io/blockchain/channel"
)

const (
	TypeMsgCreateHowIs = "create_how_is"
	TypeMsgUpdateHowIs = "update_how_is"
	TypeMsgDeleteHowIs = "delete_how_is"
)

var _ sdk.Msg = &MsgCreateHowIs{}

func NewMsgCreateHowIs(
	creator string,
	did string,
	c *ChannelDoc,
) *MsgCreateHowIs {
	return &MsgCreateHowIs{
		Creator: creator,
		Did:     did,
		Channel: c,
	}
}

func NewMsgCreateHowIsFromBuf(msg *ct.MsgCreateHowIs) *MsgCreateHowIs {
	return &MsgCreateHowIs{
		Creator: msg.GetCreator(),
		Did:     msg.GetDid(),
		Channel: NewChannelDocFromBuf(msg.GetChannel()),
	}
}

func (msg *MsgCreateHowIs) Route() string {
	return RouterKey
}

func (msg *MsgCreateHowIs) Type() string {
	return TypeMsgCreateHowIs
}

func (msg *MsgCreateHowIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHowIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHowIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHowIs{}

func NewMsgUpdateHowIs(
	creator string,
	did string,
	c *ChannelDoc,

) *MsgUpdateHowIs {
	return &MsgUpdateHowIs{
		Creator: creator,
		Did:     did,
		Channel: c,
	}
}

func (msg *MsgUpdateHowIs) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHowIs) Type() string {
	return TypeMsgUpdateHowIs
}

func (msg *MsgUpdateHowIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHowIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHowIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHowIs{}

func NewMsgDeleteHowIs(
	creator string,
	did string,
) *MsgDeleteHowIs {
	return &MsgDeleteHowIs{
		Creator: creator,
		Did:     did,
	}
}
func (msg *MsgDeleteHowIs) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHowIs) Type() string {
	return TypeMsgDeleteHowIs
}

func (msg *MsgDeleteHowIs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHowIs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHowIs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
