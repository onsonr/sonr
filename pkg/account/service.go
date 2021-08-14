package account

import (
	"context"
	"errors"

	crypto "github.com/libp2p/go-libp2p-core/crypto"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// DeviceServiceArgs ExchangeArgs are Peer, Device, and Contact
type DeviceServiceArgs struct {
	PubKeyBuf []byte
}

// SyncServiceResponse ExchangeResponse is Member protobuf
type DeviceServiceResponse struct {
	Success bool
}

// DeviceService Service Struct
type DeviceService struct {
	// Current Data
	call   Account
	room   GetRoomFunc
	device *md.Device
}

// Initialize Exchange Service by Room Type
func (rm *accountLinker) initService() *md.SonrError {
	// Start Exchange RPC Server
	verifyServer := rpc.NewServer(rm.host.Host(), util.ACCOUNT_PROTOCOL)
	verifyService := DeviceService{
		device: rm.account.Current,
		call:   rm,
		room:   rm.Room,
	}

	// Register Service
	err := verifyServer.RegisterName(util.DEVICE_RPC_SERVICE, &verifyService)
	if err != nil {
		return md.NewError(err, md.ErrorEvent_ROOM_RPC)
	}

	// Set Service
	rm.service = &verifyService
	go rm.handleVerifyEvents(rm.ctx)
	go rm.handleVerifyMessages(rm.ctx)
	return nil
}

// Exchange @ Starts Exchange on Local Peer Join
func (rm *accountLinker) Verify(id peer.ID) error {
	// Initialize RPC
	exchClient := rpc.NewClient(rm.host.Host(), util.ACCOUNT_PROTOCOL)
	var reply DeviceServiceResponse
	var args DeviceServiceArgs
	args.PubKeyBuf = rm.account.GetCurrent().DevicePubKeyBuf()

	// Verify with Peer
	err := exchClient.Call(id, util.DEVICE_RPC_SERVICE, util.DEVICE_METHOD_VERIFY, args, &reply)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Check for Success
	if !reply.Success {
		md.LogError(errors.New("Failed to Verify with Device"))
		rm.Topic.Close()
	}
	return nil
}

// ExchangeWith # Calls Exchange on Local Lobby Peer
func (ss *DeviceService) VerifyWith(ctx context.Context, args DeviceServiceArgs, reply *DeviceServiceResponse) error {
	// Unmarshal Public Key
	pubKey, err := crypto.UnmarshalPublicKey(args.PubKeyBuf)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Check if Public Keys Match
	reply.Success = ss.device.DeviceKeys().VerifyPubKey(pubKey)
	return nil
}

// # handleRoomEvents: listens to Pubsub Events for room
func (rm *accountLinker) handleVerifyEvents(ctx context.Context) {
	// Loop Events
	for {
		// Get next event
		event, err := rm.eventHandler.NextPeerEvent(ctx)
		if err != nil {
			md.LogError(err)
			rm.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if rm.isEventJoin(event) {
			err = rm.Verify(event.Peer)
			if err != nil {
				md.LogError(err)
				continue
			}
		} else if rm.isEventExit(event) {
			rm.OnRoomEvent(rm.room.NewExitEvent(event.Peer.String()))

		}
		md.GetState().NeedsWait()
	}
}

// # handleRoomMsssagss: listens for msssagss on pubsub room subscription
func (rm *accountLinker) handleVerifyMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := rm.subscription.Next(ctx)
		if err != nil {
			md.LogError(err)
			return
		}

		// Only forward msssagss delivered by others
		if rm.isValidMessage(msg) {
			// Unmarshal RoomEvent
			m := &md.RoomEvent{}
			err = proto.Unmarshal(msg.Data, m)
			if err != nil {
				md.LogError(err)
				continue
			}

			// Check Peer is Online, if not ignore
			rm.OnRoomEvent(m)
		}
		md.GetState().NeedsWait()
	}
}
