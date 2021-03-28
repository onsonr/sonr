package state

import (
	md "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnProtobuf func([]byte)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *md.TransferCard)
type OnTransmitted func(data *md.Peer)
type OnError func(err error, method string)
type ReturnPeer func() *md.Peer
type SyncLobby func(ref *md.Lobby, peer *md.Peer)
type NodeCallback struct {
	Invited     OnInvite
	Refreshed   OnProtobuf
	Event       OnProtobuf
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	GetPeer     ReturnPeer
	Transmitted OnTransmitted
	Error       OnError
}
