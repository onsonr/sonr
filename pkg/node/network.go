package node

import (
	"context"
	"fmt"
	"sync"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	mpc "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// `Network` is a struct that contains a `context.Context`, a `node.Node`, a list of `peer.ID`s, a list
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
type Network struct {
	ctx       context.Context
	selfNode  *Node
	selfParty party.ID
	nodes     []*Node

	mtx       sync.Mutex
	doneChan  chan bool
	errorChan chan error
	msgInChan chan *mpc.Message

	subscriptions map[party.ID]icore.PubSubSubscription
}

// It creates a new network object, assigns the subscriptions, and returns the network object
func NewNetwork(ctx context.Context, n *Node, peerIds []peer.ID) (*Network, error) {
	// Convert the peer IDs to party IDs.
	partyIds := make([]party.ID, len(peerIds))
	for i, id := range peerIds {
		partyIds[i] = party.ID(id)
	}
	partyIds = append(partyIds, party.ID(n.ID()))

	// Create the network object.
	net := &Network{
		ctx:       context.Background(),
		selfParty: party.ID(n.ID()),
		selfNode:  n,
		doneChan:  make(chan bool),
		errorChan: make(chan error),
		msgInChan: make(chan *mpc.Message),
	}
	return net, nil
}

// Closing the subscriptions and returning the done channel.
func (n *Network) Done() <-chan bool {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	for id, sub := range n.subscriptions {
		sub.Close()
		delete(n.subscriptions, id)
	}
	return n.doneChan
}

// Returning the channel that the network will use to send messages to the application.
func (n *Network) Next() <-chan *mpc.Message {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	return n.msgInChan
}

// Sending a message to the network.
func (n *Network) Send(msg *mpc.Message) {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	bz, err := msg.MarshalBinary()
	if err != nil {
		n.errorChan <- err
	}

	for _, peerNode := range n.nodes {
		if msg.IsFor(party.ID(peerNode.ID())) {
			err = n.selfNode.Publish(topicKey(party.ID(peerNode.ID())), bz)
			if err != nil {
				n.errorChan <- err
			}
		}
	}
}

// Starting the network.
func (n *Network) Start() {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	go func() {
		err := n.selfNode.Subscribe(n.ctx, topicKey(n.selfParty), n.handleSubscription)
		if err != nil {
			n.errorChan <- err
		}
	}()
	<-n.doneChan
}

//
// Private methods
//

// A goroutine that is listening for errors on the error channel.
func handleErrors(n *Network) {
	for {
		select {
		case err := <-n.errorChan:
			fmt.Println("Error while running network: ", err)
		case <-n.doneChan:
			return
		}
	}
}

// A goroutine that is listening for messages on the topic handler.
func (n *Network) handleSubscription(topic string, msg icore.PubSubMessage) error {
	msgIn := &mpc.Message{}
	err := msgIn.UnmarshalBinary(msg.Data())
	if err != nil {
		return err
	}
	n.msgInChan <- msgIn
	return nil
}

// topicKey returns the modified topic key for the given party ID.
func topicKey(id party.ID) string {
	return fmt.Sprintf("%s/mpc/cmp-keygen", string(id))
}
