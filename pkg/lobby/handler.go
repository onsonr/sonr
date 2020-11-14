package lobby

import (
	"encoding/json"
)

// 1. handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (lob *Lobby) handleMessages() {
	for {
		// get next msg from pub/sub
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.Messages)
			return
		}

		// only forward messages delivered by others
		if msg.ReceivedFrom.String() == lob.Self.ID {
			continue
		}

		// construct message
		cm := new(Message)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}

		// send valid messages onto the Messages channel
		lob.Messages <- cm
	}
}

// 2. handleEvents handles message content and ticker
func (lob *Lobby) handleEvents() {
	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			// Update Circle by event
			if m.Event == "Join" {
				lob.joinPeer(m.Data)
			} else if m.Event == "Update" {
				lob.updatePeer(m.Data)
			} else if m.Event == "Leave" {
				lob.removePeer(m.Data)
			}

		case <-lob.ctx.Done():
			return

		case <-lob.doneCh:
			return
		}
	}
}
