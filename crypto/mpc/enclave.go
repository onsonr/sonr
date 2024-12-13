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
	kVaultCIDKey         = "vault-cid"
)

func initKeyEnclave(valShare, userShare KeyShare) (KeyEnclave, error) {
	if valShare.Role() != RoleValidator {
		return nil, fmt.Errorf("first argument must be validator share")
	}
	if userShare.Role() != RoleUser {
		return nil, fmt.Errorf("second argument must be user share")
	}

	enclave := make(KeyEnclave)
	pubPoint, err := getKeyShareArrayPoint([]KeyShare{valShare, userShare})
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
	enclave[kValEnclaveKey] = valShare
	enclave[kUserEnclaveKey] = userShare
	
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
	return keys.NewPubKey(pp)
}

func (k KeyEnclave) Sign(data []byte) ([]byte, error) {
	uks, ok := k[kUserEnclaveKey].(KeyShare)
	if !ok {
		return nil, nil
	}
	vks, ok := k[kValEnclaveKey].(KeyShare)
	if !ok {
		return nil, nil
	}
	userSign, err := getSignFunc(uks, data)
	if err != nil {
		return nil, err
	}
	valSign, err := getSignFunc(vks, data)
	if err != nil {
		return nil, err
	}
	return ExecuteSigning(valSign, userSign)
}

func (k KeyEnclave) Verify(data []byte, sig []byte) (bool, error) {
	return k.PubKey().Verify(data, sig)
}
