package lobby

import (
	"context"
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	lf "github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128
const LobbySize = 16

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
	callback     lf.OnProtobuf
	onError      lf.OnError
	doneCh       chan struct{}
	ps           *pubsub.PubSub
	getPeer      lf.GetUserPeer
	topic        *pubsub.Topic
	topicHandler *pubsub.TopicEventHandler
	self         peer.ID
	selfPeer     *md.Peer
	sub          *pubsub.Subscription
}

// ^ Join Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Join(ctx context.Context, lobCall lf.LobbyCallbacks, ps *pubsub.PubSub, id peer.ID, sp *md.Peer, olc string) (*Lobby, error) {
	// Join the pubsub Topic
	point := "/sonr/lobby" + olc
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
		Code:  point,
		Size:  1,
		Peers: make(map[string]*md.Peer),
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:          ctx,
		onError:      lobCall.CallError,
		callback:     lobCall.CallRefresh,
		doneCh:       make(chan struct{}, 1),
		ps:           ps,
		getPeer:      lobCall.GetPeer,
		topic:        topic,
		topicHandler: topicHandler,
		sub:          sub,
		self:         id,
		selfPeer:     sp,
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
			Peer:  lob.getPeer(),
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
func (lob *Lobby) Find(q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	peer := lob.Peer(q)
	id := lob.ID(q)

	if peer == nil || id == "" {
		return "", nil, errors.New("Search Error, peer was not found in map.")
	}

	return id, peer, nil
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Update() error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		Peer:  lob.getPeer(),
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
