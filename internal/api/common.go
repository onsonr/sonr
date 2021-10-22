package api

import (
	sync "sync"
	"sync/atomic"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/wallet"
	common "github.com/sonr-io/core/pkg/common"
)

var (
	logger = golog.Child("internal/api")
)

// NodeImpl returns the NodeImpl for the Main Node
type NodeImpl interface {
	Profile() (*common.Profile, error)
	Peer() (*common.Peer, error)
	Close()

	OnRefresh(event *RefreshEvent)
	OnInvite(event *InviteEvent)
	OnMailbox(event *MailboxEvent)
	OnDecision(event *DecisionEvent)
	OnProgress(event *ProgressEvent)
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
