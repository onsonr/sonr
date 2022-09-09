package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateSchema = "create_schema"

var _ sdk.Msg = &MsgCreateSchema{}

func NewMsgCreateSchema(metadata []*MetadataDefintion, fields []*SchemaKindDefinition, creator string, label string) *MsgCreateSchema {
	return &MsgCreateSchema{
		Metadata: metadata,
		Creator:  creator,
		Label:    label,
		Fields:   fields,
	}
}
func (msg *MsgCreateSchema) Route() string {
	return RouterKey
}
func (msg *MsgCreateSchema) Type() string {
	return TypeMsgCreateSchema
}

func (msg *MsgCreateSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Label) < 1 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Label must be defined and non empty")
	}

	if len(msg.Fields) < 1 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Fields cannot be empty")
	}

	return nil
}

// GetCreatorDid returns the creator did
func (msg *MsgCreateSchema) GetCreatorDid() string {
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
