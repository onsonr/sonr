package motor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/pkg/vault"
	rt "github.com/sonr-io/sonr/x/registry/types"
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

type MotorNode struct {
	Cosmos      *client.Client
	Wallet      *crypto.MPCWallet
	Address     string
	PubKey      *secp256k1.PubKey
	DID         did.DID
	DIDDocument did.Document

	// Sharding
	deviceShard   string
	sharedShard   string
	recoveryShard string
	unusedShards  []string
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

	baseDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(bechAddr, "snr")))
	if err != nil {
		return nil, err
	}

	// Create the DID Document
	doc, err := did.NewDocument(baseDid.String())
	if err != nil {
		return nil, err
	}

	deviceShard, sharedShard, recShard, unusedShards, err := w.CreateInitialShards()
	if err != nil {
		//return rtmv1.CreateAccountResponse{}, err
		return nil, err
	}

	return &MotorNode{
		Cosmos:        c,
		Wallet:        w,
		Address:       bechAddr,
		PubKey:        pk,
		DID:           *baseDid,
		DIDDocument:   doc,
		deviceShard:   deviceShard,
		sharedShard:   sharedShard,
		recoveryShard: recShard,
		unusedShards:  unusedShards,
	}, nil
}

func (m *MotorNode) CreateAccount(requestBytes []byte) (rtmv1.CreateAccountResponse, error) {
	var request rtmv1.CreateAccountRequest
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// create Vault shards to make sure this works before creating WhoIs
	vc := vault.New()

	// ecnrypt pskShard with psk (must be generated)
	pskShard, psk, err := pskEncrypt(m.sharedShard)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// password protect the recovery shard
	pwShard, err := crypto.AesEncryptWithPassword(request.Password, []byte(m.recoveryShard))
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
		m.unusedShards,
		string(request.AesDscKey),
		string(request.GetAesDscKey()),
		pskShard,
		pwShard,
	)
	if err != nil {
		fmt.Println("[WARN] failed to create vault:", err)
	}

	// update DID Document
	m.DIDDocument.AddService(vaultService)

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
	docBz, err := m.DIDDocument.MarshalJSON()
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
	docBz, err := m.DIDDocument.MarshalJSON()
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

// GetVerificationMethod returns the VerificationMethod for the given party.
func (w *MotorNode) GetVerificationMethod(id party.ID) (*did.VerificationMethod, error) {
	vmdid, err := did.ParseDID(fmt.Sprintf("did:snr:%s#%s", strings.TrimPrefix(w.Address, "snr"), id))
	if err != nil {
		return nil, err
	}

	// Get base58 encoded public key.
	pub, err := w.Wallet.PublicKeyBase58()
	if err != nil {
		return nil, err
	}

	// Return the shares VerificationMethod
	return &did.VerificationMethod{
		ID:              *vmdid,
		Type:            ssi.ECDSASECP256K1VerificationKey2019,
		Controller:      w.DID,
		PublicKeyBase58: pub,
	}, nil
}
