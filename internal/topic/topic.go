package topic

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
const LOCAL_SERVICE_PID = protocol.ID("/sonr/local-service/0.2")
const REMOTE_SERVICE_PID = protocol.ID("/sonr/remote-service/0.2")

type ClientHandler interface {
	OnLocalEvent(*md.LocalEvent)
	OnRefresh(*md.Lobby)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnResponded(inv *md.InviteRequest)
}

type TopicManager struct {
	ctx          context.Context
	host         *net.HostNode
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	eventHandler *pubsub.TopicEventHandler
	user         *md.User
	lobby        *md.Lobby

	service     *LocalService
	localEvents chan *md.LocalEvent
	handler     ClientHandler
	lobbyType   md.Lobby_Type
}

// ^ Helper: Find returns Pointer to Peer.ID and Peer ^
func (tm *TopicManager) FindPeerInTopic(q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	var p *md.Peer
	var i peer.ID

	// Iterate Through Peers, Return Matched Peer
	for _, peer := range tm.lobby.Peers {
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

// ^ Check if Local Topic
func (tm *TopicManager) IsLocal() bool {
	if tm.lobbyType == md.Lobby_LOCAL {
		return true
	}
	return false
}

// ^ Leave Current Topic ^
func (tm *TopicManager) LeaveTopic() error {
	tm.eventHandler.Cancel()
	tm.subscription.Cancel()
	return tm.topic.Close()
}

// ^ handleTopicEvents: listens to Pubsub Events for topic  ^
func (tm *TopicManager) handleTopicEvents() {
	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := tm.eventHandler.NextPeerEvent(tm.ctx)
		if err != nil {
			tm.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if lobEvent.Type == pubsub.PeerJoin && tm.user.Peer.IsNotSamePeerID(lobEvent.Peer) {
			pbuf, err := tm.user.GetPeer().Buffer()
			if err != nil {
				continue
			}
			err = tm.Exchange(lobEvent.Peer, pbuf)
			if err != nil {
				continue
			}
			tm.RefreshLobby()
		}

		// Check Leave Eent
		if lobEvent.Type == pubsub.PeerLeave {
			tm.lobby.Delete(lobEvent.Peer)
			tm.RefreshLobby()
		}
		md.GetState().NeedsWait()
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
		if tm.user.GetPeer().IsSamePeerID(msg.ReceivedFrom) {
			continue
		}

		// Check Lobby Type
		// Construct message
		m := &md.LocalEvent{}
		err = proto.Unmarshal(msg.Data, m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if tm.HasPeer(m.Id) {
			tm.localEvents <- m
		}
		md.GetState().NeedsWait()
	}
}

// ^ processTopicMessages: pulls messages from channel that have been handled ^
func (tm *TopicManager) processTopicMessages() {
	for {
		select {
		// @ Local Event Channel Updated
		case m := <-tm.localEvents:
			if m.Subject == md.LocalEvent_UPDATE {
				// Update Peer Data
				tm.lobby.Add(m.From)
				tm.RefreshLobby()
			}
		case <-tm.ctx.Done():
			return
		}
		md.GetState().NeedsWait()
	}
}
