package lobby

import (
	"encoding/json"

	pb "github.com/sonr-io/core/pkg/proto"
)

// Notification is a for Lobby Pub/Sub Messaging, Converted To/From Json
type Notification struct {
	Data   string
	Event  string
	Sender string
	Peer   *pb.PeerInfo
}

// Bytes converts message struct to JSON bytes
func (msg *Notification) Bytes() []byte {
	// Convert to Bytes
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		println(err)
	}
	return msgBytes
}

// String converts message struct to JSON String
func (msg *Notification) String() string {
	// Convert to JSON
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		println(err)
	}
	return string(msgBytes)
}
