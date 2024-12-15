package mpc

import (
	"crypto/ecdsa"
	"encoding/json"

	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/keys"
	"golang.org/x/crypto/sha3"
)

// Enclave defines the interface for key management operations
type Enclave interface {
	Address() string
	DID() keys.DID
	IsValid() bool
	Marshal() ([]byte, error)
	PubKey() keys.PubKey
	Refresh() (Enclave, error)
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) (bool, error)
}

// keyEnclave implements the Enclave interface
type keyEnclave struct {
	Addr      string       `json:"address"`
	PubPoint  curves.Point `json:"-"`
	PubBytes  []byte       `json:"pub_key"`
	ValShare  Message      `json:"val_share"`
	UserShare Message      `json:"user_share"`
	VaultCID  string       `json:"vault_cid,omitempty"`
}

// Address returns the Sonr address of the keyEnclave
func (k *keyEnclave) Address() string {
	return k.Addr
}

// DID returns the DID of the keyEnclave
func (k *keyEnclave) DID() keys.DID {
	return keys.NewFromPubKey(k.PubKey())
}

// IsValid returns true if the keyEnclave is valid
func (k *keyEnclave) IsValid() bool {
	return k.PubPoint != nil && k.ValShare != nil && k.UserShare != nil && k.Addr != ""
}

// PubKey returns the public key of the keyEnclave
func (k *keyEnclave) PubKey() keys.PubKey {
	return keys.NewPubKey(k.PubPoint)
}

// Refresh returns a new keyEnclave
func (k *keyEnclave) Refresh() (Enclave, error) {
	refreshFuncVal, err := valRefreshFunc(k)
	if err != nil {
		return nil, err
	}
	refreshFuncUser, err := userRefreshFunc(k)
	if err != nil {
		return nil, err
	}
	return ExecuteRefresh(refreshFuncVal, refreshFuncUser)
}

// Sign returns the signature of the data
func (k *keyEnclave) Sign(data []byte) ([]byte, error) {
	userSign, err := userSignFunc(k, data)
	if err != nil {
		return nil, err
	}
	valSign, err := valSignFunc(k, data)
	if err != nil {
		return nil, err
	}
	return ExecuteSigning(valSign, userSign)
}

// Verify returns true if the signature is valid
func (k *keyEnclave) Verify(data []byte, sig []byte) (bool, error) {
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

// Marshal returns the JSON encoding of keyEnclave
func (k *keyEnclave) Marshal() ([]byte, error) {
	// Store compressed public point bytes before marshaling
	k.PubBytes = k.PubPoint.ToAffineCompressed()
	return json.Marshal(k)
}

// Unmarshal parses the JSON-encoded data and stores the result
func (k *keyEnclave) Unmarshal(data []byte) error {
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
