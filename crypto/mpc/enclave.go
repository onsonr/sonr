package mpc

import (
	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/keys"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1"
	"golang.org/x/crypto/sha3"
)

// Enclave defines the interface for key management operations
type Enclave interface {
	Address() string
	PubKey() keys.PubKey
	Refresh() (Enclave, error)
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) (bool, error)
}

// KeyEnclave implements the Enclave interface
type KeyEnclave struct {
	Addr       string `json:"address"`
	PubKeyData string `json:"pub_key"`
	ValShare   string `json:"val_share"`
	UserShare  string `json:"user_share"`
	VaultCID   string `json:"vault_cid,omitempty"`
}

func initKeyEnclave(valShare, userShare Message) (*KeyEnclave, error) {
	pubPoint, err := getAlicePubPoint(valShare)
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
	valKs, err := protocol.EncodeMessage(valShare)
	if err != nil {
		return nil, err
	}
	userKs, err := protocol.EncodeMessage(userShare)
	if err != nil {
		return nil, err
	}

	return &KeyEnclave{
		Addr:       addr,
		PubKeyData: string(ppJSON),
		ValShare:   valKs,
		UserShare:  userKs,
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

func (k *KeyEnclave) Refresh() (Enclave, error) {
	refreshFuncVal, err := k.valRefreshFunc()
	if err != nil {
		return nil, err
	}
	refreshFuncUser, err := k.userRefreshFunc()
	if err != nil {
		return nil, err
	}
	return ExecuteRefresh(refreshFuncVal, refreshFuncUser)
}

func (k *KeyEnclave) Sign(data []byte) ([]byte, error) {
	userSign, err := k.userSignFunc(data)
	if err != nil {
		return nil, err
	}
	valSign, err := k.valSignFunc(data)
	if err != nil {
		return nil, err
	}
	return ExecuteSigning(valSign, userSign)
}

func (k *KeyEnclave) Verify(data []byte, sig []byte) (bool, error) {
	return k.PubKey().Verify(data, sig)
}

func (k *KeyEnclave) userShare() (Message, error) {
	return protocol.DecodeMessage(k.UserShare)
}

func (k *KeyEnclave) userSignFunc(bz []byte) (SignFunc, error) {
	curve := curves.K256()
	msg, err := k.userShare()
	if err != nil {
		return nil, err
	}
	return dklsv1.NewBobSign(curve, sha3.New256(), bz, msg, protocol.Version1)
}

func (k *KeyEnclave) userRefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	msg, err := k.userShare()
	if err != nil {
		return nil, err
	}
	return dklsv1.NewBobRefresh(curve, msg, protocol.Version1)
}

func (k *KeyEnclave) valShare() (Message, error) {
	return protocol.DecodeMessage(k.ValShare)
}

func (k *KeyEnclave) valSignFunc(bz []byte) (SignFunc, error) {
	curve := curves.K256()
	msg, err := k.valShare()
	if err != nil {
		return nil, err
	}
	return dklsv1.NewAliceSign(curve, sha3.New256(), bz, msg, protocol.Version1)
}

func (k *KeyEnclave) valRefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	msg, err := k.valShare()
	if err != nil {
		return nil, err
	}
	return dklsv1.NewAliceRefresh(curve, msg, protocol.Version1)
}
