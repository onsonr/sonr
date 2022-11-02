package motor

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/proto"
	"github.com/libp2p/go-msgio"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/host"
	dp "github.com/sonr-io/sonr/pkg/motor/x/discover"
	tp "github.com/sonr-io/sonr/pkg/motor/x/transmit"
	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	v1 "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
)

func (mtr *motorNodeImpl) Connect(request mt.ConnectRequest) (*mt.ConnectResponse, error) {
	if mtr.SonrHost != nil && !request.ResetConnection {
		mtr.log.Warn("Host already connected")
		return &mt.ConnectResponse{
			Success: true,
			Message: "Host already connected",
		}, nil
	}

	// Setup host config
	var err error
	cnfg := config.DefaultConfig(config.Role_MOTOR, config.WithAccountAddress(mtr.GetAddress()), config.WithDeviceID(mtr.DeviceID), config.WithHomePath(mtr.homeDir), config.WithSupportPath(mtr.supportDir), config.WithTempPath(mtr.tempDir))

	// Create new host
	mtr.log.Info("Starting host...")
	mtr.SonrHost, err = host.NewDefaultHost(context.Background(), cnfg, mtr.callback)
	if err != nil {
		return nil, err
	}

	// Utilize discovery protocol
	if request.GetEnableDiscovery() {
		mtr.log.Info("Enabling Discovery...")
		mtr.discover, err = dp.New(context.Background(), mtr.SonrHost, mtr.callback)
		if err != nil {
			return nil, err
		}
	}

	// Utilize transmit protocol
	if request.GetEnableTransmit() {
		mtr.log.Info("Enabling Transmit...")
		mtr.transmit, err = tp.New(context.Background(), mtr.SonrHost)
		if err != nil {
			return nil, err
		}
	}

	mtr.log.Info("✅ Motor Host Connected")
	mtr.hostInitialized = true
	return &mt.ConnectResponse{
		Success: true,
		Message: "Successfully connected host to network",
	}, nil
}

func (mtr *motorNodeImpl) OpenLinking(request mt.LinkingRequest) (*mt.LinkingResponse, error) {
	if !mtr.IsHostActive() {
		return nil, fmt.Errorf("host inactive")
	}

	// Setup stream handler
	mtr.SonrHost.Host().SetStreamHandler(proto.ProtoIDv1, mtr.handleLinking)

	peerInfo := &peer.AddrInfo{
		ID:    mtr.SonrHost.Host().ID(),
		Addrs: mtr.SonrHost.Host().Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(peerInfo)
	if err != nil {
		return nil, fmt.Errorf("convert to p2p addrs: %s", err)
	}

	encodedAddrs := make([][]byte, len(addrs))
	for i, a := range addrs {
		encodedAddrs[i] = a.Bytes()
	}

	return &mt.LinkingResponse{
		Success:  true,
		P2PAddrs: encodedAddrs,
	}, nil
}

func (m *motorNodeImpl) PairDevice(request mt.PairingRequest) (*mt.PairingResponse, error) {
	if !m.IsHostActive() {
		return nil, fmt.Errorf("host is not active")
	}
	if len(request.P2PAddrs) == 0 {
		return nil, fmt.Errorf("missing p2p addresses")
	}
	if m.encryptionKey == nil || len(m.encryptionKey) == 0 {
		return nil, fmt.Errorf("not logged in")
	}

	var err error
	addrs := make([]ma.Multiaddr, len(request.P2PAddrs))
	for i, a := range request.P2PAddrs {
		addrs[i], err = ma.NewMultiaddrBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decode p2p multiaddr: %s", err)
		}
	}

	peerInfos, err := peer.AddrInfosFromP2pAddrs(addrs...)
	if err != nil {
		return nil, fmt.Errorf("get p2p addrs: %s", err)
	}

	for _, peerInfo := range peerInfos {
		err = m.connectToPeerAndTransmit(peerInfo, request)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, fmt.Errorf("attempt to connect to all peerInfos: %s", err)
	}

	return &mt.PairingResponse{
		Success: true,
	}, nil
}

func (mtr *motorNodeImpl) connectToPeerAndTransmit(peerInfo peer.AddrInfo, request mt.PairingRequest) error {
	if err := mtr.SonrHost.Connect(peerInfo); err != nil {
		mtr.SonrHost.Host().Peerstore().ClearAddrs(peerInfo.ID)
		return errors.Wrap(err, "failed to connect to peer while attempting to pair")
	}

	str, err := mtr.SonrHost.NewStream(context.Background(), peerInfo.ID, proto.ProtoIDv1)
	if err != nil {
		return errors.Wrap(err, "failed to open stream while attempting to pair")
	}
	defer str.Close()

	strWr := msgio.NewWriter(str)
	authInfo := &ct.AuthInfo{
		Address:   mtr.Address,
		Did:       mtr.GetDID().String(),
		AesPskKey: mtr.encryptionKey,
	}
	bz, err := authInfo.Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal auth info")
	}

	if _, err = strWr.Write(bz); err != nil {
		mtr.log.Error("failed to write auth info: %v", err)
		return errors.Wrap(err, "failed to write auth info")
	}
	return nil
}

func (mtr *motorNodeImpl) handleLinking(stream network.Stream) {
	defer stream.Close()
	r := msgio.NewReader(stream)
	bz, err := r.ReadMsg()
	if err != nil {
		mtr.log.Error("failed to read auth info: %v", err)
		return
	}

	authInfo := &ct.AuthInfo{}
	if err = authInfo.Unmarshal(bz); err != nil {
		mtr.log.Error("failed to unmarshal auth info: %v", err)
		ev := v1.LinkingEvent{
			Type: v1.LinkingEventType_LINKING_EVENT_TYPE_LINKING_FAILED,
		}
		bz, err := ev.Marshal()
		if err != nil {
			mtr.log.Error("failed to marshal linking event: %v", err)
			err = r.Close()
			if err != nil {
				mtr.log.Error("failed to close reader: %s", err)
			}
			return
		}

		mtr.callback.OnLinking(bz)
		return
	}

	mtr.log.Info("🎉 Successfully received AuthInfo!\n")
	mtr.log.Debug("%+v\n", authInfo)
	ev := v1.LinkingEvent{
		Type:     v1.LinkingEventType_LINKING_EVENT_TYPE_LINKING_COMPLETE,
		AuthInfo: authInfo,
	}

	evbz, err := ev.Marshal()
	if err != nil {
		mtr.log.Error("failed to marshal linking event: %v", err)
		return
	}
	mtr.callback.OnLinking(evbz)
	stream.Close()
}
