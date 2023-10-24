package btcr

import (
	"github.com/sonr-io/core/pkg/crypto"
	"github.com/sonr-io/core/services/did/types"
	"github.com/sonr-io/core/services/mpc"
)

const Method = types.DIDMethod("btcr")

type BitcoinAccount struct {
	method    types.DIDMethod
	ID        types.DIDIdentifier
	Resources []types.DIDResource
	acc       *mpc.AccountV1
	kss       mpc.KeyshareSet
}

// NewBitcoinAccount creates a new Bitcoin Wallet Actor DID
func NewBitcoinAccount(key types.DIDSecretKey) (types.WalletAccount, error) {
	ct := crypto.BTCCoinType
	m := types.DIDMethod(ct.DIDMethod())
	acc, pks, err := mpc.GenerateV2("primary", ct)
	if err != nil {
		return nil, err
	}
	id := types.DIDIdentifier(acc.Address)
	pbz, err := acc.Marshal()
	if err != nil {
		return nil, err
	}
	pr, err := id.AddResource("public", pbz)
	if err != nil {
		return nil, err
	}
	m.SetKey(id.String(), string(pbz))
	privBz, err := pks.EncryptUserKeyshare(key)
	if err != nil {
		return nil, err
	}
	encBz, err := privBz.Marshal()
	if err != nil {
		return nil, err
	}
	_, err = id.AddResource("private", encBz)
	if err != nil {
		return nil, err
	}
	return &BitcoinAccount{
		method:    m,
		ID:        id,
		Resources: []types.DIDResource{pr},
		acc:       acc,
		kss:       pks,
	}, nil
}

// ResolveAccount resolves a Sonr Wallet Actor DID
func ResolveAccount(didString string, key types.DIDSecretKey) (types.WalletAccount, error) {
	ct := crypto.BTCCoinType
	m := types.DIDMethod(ct.DIDMethod())

	id := types.DIDIdentifier(didString)

	// Get public resource
	pubResource, err := id.FetchResource("public")
	if err != nil {
		return nil, err
	}
	acc := &mpc.AccountV1{}
	if err := acc.Unmarshal(pubResource); err != nil {
		return nil, err
	}
	// Get private resource and decrypt it
	privResource, err := id.FetchResource("private")
	if err != nil {
		return nil, err
	}
	epks := &mpc.EncKeyshareSet{}
	if err := epks.Unmarshal(privResource); err != nil {
		return nil, err
	}
	kss, err := epks.DecryptUserKeyshare(key)
	if err != nil {
		return nil, err
	}
	return &BitcoinAccount{
		method: m,
		ID:     id,
		acc:    acc,
		kss:    kss,
	}, nil
}

// Address returns the address of the account
func (a *BitcoinAccount) Address() string {
	return a.acc.Address
}

// Info returns the account data
func (a *BitcoinAccount) Info() *crypto.AccountData {
	return a.acc.GetAccountData()
}

// Method returns the DID method
func (a *BitcoinAccount) Method() types.DIDMethod {
	return a.method
}

// Sign signs a message with the account
func (a *BitcoinAccount) Sign(msg []byte) ([]byte, error) {
	return a.kss.Sign(msg)
}

// PublicKey returns the public key of the account
func (a *BitcoinAccount) PublicKey() (*crypto.Secp256k1PubKey, error) {
	return a.acc.PublicKey(), nil
}

// Type returns the type of the account
func (a *BitcoinAccount) Type() string {
	return "secp256k1"
}

// Verify verifies a signature
func (a *BitcoinAccount) Verify(msg []byte, sig []byte) (bool, error) {
	return a.acc.Verify(msg, sig)
}
