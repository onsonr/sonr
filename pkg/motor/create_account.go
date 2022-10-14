package motor

import (
	"errors"
	"fmt"

	kr "github.com/sonr-io/sonr/internal/keyring"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/pkg/vault"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (mtr *motorNodeImpl) CreateAccount(request mt.CreateAccountRequest) (mt.CreateAccountResponse, error) {
	// create DSC and store it in keychain
	dsc, err := kr.CreateDSC()
	if err != nil {
		return mt.CreateAccountResponse{}, fmt.Errorf("create DSC: %s", err)
	}

	// create PSK and store it in keychain
	psk, err := kr.CreatePSK()
	if err != nil {
		return mt.CreateAccountResponse{}, fmt.Errorf("create PSK: %s", err)
	}

	res, err := mtr.CreateAccountWithKeys(mt.CreateAccountWithKeysRequest{
		Password:  request.Password,
		AesDscKey: dsc,
		AesPskKey: psk,
		Metadata:  request.Metadata,
	})
	if err != nil {
		return mt.CreateAccountResponse{}, err
	}

	return mt.CreateAccountResponse{
		Address: res.Address,
		WhoIs:   res.WhoIs,
	}, nil
}

// CreateAccountWithKeys allows PSK and DSC to be provided manually
func (mtr *motorNodeImpl) CreateAccountWithKeys(request mt.CreateAccountWithKeysRequest) (mt.CreateAccountWithKeysResponse, error) {
	// Create Client instance
	mtr.Cosmos = client.NewClient(mtr.clientMode)

	// create motor
	if err := initMotor(mtr); err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("initialize motor: %s", err)
	}

	// Request from Faucet
	err := mtr.Cosmos.RequestFaucet(mtr.Address)
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("request from faucet: %s", err)
	}

	// Create the DID Document
	doc, err := did.NewDocument(mtr.DID.String())
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("create DID document: %s", err)
	}
	mtr.DIDDocument = doc

	// Format DID for setting MPC as controller
	controller, err := did.ParseDID(fmt.Sprintf("%s#mpc", doc.GetID().String()))
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("parse controller DID: %s", err)
	}

	// Add MPC as a VerificationMethod for the assertion of the DID Document
	vm, err := did.NewVerificationMethodFromBytes(doc.GetID(), ssi.ECDSASECP256K1VerificationKey2019, *controller, mtr.GetPubKey().Bytes())
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, err
	}
	doc.AddAssertionMethod(vm)

	// Create Initial Shards
	deviceShard, sharedShard, recShard, unusedShards, err := mtr.Wallet.CreateInitialShards()

	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("create shards: %s", err)
	}
	mtr.deviceShard = deviceShard
	mtr.sharedShard = sharedShard
	mtr.recoveryShard = recShard
	mtr.unusedShards = unusedShards

	// create Vault shards to make sure this works before creating WhoIs
	vc := vault.New()
	if _, err := createWhoIs(mtr); err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("create account: %s", err)
	}

	// encrypt dscShard with DSC
	dscShard, err := dscEncrypt(mtr.deviceShard, request.AesDscKey)
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("encrypt backup shards: %s", err)
	}

	// encrypt pskShard with psk (must be generated)
	pskShard, err := pskEncrypt(mtr.sharedShard, request.AesPskKey)
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("encrypt psk shards: %s", err)
	}

	// password protect the recovery shard
	pwShard, err := mpc.AesEncryptWithPassword(request.Password, mtr.recoveryShard)
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("encrypt password shard: %s", err)
	}

	// create vault
	vaultService, err := vc.CreateVault(
		mtr.Address,
		mtr.unusedShards,
		mtr.DeviceID,
		dscShard,
		pskShard,
		pwShard,
	)
	fmt.Println("Response From Create Vault :", vaultService)
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("setup vault: %s", err)
	}

	// update DID Document
	mtr.DIDDocument.AddService(vaultService)

	// update whois
	resp, err := updateWhoIs(mtr)
	if err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("update WhoIs: %s", err)
	}

	return mt.CreateAccountWithKeysResponse{
		Address: mtr.Address,
		WhoIs:   resp.GetWhoIs(),
	}, err
}

func createWhoIs(m *motorNodeImpl) (*rt.MsgCreateWhoIsResponse, error) {
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

	cwir := &rt.MsgCreateWhoIsResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cwir); err != nil {
		return nil, err
	}

	return cwir, nil
}

func updateWhoIs(m *motorNodeImpl) (*rt.MsgUpdateWhoIsResponse, error) {
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

	cwir := &rt.MsgUpdateWhoIsResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cwir); err != nil {
		return nil, err
	}

	return cwir, nil
}

func pskEncrypt(shard, key []byte) ([]byte, error) {
	cipherShard, err := mpc.AesEncryptWithKey(key, shard)
	if err != nil {
		return nil, err
	}

	return cipherShard, nil
}

// dscEncrypt encrypts the shard with the DSC key
// Returns: encrypted shard, given key, error
func dscEncrypt(shard, dsc []byte) ([]byte, error) {
	// Check if the DSC is valid
	if len(dsc) != 32 {
		return nil, errors.New("dsc must be 32 bytes")
	}

	return mpc.AesEncryptWithKey(dsc, shard)
}
