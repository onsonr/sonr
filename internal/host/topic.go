package host

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// ExchangeServiceArgs ExchangeArgs is Peer protobuf
type ExchangeServiceArgs struct {
	Peer []byte
}

// ExchangeServiceResponse ExchangeResponse is also Peer protobuf
type ExchangeServiceResponse struct {
	Peer []byte
}

// ExchangeService Service Struct
type ExchangeService struct {
	// Current Data
	call TopicHandler
	user *md.User
}

type TopicHandler interface {
	OnEvent(*md.LobbyEvent)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnResponded(inv *md.InviteRequest)
}

type TopicManager struct {
	ctx          context.Context
	host         HostNode
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	eventHandler *pubsub.TopicEventHandler
	user         *md.User

	events    chan *md.LobbyEvent
	exchange  *ExchangeService
	handler   TopicHandler
	topicType md.TopicType
}

// NewLocal ^ Create New Contained Topic Manager ^ //
func (h *hostNode) JoinTopic(ctx context.Context, u *md.User, name string, th TopicHandler) (*TopicManager, *md.SonrError) {
	// Join Topic
	topic, err := h.pubsub.Join(name)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_JOIN)
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_SUB)
	}

	// Create Topic Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_HANDLER)
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		handler:      th,
		user:         u,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		topicType:    md.TopicType_LOCAL,
		events:       make(chan *md.LobbyEvent, util.TOPIC_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Start Exchange Server
	exchangeServer := rpc.NewServer(h.Host(), util.AUTH_PROTOCOL)
	psv := ExchangeService{
		user: u,
		call: th,
	}

	// Register Service
	err = exchangeServer.RegisterName(util.EXCHANGE_RPC_SERVICE, &psv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_RPC)
	}

	// Set Service
	mgr.exchange = &psv
	go mgr.handleTopicEvents(context.Background())
	go mgr.handleTopicMessages(context.Background())
	go mgr.processTopicMessages(context.Background())
	return mgr, nil
}

// FindPeerInTopic @ Helper: Find returns Pointer to Peer.ID and Peer
func (tm *TopicManager) FindPeerInTopic(q string) (peer.ID, error) {
	// Iterate through Topic Peers
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return id, nil
		}
	}
	return "", errors.New("Peer ID was not found in topic")
}

// PushEvent @ Send Updated Lobby
func (tm *TopicManager) PushEvent(event *md.LobbyEvent) {
	tm.handler.OnEvent(event)
}

// Publish @ Publish message to specific peer in topic
func (tm *TopicManager) Publish(msg *md.LobbyEvent) error {
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

// HasPeer @ Helper: ID returns ONE Peer.ID in Topic
func (tm *TopicManager) HasPeer(q string) bool {
	// Iterate through PubSub in topic
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// IsLocal @ Check if Local Topic
func (tm *TopicManager) IsLocal() bool {
	if tm.topicType == md.TopicType_LOCAL {
		return true
	}
	return false
}

// Exchange @ Starts Exchange on Local Peer Join
func (tm *TopicManager) Exchange(id peer.ID, peerBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host.Host(), util.EXCHANGE_PROTOCOL)
	var reply ExchangeServiceResponse
	var args ExchangeServiceArgs

	// Set Args
	args.Peer = peerBuf

	// Call to Peer
	err := exchClient.Call(id, util.EXCHANGE_RPC_SERVICE, util.EXCHANGE_METHOD_EXCHANGE, args, &reply)
	if err != nil {
		return err
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Peer, remotePeer)

	// Send Error
	if err != nil {
		return err
	}

	// Update Peer with new data
	tm.PushEvent(md.NewJoinLocalEvent(remotePeer))
	return nil
}

// ExchangeWith # Calls Exchange on Local Lobby Peer
func (ts *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeServiceArgs, reply *ExchangeServiceResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers with Lobby
	ts.call.OnEvent(md.NewJoinLocalEvent(remotePeer))

	// Set Message data and call done
	buf, err := ts.user.Peer.Buffer()
	if err != nil {
		return err
	}
	reply.Peer = buf
	return nil
}

// # handleTopicEvents: listens to Pubsub Events for topic
func (tm *TopicManager) handleTopicEvents(ctx context.Context) {
	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := tm.eventHandler.NextPeerEvent(ctx)
		if err != nil {
			tm.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if lobEvent.Type == pubsub.PeerJoin && lobEvent.Peer != tm.host.ID() {
			pbuf, err := tm.user.GetPeer().Buffer()
			if err != nil {
				continue
			}
			err = tm.Exchange(lobEvent.Peer, pbuf)
			if err != nil {
				continue
			}
		} else if lobEvent.Type == pubsub.PeerLeave {
			tm.PushEvent(md.NewExitLocalEvent(lobEvent.Peer.String()))
		}
		md.GetState().NeedsWait()
	}
}

// # handleTopicMessages: listens for messages on pubsub topic subscription
func (tm *TopicManager) handleTopicMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := tm.subscription.Next(ctx)
		if err != nil {
			return
		}

		// Only forward messages delivered by others
		if tm.user.GetPeer().IsSamePeerID(msg.ReceivedFrom) {
			continue
		}

		// Check Lobby Type
		m := &md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if tm.HasPeer(m.Id) {
			tm.events <- m
		}
		md.GetState().NeedsWait()
	}
}

// # processTopicMessages: pulls messages from channel that have been handled
func (tm *TopicManager) processTopicMessages(ctx context.Context) {
	for {
		select {
		// @ Local Event Channel Updated
		case m := <-tm.events:
			tm.PushEvent(m)
		case <-ctx.Done():
			return
		}
		md.GetState().NeedsWait()
	}
}
