package lobby

import (
	"context"
	"encoding/json"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"

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
	ctx      context.Context
	callback Callback
	doneCh   chan struct{}
	peerDB   *badger.DB
	ps       *pubsub.PubSub
	topic    *pubsub.Topic
	sub      *pubsub.Subscription
}

// Enter Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby
func Enter(ctx context.Context, call Callback, ps *pubsub.PubSub, hostID peer.ID, firstName string, lastName string, device string, profilePic string, olcCode string) (*Lobby, error) {
	// Create Peer Struct
	peer := Peer{
		ID:         hostID,
		Device:     device,
		FirstName:  firstName,
		LastName:   lastName,
		ProfilePic: profilePic,
		Direction:  0.0,
	}

	// Join the pubsub Topic
	topic, err := ps.Join(olcCode)
	if err != nil {
		return nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Initialize Badger DB
	opt := badger.DefaultOptions("").WithInMemory(true)
	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:      ctx,
		callback: call,
		doneCh:   make(chan struct{}, 1),
		peerDB:   db,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		Self:     peer,
		Code:     olcCode,
		Messages: make(chan *Message, ChatRoomBufSize),
	}

	// Publish Join Message
	msg := Message{
		Event:    "Update",
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
	// ** Initialize Variables ** //
	var peerSlice []Peer

	// ** Open Data Store Read Transaction ** //
	err := lob.peerDB.View(func(txn *badger.Txn) error {
		// @ Create Iterator
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		// @ Iterate over bucket
		for it.Rewind(); it.Valid(); it.Next() {
			// Get Item and Key
			item := it.Item()
			id := item.Key()

			// Get Item Value
			err := item.Value(func(peer []byte) error {
				// Log Key/Value
				fmt.Printf("id=%s, peer=%s\n", id, peer)

				// Convert Value to String Add to Slice
				cm := new(Peer)
				err := json.Unmarshal(peer, cm)

				// Check for Error
				if err != nil {
					fmt.Println(err)
				} else {
					// Add Item value to Slice
					peerSlice = append(peerSlice, *cm)
				}
				return nil
			})

			// Check error
			if err != nil {
				return err
			}
		}
		return nil
	})

	// Check for Error
	if err != nil {
		fmt.Println(err)
	}

	// ** Convert slice to bytes ** //
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
	lob.peerDB.Close()
	lob.doneCh <- struct{}{}
}
