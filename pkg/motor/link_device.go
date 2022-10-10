package motor

import (
	"fmt"

	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

func (m *motorNodeImpl) OpenLinking(request mt.LinkingRequest) (*mt.LinkingResponse, error) {
	if !m.IsHostActive() {
		return nil, fmt.Errorf("host is not active")
	}
	addrs := m.SonrHost.Host().Addrs()
	if len(addrs) == 0 {
		return nil, fmt.Errorf("no addresses available")
	}

	id := m.SonrHost.Host().ID()
	if id == "" {
		return nil, fmt.Errorf("no id available")
	}
	addrsStr := make([]string, len(addrs))
	for i, addr := range addrs {
		addrsStr[i] = addr.String()
	}

	return &mt.LinkingResponse{
		AddrInfo: &ct.AddrInfo{
			Id:    id.String(),
			Addrs: addrsStr,
		},
	}, nil
}
func (m *motorNodeImpl) PairDevice(request mt.PairDeviceRequest) (*mt.PairDeviceResponse, error) {
	if !m.IsHostActive() {
		return nil, fmt.Errorf("host is not active")
	}
}
