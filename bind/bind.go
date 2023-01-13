package motor

import (
	"context"

	"github.com/sonr-hq/sonr/bind/motor"
	_ "golang.org/x/mobile/bind"
)

var (
	ctx context.Context
	mtr *motor.MotorInstance
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(data []byte)
	OnLinking(data []byte)
}

func Init(cb MotorCallback) ([]byte, error) {
	ctx = context.Background()
	mtdr, err := motor.NewMotorInstance(ctx)
	if err != nil {
		return nil, err
	}
	mtr = mtdr
	return nil, nil
}
