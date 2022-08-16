package bucket

/*
	Underlying api definition of Buckets
	Higher level APIs implementing bucket features

	Version: 0.1.0
*/
type Bucket interface {
	/*
		Retrieves a piece of content by the given uri
	*/
	GetContent(id string) (*BucketContent, error)

	/*
		Checks if a given uri for existence in the given bucket
	*/
	ContentExists(id string) bool

	/*
		Checks if a given uri is for a bucket
	*/
	IsBucket(id string) bool

	/*
		Checks if a given uri is for an object
	*/
	IsContent(id string) bool

	ResolveBuckets(address string) error

	ResolveContent() error
}
