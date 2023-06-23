package types

import (
	"fmt"
	"strings"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// KeyShare is a type that interacts with a cmp.Config file located on disk.
type KeyShare interface {
	// Bytes returns the bytes of the keyshare file - the marshalled cmp.Config
	Bytes() []byte

	// Config returns the cmp.Config.
	CMPConfig() *cmp.Config

	// DeriveBip44 returns a new keyshare with the same key but a new coin type
	DeriveBip44(ct uint32, idx int, name string) (KeyShare, error)

	// Did returns the Identifier with fragment of the keyshare
	Did() string

	// IsEncrypted checks if the file at current path is encrypted.
	IsEncrypted() bool

	// PartyID returns the party id based on the keyshare file name
	PartyID() crypto.PartyID

	// CoinType returns the coin type based on the keyshare file name
	ParseCoinType() crypto.CoinType

	// PubKey returns the public key of the keyshare
	PubKey() *crypto.PubKey

	// Rename renames the keyshare file
	Rename(opts ...KeyShareOption) KeyShare

	// ToProto returns a protobuf representation of the keyshare
	ToProto() *VaultKeyshare
}

// Keyshare name format is a DID did:{coin_type}:{account_address}#ks-{account_name}-{keyshare_name}
// did:{coin_type}:{account_address}#ks-{account_name}-{keyshare_name}
func NewKeyshare(bytes []byte, coinType crypto.CoinType, opts ...KeyShareOption) (KeyShare, error) {
	// setup default options
	options := &keyshareOpts{
		fragment: "",
		modified: false,
	}
	for _, opt := range opts {
		opt(options)
	}
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	err := conf.UnmarshalBinary(bytes)
	if err != nil {
		return nil, err
	}

	ks := &VaultKeyshare{
		Config:   bytes,
		CoinType: coinType.BipPath(),
	}
	addr := coinType.FormatAddress(ks.PubKey())
	if options.fragment != "" {
		ks.Id = fmt.Sprintf("did:%s:%s#%s", coinType.DidMethod(), addr, options.fragment)
	} else {
		ks.Id = fmt.Sprintf("did:%s:%s#ks-%s", coinType.DidMethod(), addr, string(conf.ID))
	}
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
func (ks *VaultKeyshare) Bytes() []byte {
	return ks.Config
}

// CoinType returns the coin type based on the keyshare file name
func (ks *VaultKeyshare) ParseCoinType() crypto.CoinType {
	return crypto.CoinTypeFromBipPath(ks.CoinType)
}

// Config returns the cmp.Config.
func (ks *VaultKeyshare) CMPConfig() *cmp.Config {
	cnfg := cmp.EmptyConfig(curve.Secp256k1{})
	err := cnfg.UnmarshalBinary(ks.Config)
	if err != nil {
		panic(err)
	}
	return cnfg
}

// DeriveBip44 returns a derived keyshare from the current keyshare.
func (ks *VaultKeyshare) DeriveBip44(ct uint32, idx int, name string) (KeyShare, error) {
	cnfg, err := ks.CMPConfig().DeriveBIP32(ct)
	if err != nil {
		return nil, err
	}

	bz, err := cnfg.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return NewKeyshare(bz, crypto.CoinTypeFromBipPath(ct))
}

// DID returns the cid of the keyshare
func (ks *VaultKeyshare) Did() string {
	return ks.Id
}

// A Keyshare is encrypted if its name contains an apostrophe at the end.
func (ks *VaultKeyshare) IsEncrypted() bool {
	if ks.IsVault() {
		return false
	}
	return strings.HasSuffix(ks.Id, "'")
}

func (ks *VaultKeyshare) IsVault() bool {
	return strings.Contains(ks.Id, "vault")
}

// PartyID returns the party id based on the keyshare file name
func (ks *VaultKeyshare) PartyID() crypto.PartyID {
	res, err := ParseKeyShareDID(ks.Id)
	if err != nil {
		panic(err)
	}
	return crypto.PartyID(res.KeyShareName)
}

// PublicKey returns the public key of the keyshare
func (ks *VaultKeyshare) PubKey() *crypto.PubKey {
	skPP, ok := ks.CMPConfig().PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil
	}
	return crypto.NewSecp256k1PubKey(bz)
}

// Rename renames the keyshare did with fragment options
func (ks *VaultKeyshare) Rename(opts ...KeyShareOption) KeyShare {
	// setup default options
	options := &keyshareOpts{
		fragment: "",
		modified: false,
	}
	for _, opt := range opts {
		opt(options)
	}
	if !options.modified {
		return ks
	}
	ks.Id = fmt.Sprintf("did:%s:%s#%s", ks.ParseCoinType().DidMethod(), ks.PubKey().Address(), options.fragment)
	return ks
}

// ToProto returns a protobuf representation of the keyshare
func (ks *VaultKeyshare) ToProto() *VaultKeyshare {
	return ks
}



// ! ||--------------------------------------------------------------------------------||
// ! ||                           Helper Methods for KeyShare                          ||
// ! ||--------------------------------------------------------------------------------||

type keyshareOpts struct {
	fragment string
	modified bool
}

type KeyShareOption func(*keyshareOpts)

func SetUnclaimed(idx int) KeyShareOption {
	return func(o *keyshareOpts) {
		o.modified = true
		o.fragment = fmt.Sprintf("ucw-%d", idx)
	}
}

func SetClaimed(fragment string) KeyShareOption {
	return func(o *keyshareOpts) {
		o.modified = true
		o.fragment = fragment
	}
}
