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
	call    RoomHandler
	linkers []*md.Peer
	user    *md.User
}

// SyncServiceArgs ExchangeArgs are Peer, Device, and Contact
type SyncServiceArgs struct {
	Contact []byte
	Device  []byte
	Peer    []byte
}

// SyncServiceResponse ExchangeResponse is Member protobuf
type SyncServiceResponse struct {
	Success bool
	Contact []byte
	Device  []byte
	Peer    []byte
}

// SyncService Service Struct
type SyncService struct {
	// Current Data
	call    RoomHandler
	linkers []*md.Peer
	user    *md.User
}

type RoomHandler interface {
	OnRoomEvent(*md.RoomEvent)
}

type RoomManager struct {
	ctx          context.Context
	host         HostNode
	Topic        *ps.Topic
	subscription *ps.Subscription
	eventHandler *ps.TopicEventHandler
	user         *md.User

	events   chan *md.RoomEvent
	exchange *ExchangeService
	handler  RoomHandler
	linkers  []*md.Peer
	room     *md.Room
}

// NewLocal ^ Create New Contained Room Manager ^ //
func (h *hostNode) JoinRoom(ctx context.Context, u *md.User, room *md.Room, th RoomHandler) (*RoomManager, *md.SonrError) {
	// Join Room
	name := room.GetName()
	topic, err := h.pubsub.Join(name)
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_TOPIC_JOIN)
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_TOPIC_SUB)
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_TOPIC_HANDLER)
	}

	// Create Lobby Manager
	mgr := &RoomManager{
		handler:      th,
		user:         u,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		room:         room,
		linkers:      make([]*md.Peer, 0),
		events:       make(chan *md.RoomEvent, util.MAX_CHAN_DATA),
		subscription: sub,
		Topic:        topic,
	}

	// Start Exchange RPC Server
	exchangeServer := rpc.NewServer(h.Host(), util.EXCHANGE_PROTOCOL)
	esv := ExchangeService{
		user:    u,
		call:    th,
		linkers: mgr.linkers,
	}

	// Register Service
	err = exchangeServer.RegisterName(util.EXCHANGE_RPC_SERVICE, &esv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorEvent_TOPIC_RPC)
	}

	// Set Service
	mgr.exchange = &esv

	// Handle Events
	go mgr.handleRoomEvents(context.Background())
	go mgr.handleRoomMessages(context.Background())
	return mgr, nil
}

// FindPeer @ Helper: Find returns Pointer to Peer.ID and Peer
func (tm *RoomManager) FindPeer(q string) (peer.ID, error) {
	// Iterate through Room Peers
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return id, nil
		}
	}
	return "", errors.New("Peer ID was not found in room")
}

// Publish @ Publish message to specific peer in room
func (tm *RoomManager) Publish(msg *md.RoomEvent) error {
	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(msg)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Publish to Room
	err = tm.Topic.Publish(tm.ctx, bytes)
	if err != nil {
		md.LogError(err)
		return err
	}
	return nil
}

// Returns RoomData Data instance
func (tm *RoomManager) RoomData() *md.Room {
	return tm.room
}

func (tm *RoomManager) HasLinker(q string) bool {
	for _, p := range tm.linkers {
		if p.PeerID() == q {
			return true
		}
	}
	return false
}

// HasPeer Method Checks if Peer ID String is Subscribed to Room
func (tm *RoomManager) HasPeer(q string) bool {
	// Iterate through PubSub in room
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// HasPeer Method Checks if Peer ID is Subscribed to Room
func (tm *RoomManager) HasPeerID(q peer.ID) bool {
	// Iterate through PubSub in room
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id == q {
			return true
		}
	}
	return false
}

// Returns List of Linkers in Room
func (tm *RoomManager) ListLinkers() *md.Linkers {
	return &md.Linkers{
		List: tm.linkers,
	}
}

// Exchange @ Starts Exchange on Local Peer Join
func (tm *RoomManager) Exchange(id peer.ID, peerBuf []byte) error {
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
	if remotePeer.Status != md.Peer_PAIRING {
		tm.handler.OnRoomEvent(md.NewJoinEvent(remotePeer))
	} else {
		// Add Linker if Not Present
		if !tm.HasLinker(remotePeer.PeerID()) {
			// Append Linkers
			tm.linkers = append(tm.linkers, remotePeer)
		}
	}
	return nil
}

// ExchangeWith # Calls Exchange on Local Lobby Peer
func (es *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeServiceArgs, reply *ExchangeServiceResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		md.LogError(err)
		return err
	}

	// Update Peers with Lobby
	if remotePeer.Status != md.Peer_PAIRING {
		es.call.OnRoomEvent(md.NewJoinEvent(remotePeer))
	} else {
		// Add Linker if Not Present
		if !es.HasLinker(remotePeer.PeerID()) {
			// Append Linkers
			es.linkers = append(es.linkers, remotePeer)
		}
	}

	// Set Message data and call done
	buf, err := es.user.GetPrimary().Buffer()
	if err != nil {
		md.LogError(err)
		return err
	}
	reply.Peer = buf
	return nil
}

