package room

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	"github.com/sonr-io/core/internal/emitter"
	ac "github.com/sonr-io/core/pkg/account"
	"github.com/sonr-io/core/pkg/data"
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
	emitter *emitter.Emitter
	linkers []*data.Peer
	room    GetRoomFunc
	account ac.Account
}

// initExchange Initializes Exchange Service by Room Type
func (rm *RoomManager) initExchange() *data.SonrError {
	// Start Exchange RPC Server
	exchangeServer := rpc.NewServer(rm.host.Host(), util.EXCHANGE_PROTOCOL)
	esv := ExchangeService{
		account: rm.account,
		emitter: rm.emitter,
		linkers: rm.linkers,
		room:    rm.Room,
	}

	// Register Service
	err := exchangeServer.RegisterName(util.EXCHANGE_RPC_SERVICE, &esv)
	if err != nil {
		return data.NewError(err, data.ErrorEvent_ROOM_RPC)
	}

	// Set Service
	rm.exchange = &esv

	// Handle Events
	go rm.handleExchangeEvents(context.Background())
	go rm.handleExchangeMessages(context.Background())
	return nil
}

// Exchange method Starts Exchange on Local Peer Join
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
		data.LogError(err)
		return err
	}

	// Received Message
	remotePeer := &data.Member{}
	err = proto.Unmarshal(reply.Member, remotePeer)

	// Send Error
	if err != nil {
		data.LogError(err)
		return err
	}

	// Update Peer with new data
	if remotePeer.Active.Status != data.Peer_PAIRING {
		rm.emitter.Emit(emitter.EMIT_ROOM_EVENT, rm.room.NewJoinEvent(remotePeer))
	} else {
		// Add Linker if Not Present
		if !rm.HasLinker(remotePeer.Active.PeerID()) {
			// Append Linkers
			rm.linkers = append(rm.linkers, remotePeer.Active)
		}
	}
	return nil
}

// ExchangeWith method Calls Exchange on Local Lobby Peer
func (es *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeServiceArgs, reply *ExchangeServiceResponse) error {
	// Peer Data
	remotePeer := &data.Member{}
	err := proto.Unmarshal(args.Member, remotePeer)
	if err != nil {
		data.LogError(err)
		return err
	}

	// Update Peers with Lobby
	if remotePeer.Active.Status != data.Peer_PAIRING {
		es.emitter.Emit(emitter.EMIT_ROOM_EVENT, es.room().NewJoinEvent(remotePeer))
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
		data.LogError(err)
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

// handleExchangeEvents method listens to Pubsub Events for room
func (rm *RoomManager) handleExchangeEvents(ctx context.Context) {
	// Loop Events
	for {
		// Get next event
		event, err := rm.eventHandler.NextPeerEvent(ctx)
		if err != nil {
			data.LogError(err)
			rm.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if rm.isEventJoin(event) {
			pbuf, err := proto.Marshal(rm.account.Member())
			if err != nil {
				data.LogError(err)
				continue
			}
			err = rm.Exchange(event.Peer, pbuf)
			if err != nil {
				data.LogError(err)
				continue
			}
		} else if rm.isEventExit(event) {
			rm.emitter.Emit(emitter.EMIT_ROOM_EVENT, rm.room.NewExitEvent(event.Peer.String()))
		}
		data.GetState().NeedsWait()
	}
}

// handleExchangeMessages method listens for messages on pubsub room subscription
func (rm *RoomManager) handleExchangeMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := rm.subscription.Next(ctx)
		if err != nil {
			data.LogError(err)
			return
		}

		// Only forward messages delivered by others
		if rm.isValidMessage(msg) {
			// Unmarshal RoomEvent
			m := &data.RoomEvent{}
			err = proto.Unmarshal(msg.Data, m)
			if err != nil {
				data.LogError(err)
				continue
			}

			// Check Peer is Online, if not ignore
			if m.Member.Active.GetStatus() == data.Peer_ONLINE {
				rm.emitter.Emit(emitter.EMIT_ROOM_EVENT, m)
			} else if m.Member.Active.GetStatus() == data.Peer_PAIRING {
				// Validate Linker not Already Set
				if !rm.HasLinker(m.Member.Active.PeerID()) {
					// Append Linkers
					rm.linkers = append(rm.linkers, m.Member.GetActive())
				}
			}
		}
		data.GetState().NeedsWait()
	}
}
