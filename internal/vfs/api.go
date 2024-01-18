package vfs

import (
	"github.com/asynkron/protoactor-go/actor"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
)

type GenerateRequest struct {
	AlicePID *actor.PID
	BobPID   *actor.PID
}

type GenerateResponse struct {
	PublicKey []byte
	Address   string
	CoinType  modulev1.CoinType
	DID       string
	Error     error
}
