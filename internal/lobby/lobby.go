package lobby

import (
	"context"
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/internal/network"
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

	// Networking
	ctx    context.Context
	call   md.LobbyCallback
	host   host.Host
	pubSub *pubsub.PubSub

	// Connection
	router       *net.ProtocolRouter
	topic        *pubsub.Topic
	topicHandler *pubsub.TopicEventHandler
	selfPeer     *md.Peer
	peersv       *ExchangeService
	sub          *pubsub.Subscription
}

// ^ Join Joins/Subscribes to pubsub topic and returns Lobby ^
func Join(ctx context.Context, lobCall md.LobbyCallback, h host.Host, ps *pubsub.PubSub, sp *md.Peer, pr *net.ProtocolRouter) (*Lobby, error) {
	// Join the pubsub Topic
	topic, err := ps.Join(pr.Topic(net.SetIDForLocal()))
	if err != nil {
		return nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Top Handler
	topicHandler, err := topic.EventHandler()
	if err != nil {
		return nil, err
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:          ctx,
		call:         lobCall,
		pubSub:       ps,
		host:         h,
		router:       pr,
		topic:        topic,
		topicHandler: topicHandler,
		sub:          sub,
		selfPeer:     sp,

		messages: make(chan *md.LobbyEvent, ChatRoomBufSize),
		data: &md.Lobby{
			Olc:    pr.OLC,
			Size:   1,
			Peers:  make(map[string]*md.Peer),
			Groups: make(map[string]*md.Group),
		},
	}

	// Create PeerService
	peersvServer := gorpc.NewServer(h, pr.Exchange())
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

// ^ Helper: ID returns ONE Peer.ID in PubSub ^
func (lob *Lobby) HasPeer(q string) bool {
	// Iterate through PubSub in topic
	for _, id := range lob.pubSub.ListPeers(lob.router.Topic(net.SetIDForLocal())) {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// ^ Helper: ID returns ONE Peer.ID in PubSub ^
func (lob *Lobby) ID(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range lob.pubSub.ListPeers(lob.router.Topic(net.SetIDForLocal())) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	return ""
}

// ^ Helper: Peer returns ONE Peer in Lobby ^
func (lob *Lobby) Peer(q string) *md.Peer {
	// Iterate Through Peers, Return Matched Peer
	for _, peer := range lob.data.Peers {
		// If Found Match
		if peer.Id.Peer == q {
			return peer
		}
	}
	return nil
}

// ^ Helper: Find returns Pointer to Peer.ID and Peer ^
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
func (lob *Lobby) Resume() error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_RESUME,
		From:  lob.call.Peer(),
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
		From:  lob.call.Peer(),
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

// ^ Send updates lobby ^
func (lob *Lobby) Update() error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		From:  lob.call.Peer(),
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

// ^ Send publishes a message to specific peer in lobby ^
func (lob *Lobby) Message(msg string, to string) error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event:   md.LobbyEvent_MESSAGE,
		From:    lob.call.Peer(),
		Id:      lob.call.Peer().Id.Peer,
		Message: msg,
		To:      to,
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
