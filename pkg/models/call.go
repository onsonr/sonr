package models

// Define Function Types
type OnProtobuf func([]byte)
type OnQueued func(card *TransferCard, req *InviteRequest)
type OnMultiQueued func(card *TransferCard, req *InviteRequest, count int)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *TransferCard)
type OnTransmitted func(data *Peer)
type OnError func(err error, method string)
type ReturnPeer func() *Peer
type UpdatePeer func(peer *Peer)

type LobbyCallback struct {
	Event   OnProtobuf
	Refresh OnProtobuf
	Error   OnError
	Peer    ReturnPeer
}

// ^ Creates New Lobby Callback ^ //
func NewLobbyCallback(callEvent OnProtobuf, callRefresh OnProtobuf, callError OnError, getPeer ReturnPeer) LobbyCallback {
	return LobbyCallback{
		Event:   callEvent,
		Refresh: callRefresh,
		Error:   callError,
		Peer:    getPeer,
	}
}

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
	Queued      OnQueued
	MultiQueued OnMultiQueued
	Error       OnError
}

// ^ Creates New File Callback ^ //
func NewFileCallback(callQueued OnQueued, callMultiQueued OnMultiQueued, callError OnError) FileCallback {
	return FileCallback{
		Queued:      callQueued,
		MultiQueued: callMultiQueued,
		Error:       callError,
	}
}
