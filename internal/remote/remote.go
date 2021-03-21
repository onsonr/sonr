package remote

import (
	"context"
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/transfer"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/net"
)

type RemotePoint struct {
	// Networking
	ctx    context.Context
	call   md.TransferCallback
	host   host.Host
	point  string
	pubSub *pubsub.PubSub

	// Connection
	router       *net.ProtocolRouter
	topic        *pubsub.Topic
	topicHandler *pubsub.TopicEventHandler
	subscription *pubsub.Subscription
	transfer     *transfer.TransferController

	// Data
	invite   *md.AuthInvite
	selfPeer *md.Peer
}

func StartRemotePoint(ctx context.Context, h host.Host, ps *pubsub.PubSub, sp *md.Peer, pr *net.ProtocolRouter, authInv *md.AuthInvite, tr *transfer.TransferController, lobCall md.TransferCallback) (*RemotePoint, error) {
	// Return Default Option
	_, w, err := net.RandomWords("english", 3)
	if err != nil {
		return nil, err
	}

	// Return Split Words Join Group in Lobby
	words := fmt.Sprintf("%s-%s-%s", w[0], w[1], w[2])
	if err != nil {
		sentry.CaptureException(err)
	}

	// Join the local pubsub Topic
	topic, err := ps.Join(pr.Remote(words))
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
		call:         lobCall,
		host:         h,
		invite:       authInv,
		point:        words,
		pubSub:       ps,
		selfPeer:     sp,
		router:       pr,
		subscription: sub,
		transfer:     tr,
		topic:        topic,
		topicHandler: topicHandler,
	}
	go rp.handleEvents()
	return rp, nil
}

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (rp *RemotePoint) handleEvents() {
	// @ Create Topic Handler
	topicHandler, err := rp.topic.EventHandler()
	if err != nil {
		log.Println(err)
		return
	}

	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := topicHandler.NextPeerEvent(rp.ctx)
		if err != nil {
			topicHandler.Cancel()
			return
		}

		// Peer Has Joined
		if lobEvent.Type == pubsub.PeerJoin {
			// Get Peer's ID
			rp.Direct(lobEvent.Peer)

		}
		md.GetState().NeedsWait()
	}
}
