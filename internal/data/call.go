package data

import (
	md "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnBool func(bool)
type OnProtobuf func([]byte)
type OnQueued func(card *md.TransferCard, req *md.InviteRequest)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *md.TransferCard)
type OnTransmitted func(data *md.Peer)
type OnError func(err error, method string)
type ReturnPeer func() *md.Peer
type ReturnBuf func() []byte
type SyncLobby func(ref *md.Lobby, peer *md.Peer)

type TransferCallback struct {
	Invited     OnInvite
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	Transmitted OnTransmitted
	Error       OnError
}

// ^ Creates New Transfer Callback ^ //
func NewTransferCallback(callInvited OnInvite, callRemote OnProtobuf, callResponded OnProtobuf, callProgress OnProgress, callReceived OnReceived, callTransmitted OnTransmitted, callError OnError) TransferCallback {
	return TransferCallback{
		Invited:     callInvited,
		RemoteStart: callRemote,
		Responded:   callResponded,
		Progressed:  callProgress,
		Received:    callReceived,
		Transmitted: callTransmitted,
		Error:       callError,
	}
}

type FileCallback struct {
	Queued OnQueued
	Error  OnError
}

type NodeCallback struct {
	Connected   OnBool
	Ready       OnBool
	Invited     OnInvite
	Refreshed   OnProtobuf
	Event       OnProtobuf
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	Transmitted OnTransmitted
	Queued      OnQueued
	Error       OnError
}
