package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/motor/x/object"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

var (
	objectBuilders map[string]*object.ObjectBuilder
)

type BucketObjectItem struct {
	Uri        string
	Name       string
	BucketDid  string
	ObjectData []byte
	SchemaDid  string
}

type ObjectResolver struct {
	name      string
	schemaDid string
	LastSync  int64
}

func NewObjectResolver(name, schemaDid string) (*ObjectResolver, error) {
	if _, ok := objectBuilders[name]; ok {
		return nil, fmt.Errorf("object resolver exists with name '%s'", name)
	}

	builder, err := instance.NewObjectBuilder(schemaDid)
	if err != nil {
		return nil, err
	}

	resolver := &ObjectResolver{
		schemaDid: schemaDid,
		name:      name,
	}

	objectBuilders[name] = builder
	return resolver, nil
}

func (r *ObjectResolver) SetLabel(label string) error {
	if instance == nil {
		return errWalletNotExists
	}
	return SetObjectLabel(r.name, label)
}

func (r *ObjectResolver) SetBool(fieldName string, v int) error {
	if instance == nil {
		return errWalletNotExists
	}
	return SetBool(r.name, fieldName, v)
}

func (r *ObjectResolver) SetInt(fieldName string, v int) error {
	if instance == nil {
		return errWalletNotExists
	}
	return SetInt(r.name, fieldName, v)
}

func (r *ObjectResolver) SetFloat(name, fieldName string, v float32) error {
	if instance == nil {
		return errWalletNotExists
	}
	return SetFloat(name, fieldName, v)
}

func (r *ObjectResolver) SetString(fieldName string, v string) error {
	if instance == nil {
		return errWalletNotExists
	}
	return SetString(r.name, fieldName, v)
}

func (r *ObjectResolver) SetBytes(fieldName string, v []byte) error {
	if instance == nil {
		return errWalletNotExists
	}
	return SetBytes(r.name, fieldName, v)
}

func (r *ObjectResolver) SetLink(fieldName, did string) error {
	if instance == nil {
		return errWalletNotExists
	}
	return SetLink(r.name, fieldName, did)
}

func (r *ObjectResolver) RemoveField(fieldName string) error {
	if instance == nil {
		return errWalletNotExists
	}
	return RemoveObjectField(r.name, fieldName)
}

func (r *ObjectResolver) Publish() ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	return UploadObject(r.name)
}

func (r *ObjectResolver) Validate() ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	return BuildObject(r.name)
}

type BucketResolver struct {
	did               string
	Initialized       bool
	LastSync          int64
	bucketObjectsChan chan *bt.BucketItem
	bucketContentMap  map[string]*BucketObjectItem
}

func NewBucketResolver(did string) (*BucketResolver, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	if err := ResolveBucket(did); err != nil {
		return nil, err
	}

	resolver := &BucketResolver{
		did: did,
	}
	go resolver.backgroundBucketObjectsRefresh()
	return resolver, nil
}

func (r *BucketResolver) ListObjects() ([]string, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	if !r.Initialized {
		go r.backgroundBucketObjectsRefresh()
		return nil, fmt.Errorf("bucket '%s' not initialized", r.did)
	}

	jsonList := make([]string, 0)
	for _, item := range r.bucketContentMap {
		jsonList = append(jsonList, string(item.ObjectData))
	}
	return jsonList, nil
}

// backgroundBucketObjectsRefresh refreshes the bucket objects in the background
func (r *BucketResolver) backgroundBucketObjectsRefresh() {
	r.bucketObjectsChan = make(chan *bt.BucketItem, 100)
	r.bucketContentMap = make(map[string]*BucketObjectItem)
	doneChan := make(chan bool)
	bucket, err := instance.GetBucket(r.did)
	if err != nil {
		return
	}

	go func() {
		for {
			select {
			case item := <-r.bucketObjectsChan:
				if item == nil {
					return
				}

				c, err := bucket.GetContentById(item.Uri)
				if err != nil {
					continue
				}

				if c.ContentType != bt.ResourceIdentifier_CID {
					continue
				}

				objbz, err := GetObject(item.Uri)
				if err != nil {
					continue
				}

				buo := &BucketObjectItem{
					Uri:        item.Uri,
					Name:       item.Name,
					BucketDid:  r.did,
					SchemaDid:  item.SchemaDid,
					ObjectData: objbz,
				}
				r.bucketContentMap[item.Uri] = buo
			case <-doneChan:
				r.Initialized = true
				return
			}
		}
	}()

	go func() {
		for {
			items := bucket.GetBucketItems()
			for _, item := range items {
				r.bucketObjectsChan <- item
			}
			doneChan <- true
		}
	}()
}
