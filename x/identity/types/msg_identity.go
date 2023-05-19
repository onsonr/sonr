package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterIdentity = "register_identity"

var _ sdk.Msg = &MsgRegisterIdentity{}

func NewMsgRegisterIdentity(creator string, wallet_id uint32, alias string, identity *Identity, relationships ...*VerificationRelationship) *MsgRegisterIdentity {
	msg := &MsgRegisterIdentity{
		Creator:              creator,
		WalletId:             wallet_id,
		Alias:                alias,
		Identity:             identity,
		Authentication:       make([]*VerificationRelationship, 0),
		Assertion:            make([]*VerificationRelationship, 0),
		KeyAgreement:         make([]*VerificationRelationship, 0),
		CapabilityInvocation: make([]*VerificationRelationship, 0),
		CapabilityDelegation: make([]*VerificationRelationship, 0),
	}

	for _, relationship := range relationships {
		if relationship.Type == AuthenticationRelationshipName {
			msg.Authentication = append(msg.Authentication, relationship)
		}
		if relationship.Type == AssertionRelationshipName {
			msg.Assertion = append(msg.Assertion, relationship)
		}
		if relationship.Type == KeyAgreementRelationshipName {
			msg.KeyAgreement = append(msg.KeyAgreement, relationship)
		}
		if relationship.Type == CapabilityInvocationRelationshipName {
			msg.CapabilityInvocation = append(msg.CapabilityInvocation, relationship)
		}
		if relationship.Type == CapabilityDelegationRelationshipName {
			msg.CapabilityDelegation = append(msg.CapabilityDelegation, relationship)
		}
	}
	return msg
}

func (msg *MsgRegisterIdentity) Route() string {
	return RouterKey
}

func (msg *MsgRegisterIdentity) Type() string {
	return TypeMsgRegisterIdentity
}

func (msg *MsgRegisterIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
