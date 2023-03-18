package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDidDocument{}, "identity/CreateDidDocument", nil)
	cdc.RegisterConcrete(&MsgUpdateDidDocument{}, "identity/UpdateDidDocument", nil)
	cdc.RegisterConcrete(&MsgDeleteDidDocument{}, "identity/DeleteDidDocument", nil)
	cdc.RegisterConcrete(&MsgRegisterService{}, "identity/RegisterService", nil)
	cdc.RegisterConcrete(&MsgRegisterAccount{}, "identity/RegisterAccount", nil)
	cdc.RegisterConcrete(&MsgImportPublicKey{}, "identity/ImportPublicKey", nil)
	cdc.RegisterConcrete(&MsgDeletePublicKey{}, "identity/DeletePublicKey", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDidDocument{},
		&MsgUpdateDidDocument{},
		&MsgDeleteDidDocument{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgImportPublicKey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeletePublicKey{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

const contextKey = "@context"
const controllerKey = "controller"
const authenticationKey = "authentication"
const assertionMethodKey = "assertionMethod"
const keyAgreementKey = "keyAgreement"
const capabilityInvocationKey = "capabilityInvocation"
const capabilityDelegationKey = "capabilityDelegation"
const verificationMethodKey = "verificationMethod"
const serviceEndpointKey = "serviceEndpoint"
