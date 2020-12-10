package lobby

import (
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
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
		if msg.ReceivedFrom == lob.self {
			continue
		}

		// Construct message
		notif := md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, &notif)
		if err != nil {
			continue
		}

		// Send valid messages onto the Messages channel
		lob.Messages <- &notif
	}
}

// ^ 1a. processMessages handles message content and ticker ^
func (lob *Lobby) processMessages() {
	// Timer checks to dispose of peers
	dur := (time.Second * 2) + (time.Millisecond * 500)
	peerRefreshTicker := time.NewTicker(dur)
	defer peerRefreshTicker.Stop()

	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			// Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				lob.updatePeer(m.Peer)
			}

		// ** Refresh and Validate Lobby Peers Periodically ** //
		case <-peerRefreshTicker.C:
			//lob.filterLobby()

		case <-lob.ctx.Done():
			return

		case <-lob.doneCh:
			return
		}
	}
}

// // ^ 2. handleEvents listens for topicEvents ^
// func (lob *Lobby) handleEvents() {
// 	for {
// 		// handle topic events
// 		event, err := lob.topicHandler.NextPeerEvent(lob.ctx)
// 		if err != nil {
// 			return
// 		}

// 		lob.Events <- &event
// 	}
// }

// // ^ 2a. processEvents handles events from topic event channel
// func (lob *Lobby) processEvents() {
// 	for {
// 		select {
// 		// ** Event for Peer Left ** //
// 		case e := <-lob.Events:
// 			if e.Type == pubsub.PeerLeave {
// 				log.Println(e.Peer.ShortString(), " has exited.")
// 			}
// 			if e.Type == pubsub.PeerJoin {
// 				log.Println(e.Peer.ShortString(), " has joined.")
// 			}

// 		case <-lob.ctx.Done():
// 			break
// 		}
// 	}
// }

// ** ID returns ONE Peer.ID in PubSub **
func (lob *Lobby) ID(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range lob.ps.ListPeers(lob.Data.Code) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	return ""
}

// ** Peer returns ONE Peer in Lobby **
func (lob *Lobby) Peer(q string) *md.Peer {
	// Iterate Through Peers, Return Matched Peer
	for _, peer := range lob.Data.Peers {
		// If Found Match
		if peer.Id == q {
			return peer
		}
	}
	return nil
}

// ** filterLobby updates lobby and removes peers that arent subscribed ** //
func (lob *Lobby) refreshLobby() {
	// Loop through Subscribed Peers
	for _, id := range lob.ps.ListPeers(lob.Data.Code) {
		// Find Peer that is not found
		if peer, found := lob.Data.Peers[id.String()]; !found {
			log.Println(peer.String(), " is not subscribed anymore.")

			// Remove Unsubscribed Peer from Map
			delete(lob.Data.Peers, id.String())
		}
	}

	// Marshal data to bytes
	bytes, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Return Callback
	lob.callback(bytes)
}

// ** updatePeer changes peer values in Lobby **
func (lob *Lobby) updatePeer(peer *md.Peer) {
	// Update Peer with new data
	id := peer.Id
	lob.Data.Peers[id] = peer
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Marshal data to bytes
	bytes, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.callback(bytes)
}
