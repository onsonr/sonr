package mpc

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-hq/sonr/pkg/node"

	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	mpc "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

const (
	K_RENDENZVOUS = "/sonr/v0.2.0/mpc"
	K_PROTOCOL_ID = protocol.ID("mpc/cmp-keygen/1.0.0")
	K_PUBSUB      = "#sonr-mpc-keygen"
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
	ctx              context.Context
	id               party.ID
	vaultId          party.ID
	selfNode         *node.Motor
	protocolId       protocol.ID
	outboundMsgs     chan *mpc.Message
	inboundMsgs      chan *mpc.Message
	routingDiscovery *drouting.RoutingDiscovery
}

// It creates a new session object, which is a struct that contains the node, the id of the party, the
// ids of the other parties, the size of the pool, and a channel for incoming messages
func NewDualSession(n *node.Motor, vault party.ID, pid protocol.ID) (*Session, error) {
	ctx := context.Background()

	routingDiscovery := drouting.NewRoutingDiscovery(n.Node.DHTClient)
	dutil.Advertise(ctx, routingDiscovery, K_RENDENZVOUS)

	s := &Session{
		ctx:              ctx,
		id:               party.ID(n.ID()),
		vaultId:          vault,
		protocolId:       pid,
		selfNode:         n,
		outboundMsgs:     make(chan *mpc.Message),
		inboundMsgs:      make(chan *mpc.Message),
		routingDiscovery: routingDiscovery,
	}

	s.selfNode.SetStreamHandler(pid, s.readProtocolStream)
	return s, nil
}

// Running the protocol.
func (s *Session) RunProtocol(create mpc.StartFunc, sessionID []byte, leader bool) (interface{}, error) {
	// Setup Start Function
	handler, err := mpc.NewTwoPartyHandler(create, nil, false)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}

	peerChan, err := s.routingDiscovery.FindPeers(s.ctx, K_RENDENZVOUS)
	if err != nil {
		fmt.Println("Error finding peers")
		return nil, err
	}

	// Get the first peer
	for p := range peerChan {
		if p.ID == s.selfNode.ID() {
			continue
		}
		fmt.Println("Found peer:", p)
		// Connect to the peer
		err := s.selfNode.Host.Connect(s.ctx, p)
		if err != nil {
			fmt.Println("Connection failed:", err)
			continue
		}

		// Open a stream with the peer
		stream, err := s.selfNode.Host.NewStream(s.ctx, peer.ID(p.ID), s.protocolId)
		if err != nil {
			fmt.Println("Connection failed:", err)
			continue
		} else {
			go s.writeProtocolStream(stream)
			break
		}
	}

	for {
		select {
		case msg, ok := <-handler.Listen():
			if !ok {
				return handler.Result()
			}
			s.outboundMsgs <- msg
		case msg := <-s.inboundMsgs:
			handler.Accept(msg)
		}
	}
}

// handlePrivateShardStream reads messages from the stream until done or the stream is closed.
func (s *Session) readProtocolStream(stream network.Stream) {
	// Read the message from the stream
	rd := msgio.NewReader(stream)
	for {
		data, err := rd.ReadMsg()
		if err != nil {
			fmt.Println("Error reading from stream")
			return
		}

		// Unmarshal the message
		msg := &mpc.Message{}
		err = msg.UnmarshalBinary(data)
		if err != nil {
			fmt.Println("Error unmarshaling message")
			return
		}

		s.inboundMsgs <- msg
	}
}

func (s *Session) writeProtocolStream(stream network.Stream) {
	for {
		msg := <-s.outboundMsgs
		// Marshal the message
		data, err := msg.MarshalBinary()
		if err != nil {
			fmt.Println("Error marshaling message")
			return
		}

		// Write the message to the stream
		wr := msgio.NewWriter(stream)
		_, err = wr.Write(data)
		if err != nil {
			fmt.Println("Error writing to stream")
			return
		}
	}
}
