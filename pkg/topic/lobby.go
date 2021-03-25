package topic

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
		dt "github.com/sonr-io/core/pkg/data"
)
type Lobby struct {
	OLC   string
	Size  int32
	Count int32
	Peers map[string]*md.Peer

	// Private Properties
	callRefresh dt.OnProtobuf
	user        *md.Peer
}

// ^ Create New Circle  ^
func NewLobby(o string, u *md.Peer, cr dt.OnProtobuf) *Lobby {
	return &Lobby{
		user:        u,
		callRefresh: cr,

		OLC:   o,
		Size:  1,
		Count: 0,
		Peers: make(map[string]*md.Peer),
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

// ^ Add/Update Peer in Lobby without Callback ^
func (l *Lobby) AddWithoutRefresh(peer *md.Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
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
func (l *Lobby) Sync(ref *md.Lobby, remotePeer *md.Peer) {
	// Validate Lobbies are Different
	if l.Count != ref.Count {
		// Iterate Over List
		for id, peer := range ref.Peers {
			// Add all Peers NOT User
			if id != l.user.Id.Peer {
				l.AddWithoutRefresh(peer)
			}
		}
	}

	// Add Synced Peer to Lobby
	l.Add(remotePeer)
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
