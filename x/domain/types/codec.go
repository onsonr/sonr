package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateTLDRecord{}, "domain/CreateTLDRecord", nil)
	cdc.RegisterConcrete(&MsgUpdateTLDRecord{}, "domain/UpdateTLDRecord", nil)
	cdc.RegisterConcrete(&MsgDeleteTLDRecord{}, "domain/DeleteTLDRecord", nil)
	cdc.RegisterConcrete(&MsgCreateSLDRecord{}, "domain/CreateSLDRecord", nil)
	cdc.RegisterConcrete(&MsgUpdateSLDRecord{}, "domain/UpdateSLDRecord", nil)
	cdc.RegisterConcrete(&MsgDeleteSLDRecord{}, "domain/DeleteSLDRecord", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTLDRecord{},
		&MsgUpdateTLDRecord{},
		&MsgDeleteTLDRecord{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSLDRecord{},
		&MsgUpdateSLDRecord{},
		&MsgDeleteSLDRecord{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
