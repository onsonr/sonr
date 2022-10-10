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
			Creator:    suite.motorWithKeys.Address,
			Label:      "my awesome bucket",
			Visibility: types.BucketVisibility_PUBLIC,
			Role:       types.BucketRole_USER,
			Content:    make([]*types.BucketItem, 0),
		}
		_, b, err := suite.motorWithKeys.CreateBucket(createReq)

		assert.NoError(t, err)
		assert.NotNil(t, b)
	})

	suite.T().Run("create many buckets", func(t *testing.T) {
		uris := make([]*types.BucketItem, 0)
		for i := 0; i < 3; i++ {
			var createReq mt.CreateBucketRequest
			if i == 0 {
				createReq = mt.CreateBucketRequest{
					Creator:    suite.motorWithKeys.Address,
					Label:      fmt.Sprintf("my awesome bucket %d", i),
					Visibility: types.BucketVisibility_PUBLIC,
					Role:       types.BucketRole_USER,
					Content:    make([]*types.BucketItem, 0),
				}
			} else {
				createReq = mt.CreateBucketRequest{
					Creator:    suite.motorWithKeys.Address,
					Label:      fmt.Sprintf("my awesome bucket %d", i),
					Visibility: types.BucketVisibility_PUBLIC,
					Role:       types.BucketRole_USER,
					Content:    uris,
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

			uris = append(uris, &types.BucketItem{
				Name: "content",
				Uri:  b.GetDID(),
				Type: types.ResourceIdentifier_DID,
			})
		}
	})
}
