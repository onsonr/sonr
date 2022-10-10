package motor

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/pkg/errors"
	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	v1 "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
)

func (m *motorNodeImpl) OpenLinking(request mt.LinkingRequest) (*mt.LinkingResponse, error) {
	if !m.IsHostActive() {
		return nil, fmt.Errorf("host is not active")
	}

	m.log.Info("Enabling Linking StreamHandler...")
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate uuid")
	}
	protocolId := protocol.ID(fmt.Sprintf("/sonr/link/%s/%s", request.GetDeviceId(), id.String()))
	m.SonrHost.SetStreamHandler(protocolId, m.handleLinking)

	ai := m.SonrHost.AddrInfo(protocolId)
	if request.Filename != "" {
		err := ai.WriteQrCodeToFile(request.Filename)
		if err != nil {
			return nil, errors.Wrap(err, "failed to write qr code")
		}

		return &mt.LinkingResponse{
			Success: true,
		}, nil
	} else {
		qrBz, err := ai.WriteQrCode()
		if err != nil {
			return nil, errors.Wrap(err, "failed to write qr code")
		}

		return &mt.LinkingResponse{
			Success: true,
			QrCode:  qrBz,
		}, nil
	}
}
func (m *motorNodeImpl) PairDevice(request mt.PairDeviceRequest) (*mt.PairDeviceResponse, error) {
	if !m.IsHostActive() {
		return nil, fmt.Errorf("host is not active")
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}
	if request.AddrInfo != nil {
		peerInfo, err := request.AddrInfo.ToLibp2pAddrInfo()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert addr info")
		}
		err = m.connectToPeerAndTransmit(peerInfo, request)
		if err != nil {
			return nil, err
		}
	}
	if request.AddrInfoBase64 != "" {
		ai, err := ct.AddrInfoFromBase64(request.AddrInfoBase64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode addr info")
		}
		peerInfo, err := ai.ToLibp2pAddrInfo()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert addr info")
		}
		err = m.connectToPeerAndTransmit(peerInfo, request)
		if err != nil {
			return nil, err
		}
	}
	if request.PeerId != "" {
		peerId, err := peer.Decode(request.PeerId)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode peer id")
		}
		peerInfo := peer.AddrInfo{ID: peerId}
		err = m.connectToPeerAndTransmit(peerInfo, request)
		if err != nil {
			return nil, err
		}
	}
	return &mt.PairDeviceResponse{
		Success: true,
		Message: "Succesfully paired device via Key Exchange Linking",
	}, nil
}

func (m *motorNodeImpl) connectToPeerAndTransmit(peerInfo peer.AddrInfo, request mt.PairDeviceRequest) error {
	err := m.SonrHost.Connect(peerInfo)
	if err != nil {
		m.SonrHost.Host().Peerstore().ClearAddrs(peerInfo.ID)
		return errors.Wrap(err, "failed to connect to peer while attempting to pair")
	}
	linkPid, err := request.AddrInfo.GetLinkProtocolId()
	if err != nil {
		return errors.Wrap(err, "failed to get link protocol id of peer while attempting to pair")
	}

	str, err := m.SonrHost.NewStream(context.Background(), peerInfo.ID, linkPid)
	if err != nil {
		return errors.Wrap(err, "failed to open stream while attempting to pair")
	}
	defer str.Close()
	strWr := msgio.NewWriter(str)
	authInfo := &ct.AuthInfo{
		Address:   m.Address,
		Did:       m.GetDID().String(),
		AesPskKey: request.AesDscKey,
		AesDscKey: request.AesPskKey,
	}
	bz, err := authInfo.Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal auth info")
	}

	_, err = strWr.Write(bz)
	if err != nil {
		m.log.Error("failed to write auth info: %v", err)
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
	err = authInfo.Unmarshal(bz)
	if err != nil {
		mtr.log.Error("failed to unmarshal auth info: %v", err)
		ev := v1.LinkingEvent{
			Type: v1.LinkingEventType_LINKING_EVENT_TYPE_LINKING_FAILED,
		}
		bz, err := ev.Marshal()
		if err != nil {
			mtr.log.Error("failed to marshal linking event: %v", err)
			err = r.Close()
			if err != nil {
				mtr.log.Error("failed to close reader", err)
			}
			return
		}
		mtr.callback.OnLinking(bz)
		return
	}

	fmt.Print("Successfully received AuthInfo!")
	ev := v1.LinkingEvent{
		Type:     v1.LinkingEventType_LINKING_EVENT_TYPE_LINKING_COMPLETE,
		AuthInfo: authInfo,
	}

	evbz, err := ev.Marshal()
	if err != nil {
		fmt.Printf("failed to marshal linking event: %v", err)
		return
	}
	mtr.callback.OnLinking(evbz)
}
