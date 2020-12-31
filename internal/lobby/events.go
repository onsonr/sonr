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
			if lobEvent.Peer == lob.self {
				continue
			} else {
				if err := lob.Update(); err != nil {
					log.Println(err)
					continue
				}
			}
		}

		// Remove Peer from Lobby
		if lobEvent.Type == pubsub.PeerLeave {
			if lobEvent.Peer == lob.self {
				continue
			} else {
				lob.removePeer(lobEvent.Peer)
				continue
			}
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
			lob.setPeer(m)
		case <-lob.ctx.Done():
			return
		}

		// @ Pausable
		lifecycle.GetState().NeedsWait()
	}
}

// ^ Send Refresh calls back lobby to Node ^ //
func (lob *Lobby) sendRefresh() {
	// Set Size
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Marshal data to bytes
	bytes, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.callRefresh(bytes)
}

// // ^ Send Event calls back event to Node ^ //
// func (lob *Lobby) sendEvent(event *md.LobbyEvent) {
// 	// Marshal data to bytes
// 	bytes, err := proto.Marshal(event)
// 	if err != nil {
// 		log.Println("Cannot Marshal Error Protobuf: ", err)
// 	}

// 	// Send Callback with updated peers
// 	lob.callEvent(bytes)
// }
