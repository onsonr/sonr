package node

import (
	"time"

	sentry "github.com/getsentry/sentry-go"
	discLimit "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	disc "github.com/libp2p/go-libp2p-discovery"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	swarm "github.com/libp2p/go-libp2p-swarm"
	msgio "github.com/libp2p/go-msgio"
	"github.com/pkg/errors"
	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
	tr "github.com/sonr-io/core/pkg/transfer"
	"google.golang.org/protobuf/proto"
)

// ^ Passes node methods for TransferController ^
func (n *Node) TransferCallback() dt.TransferCallback {
	return dt.NewTransferCallback(n.call.Invited, n.call.RemoteStart, n.call.Responded, n.call.Progressed, n.call.Received, n.call.Transmitted, n.call.Error)
}

// ^ Helper Method to Handle User sent AuthInvite Response ^
func (n *Node) handleAuthInviteRPC(id peer.ID, inv *md.AuthInvite) {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Initialize Data
	rpcClient := rpc.NewClient(n.host, n.router.Auth())
	var reply AuthResponse
	var args AuthArgs
	args.Data = msgBytes

	// Call to Peer
	done := make(chan *rpc.Call, 1)
	err = rpcClient.Go(id, "AuthService", "Invited", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		sentry.CaptureException(err)
		n.call.Error(err, "Request")
	}

	// Send Callback and Reset
	n.call.Responded(reply.Data)
	// transDecs, from := n.handleAuthReply(reply.Data)

	// Check Response for Accept
	// if transDecs {
	// pc.StartOutgoing(h, id, from)
	// }
}

// ^ Helper Method to Handle User sent AuthInvite Response ^
func (n *Node) handleAuthReply(data []byte) (bool, *md.Peer) {
	// Received Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(data, &resp)
	if err != nil {
		n.call.Error(err, "handleReply")
		sentry.CaptureException(err)
		return false, nil
	}
	return resp.Decision && resp.Type == md.AuthReply_Transfer, resp.From
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

// ^ Handle Incoming Stream ^ //
func (n *Node) handleTransferIncoming(stream network.Stream) {
	// Route Data from Stream
	go func(reader msgio.ReadCloser, t *tr.IncomingFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				n.call.Error(err, "HandleIncoming:ReadMsg")
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(i, buffer)
			if err != nil {
				n.call.Error(err, "HandleIncoming:AddBuffer")
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := n.incoming.Save(); err != nil {
					n.call.Error(err, "HandleIncoming:Save")
				}
				break
			}
			dt.GetState().NeedsWait()
		}
	}(msgio.NewReader(stream), n.incoming)
}
