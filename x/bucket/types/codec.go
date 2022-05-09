package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateBucket{}, "bucket/CreateBucket", nil)
	cdc.RegisterConcrete(&MsgUpdateBucket{}, "bucket/UpdateBucket", nil)
	cdc.RegisterConcrete(&MsgDeactivateBucket{}, "bucket/DeactivateBucket", nil)
	cdc.RegisterConcrete(&MsgCreateWhichIs{}, "bucket/CreateWhichIs", nil)
	cdc.RegisterConcrete(&MsgUpdateWhichIs{}, "bucket/UpdateWhichIs", nil)
	cdc.RegisterConcrete(&MsgDeleteWhichIs{}, "bucket/DeleteWhichIs", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBucket{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateBucket{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeactivateBucket{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWhichIs{},
		&MsgUpdateWhichIs{},
		&MsgDeleteWhichIs{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
