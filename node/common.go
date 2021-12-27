package node

import (
	"fmt"
	"os"
	"sync/atomic"

	"github.com/kataras/golog"
	common "github.com/sonr-io/core/common"
	"github.com/sonr-io/core/types/go/node/motor/v1"
	"github.com/sonr-io/core/wallet"
)

var (
	TEXTILE_KEY    = os.Getenv("TEXTILE_KEY")
	TEXTILE_SECRET = os.Getenv("TEXTILE_SECRET")
	LOCATION_KEY   = os.Getenv("LOCATION_KEY")
	NB_KEY         = os.Getenv("NB_KEY")
	NB_SECRET      = os.Getenv("NB_SECRET")
	logger         = golog.Default.Child("internal/api")
)

// StubMode is the type of the node (Client, Highway)
type StubMode int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	StubMode_LIB StubMode = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	StubMode_CLI

	// StubMode_BIN is the Node utilized for Desktop background process
	StubMode_BIN

	// StubMode_FULL is the Custodian Node that manages Network
	StubMode_FULL
)

// IsLib returns true if the node is a client node.
func (m StubMode) IsLib() bool {
	return m == StubMode_LIB
}

// IsBin returns true if the node is a bin node.
func (m StubMode) IsBin() bool {
	return m == StubMode_BIN
}

// IsCLI returns true if the node is a CLI node.
func (m StubMode) IsCLI() bool {
	return m == StubMode_CLI
}

// IsFull returns true if the node is a highway node.
func (m StubMode) IsFull() bool {
	return m == StubMode_FULL
}

// Motor returns true if the node has a client.
func (m StubMode) Motor() bool {
	return m.IsLib() || m.IsBin() || m.IsCLI()
}

// Highway returns true if the node has a highway stub.
func (m StubMode) Highway() bool {
	return m.IsFull()
}

// Prefix returns golog prefix for the node.
func (m StubMode) Prefix() string {
	var name string
	switch m {
	case StubMode_LIB:
		name = "lib"
	case StubMode_CLI:
		name = "cli"
	case StubMode_BIN:
		name = "bin"
	case StubMode_FULL:
		name = "highway"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
}

// CallbackImpl is the implementation of Callback interface
type CallbackImpl interface {
	// OnRefresh is called when the LobbyProtocol is refreshed and pushes a RefreshEvent
	OnRefresh(event *motor.OnLobbyRefreshResponse)

	// OnMailbox is called when the MailboxProtocol receives a MailboxEvent
	OnMailbox(event *motor.OnMailboxMessageResponse)

	// OnInvite is called when the TransferProtocol receives InviteEvent
	OnInvite(event *motor.OnTransmitInviteResponse)

	// OnDecision is called when the TransferProtocol receives a DecisionEvent
	OnDecision(event *motor.OnTransmitInviteResponse, invite *motor.OnTransmitInviteResponse)

	// OnProgress is called when the TransferProtocol sends or receives a ProgressEvent
	OnProgress(event *motor.OnTransmitProgressResponse)

	// OnTransfer is called when the TransferProtocol completes a transfer and pushes a CompleteEvent
	OnComplete(event *motor.OnTransmitCompleteResponse)
}

// NodeImpl returns the NodeImpl for the Main Node
type NodeImpl interface {
	// GetState returns the current state of the node
	GetState() *State

	// Profile returns the profile of the node from Local Store
	Profile() (*common.Profile, error)

	// Peer returns the peer of the node
	Peer() (*common.Peer, error)

	// Close closes the node
	Close()
}

// SignedMetadataToProto converts a SignedMetadata to a protobuf.
func SignedMetadataToProto(m *wallet.SignedMetadata) *common.Metadata {
	return &common.Metadata{
		Timestamp: m.Timestamp,
		NodeId:    m.NodeId,
		PublicKey: m.PublicKey,
	}
}

// State is the internal State of the API
type State struct {
	flag uint64
	Chn  chan bool
}

// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
func (c *State) NeedsWait() {
	<-c.Chn
}

// Resume tells all of goroutines to resume execution
func (c *State) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.Chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Pause tells all of goroutines to pause execution
func (c *State) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.Chn = make(chan bool)
	}
}
