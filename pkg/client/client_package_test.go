package client_test

import (
	"testing"

	"github.com/ignite-hq/cli/ignite/pkg/cosmosaccount"
	"github.com/sonr-io/sonr/pkg/client"
	v1 "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/suite"
)

const testMnemonic = "angry fiber canoe setup moral click dove diet toddler polar tonight reduce nuclear plastic imitate range stick sick artist vast opera artefact wolf slam"

const testAccountName = "myTestAccount"

type ClientTestSuite struct {
	suite.Suite
	Account cosmosaccount.Account
	Client  *client.Client
}

func (suite *ClientTestSuite) GetAddr() string {
	//"snr1gplscgegppys9ss8m9ykkp25x49ygef62rxtlu"
	return suite.Account.Address("snr")
}

func (suite *ClientTestSuite) SetupSuite() {
	suite.Client = client.NewClient(v1.ClientMode_ENDPOINT_LOCAL)
	tmpDir := suite.T().TempDir()
	secondRegistry, _ := cosmosaccount.New(cosmosaccount.WithHome(tmpDir))
	importedAccount, _ := secondRegistry.Import(testAccountName, testMnemonic, "")
	suite.Account = importedAccount
}

func (suite *ClientTestSuite) TearDownSuite() {}

func Test_ClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
