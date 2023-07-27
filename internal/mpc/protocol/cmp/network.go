package cmp

import (
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	pv0 "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// A Network is a channel that sends messages to parties and receives messages from parties.
type Network interface {
	// Ls returns a list of peers that are connected to the network.
	Ls() []party.ID

	// A function that takes in a party ID and returns a channel of protocol messages.
	Next(id party.ID) <-chan *pv0.Message

	// Sending a message to the network.
	Send(msg *pv0.Message)

	// A channel that is closed when the party is done with the protocol.
	Done(id party.ID) chan struct{}

	// A function that is called when a party is done with the protocol.
	Quit(id party.ID)

	// IsOnlineNetwork returns true if the network is an online network.
	IsOnlineNetwork() bool
}

// EnsureSelfIDInGroup ensures that the given self ID is in the given group.
func EnsureSelfIDInGroup(selfID party.ID, group []party.ID) []party.ID {
	if len(selfID) == 0 {
		selfID = party.ID(peer.ID("user1"))
	}
	for _, id := range group {
		if id == selfID {
			return group
		}
	}
	return append(group, selfID)
}

// HandleNetworkProtocol is a helper function that loops over all the parties and calls the given handler.
func HandleNetworkProtocol(id party.ID, h pv0.Handler, network Network) {
	for {
		select {

		// outgoing messages
		case msg, ok := <-h.Listen():
			if !ok {
				<-network.Done(id)
				// the channel was closed, indicating that the protocol is done executing.
				return
			}
			go network.Send(msg)

			// incoming messages
		case msg := <-network.Next(id):
			h.Accept(msg)
		}
	}
}

// It's a network that can be used to simulate offline parties.
// @property parties - a slice of party IDs that are participating in the protocol.
// @property listenChannels - a map of party IDs to channels that will be used to send messages to the
// party.
// @property done - a channel that is closed when the network is closed.
// @property closedListenChan - This channel is used to signal that the network has been closed.
// @property mtx - a mutex to protect the listenChannels map
type offlineNetwork struct {
	parties          party.IDSlice
	listenChannels   map[party.ID]chan *protocol.Message
	done             chan struct{}
	closedListenChan chan *protocol.Message
	mtx              sync.Mutex
}

// It creates a new `OfflineNetwork` object, and initializes it with a list of parties, and a map of
// channels
func NewOfflineNetwork(parties ...party.ID) Network {
	closed := make(chan *protocol.Message)
	close(closed)
	c := &offlineNetwork{
		parties:          parties,
		listenChannels:   make(map[party.ID]chan *protocol.Message, 2*len(parties)),
		closedListenChan: closed,
	}
	return c
}

// Initializing the network.
func (n *offlineNetwork) init() {
	N := len(n.parties)
	for _, id := range n.parties {
		n.listenChannels[id] = make(chan *protocol.Message, N*N)
	}
	n.done = make(chan struct{})
}

// Ls returns a list of parties that are participating in the protocol.
func (n *offlineNetwork) Ls() []party.ID {
	return n.parties
}

// Returning a channel that is used to send messages to the party.
func (n *offlineNetwork) Next(id party.ID) <-chan *protocol.Message {
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

// Sending the message to all the parties.
func (n *offlineNetwork) Send(msg *protocol.Message) {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	for id, c := range n.listenChannels {
		if msg.IsFor(id) && c != nil {
			n.listenChannels[id] <- msg
		}
	}
}

// IsOnlineNetwork returns false.
func (n *offlineNetwork) IsOnlineNetwork() bool {
	return false
}

// Closing the channel that is used to send messages to the party.
func (n *offlineNetwork) Done(id party.ID) chan struct{} {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	if _, ok := n.listenChannels[id]; ok {
		close(n.listenChannels[id])
		delete(n.listenChannels, id)
	}
	if len(n.listenChannels) == 0 {
		close(n.done)
	}
	return n.done
}

// Removing the party from the network.
func (n *offlineNetwork) Quit(id party.ID) {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.parties = n.parties.Remove(id)
}
