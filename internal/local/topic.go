package local

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

const K_MAX_MESSAGES = 256
const K_SERVICE_PID = protocol.ID("/sonr/local-service/0.2")
const K_RPC_SERVICE = "LocalService"

type ClientCallback interface {
	OnEvent(*md.LobbyEvent)
	OnRefresh(*md.Lobby)
	OnInvite([]byte)
	OnLink([]byte)
	OnReply(id peer.ID, data []byte)
	OnResponded(inv *md.AuthInvite)
}

type LocalManager struct {
	ctx      context.Context
	host     *net.HostNode
	user     *md.User
	geoTopic *pubsub.Topic
	ipTopic  *pubsub.Topic
	lobby    *md.SyncLobby
	service  *TopicService
	messages chan *md.LobbyEvent
	callback ClientCallback
}

// ^ Create New Contained Topic Manager ^ //
func NewLocal(ctx context.Context, h *net.HostNode, u *md.User, th ClientCallback) (*LocalManager, *md.SonrError) {
	// Create Local Manager
	mgr := &LocalManager{
		callback: th,
		user:     u,
		ctx:      ctx,
		host:     h,
		lobby:    md.NewLocalLobby(u),
		messages: make(chan *md.LobbyEvent, K_MAX_MESSAGES),
	}

	// Start Exchange Server
	err := mgr.registerService()
	if err != nil {
		return nil, err
	}

	// Start IP Topic
	ipSub, ipHandler, err := mgr.joinIPTopic()
	if err != nil {
		return nil, err
	}

	// Handle IP Topic
	go mgr.handleTopicEvents(ipHandler)
	go mgr.handleTopicMessages(ipSub)

	// Check Geo Topic
	if u.HasGeo() {
		// Join Geo Topic
		geoSub, geoHandler, err := mgr.joinGeoTopic()
		if err == nil {
			// Handle Geo Topic
			go mgr.handleTopicEvents(geoHandler)
			go mgr.handleTopicMessages(geoSub)
		}

	}

	// Process Messages Return Manager
	go mgr.processTopicMessages()
	return mgr, nil
}

// ^ Create New Contained Topic Manager ^ //
func NewLocalLink(ctx context.Context, h *net.HostNode, l *md.Linker, name string, th ClientCallback) (*LocalManager, *md.SonrError) {
	// Create Lobby Manager
	mgr := &LocalManager{
		callback: th,
		ctx:      ctx,
		host:     h,
		lobby:    md.NewLocalLinkLobby(l),
		messages: make(chan *md.LobbyEvent, K_MAX_MESSAGES),
	}

	// Start Exchange Server
	err := mgr.registerService()
	if err != nil {
		return nil, err
	}

	// Start IP Topic
	ipSub, ipHandler, err := mgr.joinIPTopic()
	if err != nil {
		return nil, err
	}

	// Set Service
	go mgr.handleTopicEvents(ipHandler)
	go mgr.handleTopicMessages(ipSub)
	go mgr.processTopicMessages()
	return mgr, nil
}

// ^ Helper: Find returns Pointer to Peer.ID and Peer ^
func (tm *LocalManager) FindPeerInTopic(q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	var p *md.Peer
	var i peer.ID

	// Iterate Through Peers, Return Matched Peer
	p, err := tm.findPeer(q)
	if err != nil {
		return "", nil, err
	}

	// Iterate through Topic Peers
	id, err := tm.searchIPTopic(q)
	if err == nil {
		i = id
	}

	// Check Geo Topic
	id, err = tm.searchGeoTopic(q)
	if err == nil {
		i = id
	}

	// Validate ID
	if i == "" {
		return "", nil, errors.New("Peer ID was not found in topic.")
	}
	return i, p, nil
}

// ^ Helper: ID returns ONE Peer.ID in Topic ^
func (tm *LocalManager) HasPeer(q string) bool {
	// Iterate through PubSub in topic
	_, err := tm.searchIPTopic(q)
	if err == nil {
		return true
	}

	// Check Geo Topic
	_, err = tm.searchGeoTopic(q)
	if err == nil {
		return true
	}
	return false
}

// ^ Send message to specific peer in topic ^
func (tm *LocalManager) Send(msg *md.LobbyEvent) error {
	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = tm.ipTopic.Publish(tm.ctx, bytes)
	if err != nil {
		return err
	}

	// Check Geo Topic
	if tm.geoTopic != nil {
		// Publish to Topic
		err = tm.geoTopic.Publish(tm.ctx, bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

// ^ Leave Current Topic ^
func (tm *LocalManager) LeaveTopic() {
	tm.ipTopic.Close()

	// Check Geo Topic
	if tm.geoTopic != nil {
		tm.geoTopic.Close()
	}
}
