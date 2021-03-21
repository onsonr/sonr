package remote

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/net"
)

type RemotePoint struct {
	// Networking
	ctx    context.Context
	host   host.Host
	pubSub *pubsub.PubSub

	// Connection
	router       *net.ProtocolRouter
	topic        *pubsub.Topic
	topicHandler *pubsub.TopicEventHandler
	selfPeer     *md.Peer
	subscription *pubsub.Subscription
}

func NewRemotePoint(ctx context.Context, h host.Host, ps *pubsub.PubSub, sp *md.Peer, pr *net.ProtocolRouter, lobCall md.LobbyCallback) (*RemotePoint, error) {
	// Return Default Option
	_, w, err := net.RandomWords("english", 4)
	if err != nil {
		return nil, err
	}

	// Return Split Words Join Group in Lobby
	words := fmt.Sprintf("%s-%s-%s-%s", w[0], w[1], w[2], w[3])
	if err != nil {
		sentry.CaptureException(err)
	}

	// Join the local pubsub Topic
	topic, err := ps.Join(pr.Topic(net.SetIDForGroup(words)))
	if err != nil {
		return nil, err
	}

	// Subscribe to local Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Top Handler
	topicHandler, err := topic.EventHandler()
	if err != nil {
		return nil, err
	}

	rp := &RemotePoint{
		ctx:          ctx,
		host:         h,
		pubSub:       ps,
		selfPeer:     sp,
		router:       pr,
		subscription: sub,
		topic:        topic,
		topicHandler: topicHandler,
	}

	return rp, nil
}
