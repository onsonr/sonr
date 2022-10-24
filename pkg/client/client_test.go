package client_test

import (
	"fmt"
	"os"

	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/third_party/types/common"

	"github.com/stretchr/testify/assert"
)

func (suite *ClientTestSuite) Test_FaucetCheckBalance() {
	// Create Client instance and Generate wallet
	w, err := mpc.GenerateWallet(common.DefaultCallback())
	assert.NoError(suite.T(), err, "wallet generation succeeds")

	// Get Wallet Address
	addr, err := w.Address()
	assert.NoError(suite.T(), err, "Bech32Address successfully created")
	fmt.Println("Address:", addr)

	// Check Balance
	resp, err := suite.motorNode.GetClient().CheckBalance(addr)
	assert.NoError(suite.T(), err, "Check Balance succeeds")
	fmt.Printf("-- Get Balances (1) --\n%+v\n", resp)

	// Request Faucet and Check Again
	err = suite.motorNode.GetClient().RequestFaucet(addr)
	assert.NoError(suite.T(), err, "faucet request succeeds")
	resp2, err := suite.motorNode.GetClient().CheckBalance(addr)
	assert.NoError(suite.T(), err, "Check Balance succeeds")
	fmt.Printf("-- Get Balances (2) --\n%+v\n", resp2)
}

func (suite *ClientTestSuite) Test_QueryWhoIs() {
	acc, err := suite.motorNode.GetClient().QueryWhoIs(suite.motorNode.GetAddress())
	assert.NoError(suite.T(), err, "QueryAccount succeeds")
	fmt.Printf("-- Get Account --\n%+v\n", acc)
}

func (suite *ClientTestSuite) Test_GetFaucetAddress() {
	assert.Equal(suite.T(), os.Getenv("BLOCKCHAIN_FAUCET"), suite.motorNode.GetClient().GetFaucetAddress(), "Faucet address should be the same")
}

func (suite *ClientTestSuite) Test_GetRPCAddress() {
	assert.Equal(suite.T(), os.Getenv("BLOCKCHAIN_RPC"), suite.motorNode.GetClient().GetRPCAddress(), "RPC address should be the same")
}

func (suite *ClientTestSuite) Test_GetAPIAddress() {
	assert.Equal(suite.T(), os.Getenv("BLOCKCHAIN_REST"), suite.motorNode.GetClient().GetAPIAddress(), "API address should be the same")
}

func (suite *ClientTestSuite) Test_GetIPFSAddress() {
	assert.Equal(suite.T(), os.Getenv("IPFS_ADDRESS"), suite.motorNode.GetClient().GetIPFSAddress(), "IPFS address should be the same")
}

func (suite *ClientTestSuite) Test_GetIPFSApiAddress() {
	assert.Equal(suite.T(), os.Getenv("IPFS_API_ADDRESS"), suite.motorNode.GetClient().GetIPFSApiAddress(), "IPFS API address should be the same")
}
