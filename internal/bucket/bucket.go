package bucket

import (
	"errors"

	shell "github.com/ipfs/go-ipfs-api"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	"google.golang.org/grpc"
)

var (
	errContentNotFound = errors.New("content with id not found")
)

type Bucket interface {
	GetContent(id string) (interface{}, error)
	ContentExists(id string) error
	IsBucket(id string) bool
	IsContent(id string) bool
}

type BucketContent struct {
	Item        interface{}
	Id          string
	ContentType bt.ResourceIdentifier
}
type bucketImpl struct {
	adress       string
	whereIs      *bt.WhereIs
	contentCache map[string]*BucketContent
	shell        *shell.Shell
	queryClient  bt.QueryClient
}

func New(address string, whereIs *bt.WhereIs, shell *shell.Shell, grpc *grpc.ClientConn) Bucket {
	queryClient := bt.NewQueryClient(grpc)
	return &bucketImpl{
		adress:       address,
		whereIs:      whereIs,
		shell:        shell,
		contentCache: make(map[string]*BucketContent),
		queryClient:  queryClient,
	}
}
