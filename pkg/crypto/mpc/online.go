package mpc

import (
	"context"
	"fmt"
	"sync"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	mpc "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// `onlineNetwork` is a struct that contains a `context.Context`, a `node.Node`, a list of `peer.ID`s, a list
// of `party.ID`s, a `sync.Mutex`, a `chan bool`, a `chan error`, a `chan *mpc.Message`, and a
// `map[party.ID]icore.PubSubSubscription`.
// @property ctx - the context of the network
// @property selfNode - The node that represents the current party.
// @property {[]peer.ID} peerIds - a list of peer IDs that are connected to this node
// @property {[]party.ID} partyIds - a list of party IDs that this node is connected to
// @property mtx - a mutex to protect the network from concurrent access
// @property doneChan - a channel that will be closed when the network is stopped
// @property errorChan - This channel is used to send errors to the caller.
// @property msgInChan - This is the channel that the network will use to send messages to the
// application.
// @property subscriptions - a map of party IDs to PubSub subscriptions.
type onlineNetwork struct {
	nodes   []*node.Node
	parties party.IDSlice

	mtx  sync.Mutex
	done chan struct{}

	closedListenChan chan *mpc.Message
	listenChannels   map[party.ID]chan *mpc.Message
	subscriptions    map[party.ID]icore.PubSubSubscription
}

// It creates a new network object, assigns the subscriptions, and returns the network object
func NewOnlineNetwork(ctx context.Context, nodes ...*node.Node) (Network, error) {
	// Convert the peer IDs to party IDs.
	parties := make([]party.ID, 0)
	for _, node := range nodes {
		parties = append(parties, party.ID(node.ID()))
	}

	closed := make(chan *protocol.Message)
	close(closed)
	// Create the network object.
	net := &onlineNetwork{
		nodes:            nodes,
		parties:          parties,
		done:             make(chan struct{}),
		listenChannels:   make(map[party.ID]chan *mpc.Message, 2*len(parties)),
		subscriptions:    make(map[party.ID]icore.PubSubSubscription, 2*len(parties)),
		closedListenChan: closed,
	}
	return net, nil
}

// Initializing the network.
func (n *onlineNetwork) init() {
	N := len(n.parties)
	for _, id := range n.parties {
		n.listenChannels[id] = make(chan *protocol.Message, N*N)
	}
	n.done = make(chan struct{})

	for _, node := range n.nodes {
		sub, err := node.PubSub().Subscribe(context.Background(), topicKey(node.PartyID()))
		if err != nil {
			panic(err)
		}
		n.subscriptions[node.PartyID()] = sub
		go handleSubscription(node.PartyID(), sub, n)
	}
}

// IsOnlineNetwork returns true.
func (n *onlineNetwork) IsOnlineNetwork() bool {
	return true
}

// Ls returns the list of parties that are connected to the network.
func (n *onlineNetwork) Ls() []party.ID {
	return n.parties
}

// Returning the channel that the network will use to send messages to the application.
func (n *onlineNetwork) Next(id party.ID) <-chan *mpc.Message {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	if len(n.listenChannels) == 0 {
		n.init()
	}
	c, ok := n.listenChannels[id]
	if !ok {
		return n.closedListenChan
	}
	return c
}

// Sending a message to the network.
func (n *onlineNetwork) Send(msg *mpc.Message) {
	bz, err := msg.MarshalBinary()
	if err != nil {
		fmt.Printf("error while marshaling message: %e", err)
		return
	}

	n.mtx.Lock()
	defer n.mtx.Unlock()

	for _, node := range n.nodes {
		if msg.IsFor(node.PartyID()) {
			if err := node.Publish(topicKey(node.PartyID()), bz); err != nil {
				fmt.Printf("error while publishing message: %e", err)
			}
		}
	}
}

// Closing the subscriptions and returning the done channel.
func (n *onlineNetwork) Done(id party.ID) chan struct{} {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	for id, sub := range n.subscriptions {
		sub.Close()
		delete(n.listenChannels, id)
		delete(n.subscriptions, id)
	}
	if len(n.listenChannels) == 0 {
		close(n.done)
	}
	return n.done
}

// Removing the party from the network.
func (n *onlineNetwork) Quit(id party.ID) {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.parties = n.parties.Remove(id)
}

//
// Private methods
//

func (n *onlineNetwork) findOutTopic(msg *mpc.Message) string {
	for _, node := range n.nodes {
		if msg.IsFor(node.PartyID()) {
			return topicKey(node.PartyID())
		}
	}
	return ""
}

func (n *onlineNetwork) getFromNode(msg *mpc.Message) *node.Node {
	for _, node := range n.nodes {
		if msg.From == node.PartyID() {
			return node
		}
	}
	return nil
}

// A goroutine that is listening for messages on the topic handler.
func handleSubscription(id party.ID, sub icore.PubSubSubscription, n *onlineNetwork) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			return err
		}
		m := &mpc.Message{}
		err = m.UnmarshalBinary(msg.Data())
		if err != nil {
			return err
		}
		n.listenChannels[id] <- m
		select {
		case <-n.done:
			return nil
		}
	}
}

// topicKey returns the modified topic key for the given party ID.
func topicKey(id party.ID) string {
	return fmt.Sprintf("%s/mpc/cmp-keygen", string(id))
}
