package lobby

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleEvents() {
	// @ Create Topic Handler
	topicHandler, err := lob.topic.EventHandler()
	if err != nil {
		log.Println(err)
		return
	}

	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := topicHandler.NextPeerEvent(lob.ctx)
		if err != nil {
			topicHandler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			lob.Exchange(lobEvent.Peer)
		}

		if lobEvent.Type == pubsub.PeerLeave {
			lob.removePeer(lobEvent.Peer)
		}

		md.GetState().NeedsWait()
	}
}

// ^ 1. handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (lob *Lobby) handleMessages() {
	for {
		// Get next msg from pub/sub
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.messages)
			return
		}

		// Only forward messages delivered by others
		if msg.ReceivedFrom == lob.host.ID() {
			continue
		}

		// Construct message
		m := md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, &m)
		if err != nil {
			continue
		}

		// Update Circle by event
		lob.messages <- &m
		md.GetState().NeedsWait()
	}
}

// ^ 1a. processMessages handles message content and ticker ^
func (lob *Lobby) processMessages() {
	for {
		select {
		// @ when we receive a message from the lobby room
		case m := <-lob.messages:
			// Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				lob.updatePeer(m.Data)
			}

		case <-lob.ctx.Done():
			return
		}
		md.GetState().NeedsWait()
	}
}

// ^ removePeer removes Peer from Map ^
func (lob *Lobby) removePeer(id peer.ID) {
	// Update Peer with new data
	delete(lob.data.Peers, id.String())
	lob.data.Count = int32(len(lob.data.Peers))
	lob.data.Size = int32(len(lob.data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}

// ^ updatePeer changes Peer values in Lobby ^
func (lob *Lobby) updatePeer(peer *md.Peer) {
	// Update Peer with new data
	lob.data.Peers[peer.Id.Peer] = peer
	lob.data.Count = int32(len(lob.data.Peers))
	lob.data.Size = int32(len(lob.data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}
