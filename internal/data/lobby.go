package data

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type Lobby struct {
	OLC   string
	Size  int32
	Count int32
	Peers map[string]*md.Peer

	// Private Properties
	callRefresh OnProtobuf
	user        *md.Peer
}

// ^ Create New Circle  ^
func NewLobby(u *md.Peer, cr OnProtobuf) *Lobby {
	return &Lobby{
		user:        u,
		callRefresh: cr,
	}
}

// ^ Returns as Lobby Buffer ^
func (l *Lobby) Buffer() []byte {
	bytes, err := proto.Marshal(&md.Lobby{
		Olc:   l.OLC,
		Size:  l.Size,
		Count: l.Count,
		Peers: l.Peers,
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	return bytes
}

// ^ Add/Update Peer in Lobby ^
func (l *Lobby) Add(peer *md.Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
	l.Refresh()
}

// ^ Remove Peer from Lobby ^
func (l *Lobby) Remove(id peer.ID) {
	// Update Peer with new data
	delete(l.Peers, id.String())
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
	l.Refresh()
}

// ^ Sync Between Remote Peers Lobby ^
func (l *Lobby) Sync(ref *md.Lobby, peer *md.Peer) {
	// Validate Lobbies are Different
	if l.Count != ref.Count {
		// Iterate Over List
		for id, peer := range ref.Peers {
			// Add all Peers NOT User
			if id != l.user.Id.Peer {
				l.Peers[id] = peer
			}
		}
	}

	// Add Peer to Lobby, Refreshes Automatically
	l.Add(peer)
}

// ^ Send Updated Lobby ^
func (l *Lobby) Refresh() {
	bytes, err := proto.Marshal(&md.Lobby{
		Olc:   l.OLC,
		Size:  l.Size,
		Count: l.Count,
		Peers: l.Peers,
	})
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
		return
	}
	l.callRefresh(bytes)
}
