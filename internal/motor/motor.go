package motor

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/pkg/vault"
	rt "github.com/sonr-io/sonr/x/registry/types"
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
	txRaw, err := tx.SignTxWithWallet(m.Wallet, "/sonrio.sonr.registry.MsgCreateWhoIs", msg1)
	if err != nil {
		return nil, err
	}

	resp, err := m.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	if resp.TxResponse.RawLog != "[]" {
		return nil, errors.New(resp.TxResponse.RawLog)
	}
	return resp.TxResponse, nil
}

func updateWhoIs(m *MotorNode) (*sdk.TxResponse, error) {
	docBz, err := m.DIDDoc.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgUpdateWhoIs(m.Address, docBz)
	txRaw, err := tx.SignTxWithWallet(m.Wallet, "/sonrio.sonr.registry.MsgUpdateWhoIs", msg1)
	if err != nil {
		return nil, err
	}

	resp, err := m.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	if resp.TxResponse.RawLog != "[]" {
		return nil, errors.New(resp.TxResponse.RawLog)
	}
	return resp.TxResponse, nil
}
