package service

import (
	rpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
)

// LocalServiceArgs ExchangeArgs is Peer protobuf
type DeviceServiceArgs struct {
	Direct []byte
}

// LocalServiceResponse ExchangeResponse is also Peer protobuf
type DeviceServiceResponse struct {
	Result []byte
}

type DeviceService struct {
	ServiceClient
	handler ServiceHandler
	device  *md.Device
	respCh  chan *md.InviteResponse
	invite  *md.InviteRequest
}

// Starts New Auth Instance
func (sc *serviceClient) StartDevices() *md.SonrError {
	// Start Exchange Server
	localServer := rpc.NewServer(sc.host.Host(), util.AUTH_PROTOCOL)
	dsv := DeviceService{
		device:  sc.device,
		handler: sc.handler,
		respCh:  make(chan *md.InviteResponse, util.MAX_CHAN_DATA),
	}

	// Register Service
	err := localServer.RegisterName(util.DEVICE_RPC_SERVICE, &dsv)
	if err != nil {
		return md.NewError(err, md.ErrorEvent_ROOM_RPC)
	}
	sc.Device = &dsv
	return nil
}
