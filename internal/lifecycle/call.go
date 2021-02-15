package lifecycle

import (
	md "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnProtobuf func([]byte)
type OnInvite func(data *md.AuthInvite)
type OnProgress func(data float32)
type OnReceived func(data *md.TransferCard)
type OnTransmitted func(data *md.Peer)
type OnError func(err error, method string)
type ReturnPeer func() *md.Peer

type LobbyCallbacks struct {
	CallEvent   OnProtobuf
	CallRefresh OnProtobuf
	CallError   OnError
	GetPeer     ReturnPeer
}

type TransferCallbacks struct {
	CallInvited     OnInvite
	CallResponded   OnProtobuf
	CallProgress    OnProgress
	CallReceived    OnReceived
	CallTransmitted OnTransmitted
	CallError       OnError
}
