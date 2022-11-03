package vault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

type issueShardRequest struct {
	DscPub      string `json:"dsc_pub"`
	DscShard    string `json:"dsc_shard"`
	ShardSuffix string `json:"shard_suffix"`
}

type issueShardResponse struct {
	VaultCID string `json:"vault_cid"`
}

func (v *vaultImpl) IssueShard(d, shardSuffix, dscPub, dscShard string) (did.Service, error) {
	reqBody, err := json.Marshal(issueShardRequest{
		DscPub:      dscPub,
		DscShard:    dscShard,
		ShardSuffix: shardSuffix,
	})
	if err != nil {
		return did.Service{}, fmt.Errorf("marshal request: %s", err)
	}
	issueShardFn := func() ([]byte, error) {
		res, err := http.Post(
			fmt.Sprintf("%s/did/%s/shard/issue", v.vaultEndpoint, d),
			"application/json",
			bytes.NewBuffer(reqBody),
		)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		return ioutil.ReadAll(res.Body)
	}

	body, err := retryBuf(3, time.Second*4, issueShardFn)
	if err != nil {
		return did.Service{}, err
	}

	var isr issueShardResponse
	if err := json.Unmarshal(body, &isr); err != nil {
		return did.Service{}, err
	}

	uri, err := getVaultUri()
	if err != nil {
		fmt.Printf("Error when retrieving vault uri: %s\n", err)
	}

	return did.Service{
		ID:   ssi.MustParseURI(uri),
		Type: "vault",
		ServiceEndpoint: map[string]string{
			"cid": isr.VaultCID,
		},
	}, nil
}
