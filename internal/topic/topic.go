package topic

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	net "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

type ClientHandler interface {
	OnEvent(*md.LocalEvent)
	OnRefresh(*md.Lobby)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnResponded(inv *md.InviteRequest)
}

type TopicManager struct {
	ctx          context.Context
	host         net.HostNode
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	eventHandler *pubsub.TopicEventHandler
	user         *md.User

	service     *LocalService
	localEvents chan *md.LocalEvent
	handler     ClientHandler
	lobbyType   md.Lobby_Type
}

// @ Helper: Find returns Pointer to Peer.ID and Peer
func (tm *TopicManager) FindPeerInTopic(q string) (peer.ID, error) {
	// Retreive Data
	var i peer.ID

	// Iterate through Topic Peers
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			i = id
		}
	}

	// Validate ID
	if i == "" {
		return "", errors.New("Peer ID was not found in topic.")
	}
	return i, nil
}

// @ Helper: ID returns ONE Peer.ID in Topic
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

// @ Check if Local Topic
func (tm *TopicManager) IsLocal() bool {
	if tm.lobbyType == md.Lobby_LOCAL {
		return true
	}
	return false
}

// # handleTopicEvents: listens to Pubsub Events for topic
func (tm *TopicManager) handleTopicEvents(ctx context.Context) {
	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := tm.eventHandler.NextPeerEvent(ctx)
		if err != nil {
			tm.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if lobEvent.Type == pubsub.PeerJoin && lobEvent.Peer != tm.host.ID() {
			pbuf, err := tm.user.GetPeer().Buffer()
			if err != nil {
				continue
			}
			err = tm.Exchange(lobEvent.Peer, pbuf)
			if err != nil {
				continue
			}
		}

		// Check Leave Eent
		if lobEvent.Type == pubsub.PeerLeave {
			tm.RefreshLobby(md.NewExitLocalEvent(lobEvent.Peer.String()))
		}
		md.GetState().NeedsWait()
	}
}

// # handleTopicMessages: listens for messages on pubsub topic subscription
func (tm *TopicManager) handleTopicMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := tm.subscription.Next(ctx)
		if err != nil {
			return
		}

		// Only forward messages delivered by others
		if tm.user.GetPeer().IsSamePeerID(msg.ReceivedFrom) {
			continue
		}

		// Check Lobby Type
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

// # processTopicMessages: pulls messages from channel that have been handled
func (tm *TopicManager) processTopicMessages(ctx context.Context) {
	for {
		select {
		// @ Local Event Channel Updated
		case m := <-tm.localEvents:
			tm.RefreshLobby(m)
		case <-ctx.Done():
			return
		}
		md.GetState().NeedsWait()
	}
}
