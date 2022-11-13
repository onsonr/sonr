package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Shard []byte

type Vault struct {
	ShardBank     []Shard          `json:"shard_bank"`
	IssuedShards  map[string]Shard `json:"issued_shards"`
	PskShard      Shard            `json:"psk_shard"`
	RecoveryShard Shard            `json:"recovery_shard"`
}

type getVaultResponse struct {
	Vault Vault `json:"vault"`
}

func (v *vaultImpl) GetVaultShards(did string) (Vault, error) {
	getVaultFunc := func() ([]byte, error) {
		fmt.Printf("%s/did/%s/get", v.vaultEndpoint, did)
		res, err := http.Get(fmt.Sprintf("%s/did/%s/get", v.vaultEndpoint, did))
		if err != nil {
			return nil, fmt.Errorf("error in /did/%s/get request: %s", did, err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading get vault response body: %s", err)
		}

		statusFamily := res.StatusCode / 100
		if statusFamily == 5 {
			return nil, fmt.Errorf("get vault shards error: %d", res.StatusCode)
		} else if statusFamily != 2 {
			var er errorResponse
			if err := json.Unmarshal(body, &er); err != nil {
				return nil, fmt.Errorf("error unmarshalling err response: %s. Response body: %s", err, body)
			}
			return nil, fmt.Errorf("error getting vault: %d: %s", res.StatusCode, er)
		}

		return body, nil
	}

	body, err := retryBuf(3, time.Second*4, getVaultFunc)
	if err != nil {
		return Vault{}, err
	}

	var gvr getVaultResponse
	if err := json.Unmarshal(body, &gvr); err != nil {
		return Vault{}, err
	}

	return gvr.Vault, nil
}
