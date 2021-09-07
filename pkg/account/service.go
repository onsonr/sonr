package account

import (
	"context"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"go.uber.org/zap"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"github.com/sonr-io/core/tools/logger"
)

// DeviceServiceArgs is request to Manage/Verify events
type DeviceServiceArgs struct {
	PubKeyBuf  []byte
	RequestBuf []byte
}

// DeviceServiceResponse is response to Manage/Verify events
type DeviceServiceResponse struct {
	Success     bool
	ResponseBuf []byte
}

// DeviceService Service Struct
type DeviceService struct {
	// Current Data
	account Account
	room    GetRoomFunc
}

// Initialize Exchange Service by Room Type
func (rm *userLinker) initService() *data.SonrError {
	// Start Exchange RPC Server
	verifyServer := rpc.NewServer(rm.host.Host(), util.ACCOUNT_PROTOCOL)
	verifyService := DeviceService{
		account: rm,
		room:    rm.Room,
	}

	// Register Service
	err := verifyServer.RegisterName(util.DEVICE_RPC_SERVICE, &verifyService)
	if err != nil {
		return data.NewError(err, data.ErrorEvent_ROOM_RPC)
	}

	// Set Service
	rm.service = &verifyService
	go rm.handleTopicEvents(rm.ctx)
	go rm.handleTopicMessages(rm.ctx)
	return nil
}

// Exchange @ Starts Exchange on Local Peer Join
func (rm *userLinker) Verify(id peer.ID) error {
	// Initialize RPC
	exchClient := rpc.NewClient(rm.host.Host(), util.ACCOUNT_PROTOCOL)
	var reply DeviceServiceResponse
	var args DeviceServiceArgs
	args.PubKeyBuf = rm.user.GetCurrent().DevicePubKeyBuf()

	// Verify with Peer
	err := exchClient.Call(id, util.DEVICE_RPC_SERVICE, util.DEVICE_METHOD_VERIFY, args, &reply)
	if err != nil {
		logger.Error("Failed to Call VerifyWith Method", zap.Error(err))
		return err
	}

	// Check for Success
	if !reply.Success {
		logger.Warn("Failed to verify device")
		rm.topic.Close()
	}
	return nil
}

// ExchangeWith # Calls Exchange on Local Lobby Peer
func (ss *DeviceService) VerifyWith(ctx context.Context, args DeviceServiceArgs, reply *DeviceServiceResponse) error {
	// Unmarshal Public Key
	pubKey, err := crypto.UnmarshalPublicKey(args.PubKeyBuf)
	if err != nil {
		logger.Error("Failed to Unmarshal PublicKey", zap.Error(err))
		return err
	}

	// Check if Public Keys Match
	reply.Success = ss.account.CurrentDeviceKeys().VerifyPubKey(pubKey)
	return nil
}
