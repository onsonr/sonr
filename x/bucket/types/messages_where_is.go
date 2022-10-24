package types

import (
	fmt "fmt"
	"strings"

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

var _ sdk.Msg = &MsgUpdateBucket{}

func NewMsgUpdateBucket(creator string, id string) *MsgUpdateBucket {
	return &MsgUpdateBucket{
		Did:     id,
		Creator: creator,
	}
}

func (msg *MsgUpdateBucket) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBucket) Type() string {
	return TypeMsgUpdateBucket
}

func (msg *MsgUpdateBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteBucket{}

func NewMsgDeleteBucket(creator string, id string) *MsgDeleteBucket {
	return &MsgDeleteBucket{
		Did:     id,
		Creator: creator,
	}
}
func (msg *MsgDeleteBucket) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBucket) Type() string {
	return TypeMsgDeleteBucket
}

func (msg *MsgDeleteBucket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBucket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBucket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// GetCreatorDid returns the creator did
func (msg *MsgDefineBucket) GetCreatorDid() string {
	rawCreator := msg.GetCreator()

	// Trim snr account prefix
	if strings.HasPrefix(rawCreator, "snr") {
		rawCreator = strings.TrimLeft(rawCreator, "snr")
	}

	// Trim cosmos account prefix
	if strings.HasPrefix(rawCreator, "cosmos") {
		rawCreator = strings.TrimLeft(rawCreator, "cosmos")
	}
	return fmt.Sprintf("did:snr:%s", rawCreator)
}
