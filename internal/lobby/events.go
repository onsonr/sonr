package lobby

import (
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

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
			err := lob.Exchange(lobEvent.Peer)
			if err != nil {
				log.Println(err)
			}
		}

		if lobEvent.Type == pubsub.PeerLeave {
			// Create Event Message
			lobEvent := &md.LobbyEvent{
				Id:    lobEvent.Peer.String(),
				Event: md.LobbyEvent_EXIT,
			}

			// Marshal data to bytes
			bytes, err := proto.Marshal(lobEvent)
			if err != nil {
				log.Println("Cannot Marshal Error Protobuf: ", err)
			}

			// Send Callback with updated peers
			lob.onEvent(bytes)
		}
		lifecycle.GetState().NeedsWait()
	}
}

// ^ 1. handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
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

		if notif.Event == md.LobbyEvent_EXCHANGE {
			// Update Peer Data
			err := lob.Exchange(lob.ID(notif.Id))
			if err != nil {
				log.Println(err)
			}
		} else {
			// Send Callback with updated peers
			lob.onEvent(msg.Data)
		}

		// Send valid messages onto the Messages channel
		// lob.Messages <- &notif
		lifecycle.GetState().NeedsWait()
	}

}

// ^ 1a. processMessages handles message content and ticker ^
func (lob *Lobby) processMessages() {
	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			// Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				lob.updatePeer(m.Peer)
			} else if m.Event == md.LobbyEvent_EXCHANGE {
				// Update Peer Data
				lob.updatePeer(m.Peer)
				err := lob.Exchange(lob.ID(m.Id))
				if err != nil {
					log.Println(err)
				}
			}

		case <-lob.ctx.Done():
			return

		case <-lob.doneCh:
			return
		}
		lifecycle.GetState().NeedsWait()
	}
}
