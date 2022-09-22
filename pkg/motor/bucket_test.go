package motor

import (
	"context"
	"fmt"

	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_CreateBucket() {
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

	createReq := mt.CreateBucketRequest{
		Creator:    suite.accountAddress,
		Label:      "my awesome bucket",
		Visibility: types.BucketVisibility_PUBLIC,
		Role:       types.BucketRole_USER,
		Content:    make([]*types.BucketItem, 0),
	}
	b, err := m.CreateBucket(context.Background(), createReq)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), b)
}

func (suite *MotorTestSuite) Test_GetBucket() {
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

	b, err := m.GetBucket("did:snr:e7360323120e4b70a5366984c994b536")
	assert.NoError(suite.T(), err, "get bucket")

	b.ResolveContent()
	for _, c := range b.GetBucketItems() {
		fmt.Println(c.Name)
	}
}
