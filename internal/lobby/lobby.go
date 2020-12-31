package lobby

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Define Function Types
type OnProtobuf func([]byte)
type OnPeerJoin func() *md.Peer
type Error func(err error, method string)

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Public Vars
	Messages chan *md.LobbyMessage
	Events   chan *pubsub.PeerEvent
	Data     *md.Lobby

	// Private Vars
	ctx         context.Context
	pushInfo    OnPeerJoin
	callEvent   OnProtobuf
	callRefresh OnProtobuf
	onError     Error
	ps          *pubsub.PubSub
	topic       *pubsub.Topic
	self        peer.ID
	sub         *pubsub.Subscription
}

// ^ Join Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Join(ctx context.Context, calle OnProtobuf, callr OnProtobuf, push OnPeerJoin, onErr Error, ps *pubsub.PubSub, id peer.ID, olc string) (*Lobby, error) {
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

	// Initialize Lobby for Peers
	lobInfo := &md.Lobby{
		Code:  olc,
		Size:  1,
		Peers: make(map[string]*md.Peer),
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:         ctx,
		onError:     onErr,
		pushInfo:    push,
		callEvent:   calle,
		callRefresh: callr,
		ps:          ps,
		topic:       topic,
		sub:         sub,
		self:        id,

		Data:     lobInfo,
		Messages: make(chan *md.LobbyMessage, ChatRoomBufSize),
	}

	// Start Reading Messages / Events
	go lob.handleEvents()
	go lob.handleMessages()
	go lob.processMessages()
	return lob, nil
}

// ^ Exchange publishes a message with Current Info to Peer that Exchanged ^
func (lob *Lobby) Exchange(msg *md.LobbyMessage) error {
	// @ Check if Exchanged, Push Data if not
	if _, ok := lob.Data.Peers[msg.Id]; ok {
		return nil
	} else {
		// Set the Peer
		lob.setPeer(msg)

		// Get Peer info
		p := lob.pushInfo()

		// Create Lobby Event
		event := md.LobbyMessage{
			Event:     md.LobbyMessage_EXCHANGE,
			Peer:      p,
			Direction: p.Direction,
			Id:        p.Id,
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
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Update() error {
	// Get Peer info
	p := lob.pushInfo()

	// Create Lobby Event
	event := md.LobbyMessage{
		Event:     md.LobbyMessage_AVAILABLE,
		Peer:      p,
		Direction: p.Direction,
		Id:        p.Id,
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
func (lob *Lobby) Busy() error {
	// Get Peer info
	p := lob.pushInfo()

	// Create Lobby Event
	event := md.LobbyMessage{
		Event: md.LobbyMessage_BUSY,
		Peer:  p,
		Id:    p.Id,
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
