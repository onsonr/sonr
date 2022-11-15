package motor

import (
	"fmt"

	kr "github.com/sonr-io/sonr/internal/keyring"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/pkg/vault"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (mtr *motorNodeImpl) CreateAccount(request mt.CreateAccountRequest, waitForVault bool) (mt.CreateAccountResponse, error) {
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
	}, waitForVault)
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
func (mtr *motorNodeImpl) CreateAccountWithKeys(request mt.CreateAccountWithKeysRequest, waitForVault bool) (mt.CreateAccountWithKeysResponse, error) {
	// Create Client instance
	mtr.Cosmos = client.NewClient(mtr.clientMode)

	// set encryption key, based on preshared key
	mtr.encryptionKey = request.AesPskKey

	// create motor
	if err := initMotor(mtr); err != nil {
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("initialize motor: %s", err)
	}
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_FAUCET_REQUEST_START, Address: mtr.Address})

	// Request from Faucet
	err := mtr.Cosmos.RequestFaucet(mtr.Address)
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_FAUCET_REQUEST_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("request from faucet: %s", err)
	}

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_FAUCET_REQUEST_END, Address: mtr.Address})
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_START, Address: mtr.Address})

	// Create the DID Document
	mtr.DIDDocument, err = did.NewDocument(mtr.DID.String())
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("create DID document: %s", err)
	}

	// Format DID for setting MPC as controller
	controller, err := did.ParseDID(fmt.Sprintf("%s#mpc", mtr.DIDDocument.GetID().String()))
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, fmt.Errorf("parse controller DID: %s", err)
	}

	// Add MPC as a VerificationMethod for the assertion of the DID Document
	vm, err := did.NewVerificationMethodFromBytes(mtr.DIDDocument.GetID(), ssi.ECDSASECP256K1VerificationKey2019, *controller, mtr.GetPubKey().Bytes())
	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return mt.CreateAccountWithKeysResponse{}, err
	}
	mtr.DIDDocument.AddAssertionMethod(vm)

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_ERROR})
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_SHARD_GENERATE_START})

	r := make(chan int)
	go createVault(mtr, request, &r)
	if waitForVault {
		<-r
	}

	// perform sharding and vault creation async
	return mt.CreateAccountWithKeysResponse{
		Address: mtr.Address,
	}, err
}

func createVault(mtr *motorNodeImpl, request mt.CreateAccountWithKeysRequest, r *chan int) {
	// Create Initial Shards
	deviceShard, sharedShard, recShard, unusedShards, err := mtr.Wallet.CreateInitialShards()

	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_SHARD_GENERATE_ERROR,
			Address:      mtr.Address,
			ErrorMessage: err.Error(),
		})
		return
	}
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_SHARD_GENERATE_END, Address: mtr.Address})

	mtr.deviceShard = deviceShard
	mtr.sharedShard = sharedShard
	mtr.recoveryShard = recShard
	mtr.unusedShards = unusedShards

	// create Vault shards to make sure this works before creating WhoIs
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_WHO_IS_CREATE_START, Address: mtr.Address})
	vc := vault.New()
	if _, err := createWhoIs(mtr); err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_WHO_IS_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return
	}

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_WHO_IS_CREATE_END, Address: mtr.Address})
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_KEY_ENCRYPT_START, Address: mtr.Address})

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

	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_KEY_ENCRYPT_END, Address: mtr.Address})
	mtr.triggerWalletEvent(common.WalletEvent{Type: common.WALLET_EVENT_TYPE_VAULT_CREATE_START, Address: mtr.Address})

	// create vault
	vaultService, err := vc.CreateVault(
		mtr.Address,
		mtr.unusedShards,
		mtr.DeviceID,
		dscShard,
		pskShard,
		pwShard,
	)

	if err != nil {
		mtr.triggerWalletEvent(common.WalletEvent{
			Type:         common.WALLET_EVENT_TYPE_VAULT_CREATE_ERROR,
			ErrorMessage: err.Error(),
		})
		return
	}

	// update DID Document
	mtr.DIDDocument.AddService(vaultService)

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
	*r <- 1
}

func createWhoIs(m *motorNodeImpl) (*rt.MsgCreateWhoIsResponse, error) {
	docBz, err := m.DIDDocument.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgCreateWhoIs(m.Address, m.PubKey, docBz, rt.WhoIsType_USER)
	resp, err := m.SendTx("registry.MsgCreateWhoIs", msg1)
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
	resp, err := m.SendTx("registry.MsgUpdateWhoIs", msg1)
	if err != nil {
		return nil, err
	}

	cwir := &rt.MsgUpdateWhoIsResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cwir); err != nil {
		return nil, err
	}

	return cwir, nil
}
