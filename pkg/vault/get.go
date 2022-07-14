package vault

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Shard struct {
	Value []byte `json:"value"`
}

type jsonShard struct {
	Base64Value string `json:"value"`
}

func (s *Shard) UnmarshalJSON(b []byte) error {
	var js jsonShard
	if err := json.Unmarshal(b, &js); err != nil {
		return err
	}

	// decode base64
	rawText, err := base64.StdEncoding.DecodeString(js.Base64Value)
	if err != nil {
		return err
	}
	*s = Shard{
		Value: []byte(rawText),
	}
	return nil
}

type vault struct {
	ShardBank     []Shard          `json:"shard_bank"`
	IssuedShards  map[string]Shard `json:"issued_shards"`
	PskShard      Shard            `json:"psk_shard"`
	RecoveryShard Shard            `json:"recovery_shard"`
}


type getVaultResponse struct {
	Vault vault `json:"vault"`
}

func (v *vaultImpl) GetVaultShards(did string) (vault, error) {
	getVaultFunc := func() ([]byte, error) {
		res, err := http.Get(fmt.Sprintf("%s/did/%s/get", v.vaultEndpoint, did))
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != http.StatusOK {
			var er errorResponse
			if err := json.Unmarshal(body, &er); err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("error getting vault: %d: %s", res.StatusCode, er)
		}

		return body, nil
	}

	body, err := retryBuf(3, time.Second*3, getVaultFunc)
	if err != nil {
		return vault{}, err
	}

	var gvr getVaultResponse
	if err := json.Unmarshal(body, &gvr); err != nil {
		return vault{}, err
	}

	return gvr.Vault, nil
}
