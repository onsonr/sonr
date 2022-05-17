package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateWhoIs{}, "registry/CreateWhoIs", nil)
	cdc.RegisterConcrete(&MsgUpdateWhoIs{}, "registry/UpdateWhoIs", nil)
	cdc.RegisterConcrete(&MsgDeactivateWhoIs{}, "registry/DeactivateWhoIs", nil)
	cdc.RegisterConcrete(&MsgBuyAlias{}, "registry/BuyAlias", nil)
	cdc.RegisterConcrete(&MsgSellAlias{}, "registry/SellAlias", nil)
	cdc.RegisterConcrete(&MsgTransferAlias{}, "registry/TransferAlias", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWhoIs{},
		&MsgUpdateWhoIs{},
		&MsgDeactivateWhoIs{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyAlias{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSellAlias{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferAlias{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
