package client_test

import (
	"errors"
	"fmt"
	"io/ioutil"
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

const TEMP_ENV_RENAME_FILE_NAME = ".env.temp.rename.client.package.test"
const ENV_FILE_CONTENT = `BLOCKCHAIN_REST="http://137.184.190.146:1317"
BLOCKCHAIN_FAUCET="http://137.184.190.146:8000"
BLOCKCHAIN_RPC="137.184.190.146:9090"
IPFS_ADDRESS="https://ipfs.sonr.ws"
IPFS_API_ADDRESS="https://api.ipfs.sonr.ws"
VAULT_API_ADDRESS="http://164.92.99.233"
VAULT_ENDPOINT="https://vault.sonr.ws"`

type Client interface {
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
	motorNode motor.MotorNode
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

	// setup test .env
	env_path := filepath.Join(projectpath.Root, ".env")
	_, err = os.Stat(env_path)
	if err == nil {
		// .env already exists rename it
		new_path := filepath.Join(projectpath.Root, TEMP_ENV_RENAME_FILE_NAME)
		err = os.Rename(env_path, new_path)
		if err != nil {
			fmt.Printf("Failed to rename .env file: %s", err)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Failed to check existence of .env file: %s", err)
	}

	// copy .env file to project root
	// test_env := filepath.Join(projectpath.Root, ".env.example")
	// input, err := ioutil.ReadFile(test_env)
	// if err != nil {
	// 	fmt.Printf("Failed to read test .env file: %s", err)
	// }

	err = ioutil.WriteFile(env_path, []byte(ENV_FILE_CONTENT), 0644)
	if err != nil {
		fmt.Printf("Failed to create .env file: %s", err)
	}
}

func (suite *ClientTestSuite) TearDownSuite() {
	testKeysPath := filepath.Join(projectpath.Root, "pkg/motor/test_keys/psksnr*")
	
	// delete created accounts
	files, err := filepath.Glob(testKeysPath)
	if err != nil {
		fmt.Printf("Failed to clean up generated test keys")
	}

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("Failed to clean up %s", file)
		}
	}

	// delete .env
	env_path := filepath.Join(projectpath.Root, ".env")
	err = os.Remove(env_path)
	if err != nil {
		fmt.Printf("Failed to clean up .env file: %s", err)
	}

	// rename old .env back to .env if it exists
	old_env := filepath.Join(projectpath.Root, TEMP_ENV_RENAME_FILE_NAME)
	_, err = os.Stat(old_env)
	if err == nil {
		err = os.Rename(old_env, env_path)
		if err != nil {
			fmt.Printf("Failed to rename old .env file: %s", err)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Failed to check existence of old .env file: %s", err)
	}

	fmt.Println("Teardown of test suite complete.")
}

func Test_ClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
 