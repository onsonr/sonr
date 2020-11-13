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
	// Messages is a channel of Messages received from other peers in the chat room
	Messages chan *Message
	callback Callback

	peers  []Peer
	ctx    context.Context
	ps     *pubsub.PubSub
	topic  *pubsub.Topic
	sub    *pubsub.Subscription
	doneCh chan struct{}

	Code string
	Self *Peer
}

// Enter tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func Enter(ctx context.Context, call Callback, ps *pubsub.PubSub, hostID peer.ID, firstName string, lastName string, device string, profilePic string, olcCode string) (*Lobby, error) {
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

	// Set Peer Info
	peer := Peer{
		ID:         hostID.String(),
		Device:     device,
		FirstName:  firstName,
		LastName:   lastName,
		ProfilePic: profilePic,
	}

	// Handle Graph
	var peers []Peer
	peers = append(peers, peer)

	// Create Lobby Type
	lob := &Lobby{
		ctx:      ctx,
		doneCh:   make(chan struct{}, 1),
		peers:    peers,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		Self:     &peer,
		Code:     olcCode,
		callback: call,
		Messages: make(chan *Message, ChatRoomBufSize),
	}

	// Publish Join Message
	msg := Message{
		Event:    "Join",
		Value:    peer.String(),
		SenderID: hostID.String(),
	}
	lob.Publish(msg)

	// start reading messages
	go lob.handleMessages()
	go lob.handleEvents()
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

// End terminates lobby loop
func (lob *Lobby) End() {
	lob.doneCh <- struct{}{}
}

// handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (lob *Lobby) handleMessages() {
	for {
		// get next msg from pub/sub
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.Messages)
			return
		}

		// only forward messages delivered by others
		if msg.ReceivedFrom.String() == lob.Self.ID {
			continue
		} else {
			// callback new message
			lob.callback.OnMessage(string(msg.Data))
		}

		// construct message
		cm := new(Message)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}

		// send valid messages onto the Messages channel
		lob.Messages <- cm
	}
}

func olcName(code string) string {
	return "olc=" + code
}
