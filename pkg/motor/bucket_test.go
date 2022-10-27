package motor

import (
	"fmt"
	"testing"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_CreateBucket() {
	suite.T().Run("create single bucket", func(t *testing.T) {
		createReq := mt.CreateBucketRequest{
			Creator: suite.motorWithKeys.GetAddress(),
			Label:   "my awesome bucket",
		}
		_, b, err := suite.motorWithKeys.CreateBucket(createReq)

		assert.NoError(t, err)
		assert.NotNil(t, b)
	})

	suite.T().Run("create many buckets", func(t *testing.T) {
		uris := make([]*types.Bucket, 0)
		for i := 0; i < 3; i++ {
			var createReq mt.CreateBucketRequest
			if i == 0 {
				createReq = mt.CreateBucketRequest{
					Creator: suite.motorWithKeys.GetAddress(),
					Label:   fmt.Sprintf("my awesome bucket %d", i),
				}
			} else {
				createReq = mt.CreateBucketRequest{
					Creator: suite.motorWithKeys.GetAddress(),
					Label:   fmt.Sprintf("my awesome bucket %d", i),
				}
			}

			_, b, err := suite.motorWithKeys.CreateBucket(createReq)

			assert.NoError(t, err)
			assert.NotNil(t, b)

			if i != 0 {
				b.ResolveBuckets()
				buckets := b.GetBuckets()
				assert.Equal(t, len(buckets), len(uris))
			}
		}
	})
}
