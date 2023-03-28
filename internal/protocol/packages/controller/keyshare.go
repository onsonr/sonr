package controller

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// KeyShare is a type that interacts with a cmp.Config file located on disk.
type KeyShare interface {
	// AccountName returns the account name based on the keyshare file name
	AccountName() string

	// Bytes returns the bytes of the keyshare file - the marshalled cmp.Config
	Bytes() []byte

	// CoinType returns the coin type based on the keyshare file name
	CoinType() crypto.CoinType

	// Config returns the cmp.Config.
	Config() *cmp.Config

	// Did returns the cid of the keyshare
	Did() string

	// KeyShareName returns the keyshare name based on the keyshare file name
	KeyShareName() string

	// PartyID returns the party id based on the keyshare file name
	PartyID() crypto.PartyID

	// PubKey returns the public key of the keyshare
	PubKey() *crypto.PubKey

	// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
	Encrypt(credential *crypto.WebauthnCredential) error

	// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
	Decrypt(credential *crypto.WebauthnCredential) error

	// IsEncrypted checks if the file at current path is encrypted.
	IsEncrypted() bool
}

// keyShare is a type that interacts with a cmp.Config file located on disk.
type keyShare struct {
	bytes    []byte
	name     string
	lastUsed uint32
}

type Foobar struct {
	Foo string
	Bar string
}

// Keyshare name format is a DID did:{coin_type}:{account_address}#ks-{account_name}-{keyshare_name}
// did:{coin_type}:{account_address}#ks-{account_name}-{keyshare_name}
func NewKeyshare(id string, bytes []byte, coinType crypto.CoinType, accName string) (KeyShare, error) {
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	err := conf.UnmarshalBinary(bytes)
	if err != nil {
		return nil, err
	}

	ks := &keyShare{
		bytes:    bytes,
		lastUsed: uint32(time.Now().Unix()),
	}
	addr := coinType.FormatAddress(ks.PubKey())
	ks.name = fmt.Sprintf("did:%s:%s#ks-%s-%s", coinType.DidMethod(), addr, accName, string(conf.ID))
	return ks, nil
}

// LoadKeyshareFromStore loads a keyshare from a store. The value can be a base64 encoded string or a []byte.
func LoadKeyshareFromStore(key string, value interface{}) (KeyShare, error) {
	var v []byte
	switch value := value.(type) {
	case []byte:
		v = value
	case string:
		bz, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return nil, err
		}
		v = bz
	default:
		return nil, fmt.Errorf("invalid value type")
	}
	ksr, err := ParseKeyShareDid(key)
	if err != nil {
		return nil, err
	}
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	err = conf.UnmarshalBinary(v)
	if err != nil {
		return nil, err
	}

	return &keyShare{
		bytes:    v,
		name:     ksr.KeyShareName,
		lastUsed: uint32(time.Now().Unix()),
	}, nil
}

// LoadKeySharePubKeyFromConfigBytes loads KeyShare from a cmp.Config buffer.
func LoadKeySharePubKeyFromConfigBytes(bytes []byte) (*crypto.PubKey, error) {
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	err := conf.UnmarshalBinary(bytes)
	if err != nil {
		return nil, err
	}
	skPP, ok := conf.PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil, fmt.Errorf("invalid public point type")
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return crypto.NewSecp256k1PubKey(bz), nil
}

// AccountName returns the account name based on the keyshare file name
func (ks *keyShare) AccountName() string {
	res, err := ParseKeyShareDid(ks.name)
	if err != nil {
		return ""
	}
	return res.AccountName
}

// Bytes returns the bytes of the keyshare file - the marshalled cmp.Config
func (ks *keyShare) Bytes() []byte {
	return ks.bytes
}

// CoinType returns the coin type based on the keyshare file name
func (ks *keyShare) CoinType() crypto.CoinType {
	res, err := ParseKeyShareDid(ks.name)
	if err != nil {
		return crypto.SONRCoinType
	}
	return res.CoinType
}

// Config returns the cmp.Config.
func (ks *keyShare) Config() *cmp.Config {
	cnfg := cmp.EmptyConfig(curve.Secp256k1{})
	err := cnfg.UnmarshalBinary(ks.bytes)
	if err != nil {
		panic(err)
	}
	ks.lastUsed = uint32(time.Now().Unix())
	return cnfg
}

// Did returns the cid of the keyshare
func (ks *keyShare) Did() string {
	return ks.name
}

// Keyshare name format is /{purpose}/{coin_type}/{account_name}/{keyshare_name}
func (ks *keyShare) KeyShareName() string {
	res, err := ParseKeyShareDid(ks.name)
	if err != nil {
		return ""
	}
	return res.KeyShareName
}

// PartyID returns the party id based on the keyshare file name
func (ks *keyShare) PartyID() crypto.PartyID {
	res, err := ParseKeyShareDid(ks.name)
	if err != nil {
		panic(err)
	}
	return crypto.PartyID(res.KeyShareName)
}

// PublicKey returns the public key of the keyshare
func (ks *keyShare) PubKey() *crypto.PubKey {
	skPP, ok := ks.Config().PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil
	}
	return crypto.NewSecp256k1PubKey(bz)
}

// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
func (ks *keyShare) Encrypt(credential *crypto.WebauthnCredential) error {
	if ks.name == "vault" {
		return nil
	}
	enc, err := credential.Encrypt(ks.bytes)
	if err != nil {
		return err
	}
	ks.lastUsed = uint32(time.Now().Unix())
	ks.bytes = enc
	ks.name += "'" // encrypted keyshares have an apostrophe at the end
	return nil
}

// Decrypt checks if the file at current path is encrypted and if not, encrypts it.
func (ks *keyShare) Decrypt(credential *crypto.WebauthnCredential) error {
	if !ks.IsEncrypted() {
		return nil
	}

	dec, err := credential.Decrypt(ks.bytes)
	if err != nil {
		return err
	}
	ks.lastUsed = uint32(time.Now().Unix())
	ks.bytes = dec
	ks.name = strings.TrimSuffix(ks.name, "'") // remove the apostrophe
	return nil
}


// A Keyshare is encrypted if its name contains an apostrophe at the end.
func (ks *keyShare) IsEncrypted() bool {
		return false
}
