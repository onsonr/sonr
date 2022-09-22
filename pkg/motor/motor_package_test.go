package motor

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/suite"
)

type MotorTestSuite struct {
	suite.Suite
	accountAddress string
}

func (suite *MotorTestSuite) SetupSuite() {
	fmt.Println()
	fmt.Println("Setting up suite")
	fmt.Println()
	var err error
	suite.accountAddress, err = setupTestAddress()
	if err != nil {
		suite.T().Errorf("Failed to setup test address")
	}

	fmt.Printf("Setup test address: %s\n", suite.accountAddress)
}

func Test_MotorTestSuite(t *testing.T) {
	suite.Run(t, new(MotorTestSuite))
}

func setupTestAddress() (string, error) {
	aesKey := loadKey("aes.key")
	if aesKey == nil || len(aesKey) != 32 {
		key, err := mpc.NewAesKey()
		if err != nil {
			return "", err
		}
		aesKey = key
	}

	psk, err := mpc.NewAesKey()
	if err != nil {
		return "", err
	}

	req := mt.CreateAccountWithKeysRequest{
		Password:  "password123",
		AesDscKey: aesKey,
		AesPskKey: psk,
	}

	motor, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err = motor.CreateAccountWithKeys(req)
	if err != nil {
		return "", err
	}

	storeKey(fmt.Sprintf("psk%s", motor.Address), psk)

	return motor.Address, nil
}