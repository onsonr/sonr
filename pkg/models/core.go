package models

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

// ** ─── CALLBACK MANAGEMENT ────────────────────────────────────────────────────────
type HTTPHandler func(http.ResponseWriter, *http.Request)
type SetStatus func(s Status)
type OnProtobuf func([]byte)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *Transfer)
type OnTransmitted func(data *Transfer)
type OnError func(err *SonrError)
type NodeCallback struct {
	Invited     OnProtobuf
	Linked      OnProtobuf
	Refreshed   OnProtobuf
	Event       OnProtobuf
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	Status      SetStatus
	Transmitted OnTransmitted
	Error       OnError
}

// ** ─── State MANAGEMENT ────────────────────────────────────────────────────────
type state struct {
	flag uint64
	chn  chan bool
}

var (
	instance *state
	once     sync.Once
)

func GetState() *state {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &state{chn: chn}
	})

	return instance
}

// Checks rather to wait or does not need
func (c *state) NeedsWait() {
	<-c.chn
}

// Says all of goroutines to resume execution
func (c *state) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Says all of goroutines to pause execution
func (c *state) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.chn = make(chan bool)
	}
}

// ** ─── Transfer MANAGEMENT ────────────────────────────────────────────────────────
// Returns Transfer for URLLink
func (u *URLLink) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_Url{
			Url: u,
		},
	}
}

// Returns URLLink as Transfer_Url Data
func (u *URLLink) ToData() *Transfer_Url {
	return &Transfer_Url{
		Url: u,
	}
}

// Returns Transfer for SonrFile
func (f *SonrFile) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_File{
			File: f,
		},
	}
}

// Returns SonrFile as Transfer_File Data
func (f *SonrFile) ToData() *Transfer_File {
	return &Transfer_File{
		File: f,
	}
}

// Returns Transfer for Contact
func (c *Contact) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_Contact{
			Contact: c,
		},
	}
}

// Returns Contact as Transfer_Contact Data
func (c *Contact) ToData() *Transfer_Contact {
	return &Transfer_Contact{
		Contact: c,
	}
}

// ** ─── Global MANAGEMENT ────────────────────────────────────────────────────────
// Returns as Lobby Buffer
func (g *Global) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(g)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes, nil
}

// Remove Peer from Lobby
func (g *Global) Delete(id peer.ID) {
	// Find Username
	u, err := g.FindUsername(id)
	if err != nil {
		return
	}

	// Remove Peer
	delete(g.Peers, u)
}

// Method Finds PeerID for Username
func (g *Global) FindPeerID(u string) (string, error) {
	for k, v := range g.Peers {
		if k == u {
			return v, nil
		}
	}
	return "", errors.New("PeerID not Found")
}

// Method Finds Username for PeerID
func (g *Global) FindUsername(id peer.ID) (string, error) {
	for _, v := range g.Peers {
		if v == id.String() {
			return v, nil
		}
	}
	return "", errors.New("Username not Found")
}

// Method Checks if Global has Username
func (g *Global) HasUsername(u string) bool {
	for k := range g.Peers {
		if k == u {
			return true
		}
	}
	return false
}

// Sync Between Remote Peers Lobby
func (g *Global) Sync(rg *Global) {
	// Iterate Over Remote Map
	for username, id := range rg.Peers {
		if g.UserName != username {
			g.Peers[username] = id
		}
	}

	// Check Self Map
	if !g.HasUsername(rg.UserName) {
		g.Peers[rg.UserName] = rg.UserPeerID
	}
}

// ** ─── Lobby MANAGEMENT ────────────────────────────────────────────────────────
// Creates Local Lobby from User Data
func NewLocalLobby(u *User) *Lobby {
	// Get Info
	topic := u.LocalIPTopic()
	loc := u.GetRouter().GetLocation()

	// Create Lobby
	return &Lobby{
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
	}
}

// Creates Local Lobby from User Data
func NewLocalLinkLobby(l *Linker) *Lobby {
	// Get Info
	topic := l.Router.LocalIPTopic
	loc := l.GetRouter().GetLocation()

	// Create Lobby
	return &Lobby{
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
	}
}

// Get Remote Point Info
func NewRemote(u *User, list []string, file *SonrFile) *Lobby {
	r := &RemoteResponse{
		Display: fmt.Sprintf("%s %s %s", list[0], list[1], list[2]),
		Topic:   fmt.Sprintf("%s-%s-%s", list[0], list[1], list[2]),
		Count:   int32(len(list)),
		IsJoin:  false,
		Words:   list,
	}

	// Create Lobby
	return &Lobby{
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
	}
}

func NewJoinedRemote(u *User, r *RemoteResponse) *Lobby {
	// Create Lobby
	return &Lobby{
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

func (l *Lobby) ToRemoteResponseBytes() []byte {
	switch l.Info.(type) {
	// @ Join Remote
	case *Lobby_Remote:
		// Convert Info to Response
		i := l.GetRemote()
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
func (l *Lobby) Topic() string {
	topic := ""
	switch l.Info.(type) {
	// @ Create Remote
	case *Lobby_Local:
		topic = l.GetLocal().GetTopic()
	// @ Join Remote
	case *Lobby_Remote:
		topic = l.GetRemote().GetTopic()
	}
	return topic
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

// Add/Update Peer in Lobby
func (l *Lobby) Add(peer *Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
}

// Remove Peer from Lobby
func (l *Lobby) Delete(id peer.ID) {
	// Update Peer with new data
	delete(l.Peers, id.String())
}

// Sync Between Remote Peers Lobby
func (l *Lobby) Sync(ref *Lobby, remotePeer *Peer) {
	// Validate Lobbies are Different
	if l.Count() != ref.Count() {
		// Iterate Over List
		for id, peer := range ref.Peers {
			if l.User.IsNotPeerIDString(id) {
				l.Add(peer)
			}
		}
	}
	l.Add(remotePeer)
}
