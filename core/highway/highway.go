package highway

import (
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/ipfs"
)

type HighwayNode struct {
	// Node is the libp2p host
	Node   *ipfs.IPFS
	Wallet common.Wallet
	//
}
