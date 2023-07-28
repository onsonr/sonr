package sonr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/mpc"
	"github.com/sonrhq/core/pkg/did/types"
)

const Method = types.DIDMethod("sonr")

// The `SonrAccount` struct is defining a custom data type in Go. It represents a Sonr Wallet Actor DID (Decentralized Identifier) account. It has several fields including `Method`, `ID`, `Resources`, `acc`, and `pks`. These fields store information related to the Sonr account, such as the DID method, identifier,
// associated resources, and cryptographic keys.
type SonrAccount struct {
	Method    types.DIDMethod
	ID        types.DIDIdentifier

	acc *mpc.AccountV1
	pks *mpc.KeyshareV1
}

// NewSonrAccount creates a new Sonr Wallet Actor DID
func NewSonrAccount(key types.DIDSecretKey) (*SonrAccount, error) {
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
	privBz, err := pks.MarshalPrivate()
	if err != nil {
		return nil, err
	}
	encBz, err := key.Encrypt(privBz)
	if err != nil {
		return nil, err
	}
	_, err = id.AddResource("private", encBz)
	if err != nil {
		return nil, err
	}
	return &SonrAccount{
		Method:    m,
		ID:        id,
		acc:       acc,
		pks:       pks,
	}, nil
}

// ResolveAccount resolves a Sonr Wallet Actor DID
func ResolveAccount(didString string, key types.DIDSecretKey) (*SonrAccount, error) {
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
	decBz, err := key.Decrypt(privResource)
	if err != nil {
		return nil, err
	}
	pks := &mpc.KeyshareV1{}
	if err := pks.UnmarshalPrivate(decBz); err != nil {
		return nil, err
	}
	return &SonrAccount{
		Method:    m,
		ID:        id,
		acc:       acc,
		pks:       pks,
	}, nil
}

// Address returns the address of the account
func (a *SonrAccount) Address() string {
	return a.acc.Address
}

// Info returns the account data
func (a *SonrAccount) Info() *crypto.AccountData {
	return a.acc.GetAccountData()
}

// Sign signs a message with the account
func (a *SonrAccount) Sign(msg []byte) ([]byte, error) {
	return a.acc.Sign(a.pks, msg)
}

// PublicKey returns the public key of the account
func (a *SonrAccount) PublicKey() (*crypto.Secp256k1PubKey, error) {
	return a.acc.PublicKey(), nil
}

// Type returns the type of the account
func (a *SonrAccount) Type() string {
	return "secp256k1"
}

// Verify verifies a signature
func (a *SonrAccount) Verify(msg []byte, sig []byte) (bool, error) {
	return a.acc.Verify(msg, sig)
}

// SendTx sends a transaction
func (a *SonrAccount) SendTx(msgs ...sdk.Msg) (*sdk.TxResponse, error) {
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
