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
	Data     string
	Event    string
	SenderID string
}

// Bytes converts message struct to JSON bytes
func (msg *Message) Bytes() []byte {
	// Convert to Bytes
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