func (es *ExchangeService) HasLinker(q string) bool {
	for _, p := range es.linkers {
		if p.PeerID() == q {
			return true
		}
	}
	return false
}

// Exchange @ Starts Exchange on Local Peer Join
func (tm *RoomManager) Sync(id peer.ID, peerBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host.Host(), util.SYNC_PROTOCOL)
	var reply ExchangeServiceResponse
	var args ExchangeServiceArgs

	// Set Args
	args.Peer = peerBuf

	// Call to Peer
	err := exchClient.Call(id, util.SYNC_RPC_SERVICE, util.EXCHANGE_METHOD_EXCHANGE, args, &reply)
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
	if remotePeer.Status != md.Peer_PAIRING {
		tm.handler.OnRoomEvent(md.NewJoinEvent(remotePeer))
	} else {
		// Add Linker if Not Present
		if !tm.HasLinker(remotePeer.PeerID()) {
			// Append Linkers
			tm.linkers = append(tm.linkers, remotePeer)
		}
	}
	return nil
}

// ExchangeWith # Calls Exchange on Local Lobby Peer
func (es *SyncService) SyncWith(ctx context.Context, args SyncServiceArgs, reply *SyncServiceResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		md.LogError(err)
		return err
	}

	es.call.OnRoomEvent(md.NewJoinEvent(remotePeer))

	// Set Message data and call done
	buf, err := es.user.GetPrimary().Buffer()
	if err != nil {
		md.LogError(err)
		return err
	}
	reply.Peer = buf
	return nil
}

// # handleRoomEvents: listens to Pubsub Events for room
func (tm *RoomManager) handleDeviceEvents(ctx context.Context) {
	// Loop Events
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
			pbuf, err := tm.user.GetPrimary().Buffer()
			if err != nil {
				md.LogError(err)
				continue
			}
			err = tm.Sync(event.Peer, pbuf)
			if err != nil {
				md.LogError(err)
				continue
			}
		} else if tm.isEventExit(event) {
			tm.handler.OnRoomEvent(md.NewExitEvent(event.Peer.String(), tm.room))

		}
		md.GetState().NeedsWait()
	}
}

// # handleRoomMessages: listens for messages on pubsub room subscription
func (tm *RoomManager) handleDeviceMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := tm.subscription.Next(ctx)
		if err != nil {
			md.LogError(err)
			return
		}

		// Only forward messages delivered by others
		if tm.isValidMessage(msg) {
			// Unmarshal RoomEvent
			m := &md.RoomEvent{}
			err = proto.Unmarshal(msg.Data, m)
			if err != nil {
				md.LogError(err)
				continue
			}

			// Check Peer is Online, if not ignore
			if m.Peer.GetStatus() == md.Peer_ONLINE {
				tm.handler.OnRoomEvent(m)
			} else if m.Peer.GetStatus() == md.Peer_PAIRING {
				// Validate Linker not Already Set
				if !tm.HasLinker(m.Peer.PeerID()) {
					// Append Linkers
					tm.linkers = append(tm.linkers, m.Peer)
				}
			}
		}
		md.GetState().NeedsWait()
	}
}

// # handleRoomEvents: listens to Pubsub Events for room
func (tm *RoomManager) handleRoomEvents(ctx context.Context) {
	// Loop Events
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
			pbuf, err := tm.user.GetPrimary().Buffer()
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
			tm.handler.OnRoomEvent(md.NewExitEvent(event.Peer.String(), tm.room))

		}
		md.GetState().NeedsWait()
	}
}

// # handleRoomMessages: listens for messages on pubsub room subscription
func (tm *RoomManager) handleRoomMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := tm.subscription.Next(ctx)
		if err != nil {
			md.LogError(err)
			return
		}

		// Only forward messages delivered by others
		if tm.isValidMessage(msg) {
			// Unmarshal RoomEvent
			m := &md.RoomEvent{}
			err = proto.Unmarshal(msg.Data, m)
			if err != nil {
				md.LogError(err)
				continue
			}

			// Check Peer is Online, if not ignore
			if m.Peer.GetStatus() == md.Peer_ONLINE {
				tm.handler.OnRoomEvent(m)
			} else if m.Peer.GetStatus() == md.Peer_PAIRING {
				// Validate Linker not Already Set
				if !tm.HasLinker(m.Peer.PeerID()) {
					// Append Linkers
					tm.linkers = append(tm.linkers, m.Peer)
				}
			}
		}
		md.GetState().NeedsWait()
	}
}

// # Check if PeerEvent is Join and NOT User
func (tm *RoomManager) isEventJoin(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != tm.host.ID()
}

// # Check if PeerEvent is Exit and NOT User
func (tm *RoomManager) isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave && ev.Peer != tm.host.ID()
}

// # Check if Message is NOT from User
func (tm *RoomManager) isValidMessage(msg *ps.Message) bool {
	return tm.host.ID() != msg.ReceivedFrom && tm.HasPeerID(msg.ReceivedFrom)
}
