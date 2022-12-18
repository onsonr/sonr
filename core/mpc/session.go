package mpc

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-hq/sonr/internal/node"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	mpc "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// A Session is a struct that contains a party ID, a list of peers, a message channel, a node, a topic
// handler, and a pool size.
// @property id - the id of the session
// @property {[]party.ID} peers - a list of all the peers in the session
// @property msgInChan - This is the channel that the session will use to receive messages from the
// network.
// @property selfNode - This is the node that represents the current party.
// @property topicHandler - This is the topic handler that will be used to handle messages received
// from the network.
// @property {int} poolSize - The number of goroutines that will be used to handle incoming messages.
type Session struct {
	id               party.ID
	peers            []party.ID
	doneChan         chan bool
	errorChan        chan error
	msgInChan        chan *mpc.Message
	selfNode         *node.Node
	protocolId       protocol.ID
	topicHandler     node.TopicHandler
	poolSize         int
	currentRound     int
	finalRoundNumber int
}

// It creates a new session object, which is a struct that contains the node, the id of the party, the
// ids of the other parties, the size of the pool, and a channel for incoming messages
func NewSession(n *node.Node, id party.ID, ids []party.ID, pid protocol.ID) (*Session, error) {
	th, err := n.Subscribe(fmt.Sprintf("/sonr/v0.2.0/mpc/keygen"))
	if err != nil {
		return nil, err
	}

	s := &Session{
		peers:        ids,
		topicHandler: th,
		id:           id,
		protocolId:   pid,
		selfNode:     n,
		poolSize:     0,
		msgInChan:    make(chan *mpc.Message),
		doneChan:     make(chan bool, 1),
		errorChan:    make(chan error, 1),
		currentRound: 0,
	}
	go s.handleTopicSubscription()
	s.selfNode.SetStreamHandler(pid, s.handlePrivateShardStream)
	return s, nil
}

// Running the protocol.
func (s *Session) RunProtocol(create mpc.StartFunc, sessionID []byte) (interface{}, error) {
	// Setup Peer-to-Peer Stream
	s.finalRoundNumber = getTotalRoundsFromCreate(create)

	// Setup Start Function
	handler, err := mpc.NewMultiHandler(create, nil)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}

	// Message handling loop
	for {
		select {
		// Message to be sent to other participants
		case msgOut, ok := <-handler.Listen():
			// a closed channel indicates that the protocol has finished executing
			fmt.Println("Message out number: ", msgOut.RoundNumber)
			if !ok {
				return handler.Result()
			}
			go s.publishOutMsg(msgOut)

			// Update the current round, and check if we are done
			// Checking if the current round is the final round.
			s.currentRound = int(msgOut.RoundNumber)
			if s.currentRound == s.finalRoundNumber {
				s.doneChan <- true
				return handler.Result()
			}
		case msgI := <-s.msgInChan:
			fmt.Println("Message in number: ", msgI.RoundNumber)
			if !handler.CanAccept(msgI) {
				return handler.Result()
			}
			handler.Accept(msgI)
		}
	}
}

// handlePrivateShardStream reads messages from the stream until done or the stream is closed.
func (s *Session) handlePrivateShardStream(stream network.Stream) {
	for {
		select {
		case <-s.doneChan:
			stream.Close()
			return
		default:
			r := msgio.NewReader(stream)

			// Read the next message
			buf, err := r.ReadMsg()
			if err != nil {
				panic(err)
			}

			// Unmarshal the message
			msg := &mpc.Message{}
			err = msg.UnmarshalBinary(buf)
			if err != nil {
				panic(err)
			}

			// Send the message to the handler
			s.msgInChan <- msg
			continue
		}
	}
}

// A goroutine that is listening for messages on the topic handler.
func (s *Session) handleTopicSubscription() {
	for {
		select {
		case msgIn := <-s.topicHandler.Messages():
			msg := &mpc.Message{}
			err := msg.UnmarshalBinary(msgIn)
			if err != nil {
				panic(err)
			}
			if msg.To != s.id {
				continue
			}
			fmt.Println("Received Broadcast Message")
			s.msgInChan <- msg
		}
	}
}

// Publishing the message to the topic handler.
func (s *Session) publishOutMsg(msgOut *mpc.Message) {
	// ensure this message is reliably broadcast
	bz, err := msgOut.MarshalBinary()
	if err != nil {
		return
	}
	if msgOut.Broadcast {
		err = s.topicHandler.Publish(bz)
		if err != nil {
			return
		}
	} else {
		for _, id := range s.peers {
			if msgOut.IsFor(id) {
				s.selfNode.Send(string(id), bz, s.protocolId)
			}
		}
	}
}
