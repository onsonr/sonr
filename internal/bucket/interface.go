package bucket

import (
	"github.com/sonr-io/sonr/pkg/did"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

/*
	Underlying api definition of Buckets
	Higher level APIs implementing bucket features

	Version: 0.1.0
*/
type Bucket interface {
	/*
		Retrieves a piece of content by the given uri
	*/
	GetContentById(id string) (*mt.BucketContent, error)

	GetCreator() string

	/*
		Access the `items` of the `WhereIs`
	*/
	GetBucketItems() []*bt.BucketItem

	GetBuckets() []Bucket

	GetContent() []*mt.BucketContent

	GetDID() string

	GetRole() bt.BucketRole

	GetVisibility() bt.BucketVisibility

	/*
		Checks if a given uri for existence in the given bucket
	*/
	ContentExists(id string) bool

	/*
		Creates a did service endpoint for querying a wrapped `WhereIs`
	*/
	CreateBucketServiceEndpoint(baseURI string) did.Service

	/*
		Checks if a given uri is for a bucket
	*/
	IsBucket(id string) bool

	/*
		Checks if a given uri is for an object
	*/
	IsContent(id string) bool

	/*
		Resolves all buckets defined within the current WhereIs by `did`
	*/
	ResolveBuckets() error

	/*
		Resolves all content within the bucket by `cid`
	*/
	ResolveContent() error

	ResolveContentBySchema(did string) ([]*mt.BucketContent, error)
}
