package lobby

import (
	"github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

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

		// Send valid messages onto the Messages channel
		lob.Messages <- &notif
		lifecycle.GetState().NeedsWait()
	}

}

// ^ 1a. processMessages handles message content and ticker ^
func (lob *Lobby) processMessages() {
	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			// @ Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				lob.updatePeer(m.Peer)
			} else if m.Event == md.LobbyEvent_EXIT {
				// Update Peer Data
				lob.updatePeer(m.Peer)
			}

		case <-lob.ctx.Done():
			return
		}
		lifecycle.GetState().NeedsWait()
	}
}
