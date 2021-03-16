package lobby

import (
	"context"
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128
const LobbySize = 16

// Lobby represents a subscription to a single PubSub topic
type Lobby struct {
	// Public Vars
	messages chan *md.LobbyEvent
	data     *md.Lobby

	// Private Vars
	ctx          context.Context
	call         md.LobbyCallback
	host         host.Host
	doneCh       chan struct{}
	pubSub       *pubsub.PubSub
	topic        *pubsub.Topic
	topicHandler *pubsub.TopicEventHandler
	selfPeer     *md.Peer
	peersv       *ExchangeService
	sub          *pubsub.Subscription
}

// ^ Join Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Join(ctx context.Context, lobCall md.LobbyCallback, h host.Host, ps *pubsub.PubSub, sp *md.Peer, olc string) (*Lobby, error) {
	// Join the pubsub Topic
	lobbyName := "/sonr/lobby/" + olc
	topic, err := ps.Join(lobbyName)
	if err != nil {
		return nil, err
	}

	// TODO: Remove after Validation, Log Peers
	for _, id := range ps.ListPeers(lobbyName) {
		if id != h.ID() {
			sentry.CaptureMessage(fmt.Sprintf("Peer in Lobby %s", id))
		}
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
		Olc:   lobbyName,
		Size:  1,
		Peers: make(map[string]*md.Peer),
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:          ctx,
		call:         lobCall,
		doneCh:       make(chan struct{}, 1),
		pubSub:       ps,
		host:         h,
		topic:        topic,
		topicHandler: topicHandler,
		sub:          sub,
		selfPeer:     sp,
		data:         lobInfo,
		messages:     make(chan *md.LobbyEvent, ChatRoomBufSize),
	}

	// Create PeerService
	peersvServer := gorpc.NewServer(h, protocol.ID("/sonr/lobby/exchange"))
	psv := ExchangeService{
		updatePeer: lob.updatePeer,
		getUser:    lob.call.Peer,
	}

	// Register Service
	err = peersvServer.Register(&psv)
	if err != nil {
		return nil, err
	}

	// Set Service
	lob.peersv = &psv

	// Start Reading Messages
	go lob.handleEvents()
	go lob.handleMessages()
	go lob.processMessages()
	return lob, nil
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Resume() error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_RESUME,
		Data:  lob.call.Peer(),
		Id:    lob.call.Peer().Id.Peer,
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
		Data:  lob.call.Peer(),
		Id:    lob.call.Peer().Id.Peer,
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
		Data:  lob.call.Peer(),
		Id:    lob.call.Peer().Id.Peer,
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

// ^ Send Updated Lobby ^
func (lob *Lobby) Refresh() {
	bytes, err := proto.Marshal(lob.data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	lob.call.Refresh(bytes)
}
