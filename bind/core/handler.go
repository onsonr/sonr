package core

import (
	"github.com/libp2p/go-libp2p-core/peer"
)

// Kill Channel
var doneCh chan struct{}

// EmitUpdate enters the room with given OLC(Open-Location-Code)
func (lob *Lobby) EmitUpdate(updateJSON string) {
	lob.Publish(updateJSON)
}

// EmitSend publishes message to lobby
func (lob *Lobby) EmitSend(content string) {
	lob.Publish(string(content))
}

// EmitExit makes host leave lobby
func (lob *Lobby) EmitExit(data string) {
	// Inform Lobby youre leaving
	lob.Publish(data)

	// Kill Event Loop
	doneCh <- struct{}{}
}

// EmitOffer informs peer about file
func (lob *Lobby) EmitOffer(data string) {
	lob.Publish(data)
}

// EmitAnswer and accept Peers offer
func (lob *Lobby) EmitAnswer(data string) {
	lob.Publish(data)
}

// EmitDecline Peers offer
func (lob *Lobby) EmitDecline(data string) {
	lob.Publish(data)
}

// EmitFailed informs Peers transfer was unsuccesful
func (lob *Lobby) EmitFailed(data string) {
	lob.Publish(data)
}

func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}
