package bucket_test

import (
	"testing"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/bucket"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	"google.golang.org/grpc"
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
	creator := ""
	s := shell.NewLocalShell()
	grpcClient, err := grpc.Dial(
		"localhost:9090",
		grpc.WithInsecure(),
	)

	if err != nil {
		panic(err.Error())
	}

	queryClient := bt.NewQueryClient(grpcClient)

	t.Run("Bucket should be defined", func(t *testing.T) {
		content := []*bt.BucketItem{
			{
				Name:      "test",
				Uri:       "bafyreihnj3feeesb6wmd46lmsvtwalvuckns647ghy44xn63lfsfed3ydm",
				Timestamp: time.Now().Unix(),
				Type:      bt.ResourceIdentifier_CID,
			},
		}
		instance := bucket.New(creator, CreateMockWhereIs(creator, content), s, queryClient)
		instance.ResolveContent()
	})
}
