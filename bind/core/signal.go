package core

import (
	"bytes"

	sonrLobby "github.com/sonr-io/p2p/pkg/lobby"
)

// Message gets converted to/from JSON and sent in the body of pubsub messages.
type Message struct {
	Message  string
	SenderID string
}

// Publish Sends message to Lobby
func Publish(lobby *sonrLobby.Lobby, testPublish string) {
	lobby.Publish(testPublish)
}

// GetMessages returns messages as
func GetMessages(lobby *sonrLobby.Lobby) {
	var buf bytes.Buffer
	lobby.Messages
}

// Update enters the room with given OLC(Open-Location-Code)
func Update(updateJSON string) string {
	nodeProfile = updateJSON
	lobbyRef.Publish(updateJSON)
	return lobbyRef.ID
}

// Exit informs lobby and closes host
func Exit() string {
	lobbyRef.Publish(nodeProfile)
	hostNode.Close()
	return lobbyRef.ID
}

// Offer informs peer about file
func Offer() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Answer and accept Peers offer
func Answer() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Decline Peers offer
func Decline() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Failed informs Peers transfer was unsuccesful
func Failed() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}
