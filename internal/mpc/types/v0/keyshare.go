package types

import (
	"encoding/json"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type Keyshare struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Config   []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	CoinType uint32 `protobuf:"varint,3,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
}

// KeyshareFromJSON returns a keyshare from a json representation of the keyshare
func KeyshareFromJSON(bz []byte) (*Keyshare, error) {
	ks := &Keyshare{}
	err := json.Unmarshal(bz, ks)
	if err != nil {
		return nil, err
	}
	return ks, nil
}

// The function CreateKeyshare creates a VaultKeyshare using the provided configuration and coin type.
func CreateKeyshare(config *crypto.MPCCmpConfig, coinType crypto.CoinType) *Keyshare {
	bytes, _ := config.MarshalBinary()
	ks := &Keyshare{
		Config:   bytes,
		CoinType: coinType.BipPath(),
	}
	ks.Address = coinType.FormatAddress(ks.PubKey())
	return ks
}

// DeriveBip44 returns a derived keyshare from the current keyshare.
func (ks *Keyshare) DeriveBip44(ct uint32) (*Keyshare, error) {
	cnfg, err := ks.ParseConfig().DeriveBIP32(ct)
	if err != nil {
		return nil, err
	}
	return CreateKeyshare(cnfg, crypto.CoinTypeFromBipPath(ct)), nil
}

// ParseCoinType returns the coin type of the keyshare
func (ks *Keyshare) ParseCoinType() crypto.CoinType {
	return crypto.CoinTypeFromBipPath(ks.CoinType)
}

// Config returns the cmp.Config.
func (ks *Keyshare) ParseConfig() *cmp.Config {
	cnfg := cmp.EmptyConfig(curve.Secp256k1{})
	err := cnfg.UnmarshalBinary(ks.Config)
	if err != nil {
		panic(err)
	}
	return cnfg
}

// PartyID returns the party id based on the keyshare file name
func (ks *Keyshare) PartyID() crypto.PartyID {
	return ks.ParseConfig().ID
}

// PublicKey returns the public key of the keyshare
func (ks *Keyshare) PubKey() *crypto.PubKey {
	skPP, ok := ks.ParseConfig().PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil
	}
	return crypto.NewPubKey(bz, crypto.Secp256k1KeyType)
}

// ToJSON returns the json representation of the keyshare
func (ks *Keyshare) ToJSON() ([]byte, error) {
	return json.Marshal(ks)
}

// Verify returns true if the signature is valid for the keyshare
func (ks *Keyshare) Verify(msg, sig []byte) bool {
	sigObj, err := DeserializeECDSASecp256k1Signature(sig)
	if err != nil {
		return false
	}
	return sigObj.Verify(ks.ParseConfig().PublicPoint(), msg)
}
