package motor

import (
	"fmt"

	kr "github.com/sonr-io/sonr/internal/keyring"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/pkg/vault"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (mtr *motorNodeImpl) CreateAccount(request mt.CreateAccountRequest) (mt.CreateAccountResponse, error) {
	// create DSC and store it in keychain
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_KEY_CREATE_START})
	dsc, err := kr.CreateDSC()
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_KEY_CREATE_END,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountResponse{}, fmt.Errorf("create DSC: %s", err)
	}

	// create PSK and store it in keychain
	psk, err := kr.CreatePSK()
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_KEY_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountResponse{}, fmt.Errorf("create PSK: %s", err)
	}

	res, err := mtr.CreateAccountWithKeys(mt.CreateAccountWithKeysRequest{
		Password:  request.Password,
		AesDscKey: dsc,
		AesPskKey: psk,
		Metadata:  request.Metadata,
	})
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_KEY_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountResponse{}, err
	}

	return mt.CreateAccountResponse{
		Address: res.Address,
	}, nil
}

// CreateAccountWithKeys allows PSK and DSC to be provided manually
func (mtr *motorNodeImpl) CreateAccountWithKeys(request mt.CreateAccountWithKeysRequest) (mt.CreateAccountWithKeysResponse, error) {
	// Create Client instance
	mtr.Cosmos = client.NewClient(mtr.clientMode)

	// set encryption key, based on preshared key
	mtr.encryptionKey = request.AesPskKey

	// create motor
	if err := initMotor(mtr); err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("initialize motor: %s", err)
	}
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_FAUCET_REQUEST_START})

	// Request from Faucet
	err := mtr.Cosmos.RequestFaucet(mtr.Address)
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_FAUCET_REQUEST_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("request from faucet: %s", err)
	}

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_FAUCET_REQUEST_END})
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_START})

	// Create the DID Document
	doc, err := did.NewDocument(mtr.DID.String())
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("create DID document: %s", err)
	}
	mtr.DIDDocument = doc

	// Format DID for setting MPC as controller
	controller, err := did.ParseDID(fmt.Sprintf("%s#jwk", doc.GetID().String()))
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("parse controller DID: %s", err)
	}

	// Add MPC as a VerificationMethod for the assertion of the DID Document
	pk, err := mtr.Wallet.CreateEcdsaFromPublicKey()
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, err
	}
	vm, err := did.NewVerificationMethod(doc.GetID(), ssi.JsonWebKey2020, *controller, pk)
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, err
	}
	doc.AddAssertionMethod(vm)

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR})
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_SHARD_GENERATE_START})

	go createVault(mtr, request)

	// perform sharding and vault creation async
	return mt.CreateAccountWithKeysResponse{
		Address: mtr.Address,
	}, err
}

func createVault(mtr *motorNodeImpl, request mt.CreateAccountWithKeysRequest) {
	// Create Initial Shards
	deviceShard, sharedShard, recShard, unusedShards, err := mtr.Wallet.CreateInitialShards()

	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_SHARD_GENERATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return
	}
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_SHARD_GENERATE_END})

	mtr.deviceShard = deviceShard
	mtr.sharedShard = sharedShard
	mtr.recoveryShard = recShard
	mtr.unusedShards = unusedShards

	// create Vault shards to make sure this works before creating WhoIs
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_WHO_IS_CREATE_START})
	vc := vault.New()
	if _, err := createWhoIs(mtr); err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_WHO_IS_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return
	}

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_WHO_IS_CREATE_END})
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_KEY_ENCRYPT_START})

	// encrypt dscShard with DSC
	dscShard, err := mpc.AesEncryptWithKey(request.AesDscKey, mtr.deviceShard)
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_KEY_ENCRYPT_ERROR,
			ErrorMessage: err.Error(),
		})
		return
	}

	// encrypt pskShard with psk (must be generated)
	pskShard, err := mpc.AesEncryptWithKey(request.AesPskKey, mtr.sharedShard)
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_KEY_ENCRYPT_ERROR,
			ErrorMessage: err.Error(),
		})
		return
	}

	// password protect the recovery shard
	pwShard, err := mpc.AesEncryptWithPassword(request.Password, mtr.recoveryShard)
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_KEY_ENCRYPT_ERROR,
			ErrorMessage: err.Error(),
		})
		return
	}

	// encrypt each of the unused shards
	for i, s := range mtr.unusedShards {
		mtr.unusedShards[i], err = mpc.AesEncryptWithKey(mtr.encryptionKey, s)
		if err != nil {
			mtr.triggerWalletEvent(common.WalletEvent{
				Type:         common.WALLET_EVENT_TYPE_KEY_ENCRYPT_ERROR,
				ErrorMessage: err.Error(),
			})

			return
		}
	}

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_KEY_ENCRYPT_END})
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_VAULT_CREATE_START})

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
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_VAULT_CREATE_END,
			ErrorMessage: err.Error(),
		})
		return
	}

	// update DID Document
	mtr.DIDDocument.AddService(vaultService)

	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR,
			ErrorMessage: err.Error(),
			Message:      "",
		})
		return
	}

	// update whois
	_, err = updateWhoIs(mtr)
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_VAULT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return
	}

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_VAULT_CREATE_END})
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

func updateWhoIs(m MotorNode) (*rt.MsgUpdateWhoIsResponse, error) {
	docBz, err := m.GetDIDDocument().MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgUpdateWhoIs(m.GetAddress(), docBz)
	txRaw, err := tx.SignTxWithWallet(m.GetWallet(), "/sonrio.sonr.registry.MsgUpdateWhoIs", msg1)
	if err != nil {
		return nil, err
	}

	resp, err := m.GetClient().BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	cwir := &rt.MsgUpdateWhoIsResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cwir); err != nil {
		return nil, err
	}

	return cwir, nil
}
