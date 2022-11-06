package motor

import (
	"fmt"

	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	_ "golang.org/x/mobile/bind"
)

func Connect(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.ConnectRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.Connect(request)
	if err != nil {
		return nil, err
	}
	return resp.Marshal()
}

func OpenLinking(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}
	var request mt.LinkingRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.OpenLinking(request)
	if err != nil {
		return nil, err
	}
	return resp.Marshal()
}

func PairDevice(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.PairingRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.PairDevice(request)
	if err != nil {
		return nil, err
	}
	return resp.Marshal()
}
