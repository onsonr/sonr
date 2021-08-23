package room

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	sh "github.com/sonr-io/core/internal/host"
	ac "github.com/sonr-io/core/pkg/account"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

type GetRoomFunc func() *data.Room

type RoomHandler interface {
	OnRoomEvent(*data.RoomEvent)
	OnSyncEvent(*data.SyncEvent)
}

type RoomManager struct {
	// General
	ctx          context.Context
	host         sh.HostNode
	Topic        *ps.Topic
	subscription *ps.Subscription
	eventHandler *ps.TopicEventHandler
	account      ac.Account
	handler      RoomHandler

	// Exchange
	exchange   *ExchangeService
	roomEvents chan *data.RoomEvent

	linkers []*data.Peer
	room    *data.Room
}

// NewLocal ^ Create New Contained Room Manager ^ //
func JoinRoom(ctx context.Context, h sh.HostNode, ac ac.Account, room *data.Room, th RoomHandler) (*RoomManager, *data.SonrError) {
	// Join Room
	name := room.GetName()
	topic, err := h.Pubsub().Join(name)
	if err != nil {
		return nil, data.NewError(err, data.ErrorEvent_ROOM_JOIN)
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, data.NewError(err, data.ErrorEvent_ROOM_SUB)
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, data.NewError(err, data.ErrorEvent_ROOM_HANDLER)
	}

	// Create Lobby Manager
	mgr := &RoomManager{
		handler:      th,
		account:      ac,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		room:         room,
		linkers:      make([]*data.Peer, 0),
		roomEvents:   make(chan *data.RoomEvent, util.MAX_CHAN_DATA),
		subscription: sub,
		Topic:        topic,
	}

	// Check Topic type
	if room.IsLocal() {
		// Start Exchange
		err := mgr.initExchange()
		if err != nil {
			return nil, err
		}

		// Return Manager
		return mgr, nil
	} else {
		return nil, data.NewError(errors.New("Invalid Room Type"), data.ErrorEvent_ROOM_JOIN)
	}
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
func (tm *RoomManager) Publish(msg *data.RoomEvent) error {
	if tm.room.IsLocal() || tm.room.IsGroup() {
		// Convert Event to Proto Binary
		bytes, err := proto.Marshal(msg)
		if err != nil {
			data.LogError(err)
			return err
		}

		// Publish to Room
		err = tm.Topic.Publish(tm.ctx, bytes)
		if err != nil {
			data.LogError(err)
			return err
		}
	}
	return nil
}

// Publish @ Publish message to specific peer in room
func (tm *RoomManager) Sync(msg *data.SyncEvent) error {
	if tm.room.IsDevices() {
		// Convert Event to Proto Binary
		bytes, err := proto.Marshal(msg)
		if err != nil {
			data.LogError(err)
			return err
		}

		// Publish to Room
		err = tm.Topic.Publish(tm.ctx, bytes)
		if err != nil {
			data.LogError(err)
			return err
		}
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
func (tm *RoomManager) ListLinkers() *data.Linkers {
	return &data.Linkers{
		List: tm.linkers,
	}
}
