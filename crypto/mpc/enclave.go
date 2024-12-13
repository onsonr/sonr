package mpc

import (
	"fmt"

	"github.com/onsonr/sonr/crypto/keys"
)

type KeyEnclave map[string]string

const (
	kUserEnclaveKey = "user"
	kValEnclaveKey  = "val" 
	kAddrEnclaveKey = "addr"
	kPubKeyKey      = "pub"
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

	enclave[kAddrEnclaveKey] = addr
	enclave[kPubKeyKey] = ppJSON
	enclave[kValEnclaveKey] = valShare.String()
	enclave[kUserEnclaveKey] = userShare.String()
	return enclave, nil
}

func (k KeyEnclave) Address() string {
	return k[kAddrEnclaveKey]
}

func (k KeyEnclave) PubKey() keys.PubKey {
	ppbz, ok := k[kPubKeyKey]
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
	userShare, ok := k[kUserEnclaveKey]
	if !ok {
		return nil, fmt.Errorf("user share not found")
	}
	uks, err := DecodeKeyshare(userShare)
	if err != nil {
		return nil, err
	}
	valShare, ok := k[kValEnclaveKey]
	if !ok {
		return nil, fmt.Errorf("validator share not found")
	}
	vks, err := DecodeKeyshare(valShare)
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
