package api

import (
	"context"

	mt "github.com/sonr-io/sonr/bind/motor/types/v1"
	"github.com/sonr-io/sonr/internal/node"
)

type MotorInstance struct {
	// Node is the libp2p host
	Node node.Node
	ctx  context.Context
	//
}

func NewMotorInstance(ctx context.Context) *MotorInstance {
	return &MotorInstance{
		ctx: ctx,
	}
}

func (mi *MotorInstance) Connect(req *mt.ConnectRequest, options ...node.NodeOption) (*mt.ConnectResponse, error) {
	n, err := node.New(mi.ctx, options...)
	if err != nil {
		return nil, err
	}
	mi.Node = n
	return &mt.ConnectResponse{}, nil
}
