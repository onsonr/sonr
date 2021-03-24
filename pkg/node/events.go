package node

import (
	"log"
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
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
	tr "github.com/sonr-io/core/pkg/transfer"
	"google.golang.org/protobuf/proto"
)

// ^ handleAuthInviteResponse: Handles User sent AuthInvite Response ^
func (n *Node) handleAuthInviteResponse(id peer.ID, inv *md.AuthInvite, p *md.Peer, cf *dt.ProcessedFile) {
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

	// Check for File
	n.handleAcceptedFileRequest(id, p, cf, reply.Data)
}

// ^ handleAcceptedFileRequest: Begins File Transfer if Accepted ^
func (n *Node) handleAcceptedFileRequest(id peer.ID, p *md.Peer, cf *dt.ProcessedFile, data []byte) {
	// AuthReply Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(data, &resp)
	if err != nil {
		n.call.Error(err, "handleReply")
		sentry.CaptureException(err)
	}

	// Check for File Transfer
	if resp.Decision && resp.Type == md.AuthReply_Transfer {
		n.NewOutgoingTransfer(id, p, cf)
	}
}

// ^ handleDHTPeers: Connects to Peers in DHT ^
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

// ^ handleTopicEvents: listens to Pubsub Events for topic  ^
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
			n.Exchange(tm, lobEvent.Peer, tm.returnPeerBuf())
		}

		if lobEvent.Type == pubsub.PeerLeave {
			tm.lobby.Remove(lobEvent.Peer)
		}

		dt.GetState().NeedsWait()
	}
}

// ^ handleTopicMessages: listens for messages on pubsub topic subscription ^
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
		m := &md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if n.HasPeer(tm, m.Id) {
			tm.messages <- m
		}
		dt.GetState().NeedsWait()
	}
}

// ^ processTopicMessages: pulls messages from channel that have been handled ^
func (n *Node) processTopicMessages(tm *TopicManager) {
	for {
		select {
		// @ when we receive a message from the lobby room
		case m := <-tm.messages:
			// Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				tm.lobby.Add(m.From)
			} else if m.Event == md.LobbyEvent_MESSAGE {
				// Check is Message For Self
				if m.To == n.ID().String() {
					// Convert Message
					bytes, err := proto.Marshal(m)
					if err != nil {
						log.Println("Cannot Marshal Error Protobuf: ", err)
					}

					// Call Event
					n.call.Event(bytes)
				}
			}

		case <-n.ctx.Done():
			return
		}
		dt.GetState().NeedsWait()
	}
}

// ^ handleTransferIncoming: Processes Incoming Data ^ //
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
