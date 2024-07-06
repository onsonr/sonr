package keeper

import (
	"github.com/onsonr/hway/x/oracle/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"

	"cosmossdk.io/log"

	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

// Keeper defines the middleware keeper.
type Keeper struct {
	cdc              codec.BinaryCodec
	msgServiceRouter *baseapp.MsgServiceRouter

	ics4Wrapper porttypes.ICS4Wrapper
}

// NewKeeper creates a new swap Keeper instance.
func NewKeeper(
	cdc codec.BinaryCodec,
	msgServiceRouter *baseapp.MsgServiceRouter,
	ics4Wrapper porttypes.ICS4Wrapper,
) Keeper {
	return Keeper{
		cdc:              cdc,
		msgServiceRouter: msgServiceRouter,
		ics4Wrapper:      ics4Wrapper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+ibcexported.ModuleName+"-"+types.ModuleName)
}

// SendPacket wraps IBC ChannelKeeper's SendPacket function.
func (k Keeper) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	return k.ics4Wrapper.SendPacket(
		ctx,
		chanCap,
		sourcePort,
		sourceChannel,
		timeoutHeight,
		timeoutTimestamp,
		data,
	)
}

// WriteAcknowledgement wraps IBC ChannelKeeper's WriteAcknowledgement function.
func (k Keeper) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	acknowledgement ibcexported.Acknowledgement,
) error {
	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, acknowledgement)
}

// GetAppVersion wraps IBC ChannelKeeper's GetAppVersion function.
func (k Keeper) GetAppVersion(ctx sdk.Context, portID string, channelID string) (string, bool) {
	return k.ics4Wrapper.GetAppVersion(ctx, portID, channelID)
}
