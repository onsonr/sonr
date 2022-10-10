package v1

import (
	"errors"

	protocol "github.com/libp2p/go-libp2p-core/protocol"
)

func (r *PairDeviceRequest) Validate() error {
	if r.ProtocolId == "" {
		return errors.New("protocol id is required")
	}
	if r.AddrInfo != nil {
		return nil
	}
	if r.AddrInfoBase64 != "" {
		return nil
	}
	if r.PeerId != "" {
		return nil
	}
	return errors.New("Request does not provide a topic name or address info")
}

func (r *PairDeviceRequest) GetLibp2pProtocolId() protocol.ID {
	return protocol.ID(r.ProtocolId)
}
