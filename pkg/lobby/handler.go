package lobby

import (
	"time"

	pb "github.com/sonr-io/core/pkg/models"
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
		if msg.ReceivedFrom.String() == lob.Self.GetId() {
			continue
		}

		// Construct message
		notif := pb.LobbyMessage{}
		err = proto.Unmarshal(msg.Data, &notif)
		if err != nil {
			continue
		}

		// Send valid messages onto the Messages channel
		lob.Messages <- &notif
	}
}

// ^ 2. handleEvents handles message content and ticker ^
func (lob *Lobby) handleEvents() {
	// Timer checks to dispose of peers
	peerRefreshTicker := time.NewTicker(time.Second * 3)
	defer peerRefreshTicker.Stop()

	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			// Update Circle by event
			if m.Subject == pb.LobbyMessage_UPDATE {
				// Update Peer Data
				lob.updatePeer(m.Id, m.Peer)

			} else if m.Subject == pb.LobbyMessage_EXIT {
				// Remove Peer Data
				lob.removePeer(m.Id)
			}

		// ** Refresh and Validate Lobby Peers Periodically ** //
		case <-peerRefreshTicker.C:
			lob.call.Refreshed(lob.Peers())

		case <-lob.ctx.Done():
			return

		case <-lob.doneCh:
			return
		}
	}
}
