package host

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	ps "github.com/libp2p/go-libp2p-pubsub"
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
	OnEvent(*md.TopicEvent)
}

type TopicManager struct {
	ctx          context.Context
	host         HostNode
	topic        *ps.Topic
	subscription *ps.Subscription
	eventHandler *ps.TopicEventHandler
	user         *md.User

	events    chan *md.TopicEvent
	exchange  *ExchangeService
	handler   TopicHandler
	topicData *md.Topic
}

// NewLocal ^ Create New Contained Topic Manager ^ //
func (h *hostNode) JoinTopic(ctx context.Context, u *md.User, topicData *md.Topic, th TopicHandler) (*TopicManager, *md.SonrError) {
	// Join Topic
	name := topicData.GetName()
	topic, err := h.pubsub.Join(name)
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_TOPIC_JOIN)
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_TOPIC_SUB)
	}

	// Create Topic Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_TOPIC_HANDLER)
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		handler:      th,
		user:         u,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		topicData:    topicData,
		events:       make(chan *md.TopicEvent, util.MAX_CHAN_DATA),
		subscription: sub,
		topic:        topic,
	}

	// Start Exchange RPC Server
	exchangeServer := rpc.NewServer(h.Host(), util.EXCHANGE_PROTOCOL)
	esv := ExchangeService{
		user: u,
		call: th,
	}

	// Register Service
	err = exchangeServer.RegisterName(util.EXCHANGE_RPC_SERVICE, &esv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_TOPIC_RPC)
	}

	// Set Service
	mgr.exchange = &esv
	h.topics = append(h.topics, mgr)

	// Handle Events
	go mgr.handleTopicEvents(context.Background())
	go mgr.handleTopicMessages(context.Background())
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

// Publish @ Publish message to specific peer in topic
func (tm *TopicManager) Publish(msg *md.TopicEvent) error {
	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(msg)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Publish to Topic
	err = tm.topic.Publish(tm.ctx, bytes)
	if err != nil {
		md.LogError(err)
		return err
	}
	return nil
}

// Returns Topic Data instance
func (tm *TopicManager) Topic() *md.Topic {
	return tm.topicData
}

// HasPeer Method Checks if Peer ID String is Subscribed to Topic
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

// HasPeer Method Checks if Peer ID is Subscribed to Topic
func (tm *TopicManager) HasPeerID(q peer.ID) bool {
	// Iterate through PubSub in topic
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id == q {
			return true
		}
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
		md.LogError(err)
		return err
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Peer, remotePeer)

	// Send Error
	if err != nil {
		md.LogError(err)
		return err
	}

	// Update Peer with new data
	tm.handler.OnEvent(md.NewJoinEvent(remotePeer))
	return nil
}

// ExchangeWith # Calls Exchange on Local Lobby Peer
func (ts *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeServiceArgs, reply *ExchangeServiceResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Update Peers with Lobby
	ts.call.OnEvent(md.NewJoinEvent(remotePeer))

	// Set Message data and call done
	buf, err := ts.user.Peer.Buffer()
	if err != nil {
		md.LogError(err)
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
		event, err := tm.eventHandler.NextPeerEvent(ctx)
		if err != nil {
			md.LogError(err)
			tm.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if tm.isEventJoin(event) {
			pbuf, err := tm.user.GetPeer().Buffer()
			if err != nil {
				md.LogError(err)
				continue
			}
			err = tm.Exchange(event.Peer, pbuf)
			if err != nil {
				md.LogError(err)
				continue
			}
		} else if tm.isEventExit(event) {
			tm.handler.OnEvent(md.NewExitEvent(event.Peer.String(), tm.topicData))

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
			md.LogError(err)
			return
		}

		// Only forward messages delivered by others
		if tm.isValidMessage(msg) {
			// Unmarshal TopicEvent
			m := &md.TopicEvent{}
			err = proto.Unmarshal(msg.Data, m)
			if err != nil {
				md.LogError(err)
				continue
			}

			// Check Peer is Online, if not ignore
			if m.Peer.GetStatus() == md.Peer_ONLINE {
				tm.handler.OnEvent(m)
			}
		}
		md.GetState().NeedsWait()
	}
}

// # Check if PeerEvent is Join and NOT User
func (tm *TopicManager) isEventJoin(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != tm.host.ID()
}

// # Check if PeerEvent is Exit and NOT User
func (tm *TopicManager) isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave && ev.Peer != tm.host.ID()
}

// # Check if Message is NOT from User
func (tm *TopicManager) isValidMessage(msg *ps.Message) bool {
	return tm.host.ID() != msg.ReceivedFrom && tm.HasPeerID(msg.ReceivedFrom)
}
