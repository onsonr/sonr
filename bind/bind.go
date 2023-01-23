package motor

import (
	"context"

	"github.com/sonrhq/core/bind/internal"
	_ "golang.org/x/mobile/bind"
)

var (
	ctx context.Context
	mtr *internal.MotorInstance
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(data []byte)
	OnLinking(data []byte)
}

func Init(cb MotorCallback) ([]byte, error) {
	ctx = context.Background()
	mtdr, err := internal.NewMotorInstance(ctx)
	if err != nil {
		return nil, err
	}
	mtr = mtdr
	return nil, nil
}
