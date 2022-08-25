package motor

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sonr-io/sonr/pkg/client"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/assert"
)

// TODO: improve test suite (make more robust for CI/CID)
const ADDR = "snr1fmmf89mtzcwju5mf227tzkuae3y42cvpr0l58k"

func Test_DecodeTxData(t *testing.T) {
	data := "0A91010A242F736F6E72696F2E736F6E722E72656769737472792E4D736743726561746557686F497312691267122A736E723134373071366D3476776D6537346A376D3573326364773939357A35796E6B747A726D377A35371A31122F6469643A736E723A3134373071366D3476776D6537346A376D3573326364773939357A35796E6B747A726D377A353730BC8FA197063801"

	mcr := &rt.MsgCreateWhoIsResponse{}
	err := client.DecodeTxResponseData(data, mcr)
	assert.NoError(t, err, "decodes tx data successfully")
	assert.Equal(t, "snr1470q6m4vwme74j7m5s2cdw995z5ynktzrm7z57", mcr.WhoIs.Owner)
}

func storeKey(n string, aesKey []byte) bool {
	name := fmt.Sprintf("./test_keys/%s", n)
	file, err := os.Create(name)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Write(aesKey)
	return err == nil
}

func loadKey(n string) []byte {
	name := fmt.Sprintf("./test_keys/%s", n)
	var file *os.File
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err = os.Create(name)
		if err != nil {
			return nil
		}
	} else if err != nil {
		fmt.Printf("load err: %s\n", err)
	} else {
		file, err = os.Open(name)
		if err != nil {
			return nil
		}
	}
	defer file.Close()

	key, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	return key
}
