package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateObject{}, "object/CreateObject", nil)
	cdc.RegisterConcrete(&MsgUpdateObject{}, "object/UpdateObject", nil)
	cdc.RegisterConcrete(&MsgDeactivateObject{}, "object/DeactivateObject", nil)
	cdc.RegisterConcrete(&MsgCreateWhatIs{}, "object/CreateWhatIs", nil)
	cdc.RegisterConcrete(&MsgUpdateWhatIs{}, "object/UpdateWhatIs", nil)
	cdc.RegisterConcrete(&MsgDeleteWhatIs{}, "object/DeleteWhatIs", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateObject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateObject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeactivateObject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWhatIs{},
		&MsgUpdateWhatIs{},
		&MsgDeleteWhatIs{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
