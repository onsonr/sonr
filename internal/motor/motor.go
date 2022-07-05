package motor

import (
	"encoding/json"
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
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

type MotorNode struct {
	Cosmos  *client.Client
	Wallet  *crypto.MPCWallet
	Address string
	PubKey  *secp256k1.PubKey
	DIDDoc  did.Document
	// Account *at.BaseAccount
}

func New() (*MotorNode, error) {
	// Create Client instance
	c := client.NewClient(client.ConnEndpointType_BETA)

	// Generate wallet
	w, err := crypto.GenerateWallet()
	if err != nil {
		return nil, err
	}
	bechAddr, err := w.Address()
	if err != nil {
		return nil, err
	}
	err = c.RequestFaucet(bechAddr)
	if err != nil {
		return nil, err
	}

	pk, err := w.PublicKeyProto()
	if err != nil {
		return nil, err
	}

	return &MotorNode{
		Cosmos:  c,
		Wallet:  w,
		Address: bechAddr,
		PubKey:  pk,
		DIDDoc:  w.DIDDocument,
	}, nil
}

func (m *MotorNode) CreateAccount(requestBytes []byte) (rtmv1.CreateAccountResponse, error) {
	var request rtmv1.CreateAccountRequest
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// create Vault shards to make sure this works before creating WhoIs
	vc := vault.New()
	deviceShard, sharedShard, recShard, unusedShards, err := m.Wallet.CreateInitialShards()
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// encrypt dscShard with dsc (i.e. webauthn)
	dscShard, err := dscEncrypt(deviceShard, request.AesDscKey)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// ecnrypt pskShard with psk (must be generated)
	pskShard, psk, err := pskEncrypt(sharedShard)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// password protect the recovery shard
	pwShard, err := crypto.AesEncryptWithPassword(request.Password, []byte(recShard))
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// create WhoIs
	resp, err := createWhoIs(m)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}
	fmt.Println(resp.String())

	// create vault
	vaultService, err := vc.CreateVault(
		m.Address,
		unusedShards,
		string(request.AesDscKey),
		dscShard,
		pskShard,
		pwShard,
	)
	if err != nil {
		fmt.Println("[WARN] failed to create vault:", err)
	}

	// update DID Document
	m.DIDDoc.AddService(vaultService)

	// update whois
	resp, err = updateWhoIs(m)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}
	fmt.Println(resp.String())

	return rtmv1.CreateAccountResponse{
		Address: m.Address,
		AesPsk:  psk,
	}, err
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

func pskEncrypt(shard string) (string, []byte, error) {
	key, err := crypto.NewAesKey()
	if err != nil {
		return "", nil, err
	}

	cipherShard, err := crypto.AesEncryptWithKey(key, []byte(shard))
	if err != nil {
		return "", key, err
	}

	return cipherShard, key, nil
}

func dscEncrypt(shard string, dsc []byte) (string, error) {
	if len(dsc) != 32 {
		return "", errors.New("dsc must be 32 bytes")
	}
	return crypto.AesEncryptWithKey(dsc, []byte(shard))
}
