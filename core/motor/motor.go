package motor

import (
	"context"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/ipfs"
	"github.com/sonr-hq/sonr/pkg/network"
	mt "github.com/sonr-hq/sonr/third_party/types/motor/bind/v1"
)

type MotorNode struct {
	// Node is the libp2p host
	Node   *ipfs.IPFS
	Wallet common.Wallet
	//
}

func NewMotorInstance(ctx context.Context, req *mt.InitializeRequest, options ...ipfs.NodeOption) (*MotorNode, error) {
	n, err := ipfs.New(ctx, options...)
	if err != nil {
		return nil, err
	}
	w, err := network.NewWallet()
	if err != nil {
		return nil, err
	}

	return &MotorNode{
		Node:   n,
		Wallet: w,
	}, nil
}

func (mi *MotorNode) Connect(req *mt.ConnectRequest) (*mt.ConnectResponse, error) {
	err := mi.Node.Connect(req.GetMultiaddr())
	if err != nil {
		return nil, err
	}
	return &mt.ConnectResponse{
		Success: true,
	}, nil
}
