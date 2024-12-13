package mpc

import (
	"crypto/ecdsa"
	"encoding/json"

	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/keys"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1"
	"golang.org/x/crypto/sha3"
)

// Enclave defines the interface for key management operations
type Enclave interface {
	Address() string
	IsValid() bool
	PubKey() keys.PubKey
	Refresh() (Enclave, error)
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) (bool, error)
}

// KeyEnclave implements the Enclave interface
type KeyEnclave struct {
	Addr      string       `json:"address"`
	PubPoint  curves.Point `json:"-"`
	PubBytes  []byte       `json:"pub_key"`
	ValShare  Message      `json:"val_share"`
	UserShare Message      `json:"user_share"`
	VaultCID  string       `json:"vault_cid,omitempty"`
}

// Marshal returns the JSON encoding of KeyEnclave
func (k *KeyEnclave) Marshal() ([]byte, error) {
	// Store compressed public point bytes before marshaling
	k.PubBytes = k.PubPoint.ToAffineCompressed()
	return json.Marshal(k)
}

// Unmarshal parses the JSON-encoded data and stores the result
func (k *KeyEnclave) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, k); err != nil {
		return err
	}
	// Reconstruct Point from bytes
	curve := curves.K256()
	point, err := curve.NewIdentityPoint().FromAffineCompressed(k.PubBytes)
	if err != nil {
		return err
	}
	k.PubPoint = point
	return nil
}

func (k *KeyEnclave) IsValid() bool {
	return k.PubPoint != nil && k.ValShare != nil && k.UserShare != nil && k.Addr != ""
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
	return &KeyEnclave{
		Addr:      addr,
		PubPoint:  pubPoint,
		ValShare:  valShare,
		UserShare: userShare,
	}, nil
}

func (k *KeyEnclave) Address() string {
	return k.Addr
}

func (k *KeyEnclave) PubKey() keys.PubKey {
	return keys.NewPubKey(k.PubPoint)
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
	edSig, err := deserializeSignature(sig)
	if err != nil {
		return false, err
	}
	ePub, err := getEcdsaPoint(k.PubPoint.ToAffineUncompressed())
	if err != nil {
		return false, err
	}
	pk := &ecdsa.PublicKey{
		Curve: ePub.Curve,
		X:     ePub.X,
		Y:     ePub.Y,
	}
	
	// Hash the message using SHA3-256
	hash := sha3.New256()
	hash.Write(data)
	digest := hash.Sum(nil)
	
	return ecdsa.Verify(pk, digest, edSig.R, edSig.S), nil
}

func (k *KeyEnclave) userSignFunc(bz []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobSign(curve, sha3.New256(), bz, k.UserShare, protocol.Version1)
}

func (k *KeyEnclave) userRefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobRefresh(curve, k.UserShare, protocol.Version1)
}

func (k *KeyEnclave) valSignFunc(bz []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceSign(curve, sha3.New256(), bz, k.ValShare, protocol.Version1)
}

func (k *KeyEnclave) valRefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceRefresh(curve, k.ValShare, protocol.Version1)
}
