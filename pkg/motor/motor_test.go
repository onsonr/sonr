package motor

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/assert"
)

// TODO: improve test suite (make more robust for CI/CID)
const ADDR = "snr19c99rqjsts86mm4t6u8qzy2al3ghkfgu7f2zua"

func Test_DecodeTxData(t *testing.T) {
	data := "0A91010A242F736F6E72696F2E736F6E722E72656769737472792E4D736743726561746557686F497312691267122A736E723134373071366D3476776D6537346A376D3573326364773939357A35796E6B747A726D377A35371A31122F6469643A736E723A3134373071366D3476776D6537346A376D3573326364773939357A35796E6B747A726D377A353730BC8FA197063801"

	mcr := &rt.MsgCreateWhoIsResponse{}
	err := client.DecodeTxResponseData(data, mcr)
	assert.NoError(t, err, "decodes tx data successfully")
	assert.Equal(t, "snr1470q6m4vwme74j7m5s2cdw995z5ynktzrm7z57", mcr.WhoIs.Owner)
}

func Test_GetAddress(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		AccountId: ADDR,
		Password:  "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	assert.Equal(t, ADDR, m.GetAddress())
}
