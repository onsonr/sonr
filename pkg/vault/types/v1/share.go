package v1

import (
	"errors"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// It takes a `cmp.Config` and returns a `common.WalletShare` that can be used to create a wallet
func NewShareConfig(network string, c interface{}) *ShareConfig {
	conf := c.(*cmp.Config)
	confBz, err := conf.MarshalBinary()
	if err != nil {
		return nil
	}
	partyIds := make([]string, len(conf.PartyIDs()))
	for i, id := range conf.PartyIDs() {
		partyIds[i] = string(id)
	}
	buf, err := conf.PublicPoint().(*curve.Secp256k1Point).MarshalBinary()
	if err != nil {
		return nil
	}
	if len(buf) != 33 {
		return nil
	}
	return &ShareConfig{
		SelfId:     string(conf.ID),
		PublicKey:  buf,
		Network:    network,
		CreatedAt:  time.Now().Unix(),
		ConfigData: confBz,
	}
}

// Unmarshalling the config data and returning the config.
func (s *ShareConfig) GetCMPConfig() (*cmp.Config, error) {
	conf := &cmp.Config{}
	if err := conf.UnmarshalBinary(s.ConfigData); err != nil {
		return nil, err
	}
	return conf, nil
}

// Converting the public key from the ShareConfig to a secp256k1.PubKey.
func (s *ShareConfig) GetPubKeySecp256k1() (*secp256k1.PubKey, error) {
	if len(s.PublicKey) != 33 {
		return nil, errors.New("invalid public key length")
	}
	return &secp256k1.PubKey{Key: s.PublicKey}, nil
}

// A method that returns the party ID of the share config.
func (s *ShareConfig) PartyID() party.ID {
	return party.ID(s.SelfId)
}
