package didwallets

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sonrhq/sonr/internal/crypto"
	"github.com/sonrhq/sonr/internal/mpc"
	"github.com/sonrhq/sonr/pkg/didcommon"
	"github.com/sonrhq/sonr/pkg/walletsigner"
)

// Method is the DID method for Sonr Wallet Actor DIDs
const SonrMethod = didcommon.Method("sonr")

// SonrAccount struct is defining a custom data type in Go. It represents a Sonr Wallet Actor DID (Decentralized Identifier) account. It has several fields including `Method`, `ID`, `Resources`, `acc`, and `pks`. These fields store information related to the Sonr account, such as the DID method, identifier,
// associated resources, and cryptographic keys.
type SonrAccount struct {
	Method didcommon.Method
	ID     didcommon.Identifier

	acc *mpc.AccountV1
	kss mpc.KeyshareSet
}

// NewSonrAccount creates a new Sonr Wallet Actor DID
func NewSonrAccount(key didcommon.SecretKey) (*SonrAccount, error) {
	ct := crypto.SONRCoinType
	m := didcommon.Method(ct.DIDMethod())
	acc, pks, err := mpc.GenerateV2("primary", ct)
	if err != nil {
		return nil, err
	}
	id := didcommon.Identifier(acc.Address)
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
	return &SonrAccount{
		Method: m,
		ID:     id,
		acc:    acc,
		kss:    pks,
	}, nil
}

// ResolveSonrAccount resolves a Sonr Wallet Actor DID
func ResolveSonrAccount(didString string, key didcommon.SecretKey) (*SonrAccount, error) {
	ct := crypto.SONRCoinType
	m := didcommon.Method(ct.DIDMethod())

	id := didcommon.Identifier(didString)

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
	return &SonrAccount{
		Method: m,
		ID:     id,
		acc:    acc,
		kss:    kss,
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
	return a.kss.Sign(msg)
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
	rawBz, err := walletsigner.SignCosmosTx(a, msgs...)
	if err != nil {
		return nil, err
	}
	resp, err := walletsigner.BroadcastCosmosTx(rawBz)
	if err != nil {
		return nil, err
	}
	return resp.TxResponse, nil
}
