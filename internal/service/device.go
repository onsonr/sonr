package service

import (
	rpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
)

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

// @ Starts New Auth Instance
func (sc *serviceClient) StartDevices() *md.SonrError {
	// Start Exchange Server
	localServer := rpc.NewServer(sc.host.Host(), util.AUTH_PROTOCOL)
	psv := AuthService{
		user:    sc.user,
		handler: sc.handler,
		respCh:  make(chan *md.InviteResponse, util.MAX_CHAN_DATA),
	}

	// Register Service
	err := localServer.RegisterName(util.AUTH_RPC_SERVICE, &psv)
	if err != nil {
		return md.NewError(err, md.ErrorEvent_TOPIC_RPC)
	}
	sc.Auth = &psv
	return nil
}
