package motor

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	// at "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
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
	txRaw, err := m.Wallet.SignTx(msg1)
	if err != nil {
		return nil, err
	}

	resp, err := m.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())
	return m, nil
}
