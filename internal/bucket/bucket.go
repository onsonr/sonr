package bucket

import (
	"errors"

	bt "github.com/sonr-io/sonr/x/bucket/types"
)

var (
	errContentNotFound = errors.New("content with id not found")
)

type Bucket interface {
	GetCotent(id string) (interface{}, error)
	ContentExists(id string) error
	IsBucket(id string) bool
	IsContent(id string) bool
}

type BucketContent struct {
	Item        interface{}
	Id          string
	ContentType string
}
type bucketImpl struct {
	whereIs      bt.WhereIs
	contentCache map[string]*BucketContent
}
