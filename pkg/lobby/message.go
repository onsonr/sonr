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

// AuthRequestMessage is for Auth Stream Request
type AuthRequestMessage struct {
	PeerInfo  string
	FileInfo  string
	Thumbnail []byte
}

// AuthInviteMessage is for Auth Stream Request
type AuthResponseMessage struct {
	decision bool
	peerID   string
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
