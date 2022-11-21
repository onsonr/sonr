package mpc

import (
	"sync"

	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/protocol"
	"github.com/sonr-io/sonr/internal/node"
)

// Network simulates a point-to-point network between different parties using Go channels.
// The same network is used by all processes, and can be reused for different protocols.
// When used with test.Handler, no interaction from the user is required beyond creating the network.
type Network struct {
	parties          party.IDSlice
	listenChannels   map[party.ID]chan *protocol.Message
	done             chan struct{}
	closedListenChan chan *protocol.Message
	mtx              sync.Mutex
}

func getNetwork(pids party.IDSlice) *Network {
	closed := make(chan *protocol.Message)
	close(closed)
	c := &Network{
		parties:          pids,
		listenChannels:   make(map[party.ID]chan *protocol.Message, 2*len(pids)),
		closedListenChan: closed,
	}
	return c
}

func (n *Network) init() {
	N := len(n.parties)
	for _, id := range n.parties {
		n.listenChannels[id] = make(chan *protocol.Message, N*N)
	}
	n.done = make(chan struct{})
}

func (n *Network) Next(id party.ID) <-chan *protocol.Message {
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

func (n *Network) Send(msg *protocol.Message) {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	for id, c := range n.listenChannels {
		if msg.IsFor(id) && c != nil {
			n.listenChannels[id] <- msg
		}
	}
}

func (n *Network) Done(id party.ID) chan struct{} {
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

func (n *Network) Quit(id party.ID) {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.parties = n.parties.Remove(id)
}

// handlerLoop is a helper function that loops over all the parties and calls the given handler.
func handlerLoop(id party.ID, h protocol.Handler, network *Network) {
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

func handlerLoopChannel(id party.ID, h protocol.Handler, channel *node.Channel) {
	for {
		select {

		// outgoing messages
		case msg, ok := <-h.Listen():
			if !ok {
				channel.Close()
				// the channel was closed, indicating that the protocol is done executing.
				return
			}
			buf, _ := msg.MarshalBinary()
			go channel.Send(buf)

			// incoming messages
		case msg := <-channel.NextMessage():
			var m protocol.Message
			_ = m.UnmarshalBinary(msg.Data)
			h.Accept(&m)
		}
	}
}
