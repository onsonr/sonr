package sonr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did/types"
	"github.com/sonr-io/sonr/pkg/mpc"
)

// Method is the DID method for Sonr Wallet Actor DIDs
const Method = types.DIDMethod("sonr")

// Account struct is defining a custom data type in Go. It represents a Sonr Wallet Actor DID (Decentralized Identifier) account. It has several fields including `Method`, `ID`, `Resources`, `acc`, and `pks`. These fields store information related to the Sonr account, such as the DID method, identifier,
// associated resources, and cryptographic keys.
type Account struct {
	Method types.DIDMethod
	ID     types.DIDIdentifier

	acc *mpc.AccountV1
	kss mpc.KeyshareSet
}

// NewSonrAccount creates a new Sonr Wallet Actor DID
func NewSonrAccount(key types.DIDSecretKey) (*Account, error) {
	ct := crypto.SONRCoinType
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
	_, err = id.AddResource("public", pbz)
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
	return &Account{
		Method: m,
		ID:     id,
		acc:    acc,
		kss:    pks,
	}, nil
}

// ResolveAccount resolves a Sonr Wallet Actor DID
func ResolveAccount(didString string, key types.DIDSecretKey) (*Account, error) {
	ct := crypto.SONRCoinType
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
	return &Account{
		Method: m,
		ID:     id,
		acc:    acc,
		kss:    kss,
	}, nil
}

// Address returns the address of the account
func (a *Account) Address() string {
	return a.acc.Address
}

// Info returns the account data
func (a *Account) Info() *crypto.AccountData {
	return a.acc.GetAccountData()
}

// Sign signs a message with the account
func (a *Account) Sign(msg []byte) ([]byte, error) {
	return a.kss.Sign(msg)
}

// PublicKey returns the public key of the account
func (a *Account) PublicKey() (*crypto.Secp256k1PubKey, error) {
	return a.acc.PublicKey(), nil
}

// Type returns the type of the account
func (a *Account) Type() string {
	return "secp256k1"
}

// Verify verifies a signature
func (a *Account) Verify(msg []byte, sig []byte) (bool, error) {
	return a.acc.Verify(msg, sig)
}

// SendTx sends a transaction
func (a *Account) SendTx(msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	rawBz, err := SignCosmosTx(a, msgs...)
	if err != nil {
		return nil, err
	}
	resp, err := BroadcastCosmosTx(rawBz)
	if err != nil {
		return nil, err
	}
	return resp.TxResponse, nil
}
