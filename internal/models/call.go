package models

import (
	"log"

	"google.golang.org/protobuf/proto"
)

// Define Function Types
type OnProtobuf func([]byte)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *TransferCard)
type OnTransmitted func(data *Peer)
type OnError func(err error, method string)
type ReturnPeer func() *Peer
type UpdatePeer func(peer *Peer)

type LobbyCallback struct {
	callEvent   OnProtobuf
	callRefresh OnProtobuf
	callError   OnError
	getPeer     ReturnPeer
}

// ^ Creates New Lobby Callback ^ //
func NewLobbyCallback(callEvent OnProtobuf, callRefresh OnProtobuf, callError OnError, getPeer ReturnPeer) LobbyCallback {
	return LobbyCallback{
		callEvent:   callEvent,
		callRefresh: callRefresh,
		callError:   callError,
		getPeer:     getPeer,
	}
}

// @ -- Refresh Lobby -- //
func (lc *LobbyCallback) Refresh(data *Lobby) {
	bytes, err := proto.Marshal(data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	lc.callRefresh(bytes)
}

// @ -- Call Lobby Error -- //
func (lc *LobbyCallback) Error(err error) {
	lc.callError(err, "Lobby Error")
}

// @ -- Return Current Peer Data -- //
func (lc *LobbyCallback) Peer() *Peer {
	return lc.getPeer()
}

type TransferCallback struct {
	CallInvited     OnInvite
	CallResponded   OnProtobuf
	CallProgress    OnProgress
	CallReceived    OnReceived
	CallTransmitted OnTransmitted
	CallError       OnError
}

// @ -- Call Auth Invite -- //
func (tc *TransferCallback) Invited(data []byte) {
	tc.CallResponded(data)
}

// @ -- Call Auth Responded -- //
func (tc *TransferCallback) Responded(data []byte) {
	tc.CallInvited(data)
}

// @ -- Call Transfer Progressed -- //
func (tc *TransferCallback) Progressed(data float32) {
	tc.CallProgress(data)
}

// @ -- Call Transfer Received -- //
func (tc *TransferCallback) Received(data *TransferCard) {
	tc.CallReceived(data)
}

// @ -- Call Transfer Transmitted -- //
func (tc *TransferCallback) Transmitted(data *Peer) {
	tc.CallTransmitted(data)
}

// @ -- Call Controller Error -- //
func (tc *TransferCallback) Error(err error) {
	tc.CallError(err, "Transfer Error")
}
