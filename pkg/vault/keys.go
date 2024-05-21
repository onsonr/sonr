package vault

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/di-dao/core/pkg/did"
	"github.com/di-dao/core/x/did/types"
)

// Key is a struct that contains the information of a key
type Key struct {
	// Human readable name for the key
	Name string `json:"name"`

	// Algorithm used for the key
	Algo string `json:"algo"`

	// Public key
	PubKey []byte `json:"pubKey"`

	// Address of the key
	Address bytes.HexBytes `json:"address"`

	// Coin type
	CoinType uint32 `json:"coinType"`

	// Did
	Did string `json:"did"`
}

// NewKey creates a new key from types.PublicKey
func NewKey(name string, pubKey *types.PublicKey, coinType uint32) *Key {
	return &Key{
		Name:     name,
		Algo:     pubKey.KeyType,
		PubKey:   pubKey.Key,
		Address:  pubKey.Address(),
		CoinType: coinType,
		Did:      pubKey.Did,
	}
}

// createBtcKey creates a bitcoin key from types.PublicKey
func createBtcKey(pk *types.PublicKey, wallet *WalletData) error {
	btcAddr, err := did.CreateBitcoinAddress(pk)
	if err != nil {
		return err
	}
	btcKey := &Key{
		Name:     "btc1",
		Algo:     "secp256k1",
		PubKey:   pk.Bytes(),
		Address:  btcAddr.Bytes(),
		CoinType: 0,
		Did:      fmt.Sprintf("did:btcr:%s", btcAddr.String()),
	}
	wallet.Keys = append(wallet.Keys, btcKey)
	return nil
}

// createSonrKey creates a sonr key from types.PublicKey
func createSonrKey(pk *types.PublicKey, wallet *WalletData) error {
	sonrAddr, err := did.CreateSonrAddress(pk)
	if err != nil {
		return err
	}
	snrKey := &Key{
		Name:     "sonr1",
		Algo:     "secp256k1",
		PubKey:   pk.Bytes(),
		Address:  sonrAddr.Bytes(),
		CoinType: 1,
		Did:      fmt.Sprintf("did:sonr:%s", sonrAddr.String()),
	}
	wallet.Keys = append(wallet.Keys, snrKey)
	return nil
}
