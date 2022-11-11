package client_test

import (
	"github.com/stretchr/testify/assert"
)

func (suite *ClientTestSuite) Test_FaucetCheckBalance() {
	err := suite.Client.RequestFaucet(suite.GetAddr())
	assert.NoError(suite.T(), err, "faucet request succeeds")
	_, err = suite.Client.CheckBalance(suite.GetAddr())
	assert.NoError(suite.T(), err, "Check Balance succeeds")
}

func (suite *ClientTestSuite) Test_QueryWhoIsWithoutCreatingWhoIs() {
	_, err := suite.Client.QueryWhoIs(suite.GetAddr())
	assert.Error(suite.T(), err, "whoIs should not exist")
}
