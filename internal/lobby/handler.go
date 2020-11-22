package lobby

import (
	"errors"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ 1. handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleMessages() {
	for {
		// Get next msg from pub/sub
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.Messages)
			return
		}

		// Only forward messages delivered by others
		if msg.ReceivedFrom.String() == lob.Self.GetId() {
			continue
		}

		// Construct message
		notif := pb.LobbyMessage{}
		err = proto.Unmarshal(msg.Data, &notif)
		if err != nil {
			continue
		}

		// Send valid messages onto the Messages channel
		lob.Messages <- &notif
	}
}

// ^ 2. handleEvents handles message content and ticker ^
func (lob *Lobby) handleEvents() {
	// Timer checks to dispose of peers
	peerRefreshTicker := time.NewTicker(time.Second * 3)
	defer peerRefreshTicker.Stop()

	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			// Update Circle by event
			if m.Subject == pb.LobbyMessage_UPDATE {
				// Update Peer Data
				lob.updatePeer(m.Id, m.Peer)

			} else if m.Subject == pb.LobbyMessage_EXIT {
				// Remove Peer Data
				lob.removePeer(m.Id)
			}

		// ** Refresh and Validate Lobby Peers Periodically ** //
		case <-peerRefreshTicker.C:
			lob.call.Refreshed(lob.Peers())

		case <-lob.ctx.Done():
			return

		case <-lob.doneCh:
			return
		}
	}
}

// ** ID returns ONE Peer.ID in PubSub **
func (lob *Lobby) ID(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range lob.ps.ListPeers(lob.Data.Code) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	// Log Error
	err := errors.New("Error QueryId was not found in PubSub topic")
	lob.call.Error(err, "ID")
	return ""
}

// ** Peer returns ONE Peer in Lobby **
func (lob *Lobby) Peer(q string) *pb.Peer {
	// Iterate Through Peers, Return Matched Peer
	for _, peer := range lob.Data.Peers {
		// If Found Match
		if peer.Id == q {
			return peer
		}
	}
	return nil
}

// ** removePeer deletes a peer from Lobby ** //
func (lob *Lobby) removePeer(id string) {
	// Delete peer from Lobby Map
	delete(lob.Data.Peers, id)

	// Send Callback with updated peers
	lob.call.Refreshed(lob.Data())
}

// ** updatePeer changes peer values in Lobby **
func (lob *Lobby) updatePeer(id string, data *pb.Peer) {
	// Update Peer with new data
	lob.Data.Peers[id] = data
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Send Callback with updated peers
	lob.call.Refreshed(lob.Data())
}
