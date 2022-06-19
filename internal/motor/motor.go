package motor

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	txt "github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
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

func CreateAccount() (*MotorNode, error) {
	m, err := setupDefault()
	if err != nil {
		return nil, err
	}

	docBz, err := m.DIDDoc.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgCreateWhoIs(m.Address, m.PubKey, docBz, rt.WhoIsType_USER)
	ai, txb, err := m.newTx(msg1)
	if err != nil {
		return nil, err
	}

	sig, err := m.signTx(ai, txb)
	if err != nil {
		return nil, err
	}

	resp, err := m.Cosmos.BroadcastTx(txb, sig, ai)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())
	return m, nil
}

func (m *MotorNode) newTx(msg ...sdk.Msg) (*tx.AuthInfo, *txt.TxBody, error) {
	return crypto.BuildTx(m.Wallet, msg...)
}

func (m *MotorNode) signTx(ai *tx.AuthInfo, txb *txt.TxBody) (*ecdsa.Signature, error) {
	signDocBz, err := crypto.GetSignDocBytes(ai, txb)
	if err != nil {
		return nil, err
	}
	return m.Wallet.Sign(signDocBz)
}
