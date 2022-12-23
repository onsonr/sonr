package internal

import (
	"context"
	"fmt"

	peer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	mpc "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// A MultiSession is a struct that contains a party ID, a list of peers, a message channel, a node, a topic
// handler, and a pool size.
// @property id - the id of the session
// @property {[]party.ID} peers - a list of all the peers in the session
// @property msgInChan - This is the channel that the session will use to receive messages from the
// network.
// @property selfNode - This is the node that represents the current party.
// @property topicHandler - This is the topic handler that will be used to handle messages received
// from the network.
// @property {int} poolSize - The number of goroutines that will be used to handle incoming messages.
type MultiSession struct {
	ctx      context.Context
	id       party.ID
	selfNode *node.Node
	poolSize int
	network  *Network
}

// It creates a new session object, which is a struct that contains the node, the id of the party, the
// ids of the other parties, the size of the pool, and a channel for incoming messages
func NewMultiSession(n *node.Node, ids []peer.ID) (*MultiSession, error) {
	ctx := context.Background()
	parties := make([]party.ID, len(ids))
	for i, id := range ids {
		parties[i] = party.ID(id)
	}
	net, err := NewNetwork(n, ids)
	if err != nil {
		return nil, err
	}
	s := &MultiSession{
		ctx:      ctx,
		id:       party.ID(n.ID()),
		selfNode: n,
		poolSize: 0,
		network:  net,
	}
	return s, nil
}

// Running the protocol.
func (s *MultiSession) RunProtocol(handler *mpc.MultiHandler) {
	s.network.Await()
	// Message handling loop
	for {
		select {
		// Message to be sent to other participants
		case msgOut, ok := <-handler.Listen():
			// a closed channel indicates that the protocol has finished executing
			fmt.Println("Message out number: ", msgOut.RoundNumber)
			if !ok {
				<-s.network.Done()
				return
			}
			go s.network.Send(msgOut)
		case msgI := <-s.network.Next():
			fmt.Println("Message in number: ", msgI.RoundNumber)
			handler.Accept(msgI)
		}
	}
}

// SelfID returns the ID of the current party.
func (s *MultiSession) SelfID() party.ID {
	return s.id
}

// PartyIds returns the IDs of all the parties in the session.
func (s *MultiSession) PartyIds() []party.ID {
	return s.network.partyIds
}
