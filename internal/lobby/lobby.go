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
			Olc:    pr.LocalPoint(),
			Size:   1,
			Count:  0,
			Peers:  make(map[string]*md.Peer),
			Groups: make(map[string]*md.Group),
		},
	}

	// Create PeerService
	peersvServer := gorpc.NewServer(h, pr.Exchange())
	psv := ExchangeService{
		syncLobby: lob.syncLobby,
		getUser:   lob.call.Peer,
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

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleEvents() {
	// @ Create Topic Handler
	topicHandler, err := lob.topic.EventHandler()
	if err != nil {
		log.Println(err)
		return
	}

	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := topicHandler.NextPeerEvent(lob.ctx)
		if err != nil {
			topicHandler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			lob.Exchange(lobEvent.Peer)
		}

		if lobEvent.Type == pubsub.PeerLeave {
			lob.removePeer(lobEvent.Peer)
		}

		md.GetState().NeedsWait()
	}
}

// ^ 1. handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleMessages() {
	for {
		// Get next msg from pub/sub
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.messages)
			return
		}

		// Only forward messages delivered by others
		if msg.ReceivedFrom == lob.host.ID() {
			continue
		}

		// Construct message
		m := md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, &m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if lob.HasPeer(m.Id) {
			// Update Circle by event
			lob.messages <- &m
		}
		md.GetState().NeedsWait()
	}
}

// ^ 1a. processMessages handles message content and ticker ^
func (lob *Lobby) processMessages() {
	for {
		select {
		// @ when we receive a message from the lobby room
		case m := <-lob.messages:
			// Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				lob.updatePeer(m.From)
			} else if m.Event == md.LobbyEvent_MESSAGE {
				// Check is Message For Self
				if m.To == lob.selfPeer.Id.Peer {
					// Convert Message
					bytes, err := proto.Marshal(m)
					if err != nil {
						log.Println("Cannot Marshal Error Protobuf: ", err)
					}

					// Call Event
					lob.call.Event(bytes)
				}
			}

		case <-lob.ctx.Done():
			return
		}
		md.GetState().NeedsWait()
	}
}

// ^ updatePeer changes Peer values in Lobby ^
func (lob *Lobby) getData() []byte {
	bytes, err := proto.Marshal(lob.data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return bytes
}

// ^ removePeer removes Peer from Map ^
func (lob *Lobby) removePeer(id peer.ID) {
	// Update Peer with new data
	delete(lob.data.Peers, id.String())
	lob.data.Count = int32(len(lob.data.Peers))
	lob.data.Size = int32(len(lob.data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}

// ^ removePeer removes Peer from Map ^
func (lob *Lobby) syncLobby(ref *md.Lobby, peer *md.Peer) {
	// Validate Lobbies are Different
	if lob.data.Count != ref.Count {
		// Iterate Over List
		for id, peer := range ref.Peers {
			// Add all Peers NOT User
			if id != lob.selfPeer.Id.Peer {
				lob.data.Peers[id] = peer
			}
		}
	}

	// Add Peer to Lobby
	lob.updatePeer(peer)

	// Callback with Updated Data
	lob.Refresh()
}

// ^ updatePeer changes Peer values in Lobby ^
func (lob *Lobby) updatePeer(peer *md.Peer) {
	// Update Peer with new data
	lob.data.Peers[peer.Id.Peer] = peer
	lob.data.Count = int32(len(lob.data.Peers))
	lob.data.Size = int32(len(lob.data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}
