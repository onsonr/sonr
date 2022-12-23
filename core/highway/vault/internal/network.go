package internal

import (
	"context"
	"fmt"
	"sync"
	"time"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/pkg/node"
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
	selfNode  *node.Node
	selfParty party.ID
	peerIds   []peer.ID
	partyIds  []party.ID

	mtx       sync.Mutex
	doneChan  chan bool
	errorChan chan error
	msgInChan chan *mpc.Message

	subscriptions map[party.ID]icore.PubSubSubscription
}

// It creates a new network object, assigns the subscriptions, and returns the network object
func NewNetwork(n *node.Node, peerIds []peer.ID) (*Network, error) {
	// Convert the peer IDs to party IDs.
	partyIds := make([]party.ID, len(peerIds))
	for i, id := range peerIds {
		partyIds[i] = party.ID(id)
	}

	// Create the network object.
	net := &Network{
		ctx:       context.Background(),
		selfParty: party.ID(n.ID()),
		selfNode:  n,
		peerIds:   peerIds,
		partyIds:  partyIds,
		doneChan:  make(chan bool),
		errorChan: make(chan error),
		msgInChan: make(chan *mpc.Message),
	}

	// Assign the subscriptions.
	idStrs := make([]string, len(net.peerIds))
	for i, id := range net.peerIds {
		idStrs[i] = id.String()
	}

	// Subscribe to the topic.
	sub, err := n.Subscribe(topicKey(net.selfParty))
	if err != nil {
		return nil, err
	}
	go net.handleSubscription(net.ctx, sub)
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

	for _, id := range n.partyIds {
		if msg.IsFor(id) {
			err = n.selfNode.Publish(topicKey(id), bz)
			if err != nil {
				n.errorChan <- err
			}
		}
	}
}

func (n *Network) Await() {
	fmt.Println("Awaiting for peers to connect...")
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			peers, err := n.selfNode.ListPeers(topicKey(n.selfParty))
			if err != nil {
				n.errorChan <- err
			}
			if len(peers) == len(n.peerIds)-1 {
				fmt.Println("All peers connected.")
				return
			}
			fmt.Printf("Connected to %d peers out of %d.\n", len(peers), len(n.peerIds)-1)
		}
	}
}

// A goroutine that is listening for errors on the error channel.
func (n *Network) handleErrors() {
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
func (n *Network) handleSubscription(ctx context.Context, sub icore.PubSubSubscription) {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			n.errorChan <- err
		}
		msgIn := &mpc.Message{}
		err = msgIn.UnmarshalBinary(msg.Data())
		if err != nil {
			n.errorChan <- err
		}
		n.msgInChan <- msgIn

		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}
}

// topicKey returns the modified topic key for the given party ID.
func topicKey(id party.ID) string {
	return fmt.Sprintf("%s/mpc/cmp-keygen", string(id))
}
