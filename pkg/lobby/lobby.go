package lobby

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Callback returns message from lobby
type Callback interface {
	OnMessage(s string)
	OnRefresh(s string)
	OnRequested(s string)
	OnAccepted(s string)
	OnDenied(s string)
	OnProgress(s string)
	OnComplete(s string)
}

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Public Vars
	Messages chan *Message
	Code     string
	Self     Peer

	// Private Vars
	callback Callback
	peers    map[string]Peer
	ctx      context.Context
	ps       *pubsub.PubSub
	topic    *pubsub.Topic
	sub      *pubsub.Subscription
	doneCh   chan struct{}
}

// Enter tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func Enter(ctx context.Context, call Callback, ps *pubsub.PubSub, hostID peer.ID, firstName string, lastName string, device string, profilePic string, olcCode string) (*Lobby, error) {
	// join the pubsub topic
	topic, err := ps.Join(olcCode)
	if err != nil {
		return nil, err
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Set Peer Info
	peer := Peer{
		ID:         hostID.String(),
		Device:     device,
		FirstName:  firstName,
		LastName:   lastName,
		ProfilePic: profilePic,
		Direction:  0.0,
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:      ctx,
		callback: call,
		doneCh:   make(chan struct{}, 1),
		peers:    make(map[string]Peer),
		ps:       ps,
		topic:    topic,
		sub:      sub,
		Self:     peer,
		Code:     olcCode,
		Messages: make(chan *Message, ChatRoomBufSize),
	}

	// Publish Join Message
	msg := Message{
		Event:    "Join",
		Data:     peer.String(),
		SenderID: hostID.String(),
	}
	lob.Publish(msg)

	// start reading messages
	go lob.handleMessages()
	go lob.handleEvents()
	return lob, nil
}

// GetPeers returns peers list as string
func (lob *Lobby) GetPeers() string {
	// Initialize Variables
	var peerSlice []Peer
	peersRef := lob.peers

	// Iterate through dictionary
	for _, value := range peersRef {
		// Add to slice
		peerSlice = append(peerSlice, value)
	}

	// Convert slice to bytes
	bytes, err := json.Marshal(peerSlice)
	if err != nil {
		println("Error converting peers to json ", err)
	}

	// Return as string
	return string(bytes)
}

// ListPeers returns Pub/Sub Topic Peers
func (lob *Lobby) ListPeers() []peer.ID {
	return lob.ps.ListPeers(lob.Code)
}

// Publish sends a message to the pubsub topic.
func (lob *Lobby) Publish(m Message) error {
	// Publish to Topic
	err := lob.topic.Publish(lob.ctx, m.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// End terminates lobby loop
func (lob *Lobby) End() {
	lob.doneCh <- struct{}{}
}
