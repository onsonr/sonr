package lobby

import (
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleMessages() {
	for {
		// Get next msg from pub/sub
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.Messages)
			return
		}

		// Only forward messages delivered by others
		if msg.ReceivedFrom == lob.self {
			continue
		}

		// Construct message
		notif := md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, &notif)
		if err != nil {
			continue
		}

		// Send valid messages onto the Messages channel
		lob.Messages <- &notif
		lifecycle.GetState().NeedsWait()
	}
}

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleEvents() {
	// @ Create Topic Handler
	topicHandler, err := lob.topic.EventHandler()
	if err != nil {
		log.Println(err)
		return
	}

	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := topicHandler.NextPeerEvent(lob.ctx)
		if err != nil {
			topicHandler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			log.Println("Lobby Event: Peer Joined")
		}

		if lobEvent.Type == pubsub.PeerLeave {
			log.Println("Lobby Event: Peer Left")
		}

		lifecycle.GetState().NeedsWait()
	}
}

// ^ processMessages handles message content and ticker ^
func (lob *Lobby) processMessages() {
	for {
		select {
		// @ Message Received
		case m := <-lob.Messages:
			if m.Event == md.LobbyEvent_UPDATE {
				lob.updatePeer(m)
			} else if m.Event == md.LobbyEvent_STANDBY {
				lob.setUnavailable(m)
			} else if m.Event == md.LobbyEvent_BUSY {
				lob.setUnavailable(m)
			} else if m.Event == md.LobbyEvent_EXIT {
				lob.removePeer(m)
			}
		case <-lob.ctx.Done():
			return
		}

		// @ Pausable
		lifecycle.GetState().NeedsWait()
	}
}
