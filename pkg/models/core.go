package models

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

// ** ─── CALLBACK MANAGEMENT ────────────────────────────────────────────────────────
// Define Function Types
type GetStatus func() Status
type SetStatus func(s Status)
type GetContact func() *Contact
type OnProtobuf func([]byte)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *TransferCard)
type OnTransmitted func(data *TransferCard)
type OnError func(err *SonrError)
type NodeCallback struct {
	Contact     GetContact
	Invited     OnInvite
	Refreshed   OnProtobuf
	Event       OnProtobuf
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	Status      SetStatus
	Transmitted OnTransmitted
	Error       OnError
	GetStatus   GetStatus
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

// ** ─── Router MANAGEMENT ────────────────────────────────────────────────────────
// @ Local Lobby Topic Protocol ID
func (r *Router) LocalIPTopic() string {
	return fmt.Sprintf("/sonr/topic/%s", r.Location.IPOLC())
}

func (r *Router) LocalGeoTopic() (string, error) {
	geoOlc, err := r.Location.GeoOLC()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/sonr/topic/%s", geoOlc), nil
}

// @ User Device Lobby Topic Protocol ID
func (r *Router) DeviceTopic(p *Peer) string {
	return fmt.Sprintf("/sonr/topic/%s", p.UserID())
}

// @ Transfer Controller Data Protocol ID
func (r *Router) Transfer(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/transfer/%s", id.Pretty()))
}

// @ Lobby Topic Protocol ID
func (r *Router) Topic(name string) string {
	return fmt.Sprintf("/sonr/topic/%s", name)
}

// @ Major Rendevouz Advertising Point
func (r *Router) Rendevouz() string {
	return fmt.Sprintf("/sonr/%s", r.Location.MajorOLC())
}

// ** ─── Session MANAGEMENT ────────────────────────────────────────────────────────
// ^ Prepare for Outgoing Session ^ //
func NewOutSession(p *Peer, req *InviteRequest, tc NodeCallback) *Session {
	return &Session{
		File:      req.GetFile(),
		Receiver:  req.GetTo(),
		Sender:    p,
		Index:     0,
		Direction: Session_Outgoing,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(p *Peer, inv *AuthInvite, d *Device, c NodeCallback) *Session {
	return &Session{
		File:     inv.GetFile(),
		Sender:   inv.GetFrom(),
		Receiver: p,
		//callback:     c,
		Device:    d,
		Index:     0,
		Direction: Session_Outgoing,
		// bytesBuilder: new(bytes.Buffer),
	}
}
