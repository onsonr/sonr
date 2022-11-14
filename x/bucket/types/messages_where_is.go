package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgDefineBucket = "create_where_is"
	TypeMsgUpdateBucket = "update_where_is"
	TypeMsgDeleteBucket = "delete_where_is"
)

var _ sdk.Msg = &MsgDefineBucket{}

func NewMsgDefineBucket(creator string, label string) *MsgDefineBucket {
	return &MsgDefineBucket{
		Creator: creator,
		Label:   label,
	}
}

func (msg *MsgDefineBucket) Route() string {
	return RouterKey
}

func (msg *MsgDefineBucket) Type() string {
	return TypeMsgDefineBucket
}

func (msg *MsgDefineBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDefineBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDefineBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
