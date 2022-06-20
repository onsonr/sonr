package motor

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	txt "github.com/cosmos/cosmos-sdk/types/tx"

	// at "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/vault"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
)

type MotorNode struct {
	Cosmos  *client.Client
	Wallet  *crypto.MPCWallet
	Address string
	PubKey  *secp256k1.PubKey
	DIDDoc  did.Document
	// Account *at.BaseAccount
}

func CreateAccount(password, dscPubKey, psKey []byte) (*MotorNode, error) {
	m, err := setupDefault()
	if err != nil {
		return nil, err
	}

	// create Vault shards to make sure this works before creating WhoIs
	vc := vault.New()
	dscShard, pskShard, recShard, unusedShards, err := m.Wallet.CreateInitialShards()
	if err != nil {
		return nil, err
	}

	// create WhoIs
	resp, err := createWhoIs(m)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())

	// create vault
	vaultService, err := vc.CreateVault(
		m.Address,
		unusedShards,
		string(dscPubKey),
		dscShard,
		pskShard,
		recShard,
	)
	if err != nil {
		return nil, err
	}

	// update DID Document
	m.DIDDoc.AddService(vaultService)

	// update whois
	resp, err = updateWhoIs(m)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())

	return m, err
}

func createWhoIs(m *MotorNode) (*sdk.TxResponse, error) {
	docBz, err := m.DIDDoc.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgCreateWhoIs(m.Address, m.PubKey, docBz, rt.WhoIsType_USER)
	ai, txb, err := m.newCreateTx(msg1)
	if err != nil {
		return nil, err
	}

	sig, err := m.signTx(ai, txb)
	if err != nil {
		return nil, err
	}

	txRes, err := m.Cosmos.BroadcastTx(txb, sig, ai)
	if err != nil {
		return nil, err
	}

	if txRes.RawLog != "[]" {
		return nil, errors.New(txRes.RawLog)
	}
	return txRes, nil
}

func updateWhoIs(m *MotorNode) (*sdk.TxResponse, error) {
	docBz, err := m.DIDDoc.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgUpdateWhoIs(m.Address, docBz)
	ai, txb, err := m.newUpdateTx(msg1)
	if err != nil {
		return nil, err
	}

	sig, err := m.signTx(ai, txb)
	if err != nil {
		return nil, err
	}

	txRes, err := m.Cosmos.BroadcastTx(txb, sig, ai)
	if err != nil {
		return nil, err
	}

	if txRes.RawLog != "[]" {
		return nil, errors.New(txRes.RawLog)
	}
	return txRes, nil
}

func (m *MotorNode) newCreateTx(msg ...sdk.Msg) (*tx.AuthInfo, *txt.TxBody, error) {
	return crypto.BuildCreateWhoIsTx(m.Wallet, msg...)
}

func (m *MotorNode) newUpdateTx(msg ...sdk.Msg) (*tx.AuthInfo, *txt.TxBody, error) {
	return crypto.BuildUpdateWhoIsTx(m.Wallet, msg...)
}

func (m *MotorNode) signTx(ai *tx.AuthInfo, txb *txt.TxBody) (*ecdsa.Signature, error) {
	signDocBz, err := crypto.GetSignDocBytes(ai, txb)
	if err != nil {
		return nil, err
	}
	return m.Wallet.Sign(signDocBz)
}

// func (m *MotorNode) broadcastTx(txb *txt.TxBody, sig *ecdsa.Signature, ai *tx.AuthInfo) (*sdk.TxResponse, error) {
// 	return m.Cosmos.BroadcastTx(txb, sig, ai)
// }
