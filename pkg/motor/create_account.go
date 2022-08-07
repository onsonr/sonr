package motor

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/pkg/vault"
	rt "github.com/sonr-io/sonr/x/registry/types"
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func (mtr *motorNodeImpl) CreateAccount(request rtmv1.CreateAccountRequest) (rtmv1.CreateAccountResponse, error) {
	// create motor
	fmt.Printf("initializing motor... ")
	if err := initMotor(mtr); err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("initialize motor: %s", err)
	}
	fmt.Println("done.")

	// Request from Faucet
	fmt.Printf("requesting initial balance... ")
	err := mtr.Cosmos.RequestFaucet(mtr.Address)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("request from faucet: %s", err)
	}
	fmt.Println("done.")

	// Create Initial Shards
	fmt.Printf("creating shards... ")
	deviceShard, sharedShard, recShard, unusedShards, err := mtr.Wallet.CreateInitialShards()
	if err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("create shards: %s", err)
	}
	mtr.deviceShard = deviceShard
	mtr.sharedShard = sharedShard
	mtr.recoveryShard = recShard
	mtr.unusedShards = unusedShards
	fmt.Println("done.")

	// Create the DID Document
	doc, err := did.NewDocument(mtr.DID.String())
	if err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("create DID document: %s", err)
	}
	mtr.DIDDocument = doc

	// create Vault shards to make sure this works before creating WhoIs
	fmt.Printf("creating account... ")
	vc := vault.New()
	if _, err := createWhoIs(mtr); err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("create account: %s", err)
	}
	fmt.Println("done.")

	// ecnrypt dscShard with DSC
	fmt.Printf("encrypting shards... ")
	dscShard, err := dscEncrypt(mtr.deviceShard, request.AesDscKey)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("encrypt backup shards: %s", err)
	}

	// encrypt pskShard with psk (must be generated)
	pskShard, psk, err := pskEncrypt(mtr.sharedShard)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("encrypt psk shards: %s", err)
	}

	// password protect the recovery shard
	pwShard, err := mpc.AesEncryptWithPassword(request.Password, mtr.recoveryShard)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("encrypt password shard: %s", err)
	}
	fmt.Println("done.")

	// create vault
	fmt.Printf("setting up vault... ")
	vaultService, err := vc.CreateVault(
		mtr.Address,
		mtr.unusedShards,
		mtr.DeviceID,
		dscShard,
		pskShard,
		pwShard,
	)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("setup vault: %s", err)
	}
	fmt.Println("done.")

	// update DID Document
	fmt.Printf("updating WhoIs... ")
	mtr.DIDDocument.AddService(vaultService)

	// update whois
	if _, err = updateWhoIs(mtr); err != nil {
		return rtmv1.CreateAccountResponse{}, fmt.Errorf("update WhoIs: %s", err)
	}
	fmt.Println("done.")

	fmt.Println("account created successfully.")
	return rtmv1.CreateAccountResponse{
		AesPsk:  psk,
		Address: mtr.Address,
	}, err
}

func createWhoIs(m *motorNodeImpl) (*sdk.TxResponse, error) {
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

	return resp.TxResponse, nil
}

func updateWhoIs(m *motorNodeImpl) (*sdk.TxResponse, error) {
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

	return resp.TxResponse, nil
}

func pskEncrypt(shard []byte) ([]byte, []byte, error) {
	key, err := mpc.NewAesKey()
	if err != nil {
		return nil, nil, err
	}

	cipherShard, err := mpc.AesEncryptWithKey(key, shard)
	if err != nil {
		return nil, key, err
	}

	return cipherShard, key, nil
}

func dscEncrypt(shard, dsc []byte) ([]byte, error) {
	if len(dsc) != 32 {
		return nil, errors.New("dsc must be 32 bytes")
	}
	return mpc.AesEncryptWithKey(dsc, shard)
}
