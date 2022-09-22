package motor

import (
	"fmt"
	"log"
	"testing"

	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_CreateAccountWithKeyring() {
	req := mt.CreateAccountRequest{
		Password: "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.CreateAccount(req)
	assert.NoError(suite.T(), err, "wallet generation succeeds")

	b := m.GetBalance()
	log.Println("balance:", b)

	// Print the address of the wallet
	log.Println("address:", m.Address)
}

func (suite *MotorTestSuite) Test_CreateAccountWithKeys() {
	aesKey := loadKey("aes.key")
	if aesKey == nil || len(aesKey) != 32 {
		key, err := mpc.NewAesKey()
		assert.NoError(suite.T(), err, "generates aes key")
		aesKey = key

		// store the key
		fmt.Printf("stored key? %v\n", storeKey("aes.key", key))
	} else {
		fmt.Println("loaded key")
	}

	psk, err := mpc.NewAesKey()
	assert.NoError(suite.T(), err, "create psk")

	req := mt.CreateAccountWithKeysRequest{
		Password:  "password123",
		AesDscKey: aesKey,
		AesPskKey: psk,
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err = m.CreateAccountWithKeys(req)
	assert.NoError(suite.T(), err, "wallet generation succeeds")

	b := m.GetBalance()
	log.Println("balance:", b)

	// Print the address of the wallet
	log.Println("address:", m.Address)

	// store PSK
	storeKey(fmt.Sprintf("psk%s", m.Address), psk)
}

func (suite *MotorTestSuite) Test_LoginWithKeys() {
	suite.T().Skip()
	suite.T().Run("with password and psk", func(t *testing.T) {
		pskKey := loadKey(fmt.Sprintf("psk%s", suite.accountAddress))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginWithKeysRequest{
			Did:       suite.accountAddress,
			Password:  "password123",
			AesPskKey: pskKey,
		}

		m, _ := EmptyMotor(&mt.InitializeRequest{
			DeviceId: "test_device",
		}, common.DefaultCallback())
		_, err := m.LoginWithKeys(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", m.GetBalance())
			fmt.Println("address: ", m.Address)
		}
	})

	suite.T().Run("with DSC and PSK", func(t *testing.T) {
		aesKey := loadKey("aes.key")
		fmt.Printf("aes: %x\n", aesKey)
		if aesKey == nil || len(aesKey) != 32 {
			t.Errorf("could not load key.")
			return
		}

		pskKey := loadKey(fmt.Sprintf("psk%s", suite.accountAddress))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginWithKeysRequest{
			Did:       suite.accountAddress,
			AesDscKey: aesKey,
			AesPskKey: pskKey,
		}

		m, _ := EmptyMotor(&mt.InitializeRequest{
			DeviceId: "test_device",
		}, common.DefaultCallback())
		_, err := m.LoginWithKeys(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", m.GetBalance())
			fmt.Println("address: ", m.Address)
		}
	})
}

func (suite *MotorTestSuite) Test_LoginWithKeyring() {
	suite.T().Skip()
	req := mt.LoginRequest{
		Did: suite.accountAddress,
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(suite.T(), err, "login succeeds")

	if err == nil {
		fmt.Println("balance: ", m.GetBalance())
		fmt.Println("address: ", m.Address)
	}
}

func (suite *MotorTestSuite) Test_LoginAndMakeRequest() {
	suite.T().Skip()
	pskKey := loadKey(fmt.Sprintf("psk%s", suite.accountAddress))
	if pskKey == nil || len(pskKey) != 32 {
		suite.T().Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:      suite.accountAddress,
		Password: "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(suite.T(), err, "login succeeds")

	// do something with the logged in account
	m.DIDDocument.AddAlias("gotest.snr")
	_, err = updateWhoIs(m)
	assert.NoError(suite.T(), err, "updates successfully")
}
