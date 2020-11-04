package lobby

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// SonrCallback returns message from lobby
type SonrCallback interface {
	OnMessage(s string)
	OnNewPeer(s string)
}

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Messages is a channel of messages received from other peers in the chat room
	messages chan *Message
	Callback SonrCallback

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	Code   string
	selfID peer.ID
}

// Enter tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func Enter(ctx context.Context, call SonrCallback, ps *pubsub.PubSub, hostID peer.ID, olcCode string) (*Lobby, error) {
	// join the pubsub topic
	topic, err := ps.Join(olcName(olcCode))
	if err != nil {
		return nil, err
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:      ctx,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		selfID:   hostID,
		Code:     olcCode,
		Callback: call,
		messages: make(chan *Message, ChatRoomBufSize),
	}

	// start reading messages from the subscription in a loop
	go lob.readLoop()
	return lob, nil
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

// ListPeers returns peerids in room
func (lob *Lobby) ListPeers() []peer.ID {
	return lob.ps.ListPeers(olcName(lob.Code))
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (lob *Lobby) readLoop() {
	for {
		// get next msg from pub/sub
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.messages)
			return
		}

		// only forward messages delivered by others
		if msg.ReceivedFrom == lob.selfID {
			continue
		}

		// construct message
		cm := new(Message)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}

		lob.Callback.OnMessage(cm.String())

		// send valid messages onto the Messages channel
		lob.messages <- cm
	}
}

func olcName(code string) string {
	return "olc=" + code
}
