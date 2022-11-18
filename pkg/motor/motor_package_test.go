package motor

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sonr-io/sonr/internal/projectpath"
	mtu "github.com/sonr-io/sonr/testutil/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	v1 "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MotorTestSuite struct {
	suite.Suite
	motorWithKeys MotorNode
	callback      testSuiteCallback
}

type testSuiteCallback struct {
	discoverChannel    chan v1.RefreshEvent
	walletEventChannel chan ct.WalletEvent
	linkingChannel     chan v1.LinkingEvent
}

func (c testSuiteCallback) OnDiscover(data []byte) {
	ev := v1.RefreshEvent{}
	if err := ev.Unmarshal(data); err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	go func() {
		fmt.Printf("OnDiscover: %+v\n", ev)
		c.discoverChannel <- ev
	}()
}

func (c testSuiteCallback) OnWalletEvent(data []byte) {
	ev := ct.WalletEvent{}
	if err := ev.Unmarshal(data); err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	go func() {
		fmt.Printf("OnWalletEvent: %+v\n", ev)
		c.walletEventChannel <- ev
	}()
}

func (c testSuiteCallback) OnLinking(data []byte) {
	ev := v1.LinkingEvent{}
	if err := ev.Unmarshal(data); err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	go func() {
		fmt.Printf("OnLinking: %+v\n", ev)
		c.linkingChannel <- ev
	}()
}

func (suite *MotorTestSuite) SetupSuite() {
	fmt.Println("Setting up suite...")

	suite.callback = testSuiteCallback{
		discoverChannel:    make(chan v1.RefreshEvent),
		walletEventChannel: make(chan ct.WalletEvent),
		linkingChannel:     make(chan v1.LinkingEvent),
	}

	var err error
	// setup motor
	suite.motorWithKeys, err = EmptyMotor(&mt.InitializeRequest{
		DeviceId:   "test_device",
		ClientMode: mt.ClientMode_ENDPOINT_LOCAL,
	}, suite.callback)

	if err != nil {
		suite.T().Error("Failed to setup test suite motor with keys")
	}

	err = mtu.SetupTestAddressWithKeys(suite.motorWithKeys)
	if err != nil {
		suite.T().Error("Failed to setup test address with keys")
	}

	//Wait for Vault Creation to Finish
	for {
		walletEvent, ok := <-suite.callback.walletEventChannel
		assert.True(suite.T(), ok, "Failed to read data from walletEventChannel")
		if walletEvent.Type == common.WALLET_EVENT_TYPE_VAULT_CREATE_END {
			break
		}
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
