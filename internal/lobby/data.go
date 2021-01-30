package lobby

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

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

// ** updatePeer changes peer values in Lobby **
func (lob *Lobby) addPeer(peer *md.Peer) {
	// Validate ID doesnt Exist
	if result := Contains(lob.Data.Peers, peer.Id); !result {
		// Add Peer to List
		lob.Data.Peers = append(lob.Data.Peers, peer)

		// Marshal data to bytes
		bytes, err := proto.Marshal(lob.Data)
		if err != nil {
			log.Println("Cannot Marshal Error Protobuf: ", err)
		}

		// Send Callback with updated peers
		lob.callback(bytes)
	}
}

// ** updatePeer changes peer values in Lobby **
func (lob *Lobby) removePeer(id peer.ID) {
	// Validate ID Exists
	if result := Contains(lob.Data.Peers, string(id)); result {
		// Get Peer Index
		idx := Find(lob.Data.Peers, string(id))

		// Remove Peer from List
		lob.Data.Peers = Remove(lob.Data.Peers, idx)

		// Update Size, Account for User
		lob.Data.Size = int32(len(lob.Data.Peers)) + 1

		// Marshal data to bytes
		bytes, err := proto.Marshal(lob.Data)
		if err != nil {
			log.Println("Cannot Marshal Error Protobuf: ", err)
		}

		// Send Callback with updated peers
		lob.callback(bytes)
	}
}

// ** updatePeer changes peer values in Lobby **
func (lob *Lobby) updatePeer(peer *md.Peer) {
	// Validate ID Exists
	if result := Contains(lob.Data.Peers, string(peer.Id)); result {
		// Get Peer Index
		idx := Find(lob.Data.Peers, string(peer.Id))

		// Update Value
		lob.Data.Peers[idx] = peer

		// Update Size, Account for User
		lob.Data.Size = int32(len(lob.Data.Peers)) + 1

		// Marshal data to bytes
		bytes, err := proto.Marshal(lob.Data)
		if err != nil {
			log.Println("Cannot Marshal Error Protobuf: ", err)
		}

		// Send Callback with updated peers
		lob.callback(bytes)
	} else {
		lob.addPeer(peer)
	}
}

// @ Helper: Find returns the smallest index i at which x == a[i]
func Find(a []*md.Peer, id string) int {
	for i, n := range a {
		if id == n.Id {
			return i
		}
	}
	return len(a)
}

// @ Helper: Contains tells whether a contains x.
func Contains(a []*md.Peer, id string) bool {
	for _, n := range a {
		if id == n.Id {
			return true
		}
	}
	return false
}

// @ Helper: Removes Item at Index
func Remove(a []*md.Peer, index int) []*md.Peer {
	return append(a[:index], a[index+1:]...)
}
