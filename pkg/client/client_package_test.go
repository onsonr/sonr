package client_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/internal/projectpath"
	"github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/suite"
)

type Client interface{
	GetFaucetAddress() string
	GetRPCAddress() string
	GetAPIAddress() string
	GetIPFSAddress() string
	GetIPFSApiAddress() string
	CheckBalance(string) (types.Coins, error) 
	RequestFaucet(string) error
	QueryWhoIs(string) (*rt.WhoIs, error)
	PrintConnectionEndpoints()
}

type ClientTestSuite struct {
	suite.Suite
	motorNode *motor.MotorNodeImpl
}

func (suite *ClientTestSuite) SetupSuite() {
	var err error

	suite.motorNode, err = motor.EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	if err != nil {
		fmt.Printf("Failed to setup test suite motor: %s", err)
	}

	motor.SetupTestAddressWithKeys(suite.motorNode)
}

func (suite *ClientTestSuite) TearDownSuite() {
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

func Test_ClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
 