package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateVault(t *testing.T) {
	deviceShards := [][]byte{
		{2, 2, 2, 2, 2},
	}
	dscShard := []byte{0, 1, 2, 3, 4, 5}
	pskShard := []byte{6, 7, 8, 9, 10, 11}
	recShard := []byte{12, 13, 14, 15, 16, 17}
	s, err := New().CreateVault("testdid", deviceShards, "test_device", dscShard, pskShard, recShard)
	assert.NoError(t, err, "vault created")

	fmt.Println(s.ServiceEndpoint["cid"])
}

func Test_GetVault(t *testing.T) {
	res, err := http.Get("http://127.0.0.1:1234/cid/QmPocYs5qTF7YCri5YH9eAe7DJ84CgZaz4wogNGWKmpVBv/get")
	assert.NoError(t, err, "GET succeeds")
	defer res.Body.Close()

	var gvr getVaultResponse
	// err = json.NewDecoder(res.Body).Decode(&v)
	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err, "read body")

	json.Unmarshal(body, &gvr)
	assert.NoError(t, err, "decode json succeeds")

	fmt.Printf("%+v\n", gvr)
	fmt.Printf("%+v\n", gvr.Vault.IssuedShards["test_device"])
}
