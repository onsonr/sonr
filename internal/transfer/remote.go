package transfer

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/net"
	"google.golang.org/protobuf/proto"
)

type RemotePoint struct {
	// Connection
	topic        *pubsub.Topic
	subscription *pubsub.Subscription

	// Data
	Point  string
	invite *md.AuthInvite
}

func (tr *TransferController) StartRemotePoint(authInv *md.AuthInvite) (string, error) {
	// Return Default Option
	_, w, err := net.RandomWords("english", 3)
	if err != nil {
		return "", err
	}

	// Return Split Words Join Group in Lobby
	words := fmt.Sprintf("%s-%s-%s", w[0], w[1], w[2])
	if err != nil {
		sentry.CaptureException(err)
	}

	// Join the local pubsub Topic
	topic, err := tr.pubsub.Join(tr.router.Remote(words))
	if err != nil {
		return "", err
	}

	// Subscribe to local Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return "", err
	}

	// Create Remote Point
	tr.remote = &RemotePoint{
		Point:        words,
		invite:       authInv,
		subscription: sub,
		topic:        topic,
	}

	// Check Peer Count
	peers := topic.ListPeers()
	if len(peers) == 0 {
		go tr.handleRemoteEvents()
	}

	return words, nil
}

func (tr *TransferController) JoinRemotePoint(name string) (*pubsub.Subscription, error) {
	// Join the local pubsub Topic
	topic, err := tr.pubsub.Join(tr.router.Remote(name))
	if err != nil {
		return nil, err
	}

	// Subscribe to local Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}
	return sub, nil
}

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (tr *TransferController) handleRemoteEvents() {
	if tr.remote != nil {
		// @ Create Topic Handler
		topicHandler, err := tr.remote.topic.EventHandler()
		if err != nil {
			log.Println(err)
			return
		}

		// @ Loop Events
		for {
			// Get next event
			lobEvent, err := topicHandler.NextPeerEvent(tr.ctx)
			if err != nil {
				topicHandler.Cancel()
				return
			}

			// Peer Has Joined
			if lobEvent.Type == pubsub.PeerJoin {
				// Get Peer Data
				bytes, err := proto.Marshal(tr.remote.invite)
				if err != nil {
					tr.call.Error(err, "Direct")
				}

				// Get Peer's ID
				tr.RequestInvite(tr.host, lobEvent.Peer, bytes)

			}
			md.GetState().NeedsWait()
		}
	}
}
