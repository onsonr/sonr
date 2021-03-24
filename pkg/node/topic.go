package node

import (
	"context"
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/data"
	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type TopicManager struct {
	ctx          context.Context
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	handler      *pubsub.TopicEventHandler
	lobby        *data.Lobby
	topicPoint   string
	exchange     *ExchangeService
	exchangePID  protocol.ID
}

// ^ Create New Contained Topic Manager ^ //
func (n *Node) NewTopicManager(pointName string) (*TopicManager, error) {
	// Join Topic
	topic, err := n.pubsub.Join(pointName)
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
		ctx:          n.ctx,
		topic:        topic,
		handler:      handler,
		subscription: sub,
		lobby:        data.NewLobby(n.Peer(), n.call.Refreshed),
		topicPoint:   pointName,
		exchangePID:  n.router.Exchange(pointName),
	}

	// Start Exchange Server
	peersvServer := rpc.NewServer(n.host, n.router.Exchange(pointName))
	psv := ExchangeService{
		SyncLobby: mgr.lobby.Sync,
		GetUser:   n.Peer,
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
	return mgr, nil
}

// ^ Send Peer Update to Topic ^
func (tm *TopicManager) Update(data *md.Peer) error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		From:  data,
		Id:    data.Id.Peer,
	}

	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(&event)
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

// ^ Send message to specific peer in topic ^
func (tm *TopicManager) Message(msg string, to string, data *md.Peer) error {
	// Create Lobby Event
	event := md.LobbyEvent{
		Event:   md.LobbyEvent_MESSAGE,
		From:    data,
		Id:      data.Id.Peer,
		Message: msg,
		To:      to,
	}

	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(&event)
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

// ^ Helper: Find returns Pointer to Peer.ID and Peer ^
func (n *Node) GetPeer(tm *TopicManager, q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	peer := n.FindPeerFromTopic(tm, q)
	id := n.FindIDFromTopic(tm, q)

	if peer == nil || id == "" {
		return "", nil, errors.New("Search Error, peer was not found in map.")
	}

	return id, peer, nil
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

// ^ Helper: ID returns ONE Peer.ID in PubSub ^
func (n *Node) FindIDFromTopic(tm *TopicManager, q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range n.pubsub.ListPeers(tm.topicPoint) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	return ""
}

// ^ Helper: Peer returns ONE Peer in Lobby ^
func (n *Node) FindPeerFromTopic(tm *TopicManager, q string) *md.Peer {
	// Iterate Through Peers, Return Matched Peer
	for _, peer := range tm.lobby.Peers {
		// If Found Match
		if peer.Id.Peer == q {
			return peer
		}
	}
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (n *Node) Exchange(tm *TopicManager, id peer.ID) {
	// Initialize RPC
	rpcClient := rpc.NewClient(n.host, tm.exchangePID)
	var reply ExchangeResponse
	var args ExchangeArgs

	// Set Args
	args.Lobby = tm.lobby.Buffer()
	args.Peer = n.PeerBuf()

	// Call to Peer
	err := rpcClient.Call(id, "ExchangeService", "ExchangeWith", args, &reply)
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

// ****************** //
// ** GRPC Service ** //
// ****************** //
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

// ^ Calls Invite on Remote Peer ^ //
func (ps *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeArgs, reply *ExchangeResponse) error {
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
	ps.SyncLobby(remoteLobbyRef, remotePeer)

	// Return User Peer
	userPeer := ps.GetUser()
	replyData, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Set Message data and call done
	reply.Data = replyData
	return nil
}
