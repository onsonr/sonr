package topic

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	ac "github.com/sonr-io/core/pkg/account"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// ExchangeServiceArgs ExchangeArgs is Peer protobuf
type ExchangeServiceArgs struct {
	Member []byte
}

// ExchangeServiceResponse ExchangeResponse is also Peer protobuf
type ExchangeServiceResponse struct {
	Member []byte
}

// ExchangeService Service Struct
type ExchangeService struct {
	// Current Data
	call    RoomHandler
	linkers []*md.Peer
	room    GetRoomFunc
	account ac.Account
}

// Initialize Exchange Service by Room Type
func (rm *RoomManager) initExchange() *md.SonrError {
	// Start Exchange RPC Server
	exchangeServer := rpc.NewServer(rm.host.Host(), util.EXCHANGE_PROTOCOL)
	esv := ExchangeService{
		account: rm.account,
		call:    rm.handler,
		linkers: rm.linkers,
		room:    rm.Room,
	}

	// Register Service
	err := exchangeServer.RegisterName(util.EXCHANGE_RPC_SERVICE, &esv)
	if err != nil {
		return md.NewError(err, md.ErrorEvent_ROOM_RPC)
	}

	// Set Service
	rm.exchange = &esv

	// Handle Events
	go rm.handleExchangeEvents(context.Background())
	go rm.handleExchangeMessages(context.Background())
	return nil
}

// Exchange @ Starts Exchange on Local Peer Join
func (rm *RoomManager) Exchange(id peer.ID, peerBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(rm.host.Host(), util.EXCHANGE_PROTOCOL)
	var reply ExchangeServiceResponse
	var args ExchangeServiceArgs

	// Set Args
	args.Member = peerBuf

	// Call to Peer
	err := exchClient.Call(id, util.EXCHANGE_RPC_SERVICE, util.EXCHANGE_METHOD_EXCHANGE, args, &reply)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Received Message
	remotePeer := &md.Member{}
	err = proto.Unmarshal(reply.Member, remotePeer)

	// Send Error
	if err != nil {
		md.LogError(err)
		return err
	}

	// Update Peer with new data
	if remotePeer.Active.Status != md.Peer_PAIRING {
		rm.handler.OnRoomEvent(rm.room.NewJoinEvent(remotePeer))
	} else {
		// Add Linker if Not Present
		if !rm.HasLinker(remotePeer.Active.PeerID()) {
			// Append Linkers
			rm.linkers = append(rm.linkers, remotePeer.Active)
		}
	}
	return nil
}

// ExchangeWith # Calls Exchange on Local Lobby Peer
func (es *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeServiceArgs, reply *ExchangeServiceResponse) error {
	// Peer Data
	remotePeer := &md.Member{}
	err := proto.Unmarshal(args.Member, remotePeer)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Update Peers with Lobby
	if remotePeer.Active.Status != md.Peer_PAIRING {
		es.call.OnRoomEvent(es.room().NewJoinEvent(remotePeer))
	} else {
		// Add Linker if Not Present
		if !es.HasLinker(remotePeer.Active.PeerID()) {
			// Append Linkers
			es.linkers = append(es.linkers, remotePeer.Active)
		}
	}

	// Set Message data and call done
	buf, err := proto.Marshal(es.account.Member())
	if err != nil {
		md.LogError(err)
		return err
	}
	reply.Member = buf
	return nil
}

func (es *ExchangeService) HasLinker(q string) bool {
	for _, p := range es.linkers {
		if p.PeerID() == q {
			return true
		}
	}
	return false
}

// # handleExchangeEvents: listens to Pubsub Events for room
func (rm *RoomManager) handleExchangeEvents(ctx context.Context) {
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
			pbuf, err := rm.device.GetPeer().Buffer()
			if err != nil {
				md.LogError(err)
				continue
			}
			err = rm.Exchange(event.Peer, pbuf)
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

// # handleExchangeMessages: listens for messages on pubsub room subscription
func (rm *RoomManager) handleExchangeMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := rm.subscription.Next(ctx)
		if err != nil {
			md.LogError(err)
			return
		}

		// Only forward messages delivered by others
		if rm.isValidMessage(msg) {
			// Unmarshal RoomEvent
			m := &md.RoomEvent{}
			err = proto.Unmarshal(msg.Data, m)
			if err != nil {
				md.LogError(err)
				continue
			}

			// Check Peer is Online, if not ignore
			if m.Member.Active.GetStatus() == md.Peer_ONLINE {
				rm.handler.OnRoomEvent(m)
			} else if m.Member.Active.GetStatus() == md.Peer_PAIRING {
				// Validate Linker not Already Set
				if !rm.HasLinker(m.Member.Active.PeerID()) {
					// Append Linkers
					rm.linkers = append(rm.linkers, m.Member.GetActive())
				}
			}
		}
		md.GetState().NeedsWait()
	}
}
