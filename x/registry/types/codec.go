package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterApplication{}, "registry/RegisterApplication", nil)
	cdc.RegisterConcrete(&MsgRegisterName{}, "registry/RegisterName", nil)
	cdc.RegisterConcrete(&MsgAccessName{}, "registry/AccessName", nil)
	cdc.RegisterConcrete(&MsgUpdateName{}, "registry/UpdateName", nil)
	cdc.RegisterConcrete(&MsgAccessApplication{}, "registry/AccessApplication", nil)
	cdc.RegisterConcrete(&MsgUpdateApplication{}, "registry/UpdateApplication", nil)
	cdc.RegisterConcrete(&MsgCreateWhoIs{}, "registry/CreateWhoIs", nil)
	cdc.RegisterConcrete(&MsgUpdateWhoIs{}, "registry/UpdateWhoIs", nil)
	cdc.RegisterConcrete(&MsgDeleteWhoIs{}, "registry/DeleteWhoIs", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterApplication{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccessName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccessApplication{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateApplication{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWhoIs{},
		&MsgUpdateWhoIs{},
		&MsgDeleteWhoIs{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
