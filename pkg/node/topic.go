package node

import (
	"context"
	"errors"
	"log"

	sentry "github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/pkg/data"
	dt "github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

const ChatRoomBufSize = 128

// ExchangeArgs is Peer protobuf
type ExchangeArgs struct {
	Lobby []byte
	Peer  []byte
}

// ExchangeResponse is also Peer protobuf
type ExchangeResponse struct {
	Data []byte
}

// Service Struct
type ExchangeService struct {
	GetUser   dt.ReturnPeer
	SyncLobby dt.SyncLobby
}

type TopicManager struct {
	ctx          context.Context
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	handler      *pubsub.TopicEventHandler
	lobby        *data.Lobby

	topicPoint string
	exchange   *ExchangeService
	protocol   protocol.ID
	messages   chan *md.LobbyEvent

	returnPeer    dt.ReturnPeer
	returnPeerBuf dt.ReturnBuf
}

// ^ Create New Contained Topic Manager ^ //
func (n *Node) JoinTopic(name string, protocol protocol.ID, gp dt.ReturnPeer, gpb dt.ReturnBuf) (*TopicManager, error) {
	// Join Topic
	topic, err := n.pubsub.Join(name)
	if err != nil {
		return nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Topic Handler
	handler, err := topic.EventHandler()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		handler:       handler,
		lobby:         data.NewLobby(name, gp(), n.call.Refreshed),
		messages:      make(chan *md.LobbyEvent, ChatRoomBufSize),
		protocol:      protocol,
		subscription:  sub,
		topic:         topic,
		topicPoint:    name,
		returnPeer:    gp,
		returnPeerBuf: gpb,
	}

	// Start Exchange Server
	peersvServer := rpc.NewServer(n.host, protocol)
	psv := ExchangeService{
		SyncLobby: mgr.lobby.Sync,
		GetUser:   gp,
	}

	// Register Service
	err = peersvServer.Register(&psv)
	if err != nil {
		return nil, err
	}

	// Set Service
	mgr.exchange = &psv

	go n.handleTopicEvents(mgr)
	go n.handleTopicMessages(mgr)
	go n.processTopicMessages(mgr)
	return mgr, nil
}

// ^ Calls Invite on Remote Peer ^ //
func (n *Node) Exchange(tm *TopicManager, id peer.ID, pb []byte) {
	// Initialize RPC
	exchClient := rpc.NewClient(n.host, tm.protocol)
	var reply ExchangeResponse
	var args ExchangeArgs

	// Set Args
	args.Lobby = tm.lobby.Buffer()
	args.Peer = pb

	// Call to Peer
	err := exchClient.Call(id, "ExchangeService", "ExchangeWith", args, &reply)
	if err != nil {
		n.call.Error(err, "Exchange")
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Data, remotePeer)

	// Send Error
	if err != nil {
		n.call.Error(err, "Exchange")
	}

	// Update Peer with new data
	tm.lobby.Add(remotePeer)
}

// ^ Calls Invite on Remote Peer ^ //
func (es *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeArgs, reply *ExchangeResponse) error {
	// Peer Data
	remoteLobbyRef := &md.Lobby{}
	err := proto.Unmarshal(args.Lobby, remoteLobbyRef)
	if err != nil {
		return err
	}

	remotePeer := &md.Peer{}
	err = proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers with Lobby
	es.SyncLobby(remoteLobbyRef, remotePeer)

	// Return User Peer
	userPeer := es.GetUser()
	replyData, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Set Message data and call done
	reply.Data = replyData
	return nil
}

// ^ Helper: Find returns Pointer to Peer.ID and Peer ^
func (n *Node) FindPeerInTopic(tm *TopicManager, q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	var p *md.Peer
	var i peer.ID

	// Iterate Through Peers, Return Matched Peer
	for _, peer := range tm.lobby.Peers {
		// If Found Match
		if peer.Id.Peer == q {
			p = peer
		}
	}

	// Validate Peer
	if p == nil {
		return "", nil, errors.New("Peer data was not found in topic.")
	}

	// Iterate through Topic Peers
	for _, id := range n.pubsub.ListPeers(tm.topicPoint) {
		// If Found Match
		if id.String() == q {
			i = id
		}
	}

	// Validate ID
	if i == "" {
		return "", nil, errors.New("Peer ID was not found in topic.")
	}
	return i, p, nil
}

// ^ Helper: ID returns ONE Peer.ID in Topic ^
func (n *Node) HasPeer(tm *TopicManager, q string) bool {
	// Iterate through PubSub in topic
	for _, id := range n.pubsub.ListPeers(tm.topicPoint) {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// ^ Send Direct Message to Peer in Lobby ^ //
func (n *Node) Message(msg string, to string, p *md.Peer) {
	if n.HasPeer(n.local, to) {
		// Inform Lobby
		if err := n.local.Send(&md.LobbyEvent{
			Event:   md.LobbyEvent_MESSAGE,
			From:    p,
			Id:      p.Id.Peer,
			Message: msg,
			To:      to,
		}); err != nil {
			sentry.CaptureException(err)
		}
	}
}

// ^ Send message to specific peer in topic ^
func (tm *TopicManager) Send(msg *md.LobbyEvent) error {
	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = tm.topic.Publish(tm.ctx, bytes)
	if err != nil {
		return err
	}
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Node) Update(p *md.Peer) {
	// Inform Lobby
	if err := n.local.Send(&md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		From:  p,
		Id:    p.Id.Peer,
	}); err != nil {
		sentry.CaptureException(err)
	}
}
