package topic

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/internal/network"
	"google.golang.org/protobuf/proto"
)

const K_MAX_MESSAGES = 128

type TopicManager struct {
	ctx          context.Context
	host         host.Host
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	handler      *pubsub.TopicEventHandler
	Lobby        *Lobby

	topicPoint string
	service    *TopicService
	protocol   protocol.ID
	Messages   chan *md.LobbyEvent
	callback   TopicHandler
}

type TopicHandler interface {
	GetPeer() *md.Peer
	OnEvent(*md.LobbyEvent)
	OnRefresh(*md.Lobby)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnReceiveTransfer(inv *md.AuthInvite, fs *sf.FileSystem)
}

// ^ Create New Contained Topic Manager ^ //
func NewTopic(ctx context.Context, h host.Host, ps *pubsub.PubSub, name string, router *network.ProtocolRouter, th TopicHandler) (*TopicManager, error) {
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
		return nil, err
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		callback: th,
		ctx:      ctx,
		host:     h,
		handler:  handler,
		Lobby: &Lobby{
			callback: th,
			OLC:      name,
			Size:     1,
			Count:    0,
			Peers:    make(map[string]*md.Peer),
			user:     th.GetPeer(),
		},
		Messages:     make(chan *md.LobbyEvent, K_MAX_MESSAGES),
		protocol:     router.TopicService(),
		subscription: sub,
		topic:        topic,
		topicPoint:   name,
	}

	// Start Exchange Server
	peersvServer := rpc.NewServer(h, router.TopicService())
	psv := TopicService{
		SyncLobby: mgr.Lobby.Sync,
		GetUser:   th.GetPeer,
		call:      th,
		respCh:    make(chan *md.AuthReply, 1),
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
	go mgr.processTopicMessages()
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

// ^ Leave Current Topic ^
func (tm *TopicManager) LeaveTopic() error {
	return tm.topic.Close()
}
