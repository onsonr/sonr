package bucket_test

import (
	"testing"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/assert"
)

var (
	DEV_CHAIN_ADDR = "137.184.190.146:9090"
)

func CreateMockWhereIs(creator string, content []*bt.BucketItem) *bt.WhereIs {
	return &bt.WhereIs{
		Creator:    creator,
		Content:    content,
		Did:        "did:snr:asdasd",
		Visibility: bt.BucketVisibility_PUBLIC,
		Role:       bt.BucketRole_USER,
		Timestamp:  time.Now().Unix(),
	}
}

func Test_Bucket(t *testing.T) {

	creator := "snr1ld9u3wpq752wmqaus5rzcfanqg65sgldhnscx5"
	objectURI := "bafyreihnj3feeesb6wmd46lmsvtwalvuckns647ghy44xn63lfsfed3ydm"
	s := shell.NewLocalShell()
	c := client.NewClient(client.ConnEndpointType_DEV)
	t.Run("Bucket should be defined", func(t *testing.T) {
		content := []*bt.BucketItem{
			{
				Name:      "test",
				Uri:       objectURI,
				Timestamp: time.Now().Unix(),
				Type:      bt.ResourceIdentifier_CID,
			},
		}

		instance := bucket.New(creator, CreateMockWhereIs(creator, content), s, c)
		assert.NotNil(t, instance)
	})

	t.Run("Bucket Resolve cid should be in content cache", func(t *testing.T) {
		content := []*bt.BucketItem{
			{
				Name:      "test",
				Uri:       objectURI,
				Timestamp: time.Now().Unix(),
				Type:      bt.ResourceIdentifier_CID,
			},
		}

		instance := bucket.New(creator, CreateMockWhereIs(creator, content), s, c)
		assert.NotNil(t, instance)
		err := instance.ResolveContent()
		assert.NoError(t, err)
		item, err := instance.GetContentById(content[0].Uri)
		assert.NoError(t, err)
		assert.NotNil(t, item)
		assert.ObjectsAreEqual(item.ContentType, bt.ResourceIdentifier_CID)
	})

	t.Run("Bucket Service endpoint should be valid uri", func(t *testing.T) {
		content := []*bt.BucketItem{
			{
				Name:      "test",
				Uri:       objectURI,
				Timestamp: time.Now().Unix(),
				Type:      bt.ResourceIdentifier_CID,
			},
		}

		instance := bucket.New(creator, CreateMockWhereIs(creator, content), s, c)
		ssi := instance.CreateBucketServiceEndpoint()
		assert.NotNil(t, ssi)
		assert.ObjectsAreEqual(c.GetAPIAddress(), ssi.ID.Host)
	})
}
