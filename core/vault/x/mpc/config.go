package mpc

import (
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/sonr-hq/sonr/internal/node"
	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/pool"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
)

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

// GroupIDs returns the IDs of all parties in the group.
func (w *mpcConfigWalletImpl) GroupIDs() []party.ID {
	return w.Config.PartyIDs()
}

// Unmarshal deserializes the given byte slice into a cmp.Config
func (w *mpcConfigWalletImpl) Unmarshal(data []byte) error {
	return w.Config.UnmarshalBinary(data)
}

// Sign begins a round of the MPC protocol to sign the given message.
func (w *mpcConfigWalletImpl) Sign(data []byte, th node.TopicHandler) ([]byte, error) {
	var wg sync.WaitGroup
	pl := pool.NewPool(0)
	defer pl.TearDown()
	signature, err := CmpSign(w.Config, data, w.GroupIDs(), th, &wg, pl)
	if err != nil {
		return nil, err
	}
	return SerializeSignature(signature)
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
	return party.ID(0)
}
