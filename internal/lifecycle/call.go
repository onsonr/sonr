package lifecycle

import (
	md "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnProtobuf func([]byte)
type OnQueued func(card *md.TransferCard, req *md.InviteRequest)
type OnProgress func(data float32)
type OnError func(err error, method string)
type ReturnPeer func() *md.Peer

type LobbyCallbacks struct {
	CallEvent   OnProtobuf
	CallRefresh OnProtobuf
	CallError   OnError
	GetPeer     ReturnPeer
}

type TransferCallbacks struct {
	CallInvited     OnProtobuf
	CallReceived    OnProtobuf
	CallResponded   OnProtobuf
	CallProgress    OnProgress
	CallTransmitted OnProtobuf
	CallError       OnError
}

type ProcessCallbacks struct {
	CallQueued OnQueued
	CallError  OnError
}

type TransferFileCallbacks struct {
	CallProgress OnProgress
	CallComplete OnProtobuf
	CallError    OnError
}
