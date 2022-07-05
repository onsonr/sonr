package vault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

var (
	DefaultVaultService = did.Service{
		ID:   ssi.MustParseURI("https://vault.sonr.ws"),
		Type: "vault",
	}
)

type createVaultRequest struct {
	DeviceShards      []string `json:"device_shards"`       // a list of shards encrypted with the MPC
	DscPub            string   `json:"dsc_pub"`             // the dsc public key of the creator
	EncryptedDscShard string   `json:"encrypted_dsc_shard"` // the shard for the creator, encrypted with DscPub
	PskShard          string   `json:"psk_shard"`           // the shard encrypted with PSK
	RecoveryShard     string   `json:"recovery_shard"`      // shard encrypted with password
}

type createVaultResponse struct {
	VaultCid string `json:"vault_cid"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func (v *vaultImpl) CreateVault(d string, deviceShards []string, dscPub, encDscShard, pskShard, recShard string) (did.Service, error) {
	reqBody, err := json.Marshal(createVaultRequest{
		DeviceShards:      deviceShards,
		DscPub:            dscPub,
		EncryptedDscShard: encDscShard,
		PskShard:          pskShard,
		RecoveryShard:     recShard,
	})
	if err != nil {
		return DefaultVaultService, err
	}

	createVaultFunc := func() ([]byte, error) {
		res, err := http.Post(
			fmt.Sprintf("%s/did/%s/create", v.vaultEndpoint, d),
			"application/json",
			bytes.NewBuffer(reqBody),
		)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		return ioutil.ReadAll(res.Body)
	}

	body, err := retryBuf(3, time.Second*3, createVaultFunc)
	if err != nil {
		return DefaultVaultService, err
	}

	var cvr createVaultResponse
	err = json.Unmarshal(body, &cvr)
	if err != nil {
		return DefaultVaultService, err
	}
	if cvr.VaultCid == "" {
		var errRes errorResponse
		err = json.Unmarshal(body, &errRes)
		if err != nil {
			return DefaultVaultService, err
		}
		return DefaultVaultService, fmt.Errorf("error creating vault: %s", errRes.Message)
	}

	return did.Service{
		ID:   ssi.MustParseURI("https://vault.sonr.ws"),
		Type: "vault",
		ServiceEndpoint: map[string]string{
			"cid": cvr.VaultCid,
		},
	}, nil
}

func retryBuf(attempts int, sleep time.Duration, f func() ([]byte, error)) (buf []byte, err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		buf, err = f()
		if err == nil {
			return buf, nil
		}
	}
	return nil, fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
