package vfs

import (
	"github.com/asynkron/protoactor-go/actor"

	"github.com/sonrhq/sonr/crypto/core/protocol"
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

type ProtocolMessage struct {
	Value *protocol.Message
	Error error
	To    *actor.PID
}

func NewProtocolMessage(msg *protocol.Message, to *actor.PID) *ProtocolMessage {
	return &ProtocolMessage{
		Value: msg,
		To:    to,
	}
}
