package topic

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	ps "github.com/libp2p/go-libp2p-pubsub"
	sh "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)
type GetRoomFunc func() *md.Room

type RoomHandler interface {
	OnRoomEvent(*md.RoomEvent)
	OnSyncEvent(*md.SyncEvent)
}

type RoomManager struct {
	ctx          context.Context
	host         sh.HostNode
	Topic        *ps.Topic
	subscription *ps.Subscription
	eventHandler *ps.TopicEventHandler
	user         *md.Device

	roomEvents chan *md.RoomEvent
	syncEvents chan *md.SyncEvent
	exchange   *ExchangeService
	sync       *SyncService
	handler    RoomHandler
	linkers    []*md.Peer
	room       *md.Room
}

// NewLocal ^ Create New Contained Room Manager ^ //
func JoinRoom(ctx context.Context, h sh.HostNode, u *md.Device, room *md.Room, th RoomHandler) (*RoomManager, *md.SonrError) {
	// Join Room
	name := room.GetName()
	topic, err := h.Pubsub().Join(name)
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_ROOM_JOIN)
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_ROOM_SUB)
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_ROOM_HANDLER)
	}

	// Create Lobby Manager
	mgr := &RoomManager{
		handler:      th,
		user:         u,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		room:         room,
		linkers:      make([]*md.Peer, 0),
		roomEvents:   make(chan *md.RoomEvent, util.MAX_CHAN_DATA),
		subscription: sub,
		Topic:        topic,
	}

	// Start Exchange RPC Server
	exchangeServer := rpc.NewServer(h.Host(), util.EXCHANGE_PROTOCOL)
	esv := ExchangeService{
		user:    u,
		call:    th,
		linkers: mgr.linkers,
	}

	// Register Service
	err = exchangeServer.RegisterName(util.EXCHANGE_RPC_SERVICE, &esv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_ROOM_RPC)
	}

	// Set Service
	mgr.exchange = &esv

	// Handle Events
	go mgr.handleExchangeEvents(context.Background())
	go mgr.handleExchangeMessages(context.Background())
	return mgr, nil
}

// FindPeer @ Helper: Find returns Pointer to Peer.ID and Peer
func (tm *RoomManager) FindPeer(q string) (peer.ID, error) {
	// Iterate through Room Peers
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return id, nil
		}
	}
	return "", errors.New("Peer ID was not found in room")
}

// Publish @ Publish message to specific peer in room
func (tm *RoomManager) Publish(msg *md.RoomEvent) error {
	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(msg)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Publish to Room
	err = tm.Topic.Publish(tm.ctx, bytes)
	if err != nil {
		md.LogError(err)
		return err
	}
	return nil
}

// HasLinker Method Checks if Peer ID String is a listed Linker
func (tm *RoomManager) HasLinker(q string) bool {
	for _, p := range tm.linkers {
		if p.PeerID() == q && tm.HasPeer(q) {
			return true
		}
	}
	return false
}

// HasPeer Method Checks if Peer ID String is Subscribed to Room
func (tm *RoomManager) HasPeer(q string) bool {
	// Iterate through PubSub in room
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// HasPeer Method Checks if Peer ID is Subscribed to Room
func (tm *RoomManager) HasPeerID(q peer.ID) bool {
	// Iterate through PubSub in room
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id == q {
			return true
		}
	}
	return false
}

// Returns List of Linkers in Room
func (tm *RoomManager) ListLinkers() *md.Linkers {
	return &md.Linkers{
		List: tm.linkers,
	}
}
