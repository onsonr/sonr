package topic

import (
	"context"
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/internal/network"
	dt "github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

const ChatRoomBufSize = 128

type ContinueTransfer func(id peer.ID, p *md.Peer, cf *sf.ProcessedFile, data []byte)

type TopicManager struct {
	ctx          context.Context
	host         host.Host
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	handler      *pubsub.TopicEventHandler
	Lobby        *Lobby

	topicPoint   string
	service      *TopicService
	exchProtocol protocol.ID
	authProtocol protocol.ID
	Messages     chan *md.LobbyEvent
	call         dt.NodeCallback
	returnPeer   dt.ReturnPeer
	cont         ContinueTransfer
}

type TopicHandler interface {
	Refresh() *md.Lobby
	GotInvite() *md.AuthInvite
	GotReply() *md.AuthReply
	ReturnPeer() *md.Peer
}

// ^ Create New Contained Topic Manager ^ //
func NewTopic(h host.Host, ps *pubsub.PubSub, name string, router *network.ProtocolRouter, call dt.NodeCallback, cont ContinueTransfer) (*TopicManager, error) {
	// Join Topic
	topic, err := ps.Join(name)
	if err != nil {
		return nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Topic Handler
	handler, err := topic.EventHandler()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		call:         call,
		cont:         cont,
		host:         h,
		handler:      handler,
		Lobby:        NewLobby(name, call.GetPeer(), call.Refreshed),
		Messages:     make(chan *md.LobbyEvent, ChatRoomBufSize),
		authProtocol: router.Auth(),
		exchProtocol: router.TopicExchange(),
		subscription: sub,
		topic:        topic,
		topicPoint:   name,
		returnPeer:   call.GetPeer,
	}

	// Start Exchange Server
	peersvServer := rpc.NewServer(h, router.TopicExchange())
	psv := TopicService{
		SyncLobby: mgr.Lobby.Sync,
		GetUser:   call.GetPeer,
	}

	// Register Service
	err = peersvServer.Register(&psv)
	if err != nil {
		return nil, err
	}

	// Set Service
	mgr.service = &psv

	go mgr.handleTopicEvents()
	go mgr.handleTopicMessages()
	return mgr, nil
}

// ^ Helper: Find returns Pointer to Peer.ID and Peer ^
func (tm *TopicManager) FindPeerInTopic(q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	var p *md.Peer
	var i peer.ID

	// Iterate Through Peers, Return Matched Peer
	for _, peer := range tm.Lobby.Peers {
		// If Found Match
		if peer.Id.Peer == q {
			p = peer
		}
	}

	// Validate Peer
	if p == nil {
		return "", nil, errors.New("Peer data was not found in topic.")
	}

	// Iterate through Topic Peers
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			i = id
		}
	}

	// Validate ID
	if i == "" {
		return "", nil, errors.New("Peer ID was not found in topic.")
	}
	return i, p, nil
}

// ^ Helper: ID returns ONE Peer.ID in Topic ^
func (tm *TopicManager) HasPeer(q string) bool {
	// Iterate through PubSub in topic
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// ^ Send message to specific peer in topic ^
func (tm *TopicManager) Send(msg *md.LobbyEvent) error {
	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = tm.topic.Publish(tm.ctx, bytes)
	if err != nil {
		return err
	}
	return nil
}

// ^ handleTopicEvents: listens to Pubsub Events for topic  ^
func (tm *TopicManager) handleTopicEvents() {
	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := tm.handler.NextPeerEvent(tm.ctx)
		if err != nil {
			tm.handler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			p := tm.returnPeer()
			buf, err := proto.Marshal(p)
			if err != nil {
				continue
			}
			tm.Exchange(lobEvent.Peer, buf)
		}

		if lobEvent.Type == pubsub.PeerLeave {
			tm.Lobby.Remove(lobEvent.Peer)
		}

		dt.GetState().NeedsWait()
	}
}

// ^ handleTopicMessages: listens for messages on pubsub topic subscription ^
func (tm *TopicManager) handleTopicMessages() {
	for {
		// Get next msg from pub/sub
		msg, err := tm.subscription.Next(tm.ctx)
		if err != nil {
			return
		}

		// Only forward messages delivered by others
		if msg.ReceivedFrom.String() == tm.returnPeer().Id.Peer {
			continue
		}

		// Construct message
		m := &md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if tm.HasPeer(m.Id) {
			tm.Messages <- m
		}
		dt.GetState().NeedsWait()
	}
}
