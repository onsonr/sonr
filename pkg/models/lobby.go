package models

import (
	"fmt"
	"log"
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

// ** ─── Lobby MANAGEMENT ────────────────────────────────────────────────────────
// Creates Local Lobby from User Data
func NewLocalLobby(u *User) *SyncLobby {
	// Get Info
	topic := u.LocalIPTopic()
	loc := u.GetRouter().GetLocation()

	// Create Lobby
	return &SyncLobby{
		syncMap: sync.Map{},
		internal: &Lobby{
			// General
			Type:  Lobby_LOCAL,
			Peers: make(map[string]*Peer),
			User:  u.GetPeer(),

			// Info
			Info: &Lobby_Local{
				Local: &Lobby_LocalInfo{
					Name:     topic[12:],
					Location: loc,
					Topic:    topic,
				},
			},
		},
	}
}

// Creates Local Lobby from User Data
func NewLocalLinkLobby(l *Linker) *SyncLobby {
	// Get Info
	topic := l.Router.LocalIPTopic
	loc := l.GetRouter().GetLocation()

	// Create Lobby
	return &SyncLobby{
		syncMap: sync.Map{},
		internal: &Lobby{
			// General
			Type:  Lobby_LOCAL,
			Peers: make(map[string]*Peer),
			User:  l.GetPeer(),

			// Info
			Info: &Lobby_Local{
				Local: &Lobby_LocalInfo{
					Name:     topic[12:],
					Location: loc,
					Topic:    topic,
				},
			},
		},
	}
}

// Get Remote Point Info
func NewRemote(u *User, list []string, file *SonrFile) *SyncLobby {
	r := &RemoteResponse{
		Display: fmt.Sprintf("%s %s %s", list[0], list[1], list[2]),
		Topic:   fmt.Sprintf("%s-%s-%s", list[0], list[1], list[2]),
		Count:   int32(len(list)),
		IsJoin:  false,
		Words:   list,
	}

	// Create Lobby
	return &SyncLobby{
		syncMap: sync.Map{},
		internal: &Lobby{
			// General
			Type:  Lobby_REMOTE,
			Peers: make(map[string]*Peer),
			User:  u.GetPeer(),

			// Info
			Info: &Lobby_Remote{
				Remote: &Lobby_RemoteInfo{
					IsJoin:  r.IsJoin,
					Display: r.Display,
					Words:   r.GetWords(),
					Topic:   r.GetTopic(),
					File:    file,
					Owner:   u.GetPeer(),
				},
			},
		},
	}
}

func NewJoinedRemote(u *User, r *RemoteResponse) *SyncLobby {
	// Create Lobby
	return &SyncLobby{
		syncMap: sync.Map{},
		internal: &Lobby{
			// General
			Type:  Lobby_REMOTE,
			Peers: make(map[string]*Peer),
			User:  u.GetPeer(),

			// Info
			Info: &Lobby_Remote{
				Remote: &Lobby_RemoteInfo{
					IsJoin:  r.IsJoin,
					Display: r.Display,
					Words:   r.GetWords(),
					Topic:   r.GetTopic(),
				},
			},
		},
	}
}

// Returns Lobby Peer Count
func (l *Lobby) Count() int {
	return len(l.Peers)
}

// Returns TOTAL Lobby Size with Peer
func (l *Lobby) Size() int {
	return len(l.Peers) + 1
}

// Returns as Lobby Buffer
func (l *Lobby) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes, nil
}

func (l *Lobby) Sync(key, value interface{}) bool {
	// Add All Valid Entries
	ok, id, peer := validatePeerEntry(key, value)
	if ok {
		// Update Peer with new data
		if l.User.IsNotSame(peer) {
			l.Peers[id] = peer
		}
	}
	return true
}

// ** ─── SyncLobby MANAGEMENT ────────────────────────────────────────────────────────

type SyncLobby struct {
	internal *Lobby
	syncMap  sync.Map
}

// ^ Returns Lobby as Buffer from Internal
func (sl *SyncLobby) Buffer() ([]byte, error) {
	return sl.internal.Buffer()
}

// ^ Returns Lobby Protobuf from Internal
func (sl *SyncLobby) Data() *Lobby {
	return sl.internal
}

// ^ Method Finds Peer in Map ^ //
func (sl *SyncLobby) Find(key string) (*Peer, bool) {
	val, ok := sl.syncMap.Load(key)
	if ok {
		return validatePeerValue(val)
	}
	return nil, false
}

// ^ Retreives Remote Info from Internal
func (sl *SyncLobby) GetRemote() *Lobby_RemoteInfo {
	return sl.internal.GetRemote()
}

// ^ Add Peer to Sync Map
func (sl *SyncLobby) Add(p *Peer) {
	sl.syncMap.Store(p.PeerID(), p)
	sl.syncMap.Range(sl.internal.Sync)
}

// ^ Syncs Remote Lobby Instance with Current Instance
func (sl *SyncLobby) Sync(rl *Lobby, rp *Peer) {
	// Iterate Remote Peers
	for k, v := range rl.Peers {
		sl.syncMap.LoadOrStore(k, v)
	}

	// Add Remote Peer
	sl.syncMap.LoadOrStore(rp.PeerID(), rp)
	sl.syncMap.Range(sl.internal.Sync)
}

// ^ Remove Peer from Sync Map
func (sl *SyncLobby) Remove(id peer.ID) {
	sl.syncMap.Delete(id.String())
	sl.syncMap.Range(sl.internal.Sync)
}

// ^ Returns Remote Response Buffer from Internal
func (l *SyncLobby) ToRemoteResponseBytes() []byte {
	switch l.internal.Info.(type) {
	// @ Join Remote
	case *Lobby_Remote:
		// Convert Info to Response
		i := l.internal.GetRemote()
		resp := &RemoteResponse{
			IsJoin:  i.GetIsJoin(),
			Display: i.GetDisplay(),
			Topic:   i.GetTopic(),
			Words:   i.GetWords(),
		}

		// Marshal Bytes
		data, err := proto.Marshal(resp)
		if err != nil {
			return nil
		}
		return data
	}
	return nil
}

// Returns Lobby Topic
func (l *SyncLobby) Topic() string {
	topic := ""
	switch l.internal.Info.(type) {
	// @ Create Remote
	case *Lobby_Local:
		topic = l.internal.GetLocal().GetTopic()
	// @ Join Remote
	case *Lobby_Remote:
		topic = l.internal.GetRemote().GetTopic()
	}
	return topic
}

// ** ─── Map Validation ────────────────────────────────────────────────────────
// @ Helper: Type Assertion for Peers Map Key/Value
func validatePeerEntry(key, value interface{}) (bool, string, *Peer) {
	// Type Assertions
	skey, keyOk := validatePeerKey(key)
	pvalue, valOk := validatePeerValue(value)

	// Check Results
	if keyOk && valOk {
		return true, skey, pvalue
	}
	return false, "", nil
}

// @ Helper: Type Assertion for Peers Map ONLY Key
func validatePeerKey(key interface{}) (string, bool) {
	skey, keyOk := key.(string)
	return skey, keyOk
}

// @ Helper: Type Assertion for Peers Map ONLY Value
func validatePeerValue(value interface{}) (*Peer, bool) {
	pvalue, valOk := value.(*Peer)
	return pvalue, valOk
}
