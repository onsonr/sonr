package crypto

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	. "github.com/cosmos/cosmos-sdk/crypto/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/mr-tron/base58/base58"
	"github.com/sonr-io/core/device"

	"github.com/tendermint/tendermint/libs/bytes"
)

const (
	DEVICE_PRIVKEY_FILE = "sonr-device.privkey"
	MASTER_PRIVKEY_FILE = "sonr-master.privkey"
)

type publicKey struct {
	p2p crypto.PubKey
	cryptotypes.PubKey
}

type privateKey struct {
	crypto.PrivKey
	p2p *secp256k1.PrivKey
}

func (pk *privateKey) Marshal() ([]byte, error) {
	return crypto.MarshalPrivateKey(pk.PrivKey)
}

func (pk *privateKey) Sign(msg []byte) ([]byte, []byte, error) {
	sig, err := pk.p2p.Sign(msg)
	if err != nil {
		return nil, nil, err
	}
	return msg, sig, nil
}

// Verify that the given public key is valid.
func (pk *publicKey) Verify(msg []byte, sig []byte) bool {
	ok, err := pk.p2p.Verify(msg, sig)
	if err != nil {
		return false
	}
	return ok
}

// KeySet is the set of keys required to operate on the Sonr network.
type KeySet interface {
	// Address returns the address of the keyset.
	Address() bytes.HexBytes

	// CopyToKeyring copies the keyset to the given keyring
	CopyToKeyring(k keyring.Keyring, sname string) (keyring.Info, error)

	// CryptoPrivKey returns the Secp256k1PrivKey as a libp2p crypto private key.
	CryptoPrivKey() crypto.PrivKey

	// CryptoPubKey returns the device key as a libp2p crypto public key.
	CryptoPubKey() crypto.PubKey

	// DID returns the DID of the keyset.
	DID() string

	// LegacyAminoPubKey returns the keyset as a legacy AminoPubKey.
	LegacyAminoPubKey() *multisig.LegacyAminoPubKey

	// PublicKeyBase58 returns the public key as a base58 encoded string.
	PublicKeyBase58() (string, error)

	// PeerID returns the peer ID of the keyset.
	PeerID() (peer.ID, error)

	// Secp256k1PrivKey returns the Secp256k1PrivKey as a libp2p crypto private key.
	Secp256k1PrivKey() *secp256k1.PrivKey

	// Export writes the keyset configuration to the given folder.
	Export(folder device.Folder) error
}

// keySet implements KeySet
type keySet struct {
	KeySet
	aminoPublicKey        *multisig.LegacyAminoPubKey
	privateKey            *privateKey
	publicKey             *publicKey
	motorCryptoPrivKey    crypto.PrivKey
	motorSecp256k1PrivKey *secp256k1.PrivKey
}

// CreateKeySet creates a new key set from a given mnemonic seed
func CreateKeySet(seed string) (KeySet, error) {
	priv := secp256k1.GenPrivKeyFromSecret([]byte(seed))
	pub := priv.PubKey()
	privp2p, err := crypto.UnmarshalSecp256k1PrivateKey(priv.Bytes())
	if err != nil {
		return nil, err
	}
	privKey := &privateKey{
		PrivKey: privp2p,
		p2p:     priv,
	}

	pubKey := &publicKey{
		PubKey: priv.PubKey(),
		p2p:    privp2p.GetPublic(),
	}

	pks := []PubKey{pub, pub}

	return &keySet{
		privateKey:            privKey,
		publicKey:             pubKey,
		motorCryptoPrivKey:    privp2p,
		motorSecp256k1PrivKey: priv,
		aminoPublicKey:        multisig.NewLegacyAminoPubKey(1, pks),
	}, nil
}

func LoadKeyset(folder device.Folder) (KeySet, error) {
	buf, err := folder.ReadFile(DEVICE_PRIVKEY_FILE)
	if err != nil {
		return nil, err
	}

	privp2p, err := crypto.UnmarshalSecp256k1PrivateKey(buf)
	if err != nil {
		return nil, err
	}

	bufm, err := folder.ReadFile(MASTER_PRIVKEY_FILE)
	if err != nil {
		return nil, err
	}

	priv := &secp256k1.PrivKey{}
	err = priv.UnmarshalAmino(bufm)
	if err != nil {
		return nil, err
	}

	pub := priv.PubKey()

	privKey := &privateKey{
		PrivKey: privp2p,
		p2p:     priv,
	}

	pubKey := &publicKey{
		PubKey: priv.PubKey(),
		p2p:    privp2p.GetPublic(),
	}

	pks := []PubKey{pub, pub}
	return &keySet{
		privateKey:            privKey,
		publicKey:             pubKey,
		motorCryptoPrivKey:    privp2p,
		motorSecp256k1PrivKey: priv,
		aminoPublicKey:        multisig.NewLegacyAminoPubKey(1, pks),
	}, nil
}

// Address returns the address of the key.
func (k *keySet) Address() bytes.HexBytes {
	return k.publicKey.Address()
}

// CopyToKeyring copies the key to the given keyring
func (k *keySet) CopyToKeyring(kring keyring.Keyring, sname string) (keyring.Info, error) {
	return kring.SaveMultisig(sname, k.LegacyAminoPubKey())
}

// CryptoPrivKey returns the crypto private key.
func (k *keySet) CryptoPrivKey() crypto.PrivKey {
	return k.motorCryptoPrivKey
}

// CryptoPubKey returns the crypto public key.
func (k *keySet) CryptoPubKey() crypto.PubKey {
	return k.motorCryptoPrivKey.GetPublic()
}

// DID returns the DID of the Sonr keyset
func (k *keySet) DID() string {
	return fmt.Sprintf("did:sonr:%s", k.Address())
}

// LegacyAminoPubKey returns the legacy amino public key for multisig
func (k *keySet) LegacyAminoPubKey() *multisig.LegacyAminoPubKey {
	return k.aminoPublicKey
}

// PublicKeyBase58 returns the public key in base58 format
func (k *keySet) PublicKeyBase58() (string, error) {
	buf, err := crypto.MarshalPublicKey(k.publicKey.p2p)
	if err != nil {
		panic(err)
	}
	str := base58.Encode(buf)
	return str, nil
}

// PeerID returns the peer ID of the public key
func (k *keySet) PeerID() (peer.ID, error) {
	return peer.IDFromPublicKey(k.publicKey.p2p)
}

// Secp256k1PrivKey returns the secp256k1 private key
func (k *keySet) Secp256k1PrivKey() *secp256k1.PrivKey {
	return k.motorSecp256k1PrivKey
}

// Export exports the key to the given folder and name
func (k *keySet) Export(folder device.Folder) error {
	buf, err := crypto.MarshalPrivateKey(k.motorCryptoPrivKey)
	if err != nil {
		return err
	}
	err = folder.WriteFile(fmt.Sprintf(DEVICE_PRIVKEY_FILE), buf)
	if err != nil {
		return err
	}

	buf, err = k.motorSecp256k1PrivKey.MarshalAmino()
	if err != nil {
		return err
	}

	err = folder.WriteFile(fmt.Sprintf(MASTER_PRIVKEY_FILE), buf)
	if err != nil {
		return err
	}
	return nil
}
