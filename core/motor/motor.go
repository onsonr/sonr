package motor

import (
	"context"

	"github.com/sonr-hq/sonr/pkg/node"
	mt "github.com/sonr-hq/sonr/third_party/types/motor/bind/v1"
)

type MotorNode struct {
	// Node is the libp2p host
	Node *node.Node
	//
}

func NewMotorInstance(ctx context.Context, req *mt.InitializeRequest, options ...node.NodeOption) (*MotorNode, error) {
	n, err := node.New(ctx, options...)
	if err != nil {
		return nil, err
	}
	return &MotorNode{
		Node: n,
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
