package bucket_test

import (
	"testing"

	"github.com/sonr-io/sonr/pkg/client"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/assert"
)

func (suite *BucketTestSuite) Test_Bucket() {
	c := client.NewClient(mt.ClientMode_ENDPOINT_BETA)
	suite.T().Run("Bucket should be defined", func(t *testing.T) {
		assert.NotNil(t, suite.testBucket)
	})

	suite.T().Run("Bucket Resolve cid should be in content cache", func(t *testing.T) {
		assert.NotNil(t, suite.testBucket)
		err := suite.testBucket.ResolveContent()
		assert.NoError(t, err)
		item, err := suite.testBucket.GetContentById(suite.cid)
		assert.NoError(t, err)
		assert.NotNil(t, item)
		assert.ObjectsAreEqual(item.ContentType, bt.ResourceIdentifier_CID)
	})

	suite.T().Run("Bucket Service endpoint should be valid uri", func(t *testing.T) {
		ssi := suite.testBucket.CreateBucketServiceEndpoint()
		assert.NotNil(t, ssi)
		assert.ObjectsAreEqual(c.GetAPIAddress(), ssi.ID.Host)
	})
}
