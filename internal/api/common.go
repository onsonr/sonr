package api

import (
	"os"
	sync "sync"
	"sync/atomic"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/wallet"
	common "github.com/sonr-io/core/pkg/common"
)

var (
	TEXTILE_KEY              = os.Getenv("TEXTILE_KEY")
	TEXTILE_SECRET           = os.Getenv("TEXTILE_SECRET")
	LOCATION_KEY             = os.Getenv("LOCATION_KEY")
	NB_KEY                   = os.Getenv("NB_KEY")
	NB_SECRET                = os.Getenv("NB_SECRET")
	ROLLBAR_POST_SERVER_ITEM = os.Getenv("ROLLBAR_POST_SERVER_ITEM")
	logger                   = golog.Child("internal/api")
	instance                 *state
	once                     sync.Once
)

// NodeImpl returns the NodeImpl for the Main Node
type NodeImpl interface {
	// Profile returns the profile of the node from Local Store
	Profile() (*common.Profile, error)

	// Peer returns the peer of the node
	Peer() (*common.Peer, error)

	// Close closes the node
	Close()

	// OnRefresh is called when the LobbyProtocol is refreshed and pushes a RefreshEvent
	OnRefresh(event *RefreshEvent)

	// OnMailbox is called when the MailboxProtocol receives a MailboxEvent
	OnMailbox(event *MailboxEvent)

	// OnInvite is called when the TransferProtocol receives InviteEvent
	OnInvite(event *InviteEvent)

	// OnDecision is called when the TransferProtocol receives a DecisionEvent
	OnDecision(event *DecisionEvent)

	// OnProgress is called when the TransferProtocol sends or receives a ProgressEvent
	OnProgress(event *ProgressEvent)

	// OnTransfer is called when the TransferProtocol completes a transfer and pushes a CompleteEvent
	OnComplete(event *CompleteEvent)
}

// SignedMetadataToProto converts a SignedMetadata to a protobuf.
func SignedMetadataToProto(m *wallet.SignedMetadata) *common.Metadata {
	return &common.Metadata{
		Timestamp: m.Timestamp,
		NodeId:    m.NodeId,
		PublicKey: m.PublicKey,
	}
}

// state is the internal state of the API
type state struct {
	flag uint64
	chn  chan bool
}

// GetState returns the current state of the API
func GetState() *state {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &state{chn: chn}
	})
	return instance
}

// NeedsWait Checks rather to wait or does not need
func (c *state) NeedsWait() {
	<-c.chn
}

// Resume tells all of goroutines to resume execution
func (c *state) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Pause tells all of goroutines to pause execution
func (c *state) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.chn = make(chan bool)
	}
}
