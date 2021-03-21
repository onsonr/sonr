package transfer

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/net"
	"google.golang.org/protobuf/proto"
)

type RemotePoint struct {
	// Connection
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	peerchan     chan peer.ID

	// Data
	Point  string
	invite *md.AuthInvite
}

// ^ Start New Remote Point ^
func (tr *TransferController) StartRemotePoint(authInv *md.AuthInvite) error {
	// Return Default Option
	_, w, err := net.RandomWords("english", 3)
	if err != nil {
		return err
	}

	// Return Split Words Join Group in Lobby
	words := fmt.Sprintf("%s-%s-%s", w[0], w[1], w[2])
	if err != nil {
		return err
	}

	// Create Remote Response
	resp := md.RemoteResponse{
		First:   w[0],
		Second:  w[1],
		Third:   w[2],
		Display: fmt.Sprintf("%s %s %s", w[0], w[1], w[2]),
	}

	// Join the local pubsub Topic
	topic, err := tr.pubsub.Join(tr.router.Remote(words))
	if err != nil {
		return err
	}

	// Subscribe to local Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return err
	}

	// Create Remote Point
	tr.remote = &RemotePoint{
		Point:        words,
		peerchan:     make(chan peer.ID, 1),
		invite:       authInv,
		subscription: sub,
		topic:        topic,
	}

	// Await Peer
	byteRemoteResp, err := proto.Marshal(&resp)
	if err != nil {
		return err
	}

	// Callback and Listen
	tr.call.RemoteStart(byteRemoteResp)
	go tr.handleRemoteEvents()

	// Create Handler with Timeout
	select {
	case res := <-tr.remote.peerchan:
		// Get Peer Data
		bytes, err := proto.Marshal(tr.remote.invite)
		if err != nil {
			tr.call.Error(err, "Direct")
			return err
		}

		// Get Peer's ID
		tr.RequestInvite(tr.host, res, bytes)
	case <-time.After(50 * time.Second):
		tr.remote.subscription.Cancel()
		tr.remote.topic.Close()
		tr.remote = nil
		return nil
	}
	return nil
}

// ^ Join Existing Remote Point ^
func (tr *TransferController) JoinRemotePoint(name string) (*pubsub.Subscription, error) {
	// Join the local pubsub Topic
	topic, err := tr.pubsub.Join(tr.router.Remote(name))
	if err != nil {
		return nil, err
	}

	// Check Peer Count
	peers := topic.ListPeers()
	if len(peers) == 0 {
		topic.Close()
		return nil, errors.New("Invalid Point")
	} else {
		// Subscribe to local Topic
		sub, err := topic.Subscribe()
		if err != nil {
			return nil, err
		}
		return sub, nil
	}
}

// ^ handleRemoteEvents awaits peer to join to begin transfer ^
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
				tr.remote.peerchan <- lobEvent.Peer
				return
			}
			md.GetState().NeedsWait()
		}
	}
}
