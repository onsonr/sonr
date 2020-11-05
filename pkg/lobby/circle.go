package lobby

import (
	"encoding/json"
	"time"
)

// ListPeers returns peerids in room
func (lob *Lobby) handleEvents() {
	peerRefreshTicker := time.NewTicker(time.Second)
	defer peerRefreshTicker.Stop()

	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			println(m.Event)

		// ** refresh the list of peers in the chat room periodically **
		case <-peerRefreshTicker.C:
			// Get Peers
			peers := lob.ps.ListPeers(olcName(lob.Code))
			bytes, err := json.Marshal(peers)

			// Check for Error
			if err != nil {
				println("Error refreshing peers")
			}

			// Callback
			lob.callback.OnRefresh(string(bytes))
			continue

		case <-lob.ctx.Done():
			return

		case <-lob.doneCh:
			return
		}
	}
}

// End terminates lobby loop
func (lob *Lobby) End() {
	lob.doneCh <- struct{}{}
}
