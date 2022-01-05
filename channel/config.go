package channel

import (
	"fmt"
	"strings"
	"time"
)

type ID string

// Equals returns true if the IDs are equal.
func (id ID) Equals(o ID) bool {
	return id.String() == o.String()
}

// Key returns the key with ID prefix.
func (id ID) Key(k string) string {
	return fmt.Sprintf("%s/%s", id, strings.ToLower(k))
}

// Prefix returns Golog child for Logging
func (id ID) Prefix() string {
	return fmt.Sprintf("beam/%s", id.String())
}

// String returns the ID as a string.
func (id ID) String() string {
	return strings.ToLower(string(id))
}

// Option is a function that modifies the beam options.
type Option func(*options)

// WithTTL sets the time-to-live for the beam store entries
func WithTTL(ttl time.Duration) Option {
	return func(o *options) {
		o.ttl = ttl
	}
}

// WithCapacity sets the capacity of the beam store.
func WithCapacity(capacity int) Option {
	return func(o *options) {
		o.capacity = capacity
	}
}

// options is a collection of options for the beam.
type options struct {
	ttl      time.Duration
	capacity int
}

// defaultOptions is the default options for the beam.
func defaultOptions() *options {
	return &options{
		ttl:      time.Minute * 10,
		capacity: 4096,
	}
}
