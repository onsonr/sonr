package motor

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sonr-io/sonr/internal/projectpath"
	mtu "github.com/sonr-io/sonr/testutil/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/suite"
)

type MotorTestSuite struct {
	suite.Suite
	motorWithKeys MotorNode
}

func (suite *MotorTestSuite) SetupSuite() {
	fmt.Println("Setting up suite...")

	var err error

	// setup motor
	suite.motorWithKeys, err = EmptyMotor(&mt.InitializeRequest{
		DeviceId:   "test_device",
		ClientMode: mt.ClientMode_ENDPOINT_BETA,
	}, common.DefaultCallback())

	if err != nil {
		suite.T().Error("Failed to setup test suite motor with keys")
	}

	err = mtu.SetupTestAddressWithKeys(suite.motorWithKeys)
	if err != nil {
		suite.T().Error("Failed to setup test address with keys")
	}

	fmt.Printf("Setup test address with keys: %s\n", suite.motorWithKeys.GetAddress())
}

func (suite *MotorTestSuite) TearDownSuite() {
	testKeysPath := filepath.Join(projectpath.Root, "pkg/motor/test_keys/psksnr*")

	// delete created accounts
	files, err := filepath.Glob(testKeysPath)
	if err != nil {
		suite.T().Error("Failed to clean up generated test keys")
	}

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			suite.T().Errorf("Failed to clean up %s", file)
		}
	}

	fmt.Println("Teardown of test suite complete.")
}

func Test_MotorTestSuite(t *testing.T) {
	suite.Run(t, new(MotorTestSuite))
}
