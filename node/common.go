package node

import (
	"fmt"
	"os"

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

// SignedMetadataToProto converts a SignedMetadata to a protobuf.
func SignedMetadataToProto(m *wallet.SignedMetadata) *common.Metadata {
	return &common.Metadata{
		Timestamp: m.Timestamp,
		NodeId:    m.NodeId,
		PublicKey: m.PublicKey,
	}
}
