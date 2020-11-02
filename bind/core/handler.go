package core

import (
	"github.com/libp2p/go-libp2p-core/peer"
	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
)

// Kill Channel
var doneCh chan struct{}

// Slice of Current Messages
var messages []sonrLobby.Message

// handleEvents manages input stream
func (sn *SonrNode) handleEvents() {
	// Check for Updates
	for {
		select {
		case m := <-sn.Lobby.Messages:
			// When message received add to slice
			messages = append(messages, *m)
		case <-sn.Lobby.CTX.Done():
			return
		case <-doneCh:
			return
		}
	}
}

func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}
