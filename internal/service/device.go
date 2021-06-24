package service

import md "github.com/sonr-io/core/pkg/models"

// LocalServiceArgs ExchangeArgs is Peer protobuf
type DeviceServiceArgs struct {
	Peer   []byte
	Invite []byte
}

// LocalServiceResponse ExchangeResponse is also Peer protobuf
type DeviceServiceResponse struct {
	InvReply []byte
	Peer     []byte
}

type DeviceService struct {
	ServiceClient
	handler ServiceHandler
	user    *md.User
	respCh  chan *md.InviteResponse
	invite  *md.InviteRequest
}
