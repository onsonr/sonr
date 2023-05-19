package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterServiceRecord{}, "service/RegisterServiceRecord", nil)
	cdc.RegisterConcrete(&MsgUpdateServiceRecord{}, "service/UpdateServiceRecord", nil)
	cdc.RegisterConcrete(&MsgBurnServiceRecord{}, "service/BurnServiceRecord", nil)

	cdc.RegisterConcrete(&MsgRegisterUserEntity{}, "service/RegisterUserEntity", nil)
	cdc.RegisterConcrete(&MsgAuthenticateUserEntity{}, "service/AuthenticateUserEntity", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterServiceRecord{},
		&MsgUpdateServiceRecord{},
		&MsgBurnServiceRecord{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterUserEntity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAuthenticateUserEntity{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
