package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// KeyShare is a type that interacts with a cmp.Config file located on disk.
type KeyShare interface {
	// Bytes returns the bytes of the keyshare file - the marshalled cmp.Config
	Bytes() []byte

	// CoinType returns the coin type based on the keyshare file name
	CoinType() crypto.CoinType

	// Config returns the cmp.Config.
	Config() *cmp.Config

	// DeriveBip44 returns a new keyshare with the same key but a new coin type
	DeriveBip44(ct crypto.CoinType, idx int, name string) (KeyShare, error)

	// Did returns the cid of the keyshare
	Did() string

	// PartyID returns the party id based on the keyshare file name
	PartyID() crypto.PartyID

	// PubKey returns the public key of the keyshare
	PubKey() *crypto.PubKey

	// IsEncrypted checks if the file at current path is encrypted.
	IsEncrypted() bool
}

// keyShare is a type that interacts with a cmp.Config file located on disk.
type keyShare struct {
	bytes    []byte
	name     string
	lastUsed uint32
}

// Keyshare name format is a DID did:{coin_type}:{account_address}#ks-{account_name}-{keyshare_name}
// did:{coin_type}:{account_address}#ks-{account_name}-{keyshare_name}
func NewKeyshare(id string, bytes []byte, coinType crypto.CoinType) (KeyShare, error) {
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
	ks.name = fmt.Sprintf("did:%s:%s#ks-%s", coinType.DidMethod(), addr, string(conf.ID))
	return ks, nil
}

// GetPubKeyFromCmpConfigBytes loads KeyShare from a cmp.Config buffer.
func GetPubKeyFromCmpConfigBytes(bytes []byte) (*crypto.PubKey, error) {
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

// Bytes returns the bytes of the keyshare file - the marshalled cmp.Config
func (ks *keyShare) Bytes() []byte {
	return ks.bytes
}

// CoinType returns the coin type based on the keyshare file name
func (ks *keyShare) CoinType() crypto.CoinType {
	res, err := ParseKeyShareDID(ks.name)
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

// DeriveBip44 returns a derived keyshare from the current keyshare.
func (ks *keyShare) DeriveBip44(ct crypto.CoinType, idx int, name string) (KeyShare, error) {
	cnfg, err := ks.Config().DeriveBIP32(uint32(ct.BipPath()))
	if err != nil {
		return nil, err
	}

	bz, err := cnfg.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return NewKeyshare(ks.name, bz, ct)
}

// Did returns the cid of the keyshare
func (ks *keyShare) Did() string {
	return ks.name
}

// PartyID returns the party id based on the keyshare file name
func (ks *keyShare) PartyID() crypto.PartyID {
	res, err := ParseKeyShareDID(ks.name)
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

// // Encrypt checks if the file at current path is encrypted and if not, encrypts it.
// func (ks *keyShare) Encrypt(credential servicetypes.Credential) error {
// 	if ks.IsEncrypted() {
// 		return nil
// 	}
// 	enc, err := credential.Encrypt(ks.bytes)
// 	if err != nil {
// 		return err
// 	}
// 	ks.lastUsed = uint32(time.Now().Unix())
// 	ks.bytes = enc
// 	return nil
// }

// // Decrypt checks if the file at current path is encrypted and if not, encrypts it.
// func (ks *keyShare) Decrypt(credential servicetypes.Credential) error {
// 	if !ks.IsEncrypted() {
// 		return nil
// 	}

// 	dec, err := credential.Decrypt(ks.bytes)
// 	if err != nil {
// 		return err
// 	}
// 	ks.lastUsed = uint32(time.Now().Unix())
// 	ks.bytes = dec
// 	return nil
// }

// ! ||--------------------------------------------------------------------------------||
// ! ||                           Helper Methods for KeyShare                          ||
// ! ||--------------------------------------------------------------------------------||

// A Keyshare is encrypted if its name contains an apostrophe at the end.
func (ks *keyShare) IsEncrypted() bool {
	if ks.IsVault() {
		return false
	}
	return strings.HasSuffix(ks.name, "'")
}

func (ks *keyShare) IsVault() bool {
	return strings.Contains(ks.name, "vault")
}
