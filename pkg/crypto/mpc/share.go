package mpc

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/sonr-hq/sonr/pkg/common"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// It returns an empty wallet share
func EmptyWalletShare() common.WalletShare {
	return &mpcConfigWalletImpl{}
}

// It takes a `cmp.Config` and returns a `common.WalletShare` that can be used to create a wallet
func NewWalletShare(pfix string, c interface{}) common.WalletShare {
	conf := c.(*cmp.Config)
	confBz, err := conf.MarshalBinary()
	if err != nil {
		return nil
	}
	partyIds := make([]string, len(conf.PartyIDs()))
	for i, id := range conf.PartyIDs() {
		partyIds[i] = string(id)
	}
	walletConf := &common.WalletShareConfig{
		Algorithm:    common.WalletShareConfig_CMP,
		SelfId:       string(conf.ID),
		PartyIds:     partyIds,
		CmpConfig:    confBz,
		Timestamp:    time.Now().Unix(),
		Bech32Prefix: pfix,
	}
	return &mpcConfigWalletImpl{Config: conf, walletShareConfig: walletConf}
}

// `mpcConfigWalletImpl` is a type that implements the `MPCConfig` interface.
//
// The `MPCConfig` interface is defined in the `mpc` package.
//
// The `Config` type is defined in the `cmp` package.
//
// The `WalletShareConfig` type is defined in the `common` package.
//
// The `mpcConfigWalletImpl` type is defined in the `wallet` package.
//
// The `mpcConfigWalletImpl` type has a field of type `Config` and a field of type `WalletShare
// @property  - `Config` - the configuration of the MPC protocol.
// @property walletShareConfig - This is the configuration for the wallet share.
type mpcConfigWalletImpl struct {
	*cmp.Config
	walletShareConfig *common.WalletShareConfig
}

// Returns the Bech32 representation of the given party.
func (w *mpcConfigWalletImpl) Address() string {
	pub, err := w.PublicKey()
	if err != nil {
		return ""
	}

	str, err := bech32.ConvertAndEncode(w.walletShareConfig.Bech32Prefix, pub.Bytes())
	if err != nil {
		return ""
	}
	return str
}

// MPCConfig returns the *cmp.Config of this wallet.
func (w *mpcConfigWalletImpl) CMPConfig() *cmp.Config {
	return w.Config
}

// DID returns the DID of this wallet.
func (w *mpcConfigWalletImpl) DID() (string, error) {
	prefix := w.walletShareConfig.GetBech32Prefix()
	addrPtr := strings.Split(w.Address(), prefix)
	if len(addrPtr) != 2 {
		return "", fmt.Errorf("invalid address")
	}
	return fmt.Sprintf("did:%s:%s", prefix, addrPtr[1]), nil
}

// Marshal serializes the cmp.Config into a byte slice for local storage
func (w *mpcConfigWalletImpl) Marshal() ([]byte, error) {
	return w.walletShareConfig.Marshal()
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
	walletConf := &common.WalletShareConfig{}
	if err := walletConf.Unmarshal(data); err != nil {
		return err
	}
	w.walletShareConfig = walletConf
	conf := &cmp.Config{}
	if err := conf.UnmarshalBinary(walletConf.CmpConfig); err != nil {
		return err
	}
	w.Config = conf
	return nil
}

func (w *mpcConfigWalletImpl) Share() *common.WalletShareConfig {
	return w.walletShareConfig
}

// Verify a signature with the given wallet.
func (w *mpcConfigWalletImpl) Verify(data []byte, sig []byte) bool {
	signature, err := DeserializeSignature(sig)
	if err != nil {
		return false
	}
	return signature.Verify(w.Config.PublicPoint(), data)
}

// It returns the first element of the slice that is not equal to the given ID
func searchFirstNotId(ids party.IDSlice, id party.ID) party.ID {
	for _, v := range ids {
		if v != id {
			return v
		}
	}
	return party.ID("")
}
