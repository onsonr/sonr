package topic

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	net "github.com/sonr-io/core/internal/host"
	se "github.com/sonr-io/core/internal/session"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

const K_MAX_MESSAGES = 128
const K_SERVICE_PID = protocol.ID("/sonr/topic-service/0.1")

type TopicManager struct {
	ctx         context.Context
	host        *net.HostNode
	topic       *pubsub.Topic
	localTopics []*pubsub.Topic
	isLocal     bool
	Lobby       *md.Lobby

	service      *TopicService
	Messages     chan *md.LobbyEvent
	topicHandler TopicHandler
}

type TopicHandler interface {
	GetContact() *md.Contact
	OnEvent(*md.LobbyEvent)
	OnRefresh(*md.Lobby)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte, session *se.Session)
	OnResponded(inv *md.AuthInvite, p *md.Peer)
}

func JoinTopic(ctx context.Context, h *net.HostNode, p *md.Peer, name string, r *md.Router, lt md.Lobby_Type, th TopicHandler) (*TopicManager, *md.SonrError) {
	// Join Topic
	topic, sub, handler, serr := h.Join(name)
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
		topicHandler: th,
		ctx:          ctx,
		host:         h,
		isLocal:      false,
		Lobby: &md.Lobby{
			Name:  name[12:],
			Size:  1,
			Count: 0,
			Peers: make(map[string]*md.Peer),
			Type:  lt,
			User:  p,
		},
		Messages: make(chan *md.LobbyEvent, K_MAX_MESSAGES),
		topic:    topic,
	}

	// Start Exchange Server
	peersvServer := rpc.NewServer(h.Host, K_SERVICE_PID)
	psv := TopicService{
		lobby:  mgr.Lobby,
		peer:   p,
		call:   th,
		respCh: make(chan *md.AuthReply, 1),
	}

	// Register Service
	err := peersvServer.Register(&psv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_RPC)
	}

	// Set Service
	mgr.service = &psv
	go mgr.handleTopicEvents(p, handler)
	go mgr.handleTopicMessages(p, sub)
	go mgr.processTopicMessages(p)
	return mgr, nil
}

// ^ Create New Contained Topic Manager ^ //
func NewLocalTopic(ctx context.Context, h *net.HostNode, p *md.Peer, r *md.Router, lt md.Lobby_Type, th TopicHandler) (*TopicManager, *md.SonrError) {
	// IP Topic
	ipTopicName := r.LocalIPTopic()

	// Join IP Topic
	ipTopic, ipSub, ipHandler, serr := h.Join(ipTopicName)
	if serr != nil {
		return nil, serr
	}

	// Get Geolocation Topic
	hasGeoLocation := true
	geoTopicName, err := r.LocalGeoTopic()
	if err != nil {
		hasGeoLocation = false
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		topicHandler: th,
		ctx:          ctx,
		host:         h,
		isLocal:      true,
		Lobby: &md.Lobby{
			Name:  ipTopicName[12:],
			Size:  1,
			Count: 0,
			Peers: make(map[string]*md.Peer),
			Type:  lt,
			User:  p,
		},
		localTopics: []*pubsub.Topic{ipTopic},
		Messages:    make(chan *md.LobbyEvent, K_MAX_MESSAGES),
	}

	// Handle Geo Topic
	if hasGeoLocation {
		geoTopic, geoSub, geoHandler, serr := h.Join(geoTopicName)
		if serr != nil {
			return nil, serr
		}

		mgr.localTopics = append(mgr.localTopics, geoTopic)
		go mgr.handleTopicEvents(p, geoHandler)
		go mgr.handleTopicMessages(p, geoSub)
	}

	// Start Exchange Server
	peersvServer := rpc.NewServer(h.Host, K_SERVICE_PID)
	psv := TopicService{
		lobby:  mgr.Lobby,
		peer:   p,
		call:   th,
		respCh: make(chan *md.AuthReply, 1),
	}

	// Register Service
	err = peersvServer.Register(&psv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_RPC)
	}

	// Set Service
	mgr.service = &psv
	go mgr.handleTopicEvents(p, ipHandler)
	go mgr.handleTopicMessages(p, ipSub)
	go mgr.processTopicMessages(p)
	return mgr, nil
}

// ^ Create New Contained Topic Manager ^ //
func NewTopic(ctx context.Context, h *net.HostNode, p *md.Peer, name string, r *md.Router, lt md.Lobby_Type, th TopicHandler) (*TopicManager, *md.SonrError) {
	// Join Topic
	topic, sub, handler, serr := h.Join(name)
	if serr != nil {
		return nil, serr
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		topicHandler: th,
		isLocal:      false,
		ctx:          ctx,
		host:         h,
		Lobby: &md.Lobby{
			Name:  name[12:],
			Size:  1,
			Count: 0,
			Peers: make(map[string]*md.Peer),
			Type:  lt,
			User:  p,
		},
		Messages: make(chan *md.LobbyEvent, K_MAX_MESSAGES),
		topic:    topic,
	}

	// Start Exchange Server
	peersvServer := rpc.NewServer(h.Host, K_SERVICE_PID)
	psv := TopicService{
		lobby:  mgr.Lobby,
		peer:   p,
		call:   th,
		respCh: make(chan *md.AuthReply, 1),
	}

	// Register Service
	err := peersvServer.Register(&psv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_RPC)
	}

	// Set Service
	mgr.service = &psv
	go mgr.handleTopicEvents(p, handler)
	go mgr.handleTopicMessages(p, sub)
	go mgr.processTopicMessages(p)
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
	if tm.isLocal {
		// Iterate IP Topic
		for _, id := range tm.localTopics[0].ListPeers() {
			// If Found Match
			if id.String() == q {
				i = id
			}
		}
		// Check for Geo Topic
		if len(tm.localTopics) > 1 {
			for _, id := range tm.localTopics[1].ListPeers() {
				// If Found Match
				if id.String() == q {
					i = id
				}
			}
		}
	} else {
		for _, id := range tm.topic.ListPeers() {
			// If Found Match
			if id.String() == q {
				i = id
			}
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
	// Iterate through Topic Peers
	if tm.isLocal {
		// Iterate IP Topic
		for _, id := range tm.localTopics[0].ListPeers() {
			// If Found Match
			if id.String() == q {
				return true
			}
		}
		// Check for Geo Topic
		if len(tm.localTopics) > 1 {
			for _, id := range tm.localTopics[1].ListPeers() {
				// If Found Match
				if id.String() == q {
					return true
				}
			}
		}
	} else {
		for _, id := range tm.topic.ListPeers() {
			// If Found Match
			if id.String() == q {
				return true
			}
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

	// Iterate through Topic Peers
	if tm.isLocal {
		// Publish to Topic
		err = tm.localTopics[0].Publish(tm.ctx, bytes)
		if err != nil {
			return err
		}
		// Check for Geo Topic
		if len(tm.localTopics) > 1 {
			// Publish to Topic
			err = tm.localTopics[1].Publish(tm.ctx, bytes)
			if err != nil {
				return err
			}
		}
	} else {
		// Publish to Topic
		err = tm.topic.Publish(tm.ctx, bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

// ^ Leave Current Topic ^
func (tm *TopicManager) LeaveTopic() error {
	if tm.isLocal {
		// Check for Geo Topic
		if len(tm.localTopics) > 1 {
			tm.localTopics[1].Close()
			return tm.localTopics[0].Close()
		} else {
			return tm.localTopics[0].Close()
		}
	}
	return tm.topic.Close()
}
