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
		message := md.LobbyMessage{}
		err = proto.Unmarshal(msg.Data, &message)
		if err != nil {
			continue
		}

		// Send valid messages onto the Messages channel
		lob.Messages <- &message
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

		// Push User Data to new peer
		if lobEvent.Type == pubsub.PeerJoin {
			log.Println("Lobby Event: Peer Joined")
			err := lob.Update()
			if err != nil {
				log.Println(err)
			}
		}

		// Remove Peer from Lobby
		if lobEvent.Type == pubsub.PeerLeave {
			log.Println("Lobby Event: Peer Left")
			lob.removePeer(lobEvent.Peer.String())
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
			if m.Event == md.LobbyMessage_EXCHANGE {
				if err := lob.Exchange(m); err != nil {
					log.Println(err)
				}
			} else if m.Event == md.LobbyMessage_AVAILABLE {
				lob.setPeer(m)
			} else if m.Event == md.LobbyMessage_BUSY {
				lob.setUnavailable(m)
			}
		case <-lob.ctx.Done():
			return
		}

		// @ Pausable
		lifecycle.GetState().NeedsWait()
	}
}
