package motor

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_LoginWithKeys() {
	suite.T().Run("with password and psk", func(t *testing.T) {
		pskKey := loadKey(fmt.Sprintf("psk%s", suite.motorWithKeys.Address))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginWithKeysRequest{
			AccountId: suite.motorWithKeys.Address,
			Password:  "password123",
			AesPskKey: pskKey,
		}

		_, err := suite.motorWithKeys.LoginWithKeys(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", suite.motorWithKeys.GetBalance())
			fmt.Println("address: ", suite.motorWithKeys.Address)
		}
	})

	suite.T().Run("with DSC and PSK", func(t *testing.T) {
		aesKey := loadKey("aes.key")
		fmt.Printf("aes: %x\n", aesKey)
		if aesKey == nil || len(aesKey) != 32 {
			t.Errorf("could not load key.")
			return
		}

		pskKey := loadKey(fmt.Sprintf("psk%s", suite.motorWithKeys.Address))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginWithKeysRequest{
			AccountId: suite.motorWithKeys.Address,
			AesDscKey: aesKey,
			AesPskKey: pskKey,
		}

		_, err := suite.motorWithKeys.LoginWithKeys(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", suite.motorWithKeys.GetBalance())
			fmt.Println("address: ", suite.motorWithKeys.Address)
		}
	})
}

func Test_LoginWithKeyring(t *testing.T) {
	req := mt.LoginRequest{
		AccountId: ADDR,
		Password:  "password123",
	}

	m, err := EmptyMotor(&mt.InitializeRequest{
		DeviceId:   "test_device",
		ClientMode: mt.ClientMode_ENDPOINT_BETA,
	}, common.DefaultCallback())
	assert.NoError(t, err, "create motor")

	_, err = m.Login(req)
	assert.NoError(t, err, "login succeeds")

	if err == nil {
		fmt.Println("balance: ", m.GetBalance())
		fmt.Println("address: ", m.Address)
	}
}

func (suite *MotorTestSuite) Test_LoginAndMakeRequest() {
	aesKey := loadKey("aes.key")
	fmt.Printf("aes: %x\n", aesKey)
	if aesKey == nil || len(aesKey) != 32 {
		suite.T().Errorf("could not load key.")
		return
	}

	pskKey := loadKey(fmt.Sprintf("psk%s", suite.motorWithKeys.Address))
	if pskKey == nil || len(pskKey) != 32 {
		suite.T().Errorf("could not load psk key")
		return
	}

	req := mt.LoginWithKeysRequest{
		AccountId: suite.motorWithKeys.Address,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	_, err := suite.motorWithKeys.LoginWithKeys(req)
	assert.NoError(suite.T(), err, "login succeeds")

	// do something with the logged in account
	suite.motorWithKeys.DIDDocument.AddAlias("gotest.snr")
	_, err = updateWhoIs(suite.motorWithKeys)
	assert.NoError(suite.T(), err, "updates successfully")
}

func (suite *MotorTestSuite) Test_LoginWithAlias() {
	pskKey := loadKey(fmt.Sprintf("psk%s", suite.motorWithKeys.Address))
	if pskKey == nil || len(pskKey) != 32 {
		suite.T().Errorf("could not load psk key")
		return
	}

	// alias := fmt.Sprintf("%s", randSeq(6))
	alias := fmt.Sprintf("%s.snr", randSeq(6))
	suite.motorWithKeys.DIDDocument.AddAlias(alias)
	_, err := updateWhoIs(suite.motorWithKeys)
	assert.NoError(suite.T(), err, "buy alias successfully")

	req := mt.LoginWithKeysRequest{
		AccountId: alias,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	_, err = suite.motorWithKeys.LoginWithKeys(req)
	assert.NoError(suite.T(), err, "login succeeds")
}

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
