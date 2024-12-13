package mpc

import (
	"github.com/onsonr/sonr/crypto/keys"
)

type KeyEnclave map[string]interface{}

const (
	kUserEnclaveKey      = "user"
	kValEnclaveKey       = "val"
	kAddrEnclaveKey      = "addr"
	kIssuerEnclaveKey    = "issuer"
	kPubKeyEnclaveKey    = "pub-point"
	kChainCodeEnclaveKey = "chain-code"
)

func initKeyEnclave(shares ...KeyShare) (KeyEnclave, error) {
	enclave := make(KeyEnclave)
	pubPoint, err := getKeyShareArrayPoint(shares)
	if err != nil {
		return nil, err
	}
	addr, err := computeSonrAddr(pubPoint)
	if err != nil {
		return nil, err
	}
	ppJSON, err := marshalPointJSON(pubPoint)
	if err != nil {
		return nil, err
	}
	enclave[kAddrEnclaveKey] = addr
	enclave[kPubKeyEnclaveKey] = ppJSON
	for _, share := range shares {
		if share.Role() == RoleUnknown {
			continue
		}
		if share.Role() == RoleUser {
			enclave[kUserEnclaveKey] = share
		}
		if share.Role() == RoleValidator {
			enclave[kValEnclaveKey] = share
		}
	}
	return enclave, nil
}

func (k KeyEnclave) Address() string {
	return k[kAddrEnclaveKey].(string)
}

func (k KeyEnclave) PubKey() keys.PubKey {
	pp, err := unmarshalPointJSON(k[kPubKeyEnclaveKey].([]byte))
	if err != nil {
		return nil
	}
	return keys.NewPubKey(pp, keys.DIDMethodKey)
}
