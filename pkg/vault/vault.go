package vault

import (
	"context"

	"github.com/ipfs/kubo/client/rpc"
)

type Vault interface{}

type vault struct {
	ctx context.Context
}

func New(ctx context.Context) (Vault, error) {
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}
	ipfsC.Unixfs()
	return &vault{
		ctx: ctx,
	}, nil
}
