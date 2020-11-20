package lobby

import (
	"context"
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// LobbyCallback returns message from lobby
type LobbyCallback interface {
	OnEvent(data []byte)
	OnError(data []byte)
}

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Public Vars
	Messages chan *pb.LobbyMessage
	Code     string
	Self     *pb.PeerInfo

	// Private Vars
	ctx    context.Context
	call   LobbyCallback
	doneCh chan struct{}
	peerDB *badger.DB
	ps     *pubsub.PubSub
	topic  *pubsub.Topic
	sub    *pubsub.Subscription
}

// ^ Sends generic protobuf with subject ^
func (lob *Lobby) Callback(event pb.Callback_Event, providedData []byte) {
	// Create Callback Protobuf
	callback := &pb.Callback{
		On:   event,
		Data: providedData,
	}

	// Convert to bytes
	raw, err := proto.Marshal(callback)
	if err != nil {
		fmt.Println(err)
	}

	// Send Generic callback
	lob.call.OnEvent(raw)
}

// ^ Enter Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Enter(ctx context.Context, callback LobbyCallback, ps *pubsub.PubSub, joinEvent *pb.JoinEvent) (*Lobby, error) {
	// Join the pubsub Topic
	topic, err := ps.Join(joinEvent.Olc)
	if err != nil {
		return nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Initialize Datastore for Peers
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		return nil, err
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:      ctx,
		call:     callback,
		doneCh:   make(chan struct{}, 1),
		peerDB:   db,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		Self:     joinEvent.Peer,
		Code:     joinEvent.Olc,
		Messages: make(chan *pb.LobbyMessage, ChatRoomBufSize),
	}

	// Publish Join Message
	msg := &pb.LobbyMessage{
		Event:  "Join",
		Data:   joinEvent.Peer,
		Sender: joinEvent.Peer.GetId(),
	}

	// Send Join Message
	lob.Publish(msg)

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
	lob.peerDB.Close()
	lob.doneCh <- struct{}{}
}
