package bucket_test

import (
	"context"
	"testing"
	"time"

	"github.com/sonr-io/sonr/pkg/client"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/bucket/types"
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
		item, err := suite.testBucket.GetContentById(suite.cidDoc1)
		assert.NoError(t, err)
		assert.NotNil(t, item)
		assert.ObjectsAreEqual(item.ContentType, bt.ResourceIdentifier_CID)
	})

	suite.T().Run("Bucket Service endpoint should be valid uri", func(t *testing.T) {
		ssi := suite.testBucket.CreateBucketServiceEndpoint()
		assert.NotNil(t, ssi)
		assert.ObjectsAreEqual(c.GetAPIAddress(), ssi.ID.Host)
	})

	suite.T().Run("Should add item to existing bucket", func(t *testing.T) {
		items := suite.testBucket.GetBucketItems()
		assert.Equal(t, len(items), 1)
		item := bt.BucketItem{
			Name:      "test",
			Uri:       suite.cidDoc2,
			Timestamp: time.Now().Unix(),
			Type:      bt.ResourceIdentifier_CID,
		}
		items = append(items, &item)
		assert.Equal(t, len(items), 2)
		err := suite.testBucket.ResolveContent()
		assert.NoError(t, err)
		updatedBucket, err := suite.motorNode.UpdateBucketItems(context.Background(), suite.testBucket.GetDID(), items)
		assert.NoError(t, err)
		items = updatedBucket.GetBucketItems()
		assert.Equal(t, len(items), 2)
	})

	suite.T().Run("Should remove item from existing bucket", func(t *testing.T) {
		items := suite.testBucket.GetBucketItems()
		assert.Equal(t, len(items), 1)
		err := suite.testBucket.ResolveContent()
		assert.NoError(t, err)
		updatedBucket, err := suite.motorNode.UpdateBucketItems(context.Background(), suite.testBucket.GetDID(), []*bt.BucketItem{})
		assert.NoError(t, err)
		items = updatedBucket.GetBucketItems()
		assert.Equal(t, len(items), 0)
	})

	suite.T().Run("Should modify item in existing bucket", func(t *testing.T) {
		items := suite.testBucket.GetBucketItems()
		assert.Equal(t, len(items), 1)
		name := items[0].GetName()
		assert.Equal(t, name, "test")
		items = []*bt.BucketItem{
			{
				Name:      "test updated",
				Uri:       suite.cidDoc1,
				Timestamp: time.Now().Unix(),
				Type:      bt.ResourceIdentifier_CID,
			},
		}
		updatedBucket, err := suite.motorNode.UpdateBucketItems(context.Background(), suite.testBucket.GetDID(), items)
		assert.NoError(t, err)
		items = updatedBucket.GetBucketItems()
		assert.Equal(t, len(items), 1)
		name = items[0].GetName()
		assert.Equal(t, name, "test updated")
	})

	suite.T().Run("Should add a bucket to a bucket", func(t *testing.T) {
		// create new bucket
		_, newBucket, err := suite.motorNode.CreateBucket(mt.CreateBucketRequest{
			Creator:    suite.motorNode.GetAddress(),
			Label:      "test bucket",
			Visibility: types.BucketVisibility_PUBLIC,
			Role:       types.BucketRole_USER,
			Content:    []*bt.BucketItem{},
		})
		assert.NoError(t, err)
		// add new bucket to old bucket
		items := []*bt.BucketItem{
			{
				Name:      "test bucket",
				Uri:       newBucket.GetDID(),
				Timestamp: time.Now().Unix(),
				Type:      bt.ResourceIdentifier_DID,
			},
		}
		updatedBucket, err := suite.motorNode.UpdateBucketItems(context.Background(), suite.testBucket.GetDID(), items)
		assert.NoError(t, err)
		buckets := updatedBucket.GetBuckets()
		assert.Equal(t, len(buckets), 1)
		name := items[0].GetName()
		assert.Equal(t, name, "test bucket")
	})

	suite.T().Run("Should resolve buckets", func(t *testing.T) {
		// create first bucket
		_, firstBucket, err := suite.motorNode.CreateBucket(mt.CreateBucketRequest{
			Creator:    suite.motorNode.GetAddress(),
			Label:      "first bucket",
			Visibility: types.BucketVisibility_PUBLIC,
			Role:       types.BucketRole_USER,
			Content:    []*bt.BucketItem{},
		})
		assert.NoError(t, err)
		content := []*bt.BucketItem{
			{
				Name:      "first bucket",
				Uri:       firstBucket.GetDID(),
				Timestamp: time.Now().Unix(),
				Type:      bt.ResourceIdentifier_DID,
			},
		}
		// create second bucket with first as content
		_, secondBucket, err := suite.motorNode.CreateBucket(mt.CreateBucketRequest{
			Creator:    suite.motorNode.GetAddress(),
			Label:      "second bucket",
			Visibility: types.BucketVisibility_PUBLIC,
			Role:       types.BucketRole_USER,
			Content:    content,
		})
		assert.NoError(t, err)
		// get buckets without resolving
		buckets := secondBucket.GetBuckets()
		assert.Equal(t, len(buckets), 0)
		// resolve buckets
		err = secondBucket.ResolveBuckets()
		assert.NoError(t, err)
		// get buckets after resolving
		buckets = secondBucket.GetBuckets()
		assert.Equal(t, len(buckets), 1)
	})
}
