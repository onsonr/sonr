package api

import (
	"github.com/kataras/golog"
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/keychain"
)

var (
	logger = golog.Child("internal/api")
)

// NodeImpl returns the NodeImpl for the Main Node
type NodeImpl interface {
	GetProfile() (*common.Profile, error)
	Peer() (*common.Peer, error)

	OnRefresh(event *RefreshEvent)
	OnInvite(event *InviteEvent)
	OnMailbox(event *MailboxEvent)
	OnDecision(event *DecisionEvent)
	OnProgress(event *ProgressEvent)
	OnComplete(event *CompleteEvent)
}

// SignedMetadataToProto converts a SignedMetadata to a protobuf.
func SignedMetadataToProto(m *keychain.SignedMetadata) *common.Metadata {
	return &common.Metadata{
		Timestamp: m.Timestamp,
		NodeId:    m.NodeId,
		PublicKey: m.PublicKey,
	}
}
