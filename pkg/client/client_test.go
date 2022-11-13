package client_test

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

func (suite *ClientTestSuite) Test_FaucetCheckBalance() {
	err := suite.Client.RequestFaucet(suite.GetAddr())
	assert.NoError(suite.T(), err, "faucet request succeeds")
	_, err = suite.Client.CheckBalance(suite.GetAddr())
	assert.NoError(suite.T(), err, "Check Balance succeeds")
	acc, err := suite.Client.GetAccount(suite.GetAddr())
	assert.NoError(suite.T(), err, "Get Account succeeds")
	fmt.Println(acc.GetAddress(), acc.GetAccountNumber(), acc.GetSequence())
}

func (suite *ClientTestSuite) Test_QueryWhoIsWithoutCreatingWhoIs() {
	_, err := suite.Client.QueryWhoIs(suite.GetAddr())
	assert.Error(suite.T(), err, "whoIs should not exist")
}
