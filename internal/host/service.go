package host

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	psub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Join New Topic with Name ^
func (h *hostNode) Join(name string) (*psub.Topic, *psub.Subscription, *psub.TopicEventHandler, *md.SonrError) {
	// Join Topic
	topic, err := h.pubsub.Join(name)
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_JOIN)
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_SUB)
	}

	// Create Topic Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_HANDLER)
	}
	return topic, sub, handler, nil
}

// ^ Set Stream Handler for Host ^
func (h *hostNode) HandleStream(pid protocol.ID, handler network.StreamHandler) {
	h.host.SetStreamHandler(pid, handler)
}

// ^ Start Stream for Host ^
func (h *hostNode) StartStream(p peer.ID, pid protocol.ID) (network.Stream, error) {
	return h.host.NewStream(h.ctxHost, p, pid)
}

type GlobalTopic interface {
	FindPeerID(string) (peer.ID, error)
}

type globalTopic struct {
	GlobalTopic
	ctx          context.Context
	global       *md.Global
	topic        *psub.Topic
	handler      *psub.TopicEventHandler
	subscription *psub.Subscription
	host         host.Host
	service      *GlobalService
}

func (hn *hostNode) StartGlobal(SName string) (GlobalTopic, *md.SonrError) {
	// Join Global Topic
	t, s, h, err := hn.Join("sonr-global")
	if err != nil {
		return &globalTopic{}, err
	}

	// Create Global Struct
	gt := &globalTopic{
		ctx:          hn.ctxHost,
		topic:        t,
		subscription: s,
		handler:      h,
		host:         hn.host,
		global: &md.Global{
			Peers:      make(map[string]string),
			UserPeerID: hn.host.ID().String(),
			SName:      SName,
		},
	}

	// Start Exchange Server
	globalExServer := rpc.NewServer(hn.host, globalProtocol)
	gsv := GlobalService{
		global: gt.global,
	}

	// Register Service
	serr := globalExServer.Register(&gsv)
	if serr != nil {
		return nil, md.NewError(serr, md.ErrorMessage_TOPIC_RPC)
	}

	// Return Global
	gt.service = &gsv
	go gt.handleEvents()
	return gt, nil
}

// ExchangeArgs is Peer protobuf
type GlobalServiceArgs struct {
	Global []byte
}

// ExchangeResponse is also Peer protobuf
type GlobalServiceResponse struct {
	Global []byte
}

// Service Struct
type GlobalService struct {
	global *md.Global
}

// ^ Calls Invite on Remote Peer ^ //
func (tm *globalTopic) Exchange(id peer.ID, gloBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host, globalProtocol)
	var reply GlobalServiceResponse
	var args GlobalServiceArgs

	// Set Args
	args.Global = gloBuf

	// Call to Peer
	err := exchClient.Call(id, "GlobalService", "ExchangeWith", args, &reply)
	if err != nil {
		return err
	}

	// Unmarshal Data
	rg := &md.Global{}
	err = proto.Unmarshal(reply.Global, rg)
	if err != nil {
		return err
	}

	// Synchronize Global
	tm.global.Sync(rg)
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (ts *GlobalService) ExchangeWith(ctx context.Context, args GlobalServiceArgs, reply *GlobalServiceResponse) error {
	// Unmarshal Data
	rg := &md.Global{}
	err := proto.Unmarshal(args.Global, rg)
	if err != nil {
		return err
	}

	// Update Peers with Lobby
	ts.global.Sync(rg)

	// Set Message data and call done
	buf, err := ts.global.Buffer()
	if err != nil {
		return err
	}
	reply.Global = buf
	return nil
}

// ^ Method Finds PeerID in Topic ^ //
func (g *globalTopic) FindPeerID(u string) (peer.ID, error) {
	// Find ID from Map
	id, err := g.global.FindPeerID(u)
	if err != nil {
		return "", err
	}

	// Find PeerID from Topic
	for _, v := range g.topic.ListPeers() {
		if v.String() == id {
			return v, nil
		}
	}
	return "", errors.New("PeerID not found in Topic")
}

// ^ handleEvents: listens to Pubsub Events for topic  ^
func (g *globalTopic) handleEvents() {
	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := g.handler.NextPeerEvent(g.ctx)
		if err != nil {
			g.handler.Cancel()
			return
		}

		if lobEvent.Type == psub.PeerJoin {
			gbuf, err := g.global.Buffer()
			if err != nil {
				continue
			}
			err = g.Exchange(lobEvent.Peer, gbuf)
			if err != nil {
				continue
			}
		}

		if lobEvent.Type == psub.PeerLeave {
			g.global.Delete(lobEvent.Peer)
		}
		md.GetState().NeedsWait()
	}
}
