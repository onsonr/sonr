package actor

import (
	"context"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/di-dao/sonr/crypto/kss"
)

type SignRequest struct {
	Msg []byte
}

type SignResponse struct {
	Sig []byte
	Msg []byte
}

type RefreshRequest struct{}

type RefreshResponse struct{}

type ValidatorController interface {
	GetSignFunc(ctx context.Context, msg []byte) (kss.SignFuncVal, error)
	GetRefreshFunc(ctx context.Context) (kss.RefreshFuncVal, error)
	VerifySignature(ctx context.Context, msg []byte, sig []byte) (bool, error)
}

type validatorActor struct {
	actor.Actor
}

func (v *validatorActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	}
}
