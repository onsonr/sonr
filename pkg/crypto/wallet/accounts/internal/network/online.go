package network

import (
	"context"
	"fmt"
	"sync"

	ps "github.com/libp2p/go-libp2p-pubsub"

	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/crypto"
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
	nodes   []common.P2PNode
	parties party.IDSlice

	mtx  sync.Mutex
	done chan struct{}

	closedListenChan chan *mpc.Message
	listenChannels   map[party.ID]chan *mpc.Message
	subscriptions    map[party.ID]*ps.Subscription
}

// It creates a new network object, assigns the subscriptions, and returns the network object
func NewOnlineNetwork(ctx context.Context, nodes ...common.P2PNode) (crypto.Network, error) {
	// Convert the peer IDs to party IDs.
	parties := make([]party.ID, 0)
	for _, node := range nodes {
		parties = append(parties, party.ID(node.PeerID()))
	}

	closed := make(chan *protocol.Message)
	close(closed)
	// Create the network object.
	net := &onlineNetwork{
		nodes:            nodes,
		parties:          parties,
		done:             make(chan struct{}),
		listenChannels:   make(map[party.ID]chan *mpc.Message, 2*len(parties)),
		subscriptions:    make(map[party.ID]*ps.Subscription, 2*len(parties)),
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
		sub, err := node.Subscribe(topicKey(party.ID(node.PeerID())))
		if err != nil {
			panic(err)
		}
		n.subscriptions[party.ID(node.PeerID())] = sub
		go handleSubscription(party.ID(node.PeerID()), sub, n)
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
		if msg.IsFor(party.ID(node.PeerID())) {
			if err := node.Publish(topicKey(party.ID(node.PeerID())), bz); err != nil {
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
		sub.Cancel()
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
		if msg.IsFor(party.ID(node.PeerID())) {
			return topicKey(party.ID(node.PeerID()))
		}
	}
	return ""
}

func (n *onlineNetwork) getFromNode(msg *mpc.Message) common.P2PNode {
	for _, node := range n.nodes {
		if msg.From == party.ID(node.PeerID()) {
			return node
		}
	}
	return nil
}

// A goroutine that is listening for messages on the topic handler.
func handleSubscription(id party.ID, sub *ps.Subscription, n *onlineNetwork) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			return err
		}
		m := &mpc.Message{}
		err = m.UnmarshalBinary(msg.Data)
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
