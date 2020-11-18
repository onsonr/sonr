package lobby

import (
	"context"

	badger "github.com/dgraph-io/badger/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pb "github.com/sonr-io/core/pkg/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// LobbyCallback returns message from lobby
type LobbyCallback interface {
	OnRefreshed(s string)
}

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Public Vars
	Messages chan *Notification
	Code     string
	Self     *pb.PeerInfo

	// Private Vars
	ctx      context.Context
	callback LobbyCallback
	doneCh   chan struct{}
	peerDB   *badger.DB
	ps       *pubsub.PubSub
	topic    *pubsub.Topic
	sub      *pubsub.Subscription
}

// Enter Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby
func Enter(ctx context.Context, call LobbyCallback, ps *pubsub.PubSub, p *pb.PeerInfo, olcCode string) (*Lobby, error) {
	// Join the pubsub Topic
	topic, err := ps.Join(olcCode)
	if err != nil {
		return nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Initialize Badger DB
	opt := badger.DefaultOptions("").WithInMemory(true)
	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:      ctx,
		callback: call,
		doneCh:   make(chan struct{}, 1),
		peerDB:   db,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		Self:     p,
		Code:     olcCode,
		Messages: make(chan *Notification, ChatRoomBufSize),
	}

	// Publish Join Message
	msg := Notification{
		Event:  "Update",
		Peer:   p,
		Data:   p.String(),
		Sender: p.GetId(),
	}

	// start reading messages
	go lob.handleMessages()
	go lob.handleEvents()
	lob.Publish(msg)
	return lob, nil
}

// Publish sends a message to the pubsub topic.
func (lob *Lobby) Publish(m Notification) error {
	// Publish to Topic
	err := lob.topic.Publish(lob.ctx, m.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// End terminates lobby loop
func (lob *Lobby) End() {
	lob.peerDB.Close()
	lob.doneCh <- struct{}{}
}
