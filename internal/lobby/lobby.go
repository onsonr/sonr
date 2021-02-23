package lobby

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer
const ChatRoomBufSize = 128
const LobbySize = 16

// Lobby represents a subscription to a single PubSub topic. Messages
type Lobby struct {
	// Public Vars
	Messages chan *md.LobbyEvent
	Events   chan *pubsub.PeerEvent
	Data     *md.Lobby

	// Private Vars
	ctx          context.Context
	call         md.LobbyCallback
	doneCh       chan struct{}
	ps           *pubsub.PubSub
	topic        *pubsub.Topic
	topicHandler *pubsub.TopicEventHandler
	self         peer.ID
	sub          *pubsub.Subscription
}

// ^ Join Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Join(ctx context.Context, lobCall md.LobbyCallback, ps *pubsub.PubSub, id peer.ID, sp *md.Peer, olc string) (*Lobby, error) {
	// Join the pubsub Topic
	point := "/sonr/lobby/" + olc
	topic, err := ps.Join(point)
	if err != nil {
		return nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	topicHandler, err := topic.EventHandler()
	if err != nil {
		return nil, err
	}

	// Initialize Lobby for Peers
	lobInfo := &md.Lobby{
		Olc:   point,
		Size:  1,
		Peers: make(map[string]*md.Peer),
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:          ctx,
		call:         lobCall,
		doneCh:       make(chan struct{}, 1),
		ps:           ps,
		topic:        topic,
		topicHandler: topicHandler,
		sub:          sub,
		self:         id,
		Data:         lobInfo,
		Messages:     make(chan *md.LobbyEvent, ChatRoomBufSize),
	}

	// Start Reading Messages
	go lob.handleEvents()
	go lob.handleMessages()
	go lob.processMessages()
	return lob, nil
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Exchange(peerID peer.ID) error {
	// Check if Peer already exchanged
	if p := lob.Peer(peerID.String()); p == nil {
		// Create Lobby Event
		event := md.LobbyEvent{
			Event: md.LobbyEvent_EXCHANGE,
			Data:  lob.call.GetPeer(),
			Id:    lob.call.GetPeer().Id,
		}

		// Convert Event to Proto Binary
		bytes, err := proto.Marshal(&event)
		if err != nil {
			return err
		}

		// Publish to Topic
		err = lob.topic.Publish(lob.ctx, bytes)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Resume() error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_RESUME,
		Data:  lob.call.GetPeer(),
		Id:    lob.call.GetPeer().Id,
	}

	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = lob.topic.Publish(lob.ctx, bytes)
	if err != nil {
		return err
	}
	return nil
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Standby() error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_STANDBY,
		Data:  lob.call.GetPeer(),
		Id:    lob.call.GetPeer().Id,
	}

	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = lob.topic.Publish(lob.ctx, bytes)
	if err != nil {
		return err
	}
	return nil
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Update() error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		Data:  lob.call.GetPeer(),
		Id:    lob.call.GetPeer().Id,
	}

	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = lob.topic.Publish(lob.ctx, bytes)
	if err != nil {
		return err
	}
	return nil
}

// ^ Info returns ALL Lobby Data as Bytes^
func (lob *Lobby) Refresh() {
	lob.call.OnRefresh(lob.Data)
}
