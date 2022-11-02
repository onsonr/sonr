package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type popShardResponse struct {
	Shard string `json:"shard"`
}

func (v *vaultImpl) PopShard(did string) (Shard, error) {
	popShardFn := func() ([]byte, error) {
		res, err := http.Get(fmt.Sprintf("%s/did/%s/shard/pop", v.vaultEndpoint, did))
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		return ioutil.ReadAll(res.Body)
	}

	body, err := retryBuf(3, time.Second*4, popShardFn)
	if err != nil {
		return nil, err
	}

	var psr popShardResponse
	if err := json.Unmarshal(body, &psr); err != nil {
		return nil, err
	}

	return Shard(psr.Shard), nil
}
