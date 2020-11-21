package lobby

import (
	"context"
	"log"
	"sync"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Define Function Types
type OnRefreshed func(data []byte)
type OnError func(data []byte)

// Struct to Implement Node Callback Methods
type LobbyCallback struct {
	Refreshed OnRefreshed
	Error     OnError
}

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Public Vars
	Messages chan *pb.LobbyMessage
	Self     *pb.Peer
	Info     pb.Lobby

	// Private Vars
	ctx    context.Context
	call   LobbyCallback
	doneCh chan struct{}
	mutex  sync.Mutex
	ps     *pubsub.PubSub
	topic  *pubsub.Topic
	sub    *pubsub.Subscription
}

// ^ Enter Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Enter(ctx context.Context, callback LobbyCallback, ps *pubsub.PubSub, peer *pb.Peer, olc string) (*Lobby, error) {
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
	lobInfo := pb.Lobby{
		Code:  olc,
		Count: 1,
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:    ctx,
		call:   callback,
		doneCh: make(chan struct{}, 1),
		ps:     ps,
		topic:  topic,
		sub:    sub,

		Self:     peer,
		Info:     lobInfo,
		Messages: make(chan *pb.LobbyMessage, ChatRoomBufSize),
	}

	// start reading messages
	go lob.handleMessages()
	go lob.handleEvents()
	return lob, nil
}

// Publish sends a message to the pubsub topic.
func (lob *Lobby) Publish(m *pb.LobbyMessage) error {
	// Convert Request to Proto Binary
	data, err := proto.Marshal(m)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// Publish to Topic
	err = lob.topic.Publish(lob.ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// End terminates lobby loop
func (lob *Lobby) End() {
	lob.doneCh <- struct{}{}
}
