package base

import (
	"encoding/json"
	"fmt"

	"github.com/sonrhq/sonr/internal/crypto"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	algo "github.com/sonrhq/sonr/services/mpc/protocol/dkls"
	v1types "github.com/sonrhq/sonr/services/mpc/types"
)

type EncryptionKey = *secp256k1.PubKey

type AccountV1 struct {
	// Address is the address of the account.
	Address string `json:"address"`

	// CoinType is the coin type of the account.
	CoinType crypto.CoinType `json:"coin_type"`

	// Name is the name of the account.
	Name string `json:"name"`

	// Data is the marshalled keyshare
	Data []byte `json:"data"`

	// PublicKeyshare is the public key share of the account.
	PublicKeyshare *v1types.Keyshare `json:"public_keshare"`
}

// NewAccountV1 creates a new account with the given name and coin type.
func NewAccountV1(name string, coin crypto.CoinType) (*AccountV1, v1types.KeyshareSet, error) {
	kss, err := algo.DKLSKeygen()
	if err != nil {
		return nil, v1types.EmptyKeyshareSet(), fmt.Errorf("error generating keyshare set: %v", err)
	}
	addr := kss.FormatAddress(coin)
	pubKs := kss.Alice()
	acc := &AccountV1{
		Address:        addr,
		CoinType:       coin,
		Name:           name,
		PublicKeyshare: pubKs,
	}
	return acc, kss, nil
}

// DID returns the DID of the account.
func (a *AccountV1) DID() string {
	str, _ := a.PublicKeyshare.FormatDID(a.CoinType)
	return str
}

// DIDAlias returns the DID alias or name of the account.
func (a *AccountV1) DIDAlias() string {
	return a.Name
}

// The `PublicKey()` function is a method of the `KeyshareSet` type. It returns the public key corresponding to Alice's keyshare in the keyshare set. It does this by calling the `PubKey()` method of the `Keyshare` object corresponding to Alice's keyshare. If the keyshare set is not
// valid or if there is an error in retrieving the public key, it returns an error.
func (a *AccountV1) PublicKey() *secp256k1.PubKey {
	pub, err := a.PublicKeyshare.PubKey()
	if err != nil {
		panic(err)
	}
	return pub
}

// GetAccountData returns the proto representation of the account
func (a *AccountV1) GetAccountData() *crypto.AccountData {
	dat, err := crypto.NewDefaultAccountData(a.CoinType, crypto.NewSecp256k1PubKey(a.PublicKey()))
	if err != nil {
		panic(err)
	}
	return dat
}

// Sign returns the signature of the message using the private keyshare.
func (a *AccountV1) Sign(pks *v1types.Keyshare, msg []byte) ([]byte, error) {
	kss := v1types.NewKSS(a.PublicKeyshare, pks)
	sig, err := kss.Sign(msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// Verify returns true if the signature is valid for the message and false otherwise.
func (a *AccountV1) Verify(msg, sig []byte) (bool, error) {
	return a.PublicKeyshare.Verify(msg, sig)
}

// Marshal returns the JSON encoding of the account.
func (a *AccountV1) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

// Unmarshal parses the JSON-encoded data and stores the result in the account.
func (a *AccountV1) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}
	return nil
}
