package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	crypto "github.com/sonrhq/core/internal/crypto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDidDocument{}, "identity/CreateDidDocument", nil)
	cdc.RegisterConcrete(&MsgUpdateDidDocument{}, "identity/UpdateDidDocument", nil)
	cdc.RegisterConcrete(&MsgDeleteDidDocument{}, "identity/DeleteDidDocument", nil)
	cdc.RegisterConcrete(&MsgCreateClaimableWallet{}, "identity/CreateClaimableWallet", nil)
	cdc.RegisterConcrete(&MsgUpdateClaimableWallet{}, "identity/UpdateClaimableWallet", nil)
	cdc.RegisterConcrete(&MsgDeleteClaimableWallet{}, "identity/DeleteClaimableWallet", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDidDocument{},
		&MsgUpdateDidDocument{},
		&MsgDeleteDidDocument{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClaimableWallet{},
		&MsgUpdateClaimableWallet{},
		&MsgDeleteClaimableWallet{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	registry.RegisterInterface(
		"sonrhq.sonr.crypto.PubKey",
		(*cryptotypes.PubKey)(nil), // Fix this line
	)

	// Register the concrete implementation(s) of the custom PubKey
	registry.RegisterImplementations(
		(*cryptotypes.PubKey)(nil), // Fix this line
		&crypto.PubKey{},           // Replace with the concrete implementation of your custom PubKey
	)

	cryptocodec.RegisterInterfaces(registry)
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
