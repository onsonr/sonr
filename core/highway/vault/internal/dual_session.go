package internal

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/sonr-hq/sonr/pkg/node"

	"github.com/taurusgroup/multi-party-sig/pkg/party"
	mpc "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// A DualSession is a struct that contains a party ID, a list of peers, a message channel, a node, a topic
// handler, and a pool size.
// @property id - the id of the session
// @property {[]party.ID} peers - a list of all the peers in the session
// @property msgInChan - This is the channel that the session will use to receive messages from the
// network.
// @property selfNode - This is the node that represents the current party.
// @property topicHandler - This is the topic handler that will be used to handle messages received
// from the network.
// @property {int} poolSize - The number of goroutines that will be used to handle incoming messages.
type DualSession struct {
	id         party.ID
	vaultId    party.ID
	selfNode   *node.Node
	protocolId protocol.ID
}

// It creates a new session object, which is a struct that contains the node, the id of the party, the
// ids of the other parties, the size of the pool, and a channel for incoming messages
func NewDualSession(n *node.Node, vault party.ID, pid protocol.ID) (*DualSession, error) {
	s := &DualSession{
		id:         party.ID(n.ID()),
		vaultId:    vault,
		protocolId: pid,
		selfNode:   n,
	}
	return s, nil
}

// Running the protocol.
func (s *DualSession) RunProtocol(create mpc.StartFunc, sessionID []byte, leader bool) (interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Setup Start Function
	handler, err := mpc.NewTwoPartyHandler(create, nil, false)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}

	peerChan, err := s.selfNode.PubSub().Subscribe(ctx, "testch")
	if err != nil {
		fmt.Println("Error finding peers")
		return nil, err
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			select {
			case msg, ok := <-handler.Listen():
				if !ok {
					fmt.Println("Handler closed. Protocol finished.")
					cancel()
					return
				}
				fmt.Println("Sending message")
				bz, _ := msg.MarshalBinary()
				err = s.selfNode.PubSub().Publish(ctx, "testch", bz)
				switch err {
				case nil:
				case context.Canceled:
					return
				default:
					cancel()
					return
				}
			}
		}
	}()
	// Wait for the sender to finish before we return.
	// Otherwise, we can get random errors as publish fails.
	defer func() {
		cancel()
		<-done
	}()

	for {
		msg, err := peerChan.Next(ctx)
		switch err {
		case nil:
			fmt.Println("Received message")
			// Unmarshal the message
			fromMsg := &mpc.Message{}
			err = fromMsg.UnmarshalBinary(msg.Data())
			if err != nil {
				fmt.Println("Error unmarshaling message")
				return nil, err
			}
			handler.Accept(fromMsg)
		case context.Canceled:
			return nil, err
		default:
			cancel()
			return nil, err
		}

		select {
		case <-done:
			return handler.Result()
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
