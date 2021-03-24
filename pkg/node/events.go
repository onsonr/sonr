package node

import (
	"time"

	sentry "github.com/getsentry/sentry-go"
	discLimit "github.com/libp2p/go-libp2p-core/discovery"
	disc "github.com/libp2p/go-libp2p-discovery"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/pkg/errors"
	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Passes node methods for TransferController ^
func (n *Node) TransferCallback() dt.TransferCallback {
	return dt.NewTransferCallback(n.call.Invited, n.call.RemoteStart, n.call.Responded, n.call.Progressed, n.call.Received, n.call.Transmitted, n.call.Error)
}

// ^ Handles Peers in DHT ^
func (n *Node) handleDHTPeers(routingDiscovery *disc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			n.ctx,
			n.router.LocalPoint(),
			discLimit.Limit(16),
		)
		if err != nil {
			n.call.Error(err, "Finding DHT Peers")
			n.call.Ready(false)
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != n.host.ID() {
				// Connect to Peer
				if err := n.host.Connect(n.ctx, pi); err != nil {
					// Capture Error
					sentry.CaptureException(errors.Wrap(err, "Failed to connect to peer in namespace"))

					// Remove Peer Reference
					n.host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := n.host.Network().(*swarm.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
			dt.GetState().NeedsWait()
		}

		// Refresh table every 4 seconds
		dt.GetState().NeedsWait()
		<-time.After(time.Second * 2)
	}
}

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (n *Node) handleTopicEvents(tm *TopicManager) {
	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := tm.handler.NextPeerEvent(n.ctx)
		if err != nil {
			tm.handler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			n.Exchange(tm, lobEvent.Peer)
		}

		if lobEvent.Type == pubsub.PeerLeave {
			tm.lobby.Remove(lobEvent.Peer)
		}

		dt.GetState().NeedsWait()
	}
}

// ^ 1. handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (n *Node) handleTopicMessages(tm *TopicManager) {
	for {
		// Get next msg from pub/sub
		msg, err := tm.subscription.Next(n.ctx)
		if err != nil {
			return
		}

		// Only forward messages delivered by others
		if msg.ReceivedFrom == n.ID() {
			continue
		}

		// Construct message
		m := md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, &m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if n.HasPeer(tm, m.Id) {
			// Update Circle by event
			tm.lobby.Add(m.From)
		}
		dt.GetState().NeedsWait()
	}
}
