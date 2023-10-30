package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

const (
	AuthenticationRelationshipName       = "Authentication"
	AssertionRelationshipName            = "AssertionMethod"
	KeyAgreementRelationshipName         = "KeyAgreement"
	CapabilityInvocationRelationshipName = "CapabilityInvocation"
	CapabilityDelegationRelationshipName = "CapabilityDelegation"
)

var (
	VerificationRelationshipNames = []string{
		AuthenticationRelationshipName,
		AssertionRelationshipName,
		KeyAgreementRelationshipName,
		CapabilityInvocationRelationshipName,
		CapabilityDelegationRelationshipName,
	}
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterIdentity{}, "identity/RegisterIdentity", nil)
	cdc.RegisterConcrete(&MsgCreateControllerAccount{}, "identity/CreateControllerAccount", nil)
	cdc.RegisterConcrete(&MsgUpdateControllerAccount{}, "identity/UpdateControllerAccount", nil)
	cdc.RegisterConcrete(&MsgDeleteControllerAccount{}, "identity/DeleteControllerAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterIdentity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateControllerAccount{},
		&MsgUpdateControllerAccount{},
		&MsgDeleteControllerAccount{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
