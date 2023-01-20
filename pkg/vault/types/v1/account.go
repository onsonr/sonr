package v1

import (
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

func NewAccountConfigFromShares(name string, index uint32, addrPrefix string, shares []*ShareConfig) (*AccountConfig, error) {
	addr, err := bech32.ConvertAndEncode(addrPrefix, shares[0].PublicKey)
	if err != nil {
		return nil, err
	}
	return &AccountConfig{
		Name:    name,
		Index:   index,
		Address: addr,
		Shares:  shares,
	}, nil
}

func (a *AccountConfig) GetConfigMap() map[party.ID]*cmp.Config {
	configMap := make(map[party.ID]*cmp.Config)
	for _, s := range a.Shares {
		conf, err := s.GetCMPConfig()
		if err != nil {
			continue
		}
		configMap[s.PartyID()] = conf
	}
	return configMap
}

func (a *AccountConfig) PartyIDs() []party.ID {
	ids := make([]party.ID, 0, len(a.Shares))
	for _, share := range a.Shares {
		ids = append(ids, party.ID(share.SelfId))
	}
	return ids
}

func (a *AccountConfig) PublicPoint() (curve.Point, error) {
	conf, err := a.Shares[0].GetCMPConfig()
	if err != nil {
		return nil, err
	}
	return conf.PublicPoint(), nil
}
