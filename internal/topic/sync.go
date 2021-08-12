package topic

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// SyncServiceArgs ExchangeArgs are Peer, Device, and Contact
type SyncServiceArgs struct {
	Contact []byte
	Device  []byte
	Peer    []byte
}

// SyncServiceRssponse ExchangeRssponse is Member protobuf
type SyncServiceResponse struct {
	Succsss bool
	Contact []byte
	Device  []byte
	Peer    []byte
}

// SyncService Service Struct
type SyncService struct {
	// Current Data
	call    RoomHandler
	room    GetRoomFunc
	user    *md.Device
}

// Initialize Exchange Service by Room Type
func (rm *RoomManager) initSync() *md.SonrError {
	// Start Exchange RPC Server
	syncServer := rpc.NewServer(rm.host.Host(), util.SYNC_PROTOCOL)
	syncService := SyncService{
		user:    rm.user,
		call:    rm.handler,
		room:    rm.Room,
	}

	// Register Service
	err := syncServer.RegisterName(util.SYNC_RPC_SERVICE, &syncService)
	if err != nil {
		return md.NewError(err, md.ErrorEvent_ROOM_RPC)
	}

	// Set Service
	rm.sync = &syncService
	return nil
}

// Exchange @ Starts Exchange on Local Peer Join
func (rm *RoomManager) Sync(id peer.ID, peerBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(rm.host.Host(), util.SYNC_PROTOCOL)
	var reply SyncServiceResponse
	var args SyncServiceArgs

	// Set Args
	args.Peer = peerBuf

	// Call to Peer
	err := exchClient.Call(id, util.SYNC_RPC_SERVICE, util.EXCHANGE_METHOD_EXCHANGE, args, &reply)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Received Msssage
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Peer, remotePeer)

	// Send Error
	if err != nil {
		md.LogError(err)
		return err
	}
	return nil
}

// ExchangeWith # Calls Exchange on Local Lobby Peer
func (ss *SyncService) SyncWith(ctx context.Context, args SyncServiceArgs, reply *SyncServiceResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		md.LogError(err)
		return err
	}

	ss.call.OnRoomEvent(ss.room().NewJoinEvent(remotePeer))

	// Set Msssage data and call done
	buf, err := ss.user.GetPeer().Buffer()
	if err != nil {
		md.LogError(err)
		return err
	}
	reply.Peer = buf
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
			pbuf, err := rm.user.GetPeer().Buffer()
			if err != nil {
				md.LogError(err)
				continue
			}
			err = rm.Sync(event.Peer, pbuf)
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
