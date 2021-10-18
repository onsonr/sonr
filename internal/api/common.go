package api

import (
	"context"
	"net"
	"time"

	"github.com/kataras/golog"
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/wallet"
)

var (
	logger = golog.Child("internal/api")
)

// NodeImpl returns the NodeImpl for the Main Node
type NodeImpl interface {
	GetProfile() (*common.Profile, error)
	Peer() (*common.Peer, error)
	Close()

	OnRefresh(event *RefreshEvent)
	OnInvite(event *InviteEvent)
	OnMailbox(event *MailboxEvent)
	OnDecision(event *DecisionEvent)
	OnProgress(event *ProgressEvent)
	OnComplete(event *CompleteEvent)
}

// NodeStubImpl is the interface for the node based on mode: (client, highway)
type NodeStubImpl interface {
	Serve(ctx context.Context, listener net.Listener, ticker time.Duration)
	HasProtocols() bool
	Close() error
}

// SignedMetadataToProto converts a SignedMetadata to a protobuf.
func SignedMetadataToProto(m *wallet.SignedMetadata) *common.Metadata {
	return &common.Metadata{
		Timestamp: m.Timestamp,
		NodeId:    m.NodeId,
		PublicKey: m.PublicKey,
	}
}
