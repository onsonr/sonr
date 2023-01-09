package ipfs

import (
	"context"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-hq/sonr/pkg/node/ipfs/local"
	"github.com/sonr-hq/sonr/pkg/node/ipfs/remote"
)

type IPFS interface {
	// Add
	Add(data []byte) (string, error)

	Connect(peers ...string) error

	CoreAPI() icore.CoreAPI

	// Get
	Get(hash string) ([]byte, error)

	MultiAddr() string

	PeerID() peer.ID

	GetDecrypted(cidStr string, pubKey []byte) ([]byte, error)

	AddEncrypted(file []byte, pubKey []byte) (string, error)
}

// NewLocal creates a new local IPFS node
func NewLocal(ctx context.Context, opts ...local.NodeOption) (IPFS, error) {
	return local.New(ctx, opts...)
}

// NewRemote creates a new remote IPFS node
func NewRemote(ctx context.Context) (IPFS, error) {
	return remote.NewApi(ctx)
}
