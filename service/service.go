package service

import (
	"os"

	"github.com/sonr-io/core/common"
	v1 "github.com/sonr-io/core/service/v1"
	"google.golang.org/protobuf/proto"
)

type Service interface {
	// GetConfig returns the service configuration.
	GetConfig() *v1.ServiceConfig

	// GetName returns the service name.
	GetName() string

	// GetPath returns the service path.
	GetPath() string

	// GetDid returns the service DID.
	GetDid() *common.Did

	// AddChannel adds a channel to the service.
	AddChannel(channel *common.Did)

	// AddBucket adds a bucket to the service.
	AddBucket(bucket *common.Did)

	// AddObject adds an object to the service.
	AddObject(object *common.ObjectDoc)

	// RemoveChannel removes a channel from the service.
	RemoveChannel(channel *common.Did)

	// RemoveBucket removes a bucket from the service.
	RemoveBucket(bucket *common.Did)

	// RemoveObject removes an object from the service.
	RemoveObject(object *common.ObjectDoc)

	// GetChannels returns the service channels.
	GetChannels() []*common.Did

	// GetBuckets returns the service buckets.
	GetBuckets() []*common.Did

	// GetObjects returns the service objects.
	GetObjects() map[string]*common.ObjectDoc

	// Save saves the service configuration.
	Save() error
}

type service struct {
	Service
	config *v1.ServiceConfig
	path   string
	name   string
}

// NewService creates a new Sonr service, and creates a config for it.
// Requires name and path to be set.
func NewService(name, path string, opts ...Option) Service {
	c := &v1.ServiceConfig{}
	for _, opt := range opts {
		opt(c)
	}

	return &service{
		config: c,
		path:   path,
		name:   name,
	}
}

// GetConfig returns the service configuration.
func (s *service) GetConfig() *v1.ServiceConfig {
	return s.config
}

// GetName returns the service name.
func (s *service) GetName() string {
	return s.name
}

// GetPath returns the service path.
func (s *service) GetPath() string {
	return s.path
}

// GetDid returns the service DID.
func (s *service) GetDid() *common.Did {
	return s.config.Owner
}

// AddChannel adds a channel to the service.
func (s *service) AddChannel(d *common.Did) {
	s.config.Channels = append(s.config.Channels, d)
}

// AddBucket adds a bucket to the service.
func (s *service) AddBucket(d *common.Did) {
	s.config.Buckets = append(s.config.Buckets, d)
}

// AddObject adds an object to the service.
func (s *service) AddObject(o *common.ObjectDoc) {
	s.config.Objects[o.GetDid()] = o
}

// RemoveChannel removes a channel from the service.
func (s *service) RemoveChannel(d *common.Did) {
	for i, c := range s.config.Channels {
		if c.GetId() == d.GetId() {
			s.config.Channels = append(s.config.Channels[:i], s.config.Channels[i+1:]...)
			return
		}
	}
}

// RemoveBucket removes a bucket from the service.
func (s *service) RemoveBucket(d *common.Did) {
	for i, b := range s.config.Buckets {
		if b.GetId() == d.GetId() {
			s.config.Buckets = append(s.config.Buckets[:i], s.config.Buckets[i+1:]...)
			return
		}
	}
}

// RemoveObject removes an object from the service.
func (s *service) RemoveObject(o *common.ObjectDoc) {
	delete(s.config.Objects, o.GetDid())
}

// GetChannels returns the service channels.
func (s *service) GetChannels() []*common.Did {
	return s.config.Channels
}

// GetBuckets returns the service buckets.
func (s *service) GetBuckets() []*common.Did {
	return s.config.Buckets
}

// GetObjects returns the service objects.
func (s *service) GetObjects() map[string]*common.ObjectDoc {
	return s.config.Objects
}

// Save saves the service configuration.
func (s *service) Save() error {
	buf, err := proto.Marshal(s.config)
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, buf, 0644)
}
