package topic

import (
	"context"
	"errors"

	crypto "github.com/libp2p/go-libp2p-core/crypto"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// VerifyServiceArgs ExchangeArgs are Peer, Device, and Contact
type VerifyServiceArgs struct {
	PubKeyBuf []byte
}

// SyncServiceRssponse ExchangeRssponse is Member protobuf
type VerifyServiceResponse struct {
	Success bool
}

// VerifyService Service Struct
type VerifyService struct {
	// Current Data
	call   RoomHandler
	room   GetRoomFunc
	device *md.Device
}

// Initialize Exchange Service by Room Type
func (rm *RoomManager) initSync() *md.SonrError {
	// Start Exchange RPC Server
	verifyServer := rpc.NewServer(rm.host.Host(), util.VERIFY_PROTOCOL)
	verifyService := VerifyService{
		device: rm.device,
		call:   rm.handler,
		room:   rm.Room,
	}

	// Register Service
	err := verifyServer.RegisterName(util.VERIFY_RPC_SERVICE, &verifyService)
	if err != nil {
		return md.NewError(err, md.ErrorEvent_ROOM_RPC)
	}

	// Set Service
	rm.verify = &verifyService
	return nil
}

// Exchange @ Starts Exchange on Local Peer Join
func (rm *RoomManager) Verify(id peer.ID) error {
	// Initialize RPC
	exchClient := rpc.NewClient(rm.host.Host(), util.VERIFY_PROTOCOL)
	var reply VerifyServiceResponse
	var args VerifyServiceArgs
	args.PubKeyBuf = rm.device.DevicePubKeyBuf()

	// Verify with Peer
	err := exchClient.Call(id, util.VERIFY_RPC_SERVICE, util.VERIFY_METHOD_VERIFY, args, &reply)
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
func (ss *VerifyService) VerifyWith(ctx context.Context, args VerifyServiceArgs, reply *VerifyServiceResponse) error {
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
func (rm *RoomManager) handleSyncEvents(ctx context.Context) {
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
			rm.handler.OnRoomEvent(rm.room.NewExitEvent(event.Peer.String()))

		}
		md.GetState().NeedsWait()
	}
}

// # handleRoomMsssagss: listens for msssagss on pubsub room subscription
func (rm *RoomManager) handleSyncMessages(ctx context.Context) {
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
			if m.Peer.GetStatus() == md.Peer_ONLINE {
				rm.handler.OnRoomEvent(m)
			} else if m.Peer.GetStatus() == md.Peer_PAIRING {
				// Validate Linker not Already Set
				if !rm.HasLinker(m.Peer.PeerID()) {
					// Append Linkers
					rm.linkers = append(rm.linkers, m.Peer)
				}
			}
		}
		md.GetState().NeedsWait()
	}
}
