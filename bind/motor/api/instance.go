package api

import (
	"context"
)

type MotorInstance struct {
	// Node is the libp2p host
	ctx context.Context
	//
}

func NewMotorInstance(ctx context.Context) *MotorInstance {
	return &MotorInstance{
		ctx: ctx,
	}
}
