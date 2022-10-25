package bucket

import (
	"github.com/sonr-io/sonr/pkg/did"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

/*
	Underlying api definition of Buckets
	Higher level APIs implementing bucket features

	Version: 0.1.0
*/
type BucketClient interface {
	/*
		Returns all buckets referenced by a `BucketItem`.
		`ResolveBuckets` should be called in order to hydrate content from `BucketItems`
		Returns []Bucket

	*/
	GetBuckets() []bt.Bucket

	/*
		Returns  All `BucketContent` defined by a `BucketItem`.
		`ResolveContent` should be called in order to hydrate content from `BucketItems`
		Returns `[]BucketContent`
	*/
	GetContent() []*bt.BucketContent

	/*
		Returns the DID of the bucket
		Returns string
	*/
	GetDID() string


	GetVisibility() bt.BucketVisibility

	/*
		Checks if a given uri for existence in the given bucket
	*/
	ContentExists(id string) bool

	/*
		Creates a did service endpoint for querying a wrapped `WhereIs`
	*/
	CreateBucketServiceEndpoint() did.Service

	/*
		Checks if a given uri is for a bucket
		Returns boolean
	*/
	IsBucket(id string) bool

	/*
		Checks if a given uri is for a document
		Returns boolean
	*/
	IsContent(id string) bool

	/*
		Resolves all buckets defined within the current WhereIs by `did`
		Required before calling `GetBuckets`
		Returns error
	*/
	ResolveBuckets() error

	/*
		Resolves all content within the bucket by `cid`
		Required before calling `GetContent` or `GetContentById`
	*/
	ResolveContent() error

	/*
		Matched `schemaDid`
	*/
	ResolveContentBySchema(did string) ([]*bt.BucketContent, error)
}
