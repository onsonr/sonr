package wallet

import (
	"context"
	"encoding"
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	crypto_pb "github.com/libp2p/go-libp2p-core/crypto/pb"
	"github.com/libp2p/go-libp2p-core/peer"
)

// keyPair is a joint private/public key pair
type keyPair struct {
	pub    crypto.PubKey
	priv   crypto.PrivKey
	kpType KeyPairType
}

// PrivPubKeys returns the private and public keys for the keypair given keychain
func (kp keyPair) PrivPubKeys() (crypto.PubKey, crypto.PrivKey, error) {
	if kp.priv == nil {
		logger.Errorf("%s - Failed to Return Private Key", ErrNoPrivateKey)
		return nil, nil, ErrNoPrivateKey
	}

	if kp.pub == nil {
		logger.Errorf("%s - Failed to Return Public Key", ErrNoPublicKey)
		return nil, nil, ErrNoPublicKey
	}
	return kp.pub, kp.priv, nil
}

// newSnrKeyPair creates a new key pair
func newSnrKeyPair(p crypto.PrivKey) (crypto.PrivKey, crypto.PubKey) {
	// Create a new priv key
	priv := NewSnrPrivKey(p)
	return priv, priv.GetPublic()
}

type SnrKey interface {
	Buffer() ([]byte, error)
	PeerID() (peer.ID, error)
	String() (string, error)
	Type() crypto_pb.KeyType
}

type TextileIdentity interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	// Sign the given bytes cryptographically.
	Sign(context.Context, []byte) ([]byte, error)
	// GetPublic returns the public key paired with this identity.
	GetPublic() TextilePubKey
	// Decrypt returns decrypted data.
	Decrypt(context.Context, []byte) ([]byte, error)
	// Equals returns true if the identities are equal.
	Equals(TextileIdentity) bool
}

// Pubkey can be anything that provides a verify method.
type TextilePubKey interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	// String encodes the public key into a base32 string.
	fmt.Stringer
	// UnmarshalString decodes the public key from a base32 string.
	UnmarshalString(string) error
	// Verify that 'sig' is the signed hash of 'data'
	Verify(data []byte, sig []byte) (bool, error)
	// Encrypt data with the public key.
	Encrypt([]byte) ([]byte, error)
	// Equals returns true if the keys are equal.
	Equals(TextilePubKey) bool
}

// SnrPrivKey is Sonr wrapper around crypto.PrivKey
type SnrPrivKey struct {
	crypto.PrivKey
	SnrKey
	SName string
}

// NewSnrPrivKey creates a new private key
func NewSnrPrivKey(p crypto.PrivKey) *SnrPrivKey {
	return &SnrPrivKey{
		PrivKey: p,
	}
}

// Buffer returns PrivateKey as bytes
func (priv *SnrPrivKey) Buffer() ([]byte, error) {
	buf, err := crypto.MarshalPrivateKey(priv.PrivKey)
	if err != nil {
		logger.Errorf("%s - Failed to marshal SPubKey", err)
		return nil, err
	}
	return buf, nil
}

// GetPublic returns the public key
func (priv *SnrPrivKey) GetPublic() crypto.PubKey {
	return &SnrPubKey{
		PubKey: priv.PrivKey.GetPublic(),
	}
}

// PeerID returns the peer ID from the public key
func (priv *SnrPrivKey) PeerID() (peer.ID, error) {
	return peer.IDFromPublicKey(priv.GetPublic())
}

// String returns PublicKey as Base64 string
func (priv *SnrPrivKey) String() (string, error) {
	buf, err := priv.Buffer()
	if err != nil {
		return "", err
	}
	return crypto.ConfigEncodeKey(buf), nil
}

// Type of the private key (Ed25519).
func (priv *SnrPrivKey) Type() crypto_pb.KeyType {
	return crypto_pb.KeyType_Ed25519
}

// SnrPubKey is Sonr wrapper around crypto.PubKey
type SnrPubKey struct {
	crypto.PubKey
	SnrKey
}

// NewSnrPubKey creates a new public key
func NewSnrPubKey(p crypto.PubKey) *SnrPubKey {
	return &SnrPubKey{
		PubKey: p,
	}
}

// NewSnrPubKeyFromBuffer creates a new public key from buffer
func NewSnrPubKeyFromBuffer(buf []byte) (*SnrPubKey, error) {
	// Decode the key
	pubKey, err := crypto.UnmarshalPublicKey(buf)
	if err != nil {
		logger.Errorf("%s - Failed to unmarshal PubKey from Bytes", err)
		return nil, err
	}
	return NewSnrPubKey(pubKey), nil
}

// NewSnrPubKeyFromString creates a new public key from string
func NewSnrPubKeyFromString(str string) (*SnrPubKey, error) {
	// Decode the key
	buf, err := crypto.ConfigDecodeKey(str)
	if err != nil {
		logger.Errorf("%s - Failed to decode PubKey from String", err)
		return nil, err
	}
	return NewSnrPubKeyFromBuffer(buf)
}

// Buffer returns PublicKey as bytes
func (pub *SnrPubKey) Buffer() ([]byte, error) {
	buf, err := crypto.MarshalPublicKey(pub)
	if err != nil {
		logger.Errorf("%s - Failed to marshal SPubKey", err)
		return nil, err
	}
	return buf, nil
}

// PeerID returns the peer ID from the public key
func (pub *SnrPubKey) PeerID() (peer.ID, error) {
	return peer.IDFromPublicKey(pub)
}

// String returns PublicKey as Base64 string
func (pub *SnrPubKey) String() (string, error) {
	buf, err := pub.Buffer()
	if err != nil {
		return "", err
	}
	return crypto.ConfigEncodeKey(buf), nil
}

// Type of the PubKey (Ed25519).
func (pub *SnrPubKey) Type() crypto_pb.KeyType {
	return crypto_pb.KeyType_Ed25519
}
