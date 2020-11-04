package lobby

import (
	"encoding/json"
)

// Message is a for Lobby Pub/Sub Messaging, Converted To/From Json
type Message struct {
	Value    string
	Event    string
	SenderID string
}

// Bytes converts message struct to JSON bytes
func (msg *Message) Bytes() []byte {
	// Convert to JSON
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		println(err)
	}
	return msgBytes
}

// String converts message struct to JSON String
func (msg *Message) String() string {
	// Convert to JSON
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		println(err)
	}
	return string(msgBytes)
}

// ConnectMessage is message sent when user wants to join network
type ConnectMessage struct {
	OLC     string
	Device  string
	Profile string
}

// UpdateMessage is sent when device has state change
type UpdateMessage struct {
	Direction float64
	Status    string
}
