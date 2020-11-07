package lobby

import (
	"encoding/json"
)

// ConnectRequest is message sent when user wants to join network
type ConnectRequest struct {
	OLC     string
	Device  string
	Contact string
}

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

// UpdateNotification is sent when device has state change
type UpdateNotification struct {
	Direction float64
	Status    string
	ID        string
}
