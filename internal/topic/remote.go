package topic

import (
	"context"

	net "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
)

func JoinRemote(ctx context.Context, h *net.HostNode, u *md.User, r *md.RemoteJoinRequest, th ClientCallback) (*TopicManager, *md.RemoteJoinResponse, *md.SonrError) {
	// Join Topic
	topic, sub, handler, serr := h.Join(r.GetTopic())
	if serr != nil {
		return nil, &md.RemoteJoinResponse{
			Status: md.RemoteJoinResponse_None,
		}, serr
	}

	// Check Peers
	peers := topic.ListPeers()
	if len(peers) == 0 {
		handler.Cancel()
		sub.Cancel()
		topic.Close()
		return nil, &md.RemoteJoinResponse{
			Status: md.RemoteJoinResponse_Invalid,
		}, md.NewErrorWithType(md.ErrorMessage_TOPIC_INVALID)
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		user:         u,
		topicHandler: th,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		lobby:        r.NewJoinedRemote(u),
		messages:     make(chan *md.LocalEvent, K_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Set Service
	go mgr.handleTopicMessages()
	go mgr.processTopicMessages()
	return mgr, &md.RemoteJoinResponse{
		Status: md.RemoteJoinResponse_Pending,
		Lobby:  mgr.lobby,
	}, nil
}

// ^ Create New Contained Topic Manager ^ //
func NewRemote(ctx context.Context, h *net.HostNode, u *md.User, r *md.RemoteCreateRequest, th ClientCallback) (*TopicManager, *md.RemoteCreateResponse, *md.SonrError) {
	// Join Topic
	topic, sub, handler, serr := h.Join(r.GetTopic())
	if serr != nil {
		return nil, nil, serr
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		topicHandler: th,
		user:         u,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		lobby:        r.NewCreatedRemote(u),
		messages:     make(chan *md.LocalEvent, K_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Set Service
	go mgr.handleTopicMessages()
	go mgr.processTopicMessages()
	return mgr, &md.RemoteCreateResponse{
		Success: true,
		Topic:  mgr.lobby.Topic(),
	}, nil
}
