package lobby

import (
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Define Function Types
type OnRefreshed func(data []byte)
type OnError func(err error, method string)

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
	Info     pb.LobbyInfo

	// Private Vars
	ctx    context.Context
	call   LobbyCallback
	doneCh chan struct{}
	mutex  sync.Mutex
	ps     *pubsub.PubSub
	topic  *pubsub.Topic
	self   peer.ID
	sub    *pubsub.Subscription
}

// ^ Enter Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Enter(ctx context.Context, callback LobbyCallback, ps *pubsub.PubSub, id peer.ID, olc string) (*Lobby, error) {
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
		Size:  1,
		Peers: make(map[string]*pb.Peer),
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:    ctx,
		call:   callback,
		doneCh: make(chan struct{}, 1),
		ps:     ps,
		topic:  topic,
		sub:    sub,
		self:   id,

		Info:     lobInfo,
		Messages: make(chan *pb.LobbyMessage, ChatRoomBufSize),
	}

	// start reading messages
	go lob.handleMessages()
	go lob.handleEvents()
	return lob, nil
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Send(data []byte) error {
	// Publish to Topic
	err := lob.topic.Publish(lob.ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// ^ Find returns Pointer to Peer.ID and Peer ^
func (lob *Lobby) Find(q string) (peer.ID, *pb.Peer) {
	// Retreive Data
	peer := lob.Peer(q)
	id := lob.ID(q)

	return id, peer
}

// ^ Peers returns ALL Available in Lobby ^
func (lob *Lobby) Info() []byte {
	// Convert to bytes
	data, err := proto.Marshal(&lob.Info)
	if err != nil {
		fmt.Println("Error Marshaling Lobby Data ", err)
		return nil
	}
	return data
}

// ^ End terminates lobby loop ^
func (lob *Lobby) End() {
	lob.doneCh <- struct{}{}
}
