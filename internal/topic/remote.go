package topic

import (
	"context"

	net "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
)

func JoinRemote(ctx context.Context, h *net.HostNode, u *md.User, r *md.RemoteResponse, th ClientCallback) (*TopicManager, *md.SonrError) {
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
	mgr := &TopicManager{
		user:         u,
		topicHandler: th,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		lobby:        md.NewJoinedRemote(u, r),
		messages:     make(chan *md.LobbyEvent, K_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Set Service
	go mgr.handleTopicMessages()
	go mgr.processTopicMessages()
	return mgr, nil
}

// ^ Create New Contained Topic Manager ^ //
func NewRemote(ctx context.Context, h *net.HostNode, u *md.User, l *md.Lobby, th ClientCallback) (*TopicManager, *md.SonrError) {
	// Get Topic Name
	info := l.GetRemote()

	// Join Topic
	topic, sub, handler, serr := h.Join(info.Topic)
	if serr != nil {
		return nil, serr
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		topicHandler: th,
		user:         u,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		lobby:        l,
		messages:     make(chan *md.LobbyEvent, K_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Set Service
	go mgr.handleTopicMessages()
	go mgr.processTopicMessages()
	return mgr, nil
}
