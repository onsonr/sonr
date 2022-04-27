package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateChannel{}, "channel/CreateChannel", nil)
	cdc.RegisterConcrete(&MsgDeactivateChannel{}, "channel/DeactivateChannel", nil)
	cdc.RegisterConcrete(&MsgUpdateChannel{}, "channel/UpdateChannel", nil)
	cdc.RegisterConcrete(&MsgCreateHowIs{}, "channel/CreateHowIs", nil)
	cdc.RegisterConcrete(&MsgUpdateHowIs{}, "channel/UpdateHowIs", nil)
	cdc.RegisterConcrete(&MsgDeleteHowIs{}, "channel/DeleteHowIs", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChannel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeactivateChannel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateChannel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHowIs{},
		&MsgUpdateHowIs{},
		&MsgDeleteHowIs{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
