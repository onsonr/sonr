package motor

import (
	"encoding/json"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/pkg/vault"
	rt "github.com/sonr-io/sonr/x/registry/types"
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func (m *MotorNode) CreateAccount(requestBytes []byte) (rtmv1.CreateAccountResponse, error) {
	var request rtmv1.CreateAccountRequest
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

	// create Vault shards to make sure this works before creating WhoIs
	vc := vault.New()
	dscShard, err := dscEncrypt(m.deviceShard, request.AesDscKey)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}

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

	// create vault
	vaultService, err := vc.CreateVault(
		m.Address,
		m.unusedShards,
		string(m.DeviceID),
		dscShard,
		pskShard,
		pwShard,
	)
	if err != nil {
		fmt.Println("[WARN] failed to create vault:", err)
	}

	// update DID Document
	m.DIDDocument.AddService(vaultService)

	// update whois
	resp, err := updateWhoIs(m)
	if err != nil {
		return rtmv1.CreateAccountResponse{}, err
	}
	fmt.Println(resp.String())

	return rtmv1.CreateAccountResponse{
		Address: m.Address,
		AesPsk:  psk,
	}, err
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
