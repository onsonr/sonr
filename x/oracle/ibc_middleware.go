package oracle

import (
	"github.com/onsonr/sonr/x/oracle/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"

	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks for the middleware given the
// keeper and the underlying application.
type IBCMiddleware struct {
	app    porttypes.IBCModule
	keeper keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application.
func NewIBCMiddleware(app porttypes.IBCModule, k keeper.Keeper) IBCMiddleware {
	return IBCMiddleware{
		app:    app,
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCMiddleware interface.
func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	return im.app.OnChanOpenInit(
		ctx,
		order,
		connectionHops,
		portID,
		channelID,
		chanCap,
		counterparty,
		version,
	)
}

// OnChanOpenTry implements the IBCMiddleware interface.
func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID, channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (version string, err error) {
	return im.app.OnChanOpenTry(
		ctx,
		order,
		connectionHops,
		portID,
		channelID,
		chanCap,
		counterparty,
		counterpartyVersion,
	)
}

// OnChanOpenAck implements the IBCMiddleware interface.
func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID, channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCMiddleware interface.
func (im IBCMiddleware) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCMiddleware interface.
func (im IBCMiddleware) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCMiddleware interface.
func (im IBCMiddleware) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket implements the IBCMiddleware interface.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	return im.app.OnRecvPacket(ctx, packet, relayer)
}

// OnAcknowledgementPacket implements the IBCMiddleware interface.
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCMiddleware interface.
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}

// SendPacket implements the ICS4 Wrapper interface.
func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	return im.keeper.SendPacket(
		ctx,
		chanCap,
		sourceChannel,
		sourceChannel,
		timeoutHeight,
		timeoutTimestamp,
		data,
	)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface.
func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	return im.keeper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

// GetAppVersion implements the ICS4 Wrapper interface.
func (im IBCMiddleware) GetAppVersion(
	ctx sdk.Context,
	portID string,
	channelID string,
) (string, bool) {
	return im.keeper.GetAppVersion(ctx, portID, channelID)
}
