package topic

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Send Updated Lobby ^
func (tm *TopicManager) Refresh() {
	tm.topicHandler.OnRefresh(tm.lobby)
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

		if lobEvent.Type == pubsub.PeerJoin {
			pbuf, err := tm.user.GetPeer().Buffer()
			if err != nil {
				continue
			}
			lbuf, err := tm.lobby.Buffer()
			if err != nil {
				continue
			}
			err = tm.Exchange(lobEvent.Peer, pbuf, lbuf)
			if err != nil {
				continue
			}
			tm.Refresh()
		}

		if lobEvent.Type == pubsub.PeerLeave {
			tm.lobby.Delete(lobEvent.Peer)
			tm.Refresh()
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

		// Construct message
		m := &md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if tm.HasPeer(m.Id) {
			tm.messages <- m
		}
		md.GetState().NeedsWait()
	}
}

// ^ processTopicMessages: pulls messages from channel that have been handled ^
func (tm *TopicManager) processTopicMessages() {
	for {
		select {
		// @ when we receive a message from the lobby room
		case m := <-tm.messages:
			tm.handleMessage(m)
		case <-tm.ctx.Done():
			return
		}
		md.GetState().NeedsWait()
	}
}

// ^ handleMessage: performs action for Message Type and Event Kind ^
func (tm *TopicManager) handleMessage(e *md.LobbyEvent) {
	switch e.Event.(type) {
	// Local Event
	case *md.LobbyEvent_Local:
		event := e.GetLocal()
		if event == md.LobbyEvent_UPDATE {
			// Update Peer Data
			tm.lobby.Add(e.From)
			tm.Refresh()
		}

	// Remote Event
	case *md.LobbyEvent_Remote:
		tm.topicHandler.OnEvent(e)

	default:
		return
	}
}
