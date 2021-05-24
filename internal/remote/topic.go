package remote

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	net "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

const K_MAX_MESSAGES = 128
const K_SERVICE_PID = protocol.ID("/sonr/topic-service/0.1")

type ClientCallback interface {
	OnEvent(*md.LobbyEvent)
	OnRefresh(*md.Lobby)
	OnInvite([]byte)
	OnLink([]byte)
	OnReply(id peer.ID, data []byte)
	OnResponded(inv *md.AuthInvite)
}

type RemoteManager struct {
	ctx          context.Context
	host         *net.HostNode
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	eventHandler *pubsub.TopicEventHandler
	user         *md.User
	Lobby        *md.Lobby

	service      *TopicService
	Messages     chan *md.LobbyEvent
	topicHandler ClientCallback
}

func JoinRemote(ctx context.Context, h *net.HostNode, u *md.User, r *md.RemoteResponse, th ClientCallback) (*RemoteManager, *md.SonrError) {
	// Join Topic
	topic, sub, handler, serr := h.Join(r.Topic)
	if serr != nil {
		return nil, serr
	}

	// Check Peers
	peers := topic.ListPeers()
	if len(peers) == 0 {
		handler.Cancel()
		sub.Cancel()
		topic.Close()
		return nil, md.NewErrorWithType(md.ErrorMessage_TOPIC_INVALID)
	}

	// Create Lobby Manager
	mgr := &RemoteManager{
		user:         u,
		topicHandler: th,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		Lobby:        md.NewJoinedRemote(u, r),
		Messages:     make(chan *md.LobbyEvent, K_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Set Service
	go mgr.handleTopicMessages()
	go mgr.processTopicMessages()
	return mgr, nil
}

// ^ Create New Contained Topic Manager ^ //
func NewRemote(ctx context.Context, h *net.HostNode, u *md.User, l *md.Lobby, th ClientCallback) (*RemoteManager, *md.SonrError) {
	// Get Topic Name
	info := l.GetRemote()

	// Join Topic
	topic, sub, handler, serr := h.Join(info.Topic)
	if serr != nil {
		return nil, serr
	}

	// Create Lobby Manager
	mgr := &RemoteManager{
		topicHandler: th,
		user:         u,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		Lobby:        l,
		Messages:     make(chan *md.LobbyEvent, K_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Set Service
	go mgr.handleTopicMessages()
	go mgr.processTopicMessages()
	return mgr, nil
}

// ^ Helper: Find returns Pointer to Peer.ID and Peer ^
func (tm *RemoteManager) FindPeerInTopic(q string) (peer.ID, *md.Peer, error) {
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
func (tm *RemoteManager) HasPeer(q string) bool {
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
func (tm *RemoteManager) Send(msg *md.LobbyEvent) error {
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
func (tm *RemoteManager) LeaveTopic() error {
	tm.eventHandler.Cancel()
	tm.subscription.Cancel()
	return tm.topic.Close()
}
