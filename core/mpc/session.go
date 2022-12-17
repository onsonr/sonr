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
	id           party.ID
	peers        []party.ID
	doneChan     chan bool
	errorChan    chan error
	msgInChan    chan *mpc.Message
	selfNode     *node.Node
	topicHandler node.TopicHandler
	poolSize     int
}

// It creates a new session object, which is a struct that contains the node, the id of the party, the
// ids of the other parties, the size of the pool, and a channel for incoming messages
func NewSession(n *node.Node, id party.ID, ids []party.ID) (*Session, error) {

	return &Session{
		peers:     ids,
		id:        id,
		selfNode:  n,
		poolSize:  0,
		msgInChan: make(chan *mpc.Message),
		doneChan:  make(chan bool, 1),
		errorChan: make(chan error, 1),
	}, nil
}

// Running the protocol.
func (s *Session) RunProtocol(create mpc.StartFunc, pid protocol.ID, sessionID []byte) (interface{}, error) {
	// Setup Peer-to-Peer Stream
	s.initTopic()
	s.selfNode.HandleProtocol(pid, s.handlePrivateShardStream)

	// Setup Start Function
	handler, err := mpc.NewMultiHandler(create, nil)
	if err != nil {
		return nil, err
	}

	// Message handling loop
	go func(done chan bool) {
		for {
			select {
			// Message to be sent to other participants
			case msgOut, ok := <-handler.Listen():
				// a closed channel indicates that the protocol has finished executing
				if !ok {
					s.doneChan <- true
					return
				}
				s.publishOutMsg(msgOut, pid)
			case msgI := <-s.msgInChan:
				if !handler.CanAccept(msgI) {
					// basic header validation failed, the message may be intended for a different protocol execution.
					continue
				}
				handler.Accept(msgI)
			}
		}
	}(s.doneChan)

	// Wait for the protocol to finish
	<-s.doneChan
	return handler.Result()
}

func (s *Session) initTopic() error {
	th, err := s.selfNode.Subscribe(fmt.Sprintf("/sonr/v0.2.0/mpc/keygen"))
	if err != nil {
		return err
	}
	s.topicHandler = th
	go s.handleTopicSubscription()
	return nil
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
				continue
			}
			s.msgInChan <- msg
		}
	}
}

// Publishing the message to the topic handler.
func (s *Session) publishOutMsg(msgOut *mpc.Message, pid protocol.ID) {
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
	}
	for _, id := range s.peers {
		if msgOut.IsFor(id) {
			s.selfNode.Send(string(id), bz, pid)
		}
	}
}
