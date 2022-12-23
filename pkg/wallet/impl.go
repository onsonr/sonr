package wallet

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

func NewWalletImpl(c interface{}) WalletShare {
	return &mpcConfigWalletImpl{Config: c.(*cmp.Config)}
}

type mpcConfigWalletImpl struct {
	*cmp.Config
}

// Returns the Bech32 representation of the given party.
func (w *mpcConfigWalletImpl) Address() string {
	pub, err := w.PublicKey()
	if err != nil {
		return ""
	}

	str, err := bech32.ConvertAndEncode("snr", pub.Bytes())
	if err != nil {
		return ""
	}
	return str
}

// MPCConfig returns the *cmp.Config of this wallet.
func (w *mpcConfigWalletImpl) MPCConfig() *cmp.Config {
	return w.Config
}

// Marshal serializes the cmp.Config into a byte slice for local storage
func (w *mpcConfigWalletImpl) Marshal() ([]byte, error) {
	return w.Config.MarshalBinary()
}

// PublicKey returns the public key of this wallet.
func (w *mpcConfigWalletImpl) PublicKey() (*secp256k1.PubKey, error) {
	buf, err := w.Config.PublicPoint().(*curve.Secp256k1Point).MarshalBinary()
	if err != nil {
		return nil, err
	}
	if len(buf) != 33 {
		return nil, fmt.Errorf("invalid public key length")
	}
	return &secp256k1.PubKey{
		Key: buf,
	}, nil
}

// SelfID returns the ID of this wallet.
func (w *mpcConfigWalletImpl) SelfID() party.ID {
	return w.Config.ID
}

// PartyIDs returns the IDs of all parties in the group.
func (w *mpcConfigWalletImpl) PartyIDs() []party.ID {
	return w.Config.PartyIDs()
}

// Unmarshal deserializes the given byte slice into a cmp.Config
func (w *mpcConfigWalletImpl) Unmarshal(data []byte) error {
	return w.Config.UnmarshalBinary(data)
}

// Verify a signature with the given wallet.
func (w *mpcConfigWalletImpl) Verify(data []byte, sig []byte) bool {
	signature, err := DeserializeSignature(sig)
	if err != nil {
		return false
	}
	return signature.Verify(w.Config.PublicPoint(), data)
}

func searchFirstNotId(ids party.IDSlice, id party.ID) party.ID {
	for _, v := range ids {
		if v != id {
			return v
		}
	}
	return party.ID("")
}
