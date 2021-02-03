package lifecycle

import (
	md "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnProtobuf func([]byte)
type GetUserPeer func() *md.Peer
type OnError func(err error, method string)
type OnProgress func(data float32)

type LobbyCallbacks struct {
	CallEvent   OnProtobuf
	CallRefresh OnProtobuf
	CallError   OnError
	GetPeer     GetUserPeer
}

type TransferCallbacks struct {
	CallInvited     OnProtobuf
	CallReceived    OnProtobuf
	CallResponded   OnProtobuf
	CallProgress    OnProgress
	CallTransmitted OnProtobuf
	CallError       OnError
}
