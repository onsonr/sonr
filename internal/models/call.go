package models

import (
	"log"

	"google.golang.org/protobuf/proto"
)

// Define Function Types
type OnProtobuf func([]byte)
type OnInvite func(data *AuthInvite)
type OnProgress func(data float32)
type OnReceived func(data *TransferCard)
type OnTransmitted func(data *Peer)
type OnError func(err error, method string)
type ReturnPeer func() *Peer

type LobbyCallback struct {
	CallEvent   OnProtobuf
	CallRefresh OnProtobuf
	CallError   OnError
	GetPeer     ReturnPeer
}

func (lc *LobbyCallback) ReturnPeer() *Peer {
	return lc.GetPeer()
}

func (lc *LobbyCallback) OnEvent(data []byte) {
	lc.CallEvent(data)
}

func (lc *LobbyCallback) OnRefresh(lob *Lobby) {
	// Marshal data to bytes
	data, err := proto.Marshal(lob)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	lc.CallRefresh(data)
}

func (lc *LobbyCallback) OnError(err error, method string) {
	lc.CallError(err, method)
}

type TransferCallback struct {
	CallInvited     OnInvite
	CallResponded   OnProtobuf
	CallProgress    OnProgress
	CallReceived    OnReceived
	CallTransmitted OnTransmitted
	CallError       OnError
}

func (tc *TransferCallback) OnInvited(data *AuthInvite) {
	tc.CallInvited(data)
}

func (tc *TransferCallback) OnResponded(data []byte) {
	tc.CallResponded(data)
}

func (tc *TransferCallback) OnProgress(data float32) {
	tc.CallProgress(data)
}

func (tc *TransferCallback) OnReceived(data *TransferCard) {
	tc.CallReceived(data)
}

func (tc *TransferCallback) OnTransmitted(data *Peer) {
	tc.CallTransmitted(data)
}

func (tc *TransferCallback) OnError(err error, method string) {
	tc.CallError(err, method)
}
