package lobby

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Messages is a channel of messages received from other peers in the chat room
	Messages    chan *Message
	LastMessage string

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	ID   string
	self peer.ID
}

// Message gets converted to/from JSON and sent in the body of pubsub messages.
type Message struct {
	Value    string
	Event    string
	SenderID string
}

// JoinLobby tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func JoinLobby(ctx context.Context, h *host.Host, selfID peer.ID, olcCode string) *Lobby {
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, *h)
	if err != nil {
		panic(err)
	}

	// join the pubsub topic
	topic, err := ps.Join(olcName(olcCode))
	if err != nil {
		panic(err)
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}

	// Create Lobby Type
	cr := &Lobby{
		ctx:      ctx,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		self:     selfID,
		ID:       olcCode,
		Messages: make(chan *Message, ChatRoomBufSize),
	}

	// start reading messages from the subscription in a loop
	go cr.readLoop()
	return cr
}

// Publish sends a message to the pubsub topic.
func (cr *Lobby) Publish(message string) {
	m := Message{
		Value:    message,
		SenderID: cr.self.Pretty(),
	}
	msgBytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	cr.topic.Publish(cr.ctx, msgBytes)
}

// ListPeers returns peerids in room
func (cr *Lobby) ListPeers() []peer.ID {
	return cr.ps.ListPeers(olcName(cr.ID))
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (cr *Lobby) readLoop() {
	for {
		msg, err := cr.sub.Next(cr.ctx)
		if err != nil {
			close(cr.Messages)
			return
		}
		// Set Last message
		cr.LastMessage = string(msg.Data)

		// only forward messages delivered by others
		if msg.ReceivedFrom == cr.self {
			continue
		}
		cm := new(Message)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}
		// send valid messages onto the Messages channel
		cr.Messages <- cm
	}
}

func olcName(roomName string) string {
	return "olc=" + roomName
}
