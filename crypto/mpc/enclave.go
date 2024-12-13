package mpc

import (
	"fmt"

	"github.com/onsonr/sonr/crypto/keys"
)

// Enclave defines the interface for key management operations
type Enclave interface {
	Address() string
	PubKey() keys.PubKey
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) (bool, error)
}

// KeyEnclave implements the Enclave interface
type KeyEnclave struct {
	Addr       string `json:"address"`
	PubKeyData  string `json:"pub_key"`
	ValShare    string `json:"val_share"`
	UserShare   string `json:"user_share"`
	VaultCID    string `json:"vault_cid,omitempty"`
}

func initKeyEnclave(valShare, userShare KeyShare) (*KeyEnclave, error) {
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

	return &KeyEnclave{
		Address:    addr,
		PubKeyData: string(ppJSON),
		ValShare:   valShare.String(),
		UserShare:  userShare.String(),
	}, nil
}

func (k *KeyEnclave) Address() string {
	return k.Addr
}

func (k *KeyEnclave) PubKey() keys.PubKey {
	pp, err := unmarshalPointJSON([]byte(k.PubKeyData))
	if err != nil {
		return nil
	}
	return keys.NewPubKey(pp)
}

func (k *KeyEnclave) Sign(data []byte) ([]byte, error) {
	uks, err := DecodeKeyshare(k.UserShare)
	if err != nil {
		return nil, err
	}
	vks, err := DecodeKeyshare(k.ValShare)
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

func (k *KeyEnclave) Verify(data []byte, sig []byte) (bool, error) {
	return k.PubKey().Verify(data, sig)
}
