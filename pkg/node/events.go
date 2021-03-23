package node

import (
	"log"
	"time"

	sentry "github.com/getsentry/sentry-go"
	discLimit "github.com/libp2p/go-libp2p-core/discovery"
	disc "github.com/libp2p/go-libp2p-discovery"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Interface: Callback is implemented from Plugin to receive updates ^
type Callback interface {
	OnConnected(data bool)     // Node Host has Bootstrapped
	OnReady(data bool)         // Node Host Connection Result
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnInvited(data []byte)     // User Invited
	OnRemoteStart(data []byte) // User started remote
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// ^ Passes node methods for FS/FQ ^
func (n *Node) FSCallback() md.FileCallback {
	return md.NewFileCallback(n.queued, n.multiQueued, n.error)
}

// ^ Passes node methods for TransferController ^
func (n *Node) TransferCallback() md.TransferCallback {
	return md.NewTransferCallback(n.invited, n.call.OnRemoteStart, n.call.OnResponded, n.call.OnProgress, n.received, n.transmitted, n.error)
}

// ^ queued Callback, Sends File Invite to Peer, and Notifies Client ^
func (n *Node) queued(card *md.TransferCard, req *md.InviteRequest) {
	// Retreive Current File
	currFile := n.fs.CurrentFile()
	if currFile != nil {
		card.Status = md.TransferCard_INVITE
		n.transfer.NewOutgoing(currFile)

		// Create Invite Message
		invMsg := md.AuthInvite{
			From:    n.peer,
			Payload: card.Payload,
			Card:    card,
		}

		// @ Check for Remote
		if req.IsRemote {
			// Start Remote Point
			err := n.transfer.StartRemote(&invMsg)
			if err != nil {
				sentry.CaptureException(err)
				n.error(err, "StartRemotePoint")
			}
		} else {
			// Get PeerID
			id, _, err := n.GetPeer(n.local, req.To.Id.Peer)
			if err != nil {
				n.error(err, "Queued")
			}

			// Check if ID in PeerStore
			go func(inv *md.AuthInvite) {
				// Convert Protobuf to bytes
				msgBytes, err := proto.Marshal(inv)
				if err != nil {
					n.error(err, "Marshal")
				}

				n.transfer.RequestInvite(n.host, id, msgBytes)
			}(&invMsg)
		}
	} else {
		n.error(errors.New("No current file"), "internal:queued")
	}
}

// ^ multiQueued Callback, Sends File Invite to Peer, and Notifies Client ^
func (n *Node) multiQueued(card *md.TransferCard, req *md.InviteRequest, count int) {
	// Get PeerID
	id, _, err := n.GetPeer(n.local, req.To.Id.Peer)
	// Check error
	if err != nil {
		n.error(err, "Queued")
	}

	// Retreive Current File
	currFile := n.fs.CurrentFile()
	if currFile != nil {
		card.Status = md.TransferCard_INVITE
		n.transfer.NewOutgoing(currFile)

		// Create Invite Message
		invMsg := md.AuthInvite{
			From:    n.peer,
			Payload: card.Payload,
			Card:    card,
		}

		// @ Check for Remote
		if req.IsRemote {
			// Start Remote Point
			err := n.transfer.StartRemote(&invMsg)
			if err != nil {
				n.error(err, "StartRemotePoint")
			}
		} else {
			// Check if ID in PeerStore
			go func(inv *md.AuthInvite) {
				// Convert Protobuf to bytes
				msgBytes, err := proto.Marshal(inv)
				if err != nil {
					n.error(err, "Marshal")
				}

				n.transfer.RequestInvite(n.host, id, msgBytes)
			}(&invMsg)
		}
	} else {
		n.error(errors.New("No current file"), "internal:multiQueued")
	}
}

// ^ invite Callback with data for Lifecycle ^ //
func (sn *Node) invited(data []byte) {
	// Update Status
	sn.status = md.Status_INVITED
	// Callback with Data
	sn.call.OnInvited(data)
}

// ^ transmitted Callback middleware post transfer ^ //
func (sn *Node) transmitted(peer *md.Peer) {
	// Update Status
	sn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(peer)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Callback with Data
	sn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (sn *Node) received(card *md.TransferCard) {
	// Update Status
	sn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Callback with Data
	sn.call.OnReceived(msgBytes)
}

// ^ error Callback with error instance, and method ^
func (sn *Node) error(err error, method string) {
	// Log Error
	sentry.CaptureException(err)

	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	// Send Callback
	sn.call.OnError(bytes)
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
			n.error(err, "Finding DHT Peers")
			n.call.OnReady(false)
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
			md.GetState().NeedsWait()
		}

		// Refresh table every 4 seconds
		md.GetState().NeedsWait()
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

		md.GetState().NeedsWait()
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
		md.GetState().NeedsWait()
	}
}
