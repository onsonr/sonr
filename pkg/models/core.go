package models

import (
	"errors"
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
	u, err := g.FindSName(id)
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
func (g *Global) FindSName(id peer.ID) (string, error) {
	for _, v := range g.Peers {
		if v == id.String() {
			return v, nil
		}
	}
	return "", errors.New("Username not Found")
}

// Method Checks if Global has Username
func (g *Global) HasSName(u string) bool {
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
	for otherSname, id := range rg.Peers {
		if g.Sname != otherSname {
			g.Peers[otherSname] = id
		}
	}

	// Check Self Map
	if !g.HasSName(rg.Sname) {
		g.Peers[rg.Sname] = rg.UserPeerID
	}
}
