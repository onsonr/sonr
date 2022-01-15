package service

import (
	v1 "github.com/sonr-io/core/service/v1"
)

// NewServiceConfig creates a new ServiceConfig with the provided options.
func NewServiceConfig(opts ...Option) *v1.ServiceConfig {
	c := &v1.ServiceConfig{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
