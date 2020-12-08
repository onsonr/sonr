package lobby

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Define Function Types
type Refreshed func(call md.CallbackType, data proto.Message)
type Error func(err error, method string)

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Public Vars
	Messages chan *md.LobbyEvent
	Events   chan *pubsub.PeerEvent
	Data     *md.Lobby

	// Private Vars
	ctx          context.Context
	refresh      Refreshed
	onError      Error
	doneCh       chan struct{}
	ps           *pubsub.PubSub
	topic        *pubsub.Topic
	topicHandler *pubsub.TopicEventHandler
	self         peer.ID
	sub          *pubsub.Subscription
}

// ^ Initialize Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Initialize(ctx context.Context, callr Refreshed, onErr Error, ps *pubsub.PubSub, id peer.ID, olc string) (*Lobby, error) {
	// Join the pubsub Topic
	topic, err := ps.Join(olc)
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
		Code:  olc,
		Size:  1,
		Peers: make(map[string]*md.Peer),
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:          ctx,
		onError:      onErr,
		refresh:      callr,
		doneCh:       make(chan struct{}, 1),
		ps:           ps,
		topic:        topic,
		topicHandler: topicHandler,
		sub:          sub,
		self:         id,

		Data:     lobInfo,
		Messages: make(chan *md.LobbyEvent, ChatRoomBufSize),
	}

	// start reading messages
	go lob.handleMessages()
	go lob.handleEvents()
	return lob, nil
}

// ^ Info returns ALL Lobby Data as Bytes^
func (lob *Lobby) Info() []byte {
	// Convert to bytes
	data, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Error Marshaling Lobby Data ", err)
		return nil
	}
	return data
}

// ^ Find returns Pointer to Peer.ID and Peer ^
func (lob *Lobby) Find(q string) (peer.ID, *md.Peer) {
	// Retreive Data
	peer := lob.Peer(q)
	id := lob.ID(q)

	return id, peer
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Update(p *md.Peer) error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		Peer:  p,
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

// ^ End terminates lobby loop ^
func (lob *Lobby) Exit() {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_EXIT,
		Id:    lob.self.String(),
	}

	// Convert Event to Proto Binary, Suppress Error
	bytes, _ := proto.Marshal(&event)

	// Publish to Topic
	err := lob.topic.Publish(lob.ctx, bytes)
	if err != nil {
		lob.onError(err, "Exit")
	}
	lob.doneCh <- struct{}{}
}
