package mpc

import (
	"fmt"

	"github.com/onsonr/sonr/crypto/keys"
)

type KeyEnclave map[string]string

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
	msg, err := valShare.Message()
	if err != nil {
		return nil, err
	}
	fmt.Println(msg)
	enclave := make(KeyEnclave)
	pubPoint, err := getAlicePubPoint(msg)
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

	enclave[kAddrEnclaveKey] = [](byte)(addr)
	enclave[kPubKeyEnclaveKey] = ppJSON
	enclave[kValEnclaveKey] = valShare.Bytes()
	enclave[kUserEnclaveKey] = userShare.Bytes()
	return enclave, nil
}

func (k KeyEnclave) Address() string {
	addr := k[kAddrEnclaveKey]
	if addr == nil {
		return ""
	}
	return string(addr)
}

func (k KeyEnclave) PubKey() keys.PubKey {
	ppbz, ok := k[kPubKeyEnclaveKey]
	if !ok {
		return nil
	}
	pp, err := unmarshalPointJSON(ppbz)
	if err != nil {
		return nil
	}
	return keys.NewPubKey(pp)
}

func (k KeyEnclave) Sign(data []byte) ([]byte, error) {
	ukstr, ok := k[kUserEnclaveKey]
	if !ok {
		return nil, fmt.Errorf("user share not found")
	}
	uks, err := DecodeKeyshare(string(ukstr))
	if err != nil {
		return nil, err
	}
	vkstr, ok := k[kValEnclaveKey]
	if !ok {
		return nil, fmt.Errorf("validator share not found")
	}
	vks, err := DecodeKeyshare(string(vkstr))
	if err != nil {
		return nil, err
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
